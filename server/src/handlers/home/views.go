package handlers_home

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/yukikwi/go-nuxt-boilerplate/config"
	"github.com/yukikwi/go-nuxt-boilerplate/models"
	"github.com/yukikwi/go-nuxt-boilerplate/utils"
)

func SpeakView(c *fiber.Ctx) error {
	var body SpeakRequestSerializer
	if err := utils.ValidateDataStruct(c, &body); err != nil {
		slog.Error("homeRouter.Post('/speak')", "error", err)
		return c.JSON(utils.ValidationErrorResponse{Error: err.Error()})
	}

	// Insert Message to Model
	message := models.Message{MessageText: body.Message}
	result := config.Db.Create(&message)
	if result.Error != nil {
		slog.Error("Failed to insert message", "error", result.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(utils.GenericErrorResponse{Error: "Failed to save message"})
	}

	return c.JSON(SpeakResponseSerializer{Result: "You said: " + body.Message})
}
