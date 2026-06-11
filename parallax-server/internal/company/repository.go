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
		SELECT id, entity_id, name, legal_name, cnpj, industry, website, description, start_date, end_date
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
			&company.StartDate,
			&company.EndDate,
		)

		if err != nil {
			return nil, err
		}

		companies = append(companies, company)
	}

	return companies, nil
}
