package services

import (
	"fmt"
	"time"

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

// RequestRegister valida y registra un usuario
func (s *UsuarioService) RequestRegister(usuario *schemas.Usuario) error {
	// Validaciones usando validation con mensajes centralizados
	if err := validation.ValidarNombreError(usuario.Nombre); err != nil {
		return err
	}
	if err := validation.ValidarApellidoError(usuario.Apellido); err != nil {
		return err
	}
	if err := validation.ValidarCorreoError(usuario.Correo); err != nil {
		return err
	}
	if err := validation.ValidarPassError(usuario.Pass); err != nil {
		return err
	}

	// Hashear contraseña
	hash, err := validation.HashPass(usuario.Pass)
	if err != nil {
		return err
	}
	usuario.Pass = hash

	// Asignar valores por defecto
	usuario.IDRol = 1
	usuario.Estado = true
	usuario.FechaSolicitud = time.Now()
	usuario.FechaAceptacion = time.Now()

	// Insertar en la base de datos
	if err := s.DB.Create(usuario).Error; err != nil {
		return err
	}

	return nil
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
