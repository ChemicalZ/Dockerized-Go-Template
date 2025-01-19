package user

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
}

func (h *Handler) CreateRoutes(r chi.Router) {
	r.Get("/", h.GetUsers)
}

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
