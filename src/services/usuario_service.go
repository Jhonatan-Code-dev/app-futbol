package services

import (
	"fmt"

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
