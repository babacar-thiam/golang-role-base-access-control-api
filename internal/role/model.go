package role

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

type Role struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// RequiredFields checks for missing required fields
func (r *Role) RequiredFields() error {
	fields := map[string]string{
		"name":        r.Name,
		"description": r.Description,
	}

	// Iterate through the required fields
	for key, value := range fields {
		if value == "" {
			return fmt.Errorf("%s is required", key)
		}
	}
	return nil
}

// Validate ensures the required fields are not empty
func (r *Role) Validate() error {
	return r.RequiredFields()
}
