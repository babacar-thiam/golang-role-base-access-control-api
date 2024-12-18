package auth

import (
	"encoding/json"
	"net/http"
	"role-base-access-control-api/internal/role"
)

type Handler struct {
	service *Service
}

// NewHandler handles the user HTTP requests
func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// Register handles the registration request
// @Summary Register a new user
// @Description Register a new user with the role of CLIENT or PROVIDER
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body RegisterRequest true "Register Request"
// @Success 201 {object} RegisterResponse
// @Router /auth/register [post]
func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate role (only CLIENT and PROVIDER allowed for registration)
	if req.Role != role.CLIENT && req.Role != role.PROVIDER {
		http.Error(w, "Invalid role. Must be CLIENT or PROVIDER", http.StatusBadRequest)
		return
	}

	response, err := h.service.Register(req)
	if err != nil {
		switch err.Error() {
		case "email already exists":
			http.Error(w, err.Error(), http.StatusConflict)
		case "invalid email format":
			http.Error(w, err.Error(), http.StatusBadRequest)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if encodeErr := json.NewEncoder(w).Encode(response); encodeErr != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}

// Login handles the login HTTP request
// @Summary User login
// @Description Authenticate a user and return a JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body LoginRequest true "Login Request"
// @Success 200 {object} LoginResponse
// @Router /auth/login [post]
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate the request
	if req.Email == "" || req.Password == "" {
		http.Error(w, "email and password are required", http.StatusBadRequest)
	}

	response, err := h.service.Login(req)
	if err != nil {
		switch err.Error() {
		case "email not found":
			http.Error(w, err.Error(), http.StatusUnauthorized)
		case "invalid password":
			http.Error(w, err.Error(), http.StatusUnauthorized)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if encodeErr := json.NewEncoder(w).Encode(response); encodeErr != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}
