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
				entity_id,
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
				&person.EntityID,
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

// MARK: GET PERSON BY ID
func (r *Repository) GetById(ctx context.Context, request GetPersonByIdInput) (*Person, error) {
	var person Person
	row := r.db.QueryRow(ctx, `
    SELECT
        entity_id,
        first_name,
        last_name,
        birth_date,
        gender,
        created_at,
        updated_at
    FROM person
    WHERE entity_id = $1
`, request.ID)

	err := row.Scan(
		&person.EntityID,
		&person.FirstName,
		&person.LastName,
		&person.BirthDate,
		&person.Gender,
		&person.CreatedAt,
		&person.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &person, nil
}

// MARK: CREATE NEW PERSON
func (r *Repository) Create(ctx context.Context, request CreatePersonInput) (*Person, error) {
	var person Person

	err := r.db.QueryRow(ctx, `
		INSERT INTO person (entity_id, first_name, last_name, birth_date, gender)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING entity_id, first_name, last_name, birth_date, gender, created_at
	`,
		request.EntityID,
		request.FirstName,
		request.LastName,
		request.BirthDate,
		request.Gender,
	).Scan(
		&person.EntityID,
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

// MARK: UPDATE PERSON
func (r *Repository) Update(
	ctx context.Context,
	request UpdatePersonInput,
) (*Person, error) {

	var person Person

	err := r.db.QueryRow(
		ctx,
		`
		UPDATE person
		SET
			first_name = COALESCE($1, first_name),
			last_name = COALESCE($2, last_name),
			birth_date = COALESCE($3, birth_date),
			gender = COALESCE($4, gender),
			updated_at = $5
		WHERE id = $6
		RETURNING
			entity_id,
			first_name,
			last_name,
			birth_date,
			gender,
			created_at,
			updated_at
		`,
		request.FirstName,
		request.LastName,
		request.BirthDate,
		request.Gender,
		request.UpdatedAt,
		request.ID,
	).Scan(
		&person.EntityID,
		&person.FirstName,
		&person.LastName,
		&person.BirthDate,
		&person.Gender,
		&person.CreatedAt,
		&person.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &person, nil
}

// MARK: GET CONNECTIONS FOR PERSON
func (r *Repository) GetConnectionsForPerson(
	ctx context.Context,
	request GetPersonConnectionsInput,
) ([]GetPersonConnectionsResponse, error) {

	rows, err := r.db.Query(ctx, `
		SELECT
			ec.id,
			ec.relationship_type,
			ec.target_entity_id,
			p.first_name || ' ' || p.last_name AS target_name,
			'person' AS target_type
		FROM entity_connection ec
		JOIN person p
			ON p.entity_id = ec.target_entity_id
		WHERE ec.source_entity_id = $1
	`, request.ID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var connections []GetPersonConnectionsResponse

	for rows.Next() {
		var connection GetPersonConnectionsResponse

		err := rows.Scan(
			&connection.ID,
			&connection.RelationshipType,
			&connection.TargetEntityID,
			&connection.TargetName,
			&connection.TargetType,
		)

		if err != nil {
			return nil, err
		}

		connections = append(connections, connection)
	}

	return connections, nil
}
