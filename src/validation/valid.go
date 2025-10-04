package validation

import (
	"net/mail"
	"regexp"
	"strings"
	"unicode/utf8"
)

var nameRegex = regexp.MustCompile(`^[A-Za-zÀ-ÖØ-öø-ÿ'’\- ]{2,100}$`)

func ValidarNombre(nombre string) bool {
	return nameRegex.MatchString(strings.TrimSpace(nombre))
}

func ValidarApellido(apellido string) bool {
	return ValidarNombre(apellido)
}

func ValidarCorreo(correo string) bool {
	correo = strings.TrimSpace(strings.ToLower(correo))
	addr, err := mail.ParseAddress(correo)
	return err == nil && addr.Address == correo && strings.HasSuffix(addr.Address, "@gmail.com")
}

func ValidarPass(pass string) bool {
	l := utf8.RuneCountInString(pass)
	return l >= 4 && l <= 10
}
