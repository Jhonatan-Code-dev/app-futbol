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

// UsuarioService maneja la lógica de usuarios
type UsuarioService struct {
	DB        *gorm.DB
	validator *validation.ValidationService
}

// Constructor
func NewUsuarioService(db *gorm.DB) *UsuarioService {
	return &UsuarioService{
		DB:        db,
		validator: validation.NewValidationService(),
	}
}

// RequestRegister solicita el registro de un usuario
func (s *UsuarioService) RequestRegister(usuario *schemas.Usuario) error {
	// Validar datos con ValidationService
	errorsMap := validation.ErrorMap{}
	if err := s.validator.ValidateStructInto(usuario, errorsMap); err != nil {
		// Devolver errores amigables en formato JSON
		return err
	}

	// Verificar si el correo ya está registrado
	var existing schemas.Usuario
	if err := s.DB.Where("correo = ?", usuario.Correo).First(&existing).Error; err == nil {
		return fmt.Errorf("el correo %s ya está registrado", usuario.Correo)
	} else if err != gorm.ErrRecordNotFound {
		return err
	}

	// Encriptar contraseña
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(usuario.Pass), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	usuario.Pass = string(hashedPass)
	usuario.Estado = false
	usuario.FechaSolicitud = time.Now()
	usuario.FechaAceptacion = time.Time{}

	// Guardar en la DB
	if err := s.DB.Create(usuario).Error; err != nil {
		return err
	}

	return nil
}

// Login valida al usuario y genera un token JWT
func (s *UsuarioService) Login(correo, pass string) (string, error) {
	var usuario schemas.Usuario
	err := s.DB.Where("correo = ?", correo).First(&usuario).Error
	if err != nil {
		return "", fmt.Errorf("usuario no encontrado")
	}

	if !usuario.Estado {
		return "", fmt.Errorf("usuario no aprobado")
	}

	err = bcrypt.CompareHashAndPassword([]byte(usuario.Pass), []byte(pass))
	if err != nil {
		return "", fmt.Errorf("contraseña incorrecta")
	}

	token, err := middlewares.GenerateToken(usuario.IdUsuario)
	if err != nil {
		return "", fmt.Errorf("no se pudo generar token")
	}

	return token, nil
}
