package migrations

import (
	"log"
	"time"

	"app-futbol/src/guard"
	"app-futbol/src/schemas"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) error {
	timestamp := func() string { return time.Now().Format("2006/01/02 15:04:05") }

	// Tablas
	models := []interface{}{
		&schemas.Rol{},
		&schemas.TipoPago{},
		&schemas.Fecha{},
		&schemas.Usuario{},
		&schemas.Asistencia{},
		&schemas.Pago{},
		&schemas.Permiso{},
		&schemas.RolPermiso{},
	}

	for _, m := range models {
		if err := db.AutoMigrate(m); err != nil {
			log.Printf("%s ❌ Error migrando %T: %v", timestamp(), m, err)
			return err
		}
		log.Printf("%s ✅ Migración realizada: %T", timestamp(), m)
	}

	// Llamar a los seeds
	if err := guard.SeedRoles(db); err != nil {
		log.Printf("%s ❌ Error en seed de roles: %v", timestamp(), err)
		return err
	}

	log.Printf("%s 🎉 Todas las migraciones y seeds se realizaron con éxito", timestamp())
	return nil
}
