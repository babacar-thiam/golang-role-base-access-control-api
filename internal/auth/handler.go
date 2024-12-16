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
