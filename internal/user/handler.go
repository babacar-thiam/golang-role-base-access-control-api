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

// GetUser handles getting a user by ID
// @Summary Get user by ID
// @Description Retrieves a user by their unique ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} User "User data"
// @Router /users/{id} [get]
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

// GetAllUsers handles retrieving all users
// @Summary Get all users
// @Description Retrieves all users in the system
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {array} User "List of users"
// @Router /users [get]
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
