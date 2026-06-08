package connection

type CreateEntityConnectionRequest struct {
	SourceEntityID string `json:"sourceEntityId" validate:"required, uuid"`

	TargetEntityID string `json:"targetEntityId" validate:"required, uuid"`

	RelationshipType string `json:"relationshipType" validate:"required"`

	StartDate *string `json:"startDate"`
	EndDate   *string `json:"endDate"`
}
