package main

import (
	"log"

	_ "app-futbol/docs"
	"app-futbol/src/di"
	"app-futbol/src/migrations"
	"app-futbol/src/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func main() {
	container := di.InitializeApp()

	migrations.RunMigrations(container.DB)

	app := fiber.New()
	routes.SetupRoutes(app, container)
	app.Get("/docs/*", swagger.HandlerDefault)
	log.Printf("🚀 Servidor iniciado en http://localhost:%s", container.Config.Port)
	log.Fatal(app.Listen(":" + container.Config.Port))
}
