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
func (c *UsuarioController) SolicitarRegistro(ctx *fiber.Ctx) error {
	var usuario schemas.Usuario
	if err := ctx.BodyParser(&usuario); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Fecha de solicitud se asigna en el servicio
	if err := c.Service.RequestRegister(&usuario); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(fiber.Map{
		"mensaje":    "Solicitud de registro enviada",
		"id_usuario": usuario.IdUsuario,
	})
}

// Login maneja POST /usuarios/login
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
