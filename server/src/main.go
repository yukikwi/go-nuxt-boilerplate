package main

import (
	"log/slog"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humafiber"
	"github.com/gofiber/fiber/v2"

	// Internal packages
	"github.com/yukikwi/go-nuxt-boilerplate/config"
	handlers_doc "github.com/yukikwi/go-nuxt-boilerplate/handlers/doc"
	handlers_home "github.com/yukikwi/go-nuxt-boilerplate/handlers/home"
)

// @title Go Nuxt Boilerplate API
// @version 1.0
// @contact.name Pachara Chantawong
// @contact.email pachara.chantawong@gmail.com
// @host localhost:3000
// @BasePath  /api/v1
func main() {
	slog.Info("Starting server...")
	app := fiber.New()
	api := humafiber.New(app, huma.DefaultConfig("Book API", "1.0.0"))
	v1 := huma.NewGroup(api, "/v1")

	// Register handlers module routes
	handlers_home.RegisterHomeRoutes(v1)
	handlers_doc.RegisterHomeRoutes(api)

	app.Listen(":" + config.Config.Port)
	slog.Info("Server is running on port " + config.Config.Port)
}
