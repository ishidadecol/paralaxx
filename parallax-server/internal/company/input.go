package company

type CreateCompanyInput struct {
	EntityID    string
	Name        string
	LegalName   *string
	Cnpj        *string
	Industry    *string
	Website     *string
	Description *string
}
