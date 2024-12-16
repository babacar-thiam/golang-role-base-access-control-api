package user

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Role      string    `json:"role"`
}

// RequiredFields checks for missing required fields
func (u User) RequiredFields() error {
	fields := map[string]string{
		"name":     u.Name,
		"email":    u.Email,
		"phone":    u.Phone,
		"password": u.Password,
		"role":     u.Role,
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
func (u User) Validate() error {
	return u.RequiredFields()
}
