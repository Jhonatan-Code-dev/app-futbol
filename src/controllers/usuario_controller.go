package controllers

import (
	"app-futbol/src/helpers"
	"app-futbol/src/schemas"
	"app-futbol/src/services"

	"github.com/gofiber/fiber/v2"
)

type UsuarioController struct {
	Service *services.UsuarioService
}

func NewUsuarioController(service *services.UsuarioService) *UsuarioController {
	return &UsuarioController{Service: service}
}

func (c *UsuarioController) SolicitarRegistro(ctx *fiber.Ctx) error {
	var usuario schemas.Usuario
	if err := helpers.ParseBody(ctx, &usuario); err != nil {
		return err
	}
	return helpers.JsonSuccess(ctx, "Solicitud de registro enviada correctamente")
}

func (c *UsuarioController) Login(ctx *fiber.Ctx) error {
	var body struct {
		Correo   string `json:"correo"`
		Password string `json:"password"`
	}
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
