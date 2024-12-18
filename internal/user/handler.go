package user

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
	service *Service
}

// NewHandler handles the user HTTP requests
func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	// Extract user ID from the URL
	idStr := mux.Vars(r)["id"]
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Retrieve the user
	user, usrErr := h.service.Get(id)
	if usrErr != nil {
		http.Error(w, usrErr.Error(), http.StatusNotFound)
		return
	}

	// Respond with the user data
	w.Header().Set("Content-Type", "application/json")
	if encodeErr := json.NewEncoder(w).Encode(user); encodeErr != nil {
		http.Error(w, "error encoding roles", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetAllUsers(w http.ResponseWriter, r *http.Request) {

	users, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Respond with the user data
	w.Header().Set("Content-Type", "application/json")
	if encodeErr := json.NewEncoder(w).Encode(users); encodeErr != nil {
		http.Error(w, "error encoding roles", http.StatusInternalServerError)
		return
	}
}
