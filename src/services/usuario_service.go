package services

import (
	"fmt"
	"time"

	"app-futbol/src/middlewares"
	"app-futbol/src/schemas"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UsuarioService struct {
	DB *gorm.DB
}

func NewUsuarioService(db *gorm.DB) *UsuarioService {
	return &UsuarioService{
		DB: db,
	}
}

func (s *UsuarioService) RequestRegister(usuario *schemas.Usuario) error {

	var existing schemas.Usuario
	err := s.DB.Where("correo = ?", usuario.Correo).First(&existing).Error
	if err == nil {
		return fmt.Errorf("el correo %s ya está registrado", usuario.Correo)
	} else if err != nil && err != gorm.ErrRecordNotFound {
		return fmt.Errorf("error al verificar correo: %w", err)
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(usuario.Pass), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error al encriptar contraseña: %w", err)
	}

	usuario.Pass = string(hashedPass)
	usuario.Estado = false
	usuario.FechaSolicitud = time.Now()
	usuario.FechaAceptacion = time.Time{}

	if err := s.DB.Create(usuario).Error; err != nil {
		return fmt.Errorf("error al guardar usuario: %w", err)
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
