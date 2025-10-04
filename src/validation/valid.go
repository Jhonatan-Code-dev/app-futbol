package validation

import (
	"errors"
	"net/mail"
	"regexp"
	"strings"
	"unicode/utf8"

	"golang.org/x/crypto/bcrypt"
)

var nameRegex = regexp.MustCompile(`^[A-Za-zÀ-ÖØ-öø-ÿ'’\- ]{2,100}$`)

// ValidarNombreError retorna error si nombre inválido
func ValidarNombreError(nombre string) error {
	if !nameRegex.MatchString(strings.TrimSpace(nombre)) {
		return errors.New("nombre inválido")
	}
	return nil
}

// ValidarApellidoError retorna error si apellido inválido
func ValidarApellidoError(apellido string) error {
	if !nameRegex.MatchString(strings.TrimSpace(apellido)) {
		return errors.New("apellido inválido")
	}
	return nil
}

// ValidarCorreoError retorna error si correo inválido
func ValidarCorreoError(correo string) error {
	correo = strings.TrimSpace(strings.ToLower(correo))
	addr, err := mail.ParseAddress(correo)
	if err != nil || addr.Address != correo || !strings.HasSuffix(addr.Address, "@gmail.com") {
		return errors.New("correo inválido, solo se acepta @gmail.com")
	}
	return nil
}

// ValidarPassError retorna error si pass inválido
func ValidarPassError(pass string) error {
	l := utf8.RuneCountInString(pass)
	if l < 4 || l > 10 {
		return errors.New("contraseña inválida, debe tener entre 4 y 10 caracteres")
	}
	return nil
}

// HashPass genera hash bcrypt con costo 12
func HashPass(pass string) (string, error) {
	h, e := bcrypt.GenerateFromPassword([]byte(pass), 12)
	return string(h), e
}

// ComparePass compara pass con hash bcrypt
func ComparePass(hash, pass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass)) == nil
}
