package user

import (
	"net/http"

	"github.com/Brian-Hsieh/ecomm/pkg"
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

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// parse json payload
	var payload *pkg.UserPayload
	if err := pkg.ParseJSON(r, payload); err != nil {
		pkg.WriteError(w, http.StatusBadRequest, err)
	}

	// check if user exists
}
func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {}
