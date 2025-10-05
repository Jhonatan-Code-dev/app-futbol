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
	usuario := new(schemas.Usuario)
	if err := ctx.BodyParser(usuario); err != nil {
		return helpers.Fail(ctx, fiber.StatusBadRequest, "Datos inválidos")
	}
	usuarioData := &schemas.Usuario{
		Nombre:   usuario.Nombre,
		Apellido: usuario.Apellido,
		Correo:   usuario.Correo,
		Pass:     usuario.Pass,
	}
	if errores := c.Service.RequestRegister(usuarioData); errores != nil {
		return helpers.FailWithErrors(ctx, fiber.StatusBadRequest, errores)
	}
	return helpers.Success(ctx, fiber.StatusCreated, "Usuario registrado correctamente")
}
