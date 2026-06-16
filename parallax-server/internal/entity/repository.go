package entity

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

// MARK: CREATE NEW ENTITY
func (r *Repository) Create(ctx context.Context, entity CreateEntityRequest) (*Entity, error) {
	var newEntity Entity

	err := r.db.QueryRow(ctx, `
		INSERT INTO entity (type, display_name)
		VALUES ($1, $2)
		RETURNING id, type, display_name, created_at
	`,
		entity.Type,
	).Scan(
		&newEntity.ID,
		&newEntity.Type,
		&newEntity.DisplayName,
		&newEntity.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &newEntity, nil
}

// MARK: GET ENTITIES NAMES
func (r *Repository) GetEntityNames(ctx context.Context, input GetEntitiesNamesInput) ([]EntityLookup, error) {
	query := `
	SELECT
		id,
		display_name,
		type
	FROM entity
`

	args := []any{}

	if len(input.TypeFilter) > 0 {
		query += `
		WHERE type = ANY($1)
	`
		args = append(args, input.TypeFilter)
	}

	query += `
	ORDER BY display_name
`

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entities []EntityLookup

	for rows.Next() {
		var entity EntityLookup

		err := rows.Scan(
			&entity.ID,
			&entity.DisplayName,
			&entity.Type,
		)

		if err != nil {
			return nil, err
		}

		entities = append(entities, entity)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return entities, nil
}
