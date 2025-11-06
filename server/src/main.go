package main

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/yukikwi/go-nuxt-boilerplate/config"
	handlers_home "github.com/yukikwi/go-nuxt-boilerplate/handlers/home"
	"github.com/yukikwi/go-nuxt-boilerplate/utils"
)

func main() {
	slog.Info("Starting server...")
	app := fiber.New(fiber.Config{
		ErrorHandler: utils.ErrorHandler,
	})
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// Register handlers module routes
	handlers_home.RegisterHomeRoutes(v1)

	app.Listen(":" + config.Config.Port)
	slog.Info("Server is running on port " + config.Config.Port)
}
