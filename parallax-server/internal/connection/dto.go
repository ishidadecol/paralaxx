package connection

type CreateEntityConnectionRequest struct {
	SourceEntityID string `json:"sourceEntityId" validate:"required, uuid"`

	TargetEntityID string `json:"targetEntityId" validate:"required, uuid"`

	RelationshipType string `json:"relationshipType" validate:"required"`

	StartDate *string `json:"startDate"`
	EndDate   *string `json:"endDate"`
}

type GetEntityConnectionsRequest struct {
	EntityID   string   `json:"entityId" validate:"required, uuid"`
	TypeFilter []string `json:"typeFilter"`
}

type EntityConnectionDetail struct {
	ID string `json:"id"`

	RelationshipType string `json:"relationshipType"`

	TargetEntityID string `json:"targetEntityId"`

	TargetName string `json:"targetName"`
	TargetType string `json:"targetType"`

	StartDate *string `json:"startDate"`
	EndDate   *string `json:"endDate"`
}
