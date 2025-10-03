package migrations

import (
	"log"
	"time"

	"app-futbol/src/schemas"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) error {
	timestamp := func() string { return time.Now().Format("2006/01/02 15:04:05") }

	models := []interface{}{
		&schemas.Rol{},
		&schemas.TipoPago{},
		&schemas.Fecha{},
		&schemas.Usuario{},
		&schemas.Asistencia{},
		&schemas.Pago{},
	}
	for _, m := range models {
		if err := db.AutoMigrate(m); err != nil {
			log.Printf("%s ❌ Error migrando %T: %v", timestamp(), m, err)
			return err
		}
		log.Printf("%s ✅ Migración realizada: %T", timestamp(), m)
	}

	roles := []schemas.Rol{
		{IdRol: 1, Rol: "Usuario"},
		{IdRol: 2, Rol: "Admin"},
	}
	for _, r := range roles {
		var existing schemas.Rol
		err := db.First(&existing, "id_rol = ?", r.IdRol).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				if err = db.Create(&r).Error; err != nil {
					log.Printf("%s ❌ Error creando rol %v: %v", timestamp(), r, err)
					return err
				}
				log.Printf("%s ✅ Rol creado: %v", timestamp(), r)
			} else {
				return err
			}
		}
	}

	log.Printf("%s 🎉 Todas las migraciones y seeds se realizaron con éxito", timestamp())
	return nil
}
