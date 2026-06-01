package entity

type CreateEntityRequest struct {
	Type string `json:"type" validate:"required"`
}
