package user

// Service provides user business logic and operations.
type Service struct {
	repo ReadWriter
}

// NewService creates a new user service instance.
func NewService(repo ReadWriter) *Service {
	return &Service{repo: repo}
}
