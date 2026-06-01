package connection

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

// MARK: CREATE CONNECTION
func (r *Repository) CreateConnection(ctx context.Context, connection CreateEntityConnectionRequest) (*EntityConnection, error) {
	var entityConn EntityConnection

	err := r.db.QueryRow(ctx, `
		INSERT INTO entity_connection (
			source_type,
			source_id,
			target_type,
			target_id,
			relationship_type,
			start_date,
			end_date)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, source_type, source_id, target_type, target_id, relationship_type, start_date, end_date, created_at, updated_at
	`,
		connection.SourceType,
		connection.SourceID,
		connection.TargetType,
		connection.TargetID,
		connection.RelationshipType,
		connection.StartDate,
		connection.EndDate,
	).Scan(
		&entityConn.ID,
		&entityConn.SourceType,
		&entityConn.SourceID,
		&entityConn.TargetType,
		&entityConn.TargetID,
		&entityConn.RelationshipType,
		&entityConn.StartDate,
		&entityConn.EndDate,
	)

	if err != nil {
		return nil, err
	}
	return &entityConn, nil
}
