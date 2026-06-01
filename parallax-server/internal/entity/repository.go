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
		INSERT INTO entity (type)
		VALUES ($1)
		RETURNING id, type, created_at
	`,
		entity.Type,
	).Scan(
		&newEntity.ID,
		&newEntity.Type,
		&newEntity.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &newEntity, nil
}
