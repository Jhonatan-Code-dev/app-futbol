package migrations

import (
	"log"
	"time"

	"app-futbol/database"
	"app-futbol/src/schemas"

	"gorm.io/gorm"
)

// RunMigrations ejecuta AutoMigrate y luego inserta roles iniciales usando la DB global
func RunMigrations() error {
	db := database.GetDB() // obtenemos la conexión global directamente

	// Orden de migración de modelos
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

	// Insertar roles iniciales con ID fijo si no existen
	initialRoles := []schemas.Rol{
		{IdRol: 1, Rol: "Usuario"},
		{IdRol: 2, Rol: "Admin"},
	}

	for _, r := range initialRoles {
		var existing schemas.Rol
		if err := db.First(&existing, "id_rol = ?", r.IdRol).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				if err := db.Create(&r).Error; err != nil {
					log.Printf("%s ❌ Error creando rol %v: %v", time.Now().Format("2006/01/02 15:04:05"), r, err)
					return err
				}
				log.Printf("%s ✅ Rol creado: %v", time.Now().Format("2006/01/02 15:04:05"), r)
			} else {
				return err
			}
		}
	}

	log.Printf("%s 🎉 Todas las migraciones y seeds se realizaron con éxito", time.Now().Format("2006/01/02 15:04:05"))
	return nil
}
