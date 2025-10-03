package main

import (
	"log"

	"app-futbol/di"
	_ "app-futbol/docs"
	"app-futbol/src/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func main() {

	container := di.InitializeApp()

	// Inicializar servidor Fiber
	app := fiber.New()

	// Configurar rutas usando el container
	routes.SetupRoutes(app, container)

	// Swagger
	app.Get("/docs/*", swagger.HandlerDefault)

	// Ruta base
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("¡Hola! La API está funcionando 🚀")
	})

	// Iniciar servidor en el puerto configurado desde container.Config
	log.Printf("🚀 Servidor iniciado en http://localhost:%s", container.Config.Port)
	log.Fatal(app.Listen(":" + container.Config.Port))
}
