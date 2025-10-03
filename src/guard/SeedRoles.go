package guard

import (
	"log"
	"time"

	"app-futbol/src/schemas"

	"gorm.io/gorm"
)

func SeedRoles(db *gorm.DB) error {
	timestamp := func() string { return time.Now().Format("2006/01/02 15:04:05") }

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

	return nil
}
