package helpers

import "github.com/gofiber/fiber/v2"

// Respuesta de éxito simple
func Success(ctx *fiber.Ctx, status int, message string) error {
	return ctx.Status(status).JSON(fiber.Map{
		"success": true,
		"message": message,
	})
}

// Respuesta de error simple
func Fail(ctx *fiber.Ctx, status int, message string) error {
	return ctx.Status(status).JSON(fiber.Map{
		"success": false,
		"message": message,
	})
}

// Respuesta de error con detalles
func FailWithErrors(ctx *fiber.Ctx, status int, errores any) error {
	return ctx.Status(status).JSON(fiber.Map{
		"success": false,
		"errores": errores,
	})
}
