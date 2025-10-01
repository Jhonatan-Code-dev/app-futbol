package migrations

import (
	"log"
	"time"

	"app-futbol/src/schemas"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) error {
	orderedModels := []interface{}{
		&schemas.Rol{},
		&schemas.TipoPago{},
		&schemas.Fecha{},
		&schemas.Usuario{},
		&schemas.Asistencia{},
		&schemas.Pago{},
	}

	for _, m := range orderedModels {
		if err := db.AutoMigrate(m); err != nil {
			log.Printf("%s ❌ Error migrando %T: %v", time.Now().Format("2006/01/02 15:04:05"), m, err)
			return err
		}
		log.Printf("%s ✅ Migración realizada: %T", time.Now().Format("2006/01/02 15:04:05"), m)
	}

	log.Printf("%s 🎉 Todas las migraciones se realizaron con éxito", time.Now().Format("2006/01/02 15:04:05"))
	return nil
}
