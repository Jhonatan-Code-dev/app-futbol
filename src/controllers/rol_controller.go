package controllers

import (
	"strconv"

	"app-futbol/src/schemas"
	"app-futbol/src/services"

	"github.com/gofiber/fiber/v2"
)

type RolController struct {
	Service *services.RolService
}

func NewRolController(service *services.RolService) *RolController {
	return &RolController{Service: service}
}

// CreateRol crea un nuevo rol
func (c *RolController) CreateRol(ctx *fiber.Ctx) error {
	var rol schemas.Rol
	if err := ctx.BodyParser(&rol); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := c.Service.Create(&rol); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(rol)
}

// GetRoles lista todos los roles
func (c *RolController) GetRoles(ctx *fiber.Ctx) error {
	roles, err := c.Service.GetAll()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(roles)
}

// GetRol obtiene un rol por ID
func (c *RolController) GetRol(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID inválido"})
	}

	rol, err := c.Service.GetByID(uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Rol no encontrado"})
	}
	return ctx.JSON(rol)
}

// UpdateRol actualiza un rol por ID
func (c *RolController) UpdateRol(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID inválido"})
	}

	var rol schemas.Rol
	if err := ctx.BodyParser(&rol); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	rol.IdRol = uint(id)
	if err := c.Service.Update(&rol); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(rol)
}

// DeleteRol elimina un rol por ID
func (c *RolController) DeleteRol(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID inválido"})
	}

	if err := c.Service.Delete(uint(id)); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.SendStatus(fiber.StatusNoContent)
}
