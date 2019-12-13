package handler

import (
	"durn2/view"
	viewModel "durn2/view/model"
	"net/http"
)

type HelloWorld struct {
	viewFactory view.Factory
}

func NewHelloWorld(viewFactory view.Factory) *HelloWorld {
	return &HelloWorld{viewFactory: viewFactory}
}

func (h *HelloWorld) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	v := h.viewFactory.New("helloWorld.gohtml").
		Data(viewModel.NewHelloWorld("Hello World"))
	v.Render(w)
}
