package repository

import (
	"errors"
	"strings"

	"app-futbol/src/schemas"

	"gorm.io/gorm"
)

func ValidarCorreoExistente(db *gorm.DB, correo string) error {
	correo = strings.TrimSpace(strings.ToLower(correo))
	var dummy int64
	tx := db.Model(&schemas.Usuario{}).
		Select("1").
		Where("correo = ?", correo).
		Limit(1).
		Scan(&dummy)

	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected > 0 {
		return errors.New("correo ya registrado")
	}
	return nil
}
