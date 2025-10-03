//go:build wireinject
// +build wireinject

package di

import (
	"log"

	"app-futbol/src/config"
	"app-futbol/src/controllers"
	"app-futbol/src/database"
	"app-futbol/src/services"

	"github.com/google/wire"
	"gorm.io/gorm"
)

type AppContainer struct {
	Config        *config.Config
	DB            *gorm.DB
	RolService    *services.RolService
	RolController *controllers.RolController
}

// Función pública: la usas en main.go
func InitializeApp() *AppContainer {
	container, err := initializeApp()
	if err != nil {
		log.Fatalf("❌ Error inicializando dependencias: %v", err)
	}
	return container
}

func initializeApp() (*AppContainer, error) {
	wire.Build(
		// Dependencias base
		config.NewConfig,
		database.NewDatabase,

		// Services
		services.NewRolService,

		// Controllers
		controllers.NewRolController,

		// Construye AppContainer con todo
		wire.Struct(new(AppContainer), "*"),
	)
	return nil, nil
}
