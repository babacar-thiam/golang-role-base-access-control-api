package user

type Service struct {
	repo Repository
}

// NewService creates a new instance of Service
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}
