package role

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	service *Service
}

// NewHandler creates a new Handler instance
func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// GetRoles handles getting the roles
// @Summary Get all roles
// @Description Retrieves all roles from the system
// @Tags Roles
// @Accept json
// @Produce json
// @Success 200 {array} Role "List of roles"
// @Router /admin/roles [get]
func (h *Handler) GetRoles(w http.ResponseWriter, r *http.Request) {
	roles, err := h.service.GetAll()
	if err != nil {
		http.Error(w, "error retrieving roles", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if encodeErr := json.NewEncoder(w).Encode(roles); encodeErr != nil {
		http.Error(w, "error encoding roles", http.StatusInternalServerError)
		return
	}
}
