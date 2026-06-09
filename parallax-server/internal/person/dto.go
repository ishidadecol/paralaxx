package person

type CreatePersonRequest struct {
	FirstName string  `json:"first_name"`
	LastName  *string `json:"last_name"`
	Gender    *string `json:"gender"`
	BirthDate *string `json:"birth_date"`
}

type UpdatePersonRequest struct {
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Gender    *string `json:"gender"`
	BirthDate *string `json:"birth_date"`
}

type GetPersonConnectionsResponse struct {
	ID string `json:"id"`

	RelationshipType string `json:"relationshipType"`

	TargetEntityID string `json:"targetEntityId"`

	TargetName string `json:"targetName"`
	TargetType string `json:"targetType"`
}
