//go:build wireinject
// +build wireinject

package di

import (
	"log"

	"app-futbol/config"
	"app-futbol/database"

	"github.com/google/wire"
	"gorm.io/gorm"
)

type AppContainer struct {
	Config *config.Config
	DB     *gorm.DB
}

func InitializeApp() *AppContainer {
	container, err := initializeApp()
	if err != nil {
		log.Fatalf("❌ Error inicializando dependencias: %v", err)
	}
	return container
}

func initializeApp() (*AppContainer, error) {
	wire.Build(
		config.NewConfig,
		database.NewDatabase,
		wire.Struct(new(AppContainer), "*"),
	)
	return nil, nil
}
