package auth

import (
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"role-base-access-control-api/configs"
	"role-base-access-control-api/internal/user"
	"time"
)

type Service struct {
	userRepo *user.Repository
	jwt      *JWT
	config   *configs.Config
}

// NewService creates a new instance of Service
func NewService(config *configs.Config, userRepo *user.Repository) *Service {
	return &Service{
		userRepo: userRepo,
		jwt:      NewJWT(config),
		config:   config,
	}
}

func (s *Service) Register(req RegisterRequest) (*RegisterResponse, error) {
	// Check if email already in use
	existingUser, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("email already in use")
	}

	// Check if the phone number already in use
	existingUser, err = s.userRepo.FindByPhone(req.Phone)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("phone already in use")
	}

	// Validate password
	if len(req.Password) < 6 {
		return nil, errors.New("password must be at least 6 characters")
	}

	// Hashed Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Register new user
	newUser := &user.User{
		ID:        uuid.New(),
		Name:      req.Name,
		Email:     req.Email,
		Phone:     req.Phone,
		Password:  string(hashedPassword),
		Role:      req.Role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Validate the user
	if validate := newUser.Validate(); validate != nil {
		return nil, validate
	}

	// Save the user
	if save := s.userRepo.Save(newUser); save != nil {
		return nil, save
	}

	return &RegisterResponse{
		ID:        newUser.ID,
		Name:      newUser.Name,
		Email:     newUser.Email,
		Phone:     newUser.Phone,
		Role:      newUser.Role,
		CreatedAt: newUser.CreatedAt,
		UpdatedAt: newUser.UpdatedAt,
	}, nil
}
