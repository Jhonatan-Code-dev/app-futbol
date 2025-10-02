package controllers

import (
	"app-futbol/src/schemas"
	"app-futbol/src/services"

	"github.com/gofiber/fiber/v2"
)

// UsuarioController maneja las solicitudes HTTP relacionadas con usuarios
type UsuarioController struct {
	Service *services.UsuarioService
}

// Constructor
func NewUsuarioController(service *services.UsuarioService) *UsuarioController {
	return &UsuarioController{Service: service}
}

// SolicitarRegistro maneja POST /usuarios/solicitar
// @Summary Solicitar registro de usuario
// @Description Permite a un usuario enviar una solicitud de registro. La contraseña se encripta y el usuario queda pendiente de aprobación.
// @Tags Usuarios
// @Accept json
// @Produce json
// @Param usuario body schemas.Usuario true "Datos del usuario a registrar"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string "Error de parsing del cuerpo"
// @Failure 500 {object} map[string]string "Error interno del servidor"
// @Router /api/v1/usuarios/solicitar [post]
func (c *UsuarioController) SolicitarRegistro(ctx *fiber.Ctx) error {
	var usuario schemas.Usuario
	if err := ctx.BodyParser(&usuario); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := c.Service.RequestRegister(&usuario); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(fiber.Map{
		"mensaje":    "Solicitud de registro enviada",
		"id_usuario": usuario.IdUsuario,
	})
}

// Login maneja POST /usuarios/login
// @Summary Login de usuario
// @Description Permite a un usuario autenticarse y obtener un token JWT. Solo usuarios aprobados pueden iniciar sesión.
// @Tags Usuarios
// @Accept json
// @Produce json
// @Param credentials body object{correo=string,password=string} true "Correo y contraseña"
// @Success 200 {object} map[string]string "Token JWT generado"
// @Failure 400 {object} map[string]string "Error de parsing del cuerpo"
// @Failure 401 {object} map[string]string "Usuario no autorizado o contraseña incorrecta"
// @Router /api/v1/usuarios/login [post]
func (c *UsuarioController) Login(ctx *fiber.Ctx) error {
	var body struct {
		Correo   string `json:"correo"`
		Password string `json:"password"`
	}

	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	token, err := c.Service.Login(body.Correo, body.Password)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(fiber.Map{"token": token})
}
