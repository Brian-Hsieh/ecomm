package user

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/register", h.handleRegister)
	r.HandleFunc("/login", h.handleLogin)
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {}
func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request)    {}
