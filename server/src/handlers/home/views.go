package handlers_home

import (
	"context"
	"errors"
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
		return nil, errors.New("failed to insert message")
	}

	resp := &SpeakPostResponseSerializer{}
	resp.Body.Message = "You said: " + input.Body.Message
	return resp, nil
}

func UploadPostView(ctx context.Context, input *UploadPostRequestSerializer) (*UploadPostResponseSerializer, error) {
	formData := input.RawBody.Data()
	resp := &UploadPostResponseSerializer{}
	if formData.File.IsSet {
		fileInfo := FileInfo{
			FileName: formData.File.Filename,
			Size:     formData.File.Size,
		}
		resp.Body.FileInfo = append(resp.Body.FileInfo, fileInfo)
	}
	for _, file := range formData.BulkFiles {
		fileInfo := FileInfo{
			FileName: file.Filename,
			Size:     file.Size,
		}
		resp.Body.FileInfo = append(resp.Body.FileInfo, fileInfo)
	}
	return resp, nil
}
