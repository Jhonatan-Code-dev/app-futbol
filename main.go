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

	app := fiber.New()

	routes.SetupRoutes(app, container)

	// Swagger
	app.Get("/docs/*", swagger.HandlerDefault)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("¡Hola! La API está funcionando 🚀")
	})

	log.Printf("🚀 Servidor iniciado en http://localhost:%s", container.Config.Port)
	log.Fatal(app.Listen(":" + container.Config.Port))
}
