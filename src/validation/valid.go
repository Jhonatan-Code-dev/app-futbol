package validation

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

// ValidationService centraliza todas las validaciones
type ValidationService struct {
	validate *validator.Validate
}

// Constructor
func NewValidationService() *ValidationService {
	v := validator.New()

	// Registro de validación personalizada
	v.RegisterValidation("gmail", func(fl validator.FieldLevel) bool {
		email := fl.Field().String()
		return strings.HasSuffix(email, "@gmail.com")
	})

	return &ValidationService{
		validate: v,
	}
}
