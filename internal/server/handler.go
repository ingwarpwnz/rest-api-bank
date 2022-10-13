package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	router *mux.Router
}

func NewHandler(router *mux.Router) *Handler {
	return &Handler{
		router: router,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}
