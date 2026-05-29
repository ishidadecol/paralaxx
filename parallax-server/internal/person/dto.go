package person

type CreatePersonRequest struct {
	FirstName string  `json:"first_name"`
	LastName  *string `json:"last_name"`
	Gender    *string `json:"gender"`
	BirthDate *string `json:"birth_date"`
}
