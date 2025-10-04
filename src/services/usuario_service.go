package services

import (
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
	if err := validation.ValidarNombreError(usuario.Nombre); err != nil {
		errores["nombre"] = err.Error()
	}
	if err := validation.ValidarApellidoError(usuario.Apellido); err != nil {
		errores["apellido"] = err.Error()
	}
	if err := validation.ValidarCorreoError(usuario.Correo); err != nil {
		errores["correo"] = err.Error()
	}
	if err := validation.ValidarPassError(usuario.Pass); err != nil {
		errores["pass"] = err.Error()
	}
	if err := repository.ValidarCorreoExistente(s.DB, usuario.Correo); err != nil {
		errores["correo"] = err.Error()
	}
	hash, err := validation.HashPass(usuario.Pass)
	if err != nil {
		errores["pass"] = err.Error()
		return errores
	}
	if len(errores) > 0 {
		return errores
	}
	usuario.Pass = hash
	usuario.IDRol = 1
	usuario.Estado = false
	usuario.FechaSolicitud = validation.FechaActualPeru()
	s.DB.Create(usuario)
	return nil
}
