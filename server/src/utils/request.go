package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type (
	ValidationErrorResponse struct {
		Error string `json:"error"`
	}
)

func ValidateDataStruct(ctx *fiber.Ctx, body interface{}) error {
	if err := ctx.BodyParser(body); err != nil {
		return err
	}
	validate := validator.New()
	return validate.Struct(body)
}
