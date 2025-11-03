package handlers_home

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/yukikwi/go-nuxt-boilerplate/utils"
)

func RegisterHomeRoutes(router fiber.Router) {
	homeRouter := router.Group("/home")
	homeRouter.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the Home Page!")
	})

	homeRouter.Post("/speak", func(c *fiber.Ctx) error {
		var body SpeakRequestSerializer
		if err := utils.ValidateDataStruct(c, &body); err != nil {
			slog.Error("homeRouter.Post('/speak')", "error", err)
			return c.JSON(utils.ValidationErrorResponse{Error: err.Error()})
		}

		return c.JSON(SpeakResponseSerializer{Result: "You said: " + body.Message})
	})
}
