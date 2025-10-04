package helpers

import "github.com/gofiber/fiber/v2"

// ParseBody intenta parsear el body de la request al struct dado.
// Si falla, devuelve un JsonError automáticamente.
func ParseBody(ctx *fiber.Ctx, out interface{}) error {
	if err := ctx.BodyParser(out); err != nil {
		return JsonError(ctx, "El formato de los datos enviados no es válido")
	}
	return nil
}

// ParseQuery intenta parsear los parámetros de query (?id=1&name=...).
func ParseQuery(ctx *fiber.Ctx, out interface{}) error {
	if err := ctx.QueryParser(out); err != nil {
		return JsonError(ctx, "Los parámetros de búsqueda no son válidos")
	}
	return nil
}

// ParseParams intenta parsear parámetros de la ruta (/usuarios/:id).
func ParseParams(ctx *fiber.Ctx, out interface{}) error {
	if err := ctx.ParamsParser(out); err != nil {
		return JsonError(ctx, "Los parámetros de la ruta no son válidos")
	}
	return nil
}
