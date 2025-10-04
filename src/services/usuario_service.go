package services

import (
	"fmt"

	"app-futbol/src/middlewares"
	"app-futbol/src/schemas"
	"app-futbol/src/validation"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UsuarioService struct {
	DB *gorm.DB
}

func NewUsuarioService(db *gorm.DB) *UsuarioService {
	return &UsuarioService{DB: db}
}

// services/usuario_service.go
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

	if len(errores) > 0 {
		return errores // retornamos todos los errores
	}

	// Hash de la contraseña
	hash, _ := validation.HashPass(usuario.Pass)
	usuario.Pass = hash

	// Valores por defecto
	usuario.IDRol = 1
	usuario.Estado = false
	usuario.FechaSolicitud = validation.FechaActualPeru()

	// Crear usuario en DB
	s.DB.Create(usuario)

	return nil // no hay errores
}

func (s *UsuarioService) Login(correo, pass string) (string, error) {
	var usuario schemas.Usuario
	err := s.DB.Where("correo = ?", correo).First(&usuario).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", fmt.Errorf("usuario no encontrado")
		}
		return "", fmt.Errorf("error en la base de datos: %w", err)
	}

	if !usuario.Estado {
		return "", fmt.Errorf("usuario no aprobado")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(usuario.Pass), []byte(pass)); err != nil {
		return "", fmt.Errorf("contraseña incorrecta")
	}

	token, err := middlewares.GenerateToken(usuario.IdUsuario)
	if err != nil {
		return "", fmt.Errorf("no se pudo generar token")
	}

	return token, nil
}
