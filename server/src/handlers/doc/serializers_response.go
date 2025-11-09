package handlers_doc

type DocGetResponseSerializer struct {
	ContentType string `header:"Content-Type"`
	Body        []byte `example:"<h1>Hello, world!</h1>" doc:"Greeting message"`
}
