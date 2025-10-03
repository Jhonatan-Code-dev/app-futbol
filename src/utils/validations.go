package utils

import "strings"

func NombreValido(nombre string) string {
	if nombre == "" {
		return `"nombre": "Nombre es obligatorio"`
	}
	if len(nombre) < 3 {
		return `"nombre": "Nombre debe tener al menos 3 caracteres"`
	}
	if len(nombre) > 50 {
		return `"nombre": "Nombre no debe exceder 50 caracteres"`
	}
	return ""
}

func CorreoValido(correo string) string {
	if correo == "" {
		return `"correo": "Correo es obligatorio"`
	}
	if !strings.Contains(correo, "@") {
		return `"correo": "Correo debe contener @"`
	}
	if !strings.HasSuffix(correo, "@gmail.com") {
		return `"correo": "Solo se permiten correos @gmail.com"`
	}
	return ""
}

func PassValida(pass string) string {
	if pass == "" {
		return `"pass": "Contraseña es obligatoria"`
	}
	if len(pass) < 6 {
		return `"pass": "Contraseña debe tener al menos 6 caracteres"`
	}
	if len(pass) > 50 {
		return `"pass": "Contraseña no debe exceder 50 caracteres"`
	}
	return ""
}
