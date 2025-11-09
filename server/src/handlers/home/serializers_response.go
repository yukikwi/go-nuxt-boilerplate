package handlers_home

type SpeakGetResponseSerializer struct {
	Body struct {
		Message string `json:"message" example:"Hello, world!" doc:"Greeting message"`
	}
}

type SpeakPostResponseSerializer struct {
	Body struct {
		Message string `json:"message" example:"You said: Hello!" doc:"Response message"`
	}
}
