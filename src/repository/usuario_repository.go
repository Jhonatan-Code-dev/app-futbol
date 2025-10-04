// repository/usuario_repository.go
package repository

import (
	"errors"

	"app-futbol/src/schemas"

	"gorm.io/gorm"
)

type UsuarioRepository struct {
	DB *gorm.DB
}

func NewUsuarioRepository(db *gorm.DB) *UsuarioRepository {
	return &UsuarioRepository{DB: db}
}

// Verifica si un correo ya existe y devuelve error si está registrado
func (r *UsuarioRepository) EnsureCorreoDisponible(correo string) error {
	var count int64
	if err := r.DB.Model(&schemas.Usuario{}).
		Where("correo = ?", correo).
		Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("el correo ya está registrado")
	}
	return nil
}

// Crea un nuevo usuario (pero primero asegura que el correo no exista)
func (r *UsuarioRepository) Create(usuario *schemas.Usuario) error {
	// Chequeo de correo duplicado
	if err := r.EnsureCorreoDisponible(usuario.Correo); err != nil {
		return err
	}
	return r.DB.Create(usuario).Error
}
