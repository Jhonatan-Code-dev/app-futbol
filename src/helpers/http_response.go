package helpers

import "github.com/gofiber/fiber/v2"

// jsonSuccess devuelve una respuesta JSON estándar para éxito.
func JsonSuccess(ctx *fiber.Ctx, msg string) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"estado":  true,
		"mensaje": msg,
	})
}

// jsonError devuelve una respuesta JSON estándar para error controlado.
func JsonError(ctx *fiber.Ctx, msg string) error {
	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"estado": false,
		"error":  msg,
	})
}
