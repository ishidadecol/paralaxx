package company

import "context"

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) GetCompanies(ctx context.Context) ([]Company, error) {
	return s.repository.GetAll(ctx)
}
