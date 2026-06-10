package connection

import (
	"context"
	"log"
	"time"
)

type Service struct {
	repository Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		repository: *repository,
	}
}

// MARK: CREATE CONNECTION
// TODO:
// Validate source and target entities exist before insert.
// Currently enforced by database foreign keys.
func (s *Service) CreateConnection(ctx context.Context, connection CreateEntityConnectionRequest) (*EntityConnection, error) {
	//convert start and end date to time.Time
	startDate, err := ParseDate(connection.StartDate)

	if err != nil {
		return nil, err
	}

	endDate, err := ParseDate(connection.EndDate)

	if err != nil {
		return nil, err
	}

	input := CreateEntityConnectionInput{
		SourceEntityID:   connection.SourceEntityID,
		TargetEntityID:   connection.TargetEntityID,
		RelationshipType: connection.RelationshipType,
		StartDate:        startDate,
		EndDate:          endDate,
	}

	log.Printf(
		"Source=%q Target=%q Relationship=%q",
		input.SourceEntityID,
		input.TargetEntityID,
		input.RelationshipType,
	)

	return s.repository.Create(ctx, input)
}

// MARK: GET CONNECTIONS
func (s *Service) GetConnections(ctx context.Context) ([]EntityConnection, error) {
	return s.repository.GetAll(ctx)
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

// MARK: GET CONNECTIONS FOR ENTITY
func (s *Service) GetConnectionsForEntity(ctx context.Context, request GetEntityConnectionsRequest) ([]EntityConnectionDetail, error) {
	input := GetEntityConnectionsInput{
		EntityID:   request.EntityID,
		TypeFilter: request.TypeFilter,
	}
	return s.repository.GetConnectionsForEntity(ctx, input)
}
