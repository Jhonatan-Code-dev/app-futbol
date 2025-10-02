package main

import (
	"log"

	_ "app-futbol/docs"
	"app-futbol/migrations"
	"app-futbol/src/initializer"
	"app-futbol/src/middlewares"
	"app-futbol/src/routes"

	"app-futbol/config"
	"app-futbol/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func main() {
	// Inicializar configuración y JWT
	cfg := config.GetConfig()
	middlewares.InitJWT(cfg.JWTSecret)

	// Inicializar base de datos
	db := database.InitDatabase()

	// Ejecutar migraciones con la conexión actual
	migrations.RunMigrations(db)

	// Inicializar controladores con inyección de la BD
	controllers := initializer.NewControllers(db)

	// Inicializar servidor Fiber
	app := fiber.New()

	// Configurar rutas
	routes.SetupRoutes(app, controllers)

	// Swagger
	app.Get("/docs/*", swagger.HandlerDefault)

	// Ruta base
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("¡Hola! La API está funcionando 🚀")
	})

	// Iniciar servidor
	log.Fatal(app.Listen(":3000"))
}
