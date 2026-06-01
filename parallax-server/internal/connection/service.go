package connection

import "context"

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

// MARK: CREATE CONNECTION
func (s *Service) Create(ctx context.Context, connection CreateEntityConnectionRequest) (*EntityConnection, error) {
	return s.repository.CreateConnection(ctx, connection)
}
