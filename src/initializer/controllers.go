package initializer

import (
	"app-futbol/src/controllers"
	"app-futbol/src/services"

	"gorm.io/gorm"
)

type Controllers struct {
	RolController     *controllers.RolController
	UsuarioController *controllers.UsuarioController
}

func NewControllers(db *gorm.DB) *Controllers {
	rolService := services.NewRolService(db)
	usuarioService := services.NewUsuarioService(db)

	return &Controllers{
		RolController:     controllers.NewRolController(rolService),
		UsuarioController: controllers.NewUsuarioController(usuarioService),
	}
}
