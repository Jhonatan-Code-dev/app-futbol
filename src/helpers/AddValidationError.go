package helpers

import (
	"app-futbol/src/validation"
	"log"

	"gorm.io/gorm"
)

// AddValidationError aplica una función de validación a un campo y agrega el error al mapa si existe.
func AddValidationError(errores map[string]string, fieldName, value string, validator func(string) error) {
	if err := validator(value); err != nil {
		errores[fieldName] = err.Error()
	}
}

func AddDBExistenceError(errores map[string]string, fieldName, value string, db *gorm.DB, checkFunc func(*gorm.DB, string) error) {
	if err := checkFunc(db, value); err != nil {
		errores[fieldName] = err.Error()
	}
}

// HashPassword agrega al mapa de errores si falla y retorna el hash seguro
func HashPassword(errores map[string]string, fieldName, pass string) string {
	hash, err := validation.HashPass(pass)
	if err != nil {
		errores[fieldName] = err.Error()
		return ""
	}
	return hash
}

func SafeError(errores map[string]string, fieldName string, err error) {
	if err != nil {
		log.Printf("Error interno [%s]: %v", fieldName, err)    // log interno
		errores[fieldName] = "Error interno, intente más tarde" // mensaje genérico al cliente
	}
}
