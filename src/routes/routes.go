package routes

import (
	"app-futbol/src/initializer"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes configura todas las rutas de la API, separadas por grupos
func SetupRoutes(app *fiber.App, c *initializer.Controllers) {
	api := app.Group("/api/v1")

	// Grupo de Roles
	roles := api.Group("/roles")
	{
		roles.Post("/", c.RolController.CreateRol)
		roles.Get("/", c.RolController.GetRoles)
		roles.Get("/:id", c.RolController.GetRol)
		roles.Put("/:id", c.RolController.UpdateRol)
		roles.Delete("/:id", c.RolController.DeleteRol)
	}

	// Grupo de Usuarios
	usuarios := api.Group("/usuarios")
	{
		usuarios.Post("/solicitar", c.UsuarioController.SolicitarRegistroHandler)
		usuarios.Post("/login", c.UsuarioController.LoginHandler)
	}
}
