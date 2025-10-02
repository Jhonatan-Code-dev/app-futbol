package main

import (
	"log"

	_ "app-futbol/docs"
	"app-futbol/migrations"
	"app-futbol/src/middlewares"
	"app-futbol/src/routes"
	"app-futbol/src/services"

	"app-futbol/config"
	"app-futbol/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func main() {
	// Inicializar config y JWT
	cfg := config.GetConfig()
	middlewares.InitJWT(cfg.JWTSecret)

	// Inicializar DB global
	database.InitDatabase()

	// Ejecutar migraciones sin pasar db
	if err := migrations.RunMigrations(); err != nil {
		log.Fatal("❌ Error ejecutando migraciones:", err)
	}

	app := fiber.New()

	// Inicializar servicios con DB global
	rolService := services.NewRolService(database.GetDB())
	usuarioService := services.NewUsuarioService(database.GetDB())

	// Configurar rutas
	routes.SetupRolRoutes(app, rolService)
	routes.SetupUsuarioRoutes(app, usuarioService)

	// Swagger
	app.Get("/docs/*", swagger.HandlerDefault)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("¡Hola! La API está funcionando.")
	})

	log.Fatal(app.Listen(":3000"))
}
