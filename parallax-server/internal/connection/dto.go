package connection

type CreateEntityConnectionRequest struct {
	SourceType string `json:"sourceType" validate:"required"`
	SourceID   string `json:"sourceId" validate:"required, uuid"`

	TargetType string `json:"targetType" validate:"required"`
	TargetID   string `json:"targetId" validate:"required, uuid"`

	RelationshipType string `json:"relationshipType" validate:"required"`

	StartDate *string `json:"startDate"`
	EndDate   *string `json:"endDate"`
}
