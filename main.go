package main

import (
	"log"

	"app-futbol/config"
	"app-futbol/database"
	_ "app-futbol/docs"
	"app-futbol/migrations"
	"app-futbol/src/middlewares"
	"app-futbol/src/routes"
	"app-futbol/src/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func main() {
	// Cargar configuración
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal("❌ Error cargando la configuración:", err)
	}
	middlewares.InitJWT(cfg.JWTSecret)
	// Crear conexión a la base de datos
	db, err := database.NewDatabaseConnection(cfg)
	if err != nil {
		log.Fatal("❌ Error conectando a la base de datos:", err)
	}

	// Ejecutar migraciones
	if err := migrations.RunMigrations(db); err != nil {
		log.Fatal("❌ Error ejecutando migraciones:", err)
	}

	// Inicializar Fiber
	app := fiber.New()

	// Inicializar servicios
	rolService := services.NewRolService(db)
	usuarioService := services.NewUsuarioService(db)
	// Configurar rutas desde el paquete routes
	routes.SetupRolRoutes(app, rolService)
	routes.SetupUsuarioRoutes(app, usuarioService)

	// Documentación Swagger
	app.Get("/docs/*", swagger.HandlerDefault)

	// Ruta raíz
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("¡Hola! La API está funcionando.")
	})

	// Levantar servidor
	log.Fatal(app.Listen(":3000"))
}
