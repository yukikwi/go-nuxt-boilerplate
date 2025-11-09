package handlers_home

type SpeakPostRequestSerializer struct {
	Body struct {
		Message string `json:"message" doc:"Word you say" writeOnly:"true" example:"Hello!"`
	}
}
