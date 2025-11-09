package handlers_doc

import (
	"context"
	"net/http"

	scalargo "github.com/bdpiprava/scalar-go"
	"github.com/danielgtaylor/huma/v2"
)

func RegisterHomeRoutes(api huma.API) {
	docRouter := huma.NewGroup(api, "/doc")

	huma.Register(docRouter, huma.Operation{
		OperationID: "doc-get-root",
		Method:      http.MethodGet,
		Path:        "/",
		Summary:     "Get a documentation page",
		Description: "Get a documentation page (Scalar UI).",
		Tags:        []string{"Documentation"},
	}, func(ctx context.Context, input *struct{}) (*DocGetResponseSerializer, error) {
		html, err := scalargo.NewV2(
			scalargo.WithSpecURL("http://localhost:3000/openapi.json"),
		)

		if err != nil {
			return nil, err
		}

		resp := &DocGetResponseSerializer{}
		resp.Body = []byte(html)
		resp.ContentType = "text/html"
		return resp, nil
	})

}
