package user

type Handler struct {
	service *Service
}

// NewHandler handles the user HTTP requests
func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}
