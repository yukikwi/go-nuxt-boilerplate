package handlers_home

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterHomeRoutes(router fiber.Router) {
	homeRouter := router.Group("/home")
	homeRouter.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the Home Page!")
	})

	homeRouter.Post("/speak", func(c *fiber.Ctx) error {
		return SpeakView(c)
	})
}
