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
func (r *Repository) Create(ctx context.Context, connection CreateEntityConnectionInput) (*EntityConnection, error) {
	var entityConn EntityConnection

	err := r.db.QueryRow(ctx, `
		INSERT INTO entity_connection (
			source_entity_id,
			target_entity_id,
			relationship_type,
			start_date,
			end_date)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, source_entity_id, target_entity_id, relationship_type, start_date, end_date, created_at, updated_at
	`,
		connection.SourceEntityID,
		connection.TargetEntityID,
		connection.RelationshipType,
		connection.StartDate,
		connection.EndDate,
	).Scan(
		&entityConn.ID,
		&entityConn.SourceEntityID,
		&entityConn.TargetEntityID,
		&entityConn.RelationshipType,
		&entityConn.StartDate,
		&entityConn.EndDate,
		&entityConn.CreatedAt,
		&entityConn.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}
	return &entityConn, nil
}

// MARK: GET CONNECTIONS
func (r *Repository) GetAll(ctx context.Context) ([]EntityConnection, error) {
	rows, err := r.db.Query(ctx, `
		SELECT id, source_entity_id, target_entity_id, relationship_type, start_date, end_date, created_at, updated_at
		FROM entity_connection
	`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var connections []EntityConnection

	for rows.Next() {
		var entityConn EntityConnection

		err := rows.Scan(
			&entityConn.ID,
			&entityConn.SourceEntityID,
			&entityConn.TargetEntityID,
			&entityConn.RelationshipType,
			&entityConn.StartDate,
			&entityConn.EndDate,
			&entityConn.CreatedAt,
			&entityConn.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		connections = append(connections, entityConn)
	}

	return connections, nil
}
