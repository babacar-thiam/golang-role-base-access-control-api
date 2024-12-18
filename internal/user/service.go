package user

import (
	"fmt"
	"github.com/google/uuid"
)

type Service struct {
	repo *Repository
}

// NewService creates a new instance of Service
func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

// Get gets a user by his id
func (s *Service) Get(id uuid.UUID) (*User, error) {
	user, err := s.repo.Find(id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, fmt.Errorf("user not found with id_: %s", id)
	}

	return user, nil
}

// GetAll gets all the users
func (s *Service) GetAll() ([]User, error) {
	users, err := s.repo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("users not found: %w", err)
	}

	return users, nil
}
