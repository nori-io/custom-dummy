package test

import "net/http"

type TestHandler struct {

}

func (h *TestHandler) GetMessage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Dummy API"))

}

func (h *TestHandler) PostMessage(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Query().Get("message")
	w.Write([]byte(message))
}