package role

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"time"
)

// Role types
const (
	ADMIN    = "ADMIN"
	CLIENT   = "CLIENT"
	PROVIDER = "PROVIDER"
)

var ValidRoles = map[string]bool{
	ADMIN:    true,
	CLIENT:   true,
	PROVIDER: true,
}

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

// Validate ensures the required fields are not empty and the role is valid
func (r *Role) Validate() error {
	if err := r.RequiredFields(); err != nil {
		return err
	}

	if !ValidRoles[r.Name] {
		return errors.New("invalid role type")
	}

	return nil
}

// IsValidRole checks if a role string is valid
func IsValidRole(role string) bool {
	return ValidRoles[role]
}
