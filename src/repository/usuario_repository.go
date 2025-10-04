package repository

import (
	"errors"
	"strings"

	"app-futbol/src/schemas"

	"gorm.io/gorm"
)

// ValidarCorreoExistente revisa si el correo ya está registrado.
// Retorna error si el correo existe o si ocurre un problema, con mensaje seguro.
func ValidarCorreoExistente(db *gorm.DB, correo string) error {
	correo = strings.TrimSpace(strings.ToLower(correo))
	var exists int64

	// Verificamos si ya existe
	if err := db.Model(&schemas.Usuario{}).
		Where("correo = ?", correo).
		Count(&exists).Error; err != nil {
		// Mensaje genérico seguro para errores inesperados
		return errors.New("error al procesar la solicitud, intente más tarde")
	}

	if exists > 0 {
		return errors.New("correo ya registrado")
	}

	return nil
}
