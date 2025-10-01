package routes

import (
	"app-futbol/src/controllers"
	"app-futbol/src/services"

	"github.com/gofiber/fiber/v2"
)

// SetupUsuarioRoutes define todas las rutas relacionadas a usuarios
func SetupUsuarioRoutes(app *fiber.App, usuarioService *services.UsuarioService) {
	usuarioController := controllers.NewUsuarioController(usuarioService)
	api := app.Group("/api/v1")
	usuarios := api.Group("/usuarios")
	usuarios.Post("/solicitar", usuarioController.SolicitarRegistro)
	usuarios.Post("/login", usuarioController.Login)
}
