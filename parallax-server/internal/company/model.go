package company

import "time"

type Company struct {
	ID       string `json:"id"`
	EntityID string `json:"entity_id"`

	Name      string  `json:"name"`
	LegalName *string `json:"legal_name"`
	Cnpj      *string `json:"cnpj"`

	Industry    *string `json:"industry"`
	Website     *string `json:"website"`
	Description *string `json:"description"`

	StartDate *time.Time `json:"start_date"`
	EndDate   *time.Time `json:"end_date"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
