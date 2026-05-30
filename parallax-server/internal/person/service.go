package person

import (
	"context"
	"time"

	"github.com/google/uuid"
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

// MARK: GET PERSON BY ID
func (s *Service) GetPersonById(ctx context.Context, id string) (*Person, error) {
	_, err := uuid.Parse(id)

	if err != nil {
		return nil, err
	}

	input := GetPersonByIdInput{
		ID: id,
	}
	return s.repository.GetById(
		ctx,
		input,
	)
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

// MARK: UPDATE PERSON
func (s *Service) Update(ctx context.Context, request UpdatePersonRequest, id string) (*Person, error) {
	var birthDate *time.Time

	ParseDate(request.BirthDate)

	input := UpdatePersonInput{
		ID:        id,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		BirthDate: birthDate,
		Gender:    request.Gender,
		UpdatedAt: time.Now(),
	}

	return s.repository.Update(ctx, input)
}

func ParseDate(date *string) (*time.Time, error) {

	if date == nil {
		return nil, nil
	}

	parsed, err := time.Parse(
		"2006-01-02",
		*date,
	)

	if err != nil {
		return nil, err
	}

	return &parsed, nil
}
