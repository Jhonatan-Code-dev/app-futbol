package validation

import (
	"reflect"
	"strings"
	"sync"
	"unicode"

	"github.com/go-playground/validator/v10"
)

// ValidationService: fachada mínima y rápida sobre go-playground/validator.
type ValidationService struct {
	validate *validator.Validate
}

var (
	service *ValidationService
	once    sync.Once
)

// NewValidationService construye una única instancia thread-safe.
func NewValidationService() *ValidationService {
	once.Do(func() {
		v := validator.New()
		// Usar el nombre JSON del campo en los errores.
		v.RegisterTagNameFunc(jsonTagName)

		// Validación personalizada "alphaSpace".
		_ = v.RegisterValidation("alphaSpace", alphaSpaceValidator)

		// Validación personalizada "gmail".
		_ = v.RegisterValidation("gmail", gmailValidator)

		service = &ValidationService{validate: v}
	})
	return service
}

// alphaSpaceValidator valida que el campo sea string no vacío y solo tenga letras y espacios.
func alphaSpaceValidator(fl validator.FieldLevel) bool {
	f := fl.Field()
	if f.Kind() != reflect.String {
		return false
	}
	s := strings.TrimSpace(f.String())
	if s == "" {
		return false
	}
	for _, r := range s {
		// Permitimos espacio en blanco
		if unicode.IsSpace(r) {
			continue
		}
		// Solo letras (incluye acentos, ñ, mayúsculas y Unicode)
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

// gmailValidator valida que el correo termine en @gmail.com
func gmailValidator(fl validator.FieldLevel) bool {
	f := fl.Field()
	if f.Kind() != reflect.String {
		return false
	}
	email := strings.TrimSpace(f.String())
	return strings.HasSuffix(strings.ToLower(email), "@gmail.com")
}

// jsonTagName devuelve el nombre del campo a partir del tag `json`.
func jsonTagName(sf reflect.StructField) string {
	name := sf.Tag.Get("json")
	if name == "" || name == "-" {
		return sf.Name
	}
	if i := strings.IndexByte(name, ','); i >= 0 {
		if i == 0 {
			return sf.Name
		}
		return name[:i]
	}
	return name
}

// ErrorMap representa un único error por campo.
type ErrorMap map[string]string

var _ error = ErrorMap{}

// Error devuelve un string compacto.
func (em ErrorMap) Error() string {
	if len(em) == 0 {
		return ""
	}
	var b strings.Builder
	first := true
	for field, msg := range em {
		if !first {
			b.WriteString("; ")
		} else {
			first = false
		}
		b.WriteString(field)
		b.WriteString(": ")
		b.WriteString(msg)
	}
	return b.String()
}

// ValidateStructInto valida el payload y reutiliza 'dst'.
func (s *ValidationService) ValidateStructInto(payload any, dst ErrorMap) error {
	// Limpia sin realocar
	for k := range dst {
		delete(dst, k)
	}

	if payload == nil {
		dst["_error"] = "payload nulo"
		return dst
	}

	if err := s.validate.Struct(payload); err != nil {
		ve, ok := err.(validator.ValidationErrors)
		if !ok {
			dst["_error"] = err.Error()
			return dst
		}

		for _, e := range ve {
			field := e.Field()
			var msg string
			switch e.Tag() {
			case "required":
				msg = "El campo " + field + " es obligatorio"
			case "min":
				msg = "El campo " + field + " debe tener al menos " + e.Param() + " caracteres"
			case "max":
				msg = "El campo " + field + " debe tener como máximo " + e.Param() + " caracteres"
			case "email":
				msg = "El campo " + field + " debe ser un correo válido"
			case "gmail":
				msg = "El campo " + field + " debe ser un correo @gmail.com"
			case "gte":
				msg = "El campo " + field + " debe ser mayor o igual a " + e.Param()
			case "lte":
				msg = "El campo " + field + " debe ser menor o igual a " + e.Param()
			case "alphaSpace":
				msg = "El campo " + field + " solo debe contener letras y espacios"
			default:
				msg = "Error en campo " + field + " (" + e.Tag() + ")"
			}
			// Solo un error por campo
			if _, exists := dst[field]; !exists {
				dst[field] = msg
			}
		}
		return dst
	}

	return nil
}
