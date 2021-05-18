package http

import (
	"github.com/nori-io/interfaces/nori/http"
	"github.com/nori-plugins/dummy/basic/internal/handler/http/test"
)

type Handler struct {
	R           http.Http
	TestHandler *test.TestHandler
}

type Params struct {
	R           http.Http
	TestHandler *test.TestHandler
}

func New(params Params) *Handler {
	handler := Handler{
		R:           params.R,
		TestHandler: params.TestHandler,
	}

	handler.R.Get("/dummy", handler.TestHandler.GetMessage)
	handler.R.Post("/dummy", handler.TestHandler.PostMessage)

	return &handler
}
