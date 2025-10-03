package seed

import (
	"log"
	"time"

	"app-futbol/src/schemas"

	"gorm.io/gorm"
)

func SeedPermissions(db *gorm.DB) error {
	timestamp := func() string { return time.Now().Format("2006/01/02 15:04:05") }
	permissions := []schemas.Permiso{
		{Permiso: "rol_leer", Descripcion: "Puede leer los roles"},
		{Permiso: "rol_crear", Descripcion: "Puede crear nuevos roles"},
		{Permiso: "rol_actualizar", Descripcion: "Puede actualizar roles existentes"},
		{Permiso: "rol_eliminar", Descripcion: "Puede eliminar roles"},
		{Permiso: "usuario_leer", Descripcion: "Puede leer usuarios"},
		{Permiso: "usuario_crear", Descripcion: "Puede crear usuarios"},
		{Permiso: "usuario_actualizar", Descripcion: "Puede actualizar usuarios"},
		{Permiso: "usuario_eliminar", Descripcion: "Puede eliminar usuarios"},
	}

	for _, p := range permissions {
		var existing schemas.Permiso
		err := db.First(&existing, "permiso = ?", p.Permiso).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				if err = db.Create(&p).Error; err != nil {
					log.Printf("%s ❌ Error creando permiso %v: %v", timestamp(), p, err)
					return err
				}
				log.Printf("%s ✅ Permiso creado: %v", timestamp(), p)
			} else {
				return err
			}
		}
	}

	return nil
}
