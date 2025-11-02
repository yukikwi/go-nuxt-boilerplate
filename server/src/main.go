package main

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/yukikwi/go-nuxt-boilerplate/config"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":" + config.Config["Port"])
	slog.Info("Server is running on port " + config.Config["Port"])
}
