package user

import (
	"fmt"
	"net/http"

	"github.com/Brian-Hsieh/ecomm/pkg"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	store pkg.UserStore
}

func NewHandler(store pkg.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/register", h.handleRegister)
	r.HandleFunc("/login", h.handleLogin)
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// parse json payload
	var payload pkg.UserPayload
	if err := pkg.ParseJSON(r, payload); err != nil {
		pkg.WriteError(w, http.StatusBadRequest, err)
	}

	// check if user exists
	_, err := h.store.GetUserByEmail(payload.Email)
	if err == nil {
		pkg.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", payload.Email))
		return
	}

	// hash user password
	pwd, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		pkg.WriteError(w, http.StatusInternalServerError, err)
	}

	err = h.store.CreateUser(pkg.User{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: string(pwd),
	})
	if err != nil {
		pkg.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	pkg.WriteJSON(w, http.StatusCreated, nil)
}
func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {}
