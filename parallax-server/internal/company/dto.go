package company

type CreateCompanyRequest struct {
	Name        string  `json:"name"`
	LegalName   *string `json:"legal_name"`
	Cnpj        *string `json:"cnpj"`
	Industry    *string `json:"industry"`
	Website     *string `json:"website"`
	Description *string `json:"description"`
}
