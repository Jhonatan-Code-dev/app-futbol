package controllers

import (
	"strconv"

	"app-futbol/src/schemas"
	"app-futbol/src/services"

	"github.com/gofiber/fiber/v2"
)

// RolController se encarga de recibir solicitudes HTTP
// y delegarlas al RolService para manejar la lógica de negocio.
type RolController struct {
	Service *services.RolService
}

// NewRolController crea un nuevo RolController con el servicio inyectado.
func NewRolController(service *services.RolService) *RolController {
	return &RolController{Service: service}
}

// CreateRol crea un nuevo rol
// @Summary Crear un nuevo rol
// @Description Crea un rol en la base de datos con los datos proporcionados
// @Tags Roles
// @Accept json
// @Produce json
// @Param rol body schemas.Rol true "Datos del rol"
// @Success 200 {object} schemas.Rol
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/roles [post]
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
// @Summary Listar todos los roles
// @Description Obtiene un listado completo de roles
// @Tags Roles
// @Produce json
// @Success 200 {array} schemas.Rol
// @Failure 500 {object} map[string]string
// @Router /api/v1/roles [get]
func (c *RolController) GetRoles(ctx *fiber.Ctx) error {
	roles, err := c.Service.GetAll()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(roles)
}

// GetRol obtiene un rol por ID
// @Summary Obtener rol por ID
// @Description Obtiene un rol específico usando su ID
// @Tags Roles
// @Produce json
// @Param id path int true "ID del rol"
// @Success 200 {object} schemas.Rol
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/v1/roles/{id} [get]
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
// @Summary Actualizar rol
// @Description Actualiza los datos de un rol existente
// @Tags Roles
// @Accept json
// @Produce json
// @Param id path int true "ID del rol"
// @Param rol body schemas.Rol true "Datos actualizados del rol"
// @Success 200 {object} schemas.Rol
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/roles/{id} [put]
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
// @Summary Eliminar rol
// @Description Elimina un rol de la base de datos usando su ID
// @Tags Roles
// @Param id path int true "ID del rol"
// @Success 204 "No Content"
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/roles/{id} [delete]
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
