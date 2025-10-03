package routes

import (
	"app-futbol/src/di"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, container *di.AppContainer) {

	rol := container.RolController
	api := app.Group("/api/v1")
	roles := api.Group("/roles")
	roles.Post("/", rol.CreateRol)
	roles.Get("/", rol.GetRoles)
	roles.Get("/:id", rol.GetRol)
	roles.Put("/:id", rol.UpdateRol)
	roles.Delete("/:id", rol.DeleteRol)

}
