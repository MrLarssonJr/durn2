package model

type HelloWorld struct {
	text string
}

func (h *HelloWorld) Text() string {
	return h.text
}

func NewHelloWorld(text string) *HelloWorld {
	return &HelloWorld{text: text}
}
