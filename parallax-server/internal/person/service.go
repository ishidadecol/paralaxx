package person

import (
	"context"
	"time"
)

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

// MARK: CREATE NEW PERSON
func (s *Service) Create(ctx context.Context, request CreatePersonRequest) (*Person, error) {
	var birthDate *time.Time

	//If theres a birth date, we need to parse it from string to time.Time
	if request.BirthDate != nil {

		parsed, err :=
			time.Parse(
				"2006-01-02",
				*request.BirthDate,
			)

		if err != nil {
			return nil, err
		}

		birthDate = &parsed
	}

	input := CreatePersonInput{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		BirthDate: birthDate,
		Gender:    request.Gender,
	}

	return s.repository.Create(ctx, input)
}
