package controllers

import (
	"app-futbol/src/helpers"
	"app-futbol/src/schemas"
	"app-futbol/src/services"
	"app-futbol/src/validation"

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

	// Parseamos el body
	if err := helpers.ParseBody(ctx, &usuario); err != nil {
		return err
	}

	// Lógica de negocio en el servicio
	if err := c.Service.RequestRegister(&usuario); err != nil {
		// Si es un error de validación
		if em, ok := err.(validation.ErrorMap); ok {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"estado": false,
				"errors": em,
			})
		}
		// Otro error inesperado
		return helpers.JsonInternalError(ctx, err)
	}

	return helpers.JsonSuccess(ctx, "Solicitud de registro enviada correctamente")
}

// Login maneja POST /usuarios/login
func (c *UsuarioController) Login(ctx *fiber.Ctx) error {
	var body struct {
		Correo   string `json:"correo"`
		Password string `json:"password"`
	}

	// Usamos ParseBody para no repetir lógica
	if err := helpers.ParseBody(ctx, &body); err != nil {
		return err
	}

	token, err := c.Service.Login(body.Correo, body.Password)
	if err != nil {
		return helpers.JsonError(ctx, "Credenciales inválidas")
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"estado": true,
		"token":  token,
	})
}
