package repository

import (
	"context"
	"errors"

	"app-futbol/src/schemas"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var ErrCorreoDuplicado = errors.New("usuario: correo ya registrado")

type UsuarioRepository struct{ DB *gorm.DB }

func NewUsuarioRepository(db *gorm.DB) *UsuarioRepository { return &UsuarioRepository{DB: db} }

func (r *UsuarioRepository) Create(ctx context.Context, u *schemas.Usuario) error {
	if err := r.DB.WithContext(ctx).
		Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "correo"}}, DoNothing: true}).
		Create(u).Error; err != nil {
		return err
	}
	if r.DB.RowsAffected == 0 {
		return ErrCorreoDuplicado
	}
	return nil
}
