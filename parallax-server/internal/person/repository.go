package person

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

func (r *Repository) GetAll(ctx context.Context) ([]Person, error) {

	rows, err :=
		r.db.Query(ctx, `
			SELECT
				id,
				first_name,
				last_name,
				birth_date,
				gender,
				created_at
			FROM people
		`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var people []Person

	for rows.Next() {
		var person Person

		err :=
			rows.Scan(
				&person.ID,
				&person.FirstName,
				&person.LastName,
				&person.BirthDate,
				&person.Gender,
				&person.CreatedAt,
			)

		if err != nil {
			return nil, err
		}

		people = append(people, person)
	}

	return people, nil
}
