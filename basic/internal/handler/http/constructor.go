package http

import (
	"github.com/nori-io/interfaces/nori/http/v2"
	"github.com/nori-plugins/dummy/basic/internal/handler/http/test"
)

type Handler struct {
	R           http.Router
	TestHandler *test.TestHandler
}

type Params struct {
	R           http.Router
	TestHandler *test.TestHandler
}

func New(params Params) *Handler {
	handler := Handler{
		R:           params.R,
		TestHandler: params.TestHandler,
	}

	handler.R.Get("/dummy/basic", handler.TestHandler.GetMessage)
	handler.R.Post("/dummy/basic", handler.TestHandler.PostMessage)

	return &handler
}
