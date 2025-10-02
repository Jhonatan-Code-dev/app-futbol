package services

import (
	"fmt"
	"strings"
	"time"

	"app-futbol/src/middlewares"
	"app-futbol/src/schemas"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// ValidationError para manejar errores amigables
type ValidationError struct {
	Errors map[string]string `json:"errors"`
}

func (v *ValidationError) Error() string {
	return "error de validación"
}

// UsuarioService maneja la lógica de usuarios
type UsuarioService struct {
	DB       *gorm.DB
	validate *validator.Validate
}

// Constructor
func NewUsuarioService(db *gorm.DB) *UsuarioService {
	validate := validator.New()

	// Validación personalizada para Gmail
	validate.RegisterValidation("gmail", func(fl validator.FieldLevel) bool {
		email := fl.Field().String()
		return strings.HasSuffix(email, "@gmail.com")
	})

	return &UsuarioService{
		DB:       db,
		validate: validate,
	}
}

// RequestRegister solicita el registro de un usuario
func (s *UsuarioService) RequestRegister(usuario *schemas.Usuario) error {
	// Validar datos
	if err := s.validate.Struct(usuario); err != nil {
		errorsMap := make(map[string]string)
		for _, e := range err.(validator.ValidationErrors) {
			switch e.Tag() {
			case "required":
				errorsMap[e.Field()] = "es obligatorio"
			case "min":
				errorsMap[e.Field()] = fmt.Sprintf("debe tener al menos %s caracteres", e.Param())
			case "max":
				errorsMap[e.Field()] = fmt.Sprintf("no debe exceder %s caracteres", e.Param())
			case "email":
				errorsMap[e.Field()] = "no es un correo válido"
			case "gmail":
				errorsMap[e.Field()] = "solo se permiten correos @gmail.com"
			default:
				errorsMap[e.Field()] = fmt.Sprintf("no cumple la regla %s", e.Tag())
			}
		}
		return &ValidationError{Errors: errorsMap}
	}

	// Verificar correo duplicado
	var existing schemas.Usuario
	if err := s.DB.Where("correo = ?", usuario.Correo).First(&existing).Error; err == nil {
		return fmt.Errorf("el correo %s ya está registrado", usuario.Correo)
	} else if err != gorm.ErrRecordNotFound {
		return err
	}

	// Encriptar contraseña
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(usuario.Pass), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	usuario.Pass = string(hashedPass)
	usuario.Estado = false
	usuario.FechaSolicitud = time.Now()
	usuario.FechaAceptacion = time.Time{}

	// Guardar en la DB
	if err := s.DB.Create(usuario).Error; err != nil {
		return err
	}

	return nil
}

// Login valida al usuario y genera un token JWT
func (s *UsuarioService) Login(correo, pass string) (string, error) {
	var usuario schemas.Usuario
	err := s.DB.Where("correo = ?", correo).First(&usuario).Error
	if err != nil {
		return "", fmt.Errorf("usuario no encontrado")
	}

	if !usuario.Estado {
		return "", fmt.Errorf("usuario no aprobado")
	}

	err = bcrypt.CompareHashAndPassword([]byte(usuario.Pass), []byte(pass))
	if err != nil {
		return "", fmt.Errorf("contraseña incorrecta")
	}

	token, err := middlewares.GenerateToken(usuario.IdUsuario)
	if err != nil {
		return "", fmt.Errorf("no se pudo generar token")
	}

	return token, nil
}
