package entity

import "context"

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

// MARK: CREATE ENTITY
func (s *Service) CreateEntity(ctx context.Context, request CreateEntityRequest) (*Entity, error) {
	return s.repository.Create(ctx, request)
}
