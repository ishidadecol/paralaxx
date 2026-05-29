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

// MARK: GET ALL
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
			FROM person
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

// MARK: CREATE NEW PERSON
func (r *Repository) Create(ctx context.Context, request CreatePersonInput) (*Person, error) {
	var person Person

	err := r.db.QueryRow(ctx, `
		INSERT INTO person (first_name, last_name, birth_date, gender)
		VALUES ($1, $2, $3, $4)
		RETURNING id, first_name, last_name, birth_date, gender, created_at
	`,
		request.FirstName,
		request.LastName,
		request.BirthDate,
		request.Gender,
	).Scan(
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

	return &person, nil
}
