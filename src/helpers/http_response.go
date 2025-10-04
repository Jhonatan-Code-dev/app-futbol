package helpers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

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

// JsonInternalError maneja errores inesperados.
// Se loggea el error real en consola y se devuelve un mensaje genérico al cliente.
func JsonInternalError(ctx *fiber.Ctx, err error) error {
	// Log interno para desarrolladores
	fmt.Println("ERROR INTERNO:", err.Error())

	// Respuesta genérica al cliente
	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"estado": false,
		"error":  "Ocurrió un error interno, inténtalo más tarde.",
	})
}
