package handlers_home

import (
	"context"
	"log/slog"

	"github.com/yukikwi/go-nuxt-boilerplate/config"
	"github.com/yukikwi/go-nuxt-boilerplate/models"
)

func SpeakPostView(ctx context.Context, input *SpeakPostRequestSerializer) (*SpeakPostResponseSerializer, error) {
	// Insert Message to Model
	message := models.Message{MessageText: input.Body.Message}
	result := config.Db.Create(&message)
	if result.Error != nil {
		slog.Error("Failed to insert message", "error", result.Error)
		return nil, result.Error
	}

	resp := &SpeakPostResponseSerializer{}
	resp.Body.Message = "You said: " + input.Body.Message
	return resp, nil
}
