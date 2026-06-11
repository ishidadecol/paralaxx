package company

import (
	"context"

	"github.com/ishidadecol/parallax/internal/entity"
)

type Service struct {
	repository       *Repository
	entityRepository *entity.Repository
}

func NewService(repository *Repository, entityRepository *entity.Repository) *Service {
	return &Service{
		repository:       repository,
		entityRepository: entityRepository,
	}
}

func (s *Service) GetCompanies(ctx context.Context) ([]Company, error) {
	return s.repository.GetAll(ctx)
}

func (s *Service) CreateCompany(ctx context.Context, request CreateCompanyRequest) (*Company, error) {
	//Create entity for the company
	entity, err :=
		s.entityRepository.Create(
			ctx,
			entity.CreateEntityRequest{
				Type:        "company",
				DisplayName: request.Name,
			},
		)

	if err != nil {
		return nil, err
	}

	input := CreateCompanyInput{
		EntityID:    entity.ID,
		Name:        request.Name,
		LegalName:   request.LegalName,
		Cnpj:        request.Cnpj,
		Industry:    request.Industry,
		Website:     request.Website,
		Description: request.Description,
	}

	return s.repository.Create(ctx, input)
}
