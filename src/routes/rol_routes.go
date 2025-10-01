package routes

import (
	"app-futbol/src/controllers"
	"app-futbol/src/services"

	"github.com/gofiber/fiber/v2"
)

// SetupRolRoutes configura todas las rutas relacionadas con Roles.
// Recibe la instancia de Fiber y la base de datos para inyectar el servicio.
func SetupRolRoutes(app *fiber.App, rolService *services.RolService) {
	rolController := controllers.NewRolController(rolService)

	api := app.Group("/api/v1")
	roles := api.Group("/roles")

	roles.Post("/", rolController.CreateRol)
	roles.Get("/", rolController.GetRoles)
	roles.Get("/:id", rolController.GetRol)
	roles.Put("/:id", rolController.UpdateRol)
	roles.Delete("/:id", rolController.DeleteRol)
}
