package handlers_home

import "github.com/danielgtaylor/huma/v2"

type SpeakPostRequestSerializer struct {
	Body struct {
		Message string `json:"message" doc:"Word you say" writeOnly:"true" example:"Hello!"`
	}
}

type UploadPostRequestSerializer struct {
	RawBody huma.MultipartFormFiles[struct {
		File      huma.FormFile   `form:"file"`
		BulkFiles []huma.FormFile `form:"bulk-files"`
	}]
}
