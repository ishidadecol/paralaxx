package entity

type CreateEntityRequest struct {
	Type        string `json:"type" validate:"required"`
	DisplayName string `json:"displayName" validate:"required"`
}

type EntityLookup struct {
	ID          string `json:"id"`
	Type        string `json:"type"`
	DisplayName string `json:"displayName"`
}
