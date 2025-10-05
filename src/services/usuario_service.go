package services

import (
	"app-futbol/src/helpers"
	"app-futbol/src/repository"
	"app-futbol/src/schemas"
	"app-futbol/src/validation"

	"gorm.io/gorm"
)

type UsuarioService struct {
	DB *gorm.DB
}

func NewUsuarioService(db *gorm.DB) *UsuarioService {
	return &UsuarioService{DB: db}
}

func (s *UsuarioService) RequestRegister(usuario *schemas.Usuario) map[string]string {
	errores := make(map[string]string)
	helpers.AddValidationError(errores, "nombre", usuario.Nombre, validation.ValidarNombreError)
	helpers.AddValidationError(errores, "apellido", usuario.Apellido, validation.ValidarApellidoError)
	helpers.AddValidationError(errores, "correo", usuario.Correo, validation.ValidarCorreoError)
	helpers.AddValidationError(errores, "pass", usuario.Pass, validation.ValidarPassError)
	helpers.AddDBExistenceError(errores, "correo", usuario.Correo, s.DB, repository.ValidarCorreoExistente)
	usuario.Pass = helpers.HashPassword(errores, "pass", usuario.Pass)
	if len(errores) > 0 {
		return errores
	}
	usuario.IDRol = 1
	usuario.Estado = false
	usuario.FechaSolicitud = validation.FechaActualPeru()
	if err := s.DB.Create(usuario).Error; err != nil {
		helpers.SafeError(errores, "db", err)
		return errores
	}
	return nil
}
