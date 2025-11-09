package utils

import (
	"github.com/danielgtaylor/huma/v2"
)

func BuildConfig(title string, version string, environment string) huma.Config {
	schemaPrefix := "#/components/schemas/"
	schemasPath := ""
	openAPIPath := ""

	registry := huma.NewMapRegistry(schemaPrefix, huma.DefaultSchemaNamer)

	// Customize the config based on environment
	hooks := []func(huma.Config) huma.Config{}
	if environment == "development" {
		// enable openapi schema and docs only in development
		schemasPath = "/schemas"
		openAPIPath = "/openapi"

		hooks = append(hooks,
			func(c huma.Config) huma.Config {
				// Add a link transformer to the API. This adds `Link` headers and
				// puts `$schema` fields in the response body which point to the JSON
				// Schema that describes the response structure.
				// This is a create hook so we get the latest schema path setting.
				linkTransformer := huma.NewSchemaLinkTransformer(schemaPrefix, c.SchemasPath)
				c.OnAddOperation = append(c.OnAddOperation, linkTransformer.OnAddOperation)
				c.Transformers = append(c.Transformers, linkTransformer.Transform)
				return c
			},
		)
	}

	return huma.Config{
		OpenAPI: &huma.OpenAPI{
			OpenAPI: "3.1.0",
			Info: &huma.Info{
				Title:   title,
				Version: version,
			},
			Components: &huma.Components{
				Schemas: registry,
			},
		},
		OpenAPIPath:   openAPIPath,
		DocsPath:      "",
		SchemasPath:   schemasPath,
		Formats:       huma.DefaultFormats,
		DefaultFormat: "application/json",
		CreateHooks:   hooks,
	}
}
