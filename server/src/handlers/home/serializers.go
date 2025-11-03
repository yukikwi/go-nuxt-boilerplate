package handlers_home

type SpeakRequestSerializer struct {
	Message string `json:"message" validate:"required"`
}

type SpeakResponseSerializer struct {
	Result string `json:"result"`
}
