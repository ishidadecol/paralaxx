package person

import "context"

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) GetPeople(ctx context.Context) ([]Person, error) {
	return s.repository.GetAll(ctx)
}
