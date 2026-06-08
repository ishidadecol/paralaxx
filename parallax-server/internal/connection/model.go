package connection

import "time"

type EntityConnection struct {
	ID string `json:"id"`

	SourceEntityID string `json:"source_id"`
	TargetEntityID string `json:"target_id"`

	RelationshipType string `json:"relationship_type"`

	StartDate *time.Time `json:"start_date"`
	EndDate   *time.Time `json:"end_date"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
