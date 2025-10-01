package middlewares

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var SecretKey []byte

// InitJWT inicializa el SecretKey desde la configuración
func InitJWT(secret string) {
	SecretKey = []byte(secret)
}

// -------------------------
// Generación de token
// -------------------------
func GenerateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(), // Expira en 7 días
		"iat":     time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(SecretKey)
}

// -------------------------
// Validación de token
// -------------------------
func ValidateToken(tokenStr string) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}))
}

// -------------------------
// Renovación de token
// -------------------------
func RefreshToken(tokenStr string) (string, error) {
	token, err := ValidateToken(tokenStr)
	if err != nil || !token.Valid {
		return "", errors.New("token inválido")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("no se pudo leer claims")
	}

	claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix() // Nueva expiración

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return newToken.SignedString(SecretKey)
}

// -------------------------
// Middleware Fiber para proteger rutas
// -------------------------
func JWTProtect() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenStr := c.Get("Authorization") // "Bearer <token>"
		if tokenStr == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token requerido"})
		}

		if len(tokenStr) > 7 && tokenStr[:7] == "Bearer " {
			tokenStr = tokenStr[7:]
		}

		token, err := ValidateToken(tokenStr)
		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token inválido"})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Claims inválidos"})
		}

		// Solo user_id
		c.Locals("user_id", claims["user_id"])

		return c.Next()
	}
}
