package main

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	// Internal packages
	"github.com/yukikwi/go-nuxt-boilerplate/config"
	_ "github.com/yukikwi/go-nuxt-boilerplate/docs"
	handlers_home "github.com/yukikwi/go-nuxt-boilerplate/handlers/home"
	"github.com/yukikwi/go-nuxt-boilerplate/utils"
)

// @title Go Nuxt Boilerplate API
// @version 1.0
// @contact.name Pachara Chantawong
// @contact.email pachara.chantawong@gmail.com
// @host localhost:3000
// @BasePath  /api/v1
func main() {
	slog.Info("Starting server...")
	app := fiber.New(fiber.Config{
		ErrorHandler: utils.ErrorHandler,
	})
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// Register handlers module routes
	handlers_home.RegisterHomeRoutes(v1)

	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Listen(":" + config.Config.Port)
	slog.Info("Server is running on port " + config.Config.Port)
}
