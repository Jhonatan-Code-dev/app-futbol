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
	// Validar datos con ValidationService (incluye regla "gmail")
	errorsMap := validation.ErrorMap{}
	if err := s.validator.ValidateStructInto(usuario, errorsMap); err != nil {
		return err
	}

	// Verificar si el correo ya está registrado
	var existing schemas.Usuario
	err := s.DB.Where("correo = ?", usuario.Correo).First(&existing).Error
	if err == nil {
		// Existe → error controlado
		return fmt.Errorf("el correo %s ya está registrado", usuario.Correo)
	} else if err != nil && err != gorm.ErrRecordNotFound {
		// Error inesperado de DB
		return fmt.Errorf("error al verificar correo: %w", err)
	}

	// Encriptar contraseña
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(usuario.Pass), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error al encriptar contraseña: %w", err)
	}

	// Completar datos de registro
	usuario.Pass = string(hashedPass)
	usuario.Estado = false
	usuario.FechaSolicitud = time.Now()
	usuario.FechaAceptacion = time.Time{}

	// Guardar en la DB
	if err := s.DB.Create(usuario).Error; err != nil {
		return fmt.Errorf("error al guardar usuario: %w", err)
	}

	return nil
}

// Login valida al usuario y genera un token JWT
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

	// Comparar contraseñas
	if err := bcrypt.CompareHashAndPassword([]byte(usuario.Pass), []byte(pass)); err != nil {
		return "", fmt.Errorf("contraseña incorrecta")
	}

	// Generar token
	token, err := middlewares.GenerateToken(usuario.IdUsuario)
	if err != nil {
		return "", fmt.Errorf("no se pudo generar token")
	}

	return token, nil
}
