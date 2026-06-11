package company

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
	}
}

// MARK: GET ALL COMPANIES
func (r *Repository) GetAll(ctx context.Context) ([]Company, error) {
	rows, err := r.db.Query(ctx, `
		SELECT id, entity_id, name, legal_name, cnpj, industry, website, description
		FROM company
	`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var companies []Company

	for rows.Next() {
		var company Company

		err := rows.Scan(
			&company.ID,
			&company.EntityID,
			&company.Name,
			&company.LegalName,
			&company.Cnpj,
			&company.Industry,
			&company.Website,
			&company.Description,
		)

		if err != nil {
			return nil, err
		}

		companies = append(companies, company)
	}

	return companies, nil
}

// MARK: CREATE COMPANY
func (r *Repository) Create(ctx context.Context, input CreateCompanyInput) (*Company, error) {
	var company Company

	err := r.db.QueryRow(ctx, `
		INSERT INTO company (entity_id, name, legal_name, cnpj, industry, website, description)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING entity_id, name, legal_name, cnpj, industry, website, description
	`,
		input.EntityID,
		input.Name,
		input.LegalName,
		input.Cnpj,
		input.Industry,
		input.Website,
		input.Description,
	).Scan(
		&company.EntityID,
		&company.LegalName,
		&company.Cnpj,
		&company.Industry,
		&company.Website,
		&company.Description,
	)

	if err != nil {
		return nil, err
	}

	return &company, nil
}
