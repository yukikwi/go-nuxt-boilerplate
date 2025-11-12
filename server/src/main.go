package main

import (
	"log/slog"
	"os"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humafiber"
	"github.com/danielgtaylor/huma/v2/humacli"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"

	// Internal packages
	"github.com/yukikwi/go-nuxt-boilerplate/config"
	handlers_doc "github.com/yukikwi/go-nuxt-boilerplate/handlers/doc"
	handlers_home "github.com/yukikwi/go-nuxt-boilerplate/handlers/home"
	"github.com/yukikwi/go-nuxt-boilerplate/utils"
)

func setupAPIServer(api *huma.API) {
	slog.Debug("Starting server in " + config.Config.Environment + " environment...")
	v1 := huma.NewGroup(*api, "/v1")

	// Register handlers module routes
	handlers_home.RegisterHomeRoutes(v1)

	// Register development-only routes
	if config.Config.Environment == "development" {
		handlers_doc.RegisterHomeRoutes(*api)
	}

}

func main() {
	app := fiber.New(fiber.Config{
		Prefork: config.Config.Prefork,
	})
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger = logger.With(
		slog.Group("program_info",
			slog.Int("pid", os.Getpid()),
		),
	)
	slog.SetDefault(logger)
	api := humafiber.New(app, utils.BuildConfig("Book API", "1.0.0", config.Config.Environment))
	cli := humacli.New(func(hooks humacli.Hooks, opts *struct{}) {
		// setup route handlers
		setupAPIServer(&api)

		hooks.OnStart(func() {
			// Start the server
			slog.Info("Server is running on port " + config.Config.Port)
			app.Listen(":" + config.Config.Port)
		})
	})

	cli.Root().AddCommand(&cobra.Command{
		Use:   "openapi",
		Short: "Print the OpenAPI spec",
		Run: func(cmd *cobra.Command, args []string) {
			b, err := api.OpenAPI().YAML()
			if err != nil {
				panic(err)
			}
			outputFile := "./static/openapi/docs.yaml"
			if len(args) == 1 {
				outputFile = args[0]
			}
			f, err := os.Create(outputFile)
			if err != nil {
				slog.Error("Unable to open file", "error", err)
			}
			_, err = f.Write(b)
			if err != nil {
				slog.Error("Unable to write file", "error", err)
			}
			if err := f.Close(); err != nil {
				slog.Error("Unable to close file", "error", err)
			}
		},
	})

	// Run the thing!
	cli.Run()
}
