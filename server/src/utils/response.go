package utils

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

type GenericErrorResponse struct {
	Error string `json:"error"`
}

var ErrorHandler = func(c *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError

	// Retrieve the custom status code if it's a *fiber.Error
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	// Set Content-Type: text/plain; charset=utf-8
	c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)

	// Return status code with error message as json
	return c.Status(code).JSON(GenericErrorResponse{Error: err.Error()})
}
