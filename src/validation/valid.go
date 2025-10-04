package validation

import (
	"errors"
	"regexp"
	"strings"
	"unicode"
)

// ValidarNombre verifica que el nombre tenga solo letras y espacios y longitud 2-100
func ValidarNombre(nombre string) error {
	if len(nombre) < 2 || len(nombre) > 100 {
		return errors.New("el nombre debe tener entre 2 y 100 caracteres")
	}
	for _, r := range nombre {
		if !unicode.IsLetter(r) && !unicode.IsSpace(r) {
			return errors.New("el nombre solo puede contener letras y espacios")
		}
	}
	return nil
}

// ValidarApellido verifica que el apellido tenga solo letras y espacios y longitud 2-100
func ValidarApellido(apellido string) error {
	if len(apellido) < 2 || len(apellido) > 100 {
		return errors.New("el apellido debe tener entre 2 y 100 caracteres")
	}
	for _, r := range apellido {
		if !unicode.IsLetter(r) && !unicode.IsSpace(r) {
			return errors.New("el apellido solo puede contener letras y espacios")
		}
	}
	return nil
}

// ValidarCorreo verifica que sea un correo válido usando regex simple
func ValidarCorreo(correo string) error {
	correo = strings.TrimSpace(correo)
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if !re.MatchString(correo) {
		return errors.New("correo inválido")
	}
	return nil
}

// ValidarPass verifica que la contraseña no esté vacía y tenga entre 8 y 255 caracteres
func ValidarPass(pass string) error {
	pass = strings.TrimSpace(pass)
	if len(pass) < 8 || len(pass) > 255 {
		return errors.New("la contraseña debe tener entre 8 y 255 caracteres")
	}
	return nil
}
