package connection

import "time"

type EntityConnection struct {
	ID string `json:"id"`

	SourceType string `json:"source_type"`
	SourceID   string `json:"source_id"`

	TargetType string `json:"target_type"`
	TargetID   string `json:"target_id"`

	RelationshipType string `json:"relationship_type"`

	StartDate *time.Time `json:"start_date"`
	EndDate   *time.Time `json:"end_date"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
