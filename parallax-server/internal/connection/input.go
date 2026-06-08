package connection

import "time"

type CreateEntityConnectionInput struct {
	SourceEntityID   string
	TargetEntityID   string
	RelationshipType string
	StartDate        *time.Time
	EndDate          *time.Time
}
