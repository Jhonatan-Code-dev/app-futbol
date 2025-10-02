package controllers

import (
	"app-futbol/src/schemas"
	"app-futbol/src/services"
	"errors"

	"github.com/gofiber/fiber/v2"
)

type UsuarioController struct {
	Service *services.UsuarioService
}

func NewUsuarioController(service *services.UsuarioService) *UsuarioController {
	return &UsuarioController{Service: service}
}

// SolicitarRegistro maneja POST /usuarios/solicitar
func (c *UsuarioController) SolicitarRegistro(ctx *fiber.Ctx) error {
	var usuario schemas.Usuario
	if err := ctx.BodyParser(&usuario); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"estado": false,
			"error":  "Datos inválidos: " + err.Error(),
		})
	}

	if err := c.Service.RequestRegister(&usuario); err != nil {
		var ve *services.ValidationError
		if errors.As(err, &ve) {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"estado": false,
				"errors": ve.Errors,
			})
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"estado": false,
			"error":  err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"estado":  true,
		"mensaje": "Solicitud de registro enviada correctamente",
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
