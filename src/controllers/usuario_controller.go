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

// controllers/usuario_controller.go
func (c *UsuarioController) SolicitarRegistro(ctx *fiber.Ctx) error {
	usuario := new(schemas.Usuario)
	if err := ctx.BodyParser(usuario); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Datos inválidos",
		})
	}

	// Solo considerar los campos que realmente recibimos
	usuario = &schemas.Usuario{
		Nombre:   usuario.Nombre,
		Apellido: usuario.Apellido,
		Correo:   usuario.Correo,
		Pass:     usuario.Pass,
	}

	if errores := c.Service.RequestRegister(usuario); errores != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"errores": errores,
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Usuario registrado correctamente",
	})
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
