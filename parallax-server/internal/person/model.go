package person

import "time"

type Person struct {
	EntityID string `json:"entity_id"`

	FirstName string  `json:"first_name"`
	LastName  *string `json:"last_name"`

	BirthDate *time.Time `json:"birth_date"`

	Gender *string `json:"gender"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
