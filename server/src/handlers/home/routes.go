package handlers_home

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func RegisterHomeRoutes(versionGroup *huma.Group) {
	homeRouter := huma.NewGroup(versionGroup, "/home")

	huma.Register(homeRouter, huma.Operation{
		OperationID: "home-get-root",
		Method:      http.MethodGet,
		Path:        "/",
		Summary:     "Get a home example page",
		Tags:        []string{"Home"},
	}, func(ctx context.Context, input *struct{}) (*SpeakGetResponseSerializer, error) {
		resp := &SpeakGetResponseSerializer{}
		resp.Body.Message = "Hello, world!"
		return resp, nil
	})

	huma.Register(homeRouter, huma.Operation{
		OperationID: "home-post-speak",
		Method:      http.MethodPost,
		Path:        "/speak",
		Summary:     "Post a home example speak",
		Tags:        []string{"Home"},
	}, func(ctx context.Context, input *SpeakPostRequestSerializer) (*SpeakPostResponseSerializer, error) {
		return SpeakPostView(ctx, input)
	})

	huma.Register(homeRouter, huma.Operation{
		OperationID: "home-post-upload",
		Method:      http.MethodPost,
		Path:        "/upload",
		Summary:     "Post a home example upload",
		Tags:        []string{"Home"},
	}, func(ctx context.Context, input *UploadPostRequestSerializer) (*UploadPostResponseSerializer, error) {
		return UploadPostView(ctx, input)
	})
}
