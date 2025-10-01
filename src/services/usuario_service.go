package services

import (
	"errors"
	"time"

	"app-futbol/src/middlewares"
	"app-futbol/src/schemas"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// UsuarioService maneja la lógica de usuarios
type UsuarioService struct {
	DB *gorm.DB
}

// Constructor
func NewUsuarioService(db *gorm.DB) *UsuarioService {
	return &UsuarioService{DB: db}
}

// RequestRegister solicita el registro de un usuario
func (s *UsuarioService) RequestRegister(usuario *schemas.Usuario) error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(usuario.Pass), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	usuario.Pass = string(hashedPass)
	usuario.Estado = false
	usuario.FechaSolicitud = time.Now()
	usuario.FechaAceptacion = time.Time{}

	return s.DB.Create(usuario).Error
}

// Login valida al usuario y genera un token JWT
func (s *UsuarioService) Login(correo, pass string) (string, error) {
	var usuario schemas.Usuario

	err := s.DB.Where("correo = ?", correo).First(&usuario).Error
	if err != nil {
		return "", errors.New("usuario no encontrado")
	}

	if !usuario.Estado {
		return "", errors.New("usuario no aprobado")
	}

	err = bcrypt.CompareHashAndPassword([]byte(usuario.Pass), []byte(pass))
	if err != nil {
		return "", errors.New("contraseña incorrecta")
	}

	token, err := middlewares.GenerateToken(usuario.IdUsuario)
	if err != nil {
		return "", errors.New("no se pudo generar token")
	}

	return token, nil
}
