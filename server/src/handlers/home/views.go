package handlers_home

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/yukikwi/go-nuxt-boilerplate/utils"
)

func SpeakView(c *fiber.Ctx) error {
	var body SpeakRequestSerializer
	if err := utils.ValidateDataStruct(c, &body); err != nil {
		slog.Error("homeRouter.Post('/speak')", "error", err)
		return c.JSON(utils.ValidationErrorResponse{Error: err.Error()})
	}

	return c.JSON(SpeakResponseSerializer{Result: "You said: " + body.Message})
}
