package role

type Service struct {
	repo *Repository
}

// NewService creates a new Service instance
func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

// GetAll gets all the existing roles
func (s *Service) GetAll() ([]Role, error) {
	return s.repo.FindAll()
}
