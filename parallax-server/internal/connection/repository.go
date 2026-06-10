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

// MARK: GET CONNECTIONS FOR ENTITY
func (r *Repository) GetConnectionsForEntity(
	ctx context.Context,
	input GetEntityConnectionsInput,
) ([]EntityConnectionDetail, error) {

	query := `
		SELECT
			ec.id,
			ec.relationship_type,
			ec.target_entity_id,
			e.display_name,
			e.type,
			ec.start_date,
			ec.end_date
		FROM entity_connection ec
		JOIN entity e
			ON e.id = ec.target_entity_id
		WHERE ec.source_entity_id = $1
	`

	args := []any{input.EntityID}

	if len(input.TypeFilter) > 0 {
		query += `
			AND e.type = ANY($2)
		`
		args = append(args, input.TypeFilter)
	}

	rows, err := r.db.Query(
		ctx,
		query,
		args...,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var connections []EntityConnectionDetail

	for rows.Next() {
		var connection EntityConnectionDetail

		err := rows.Scan(
			&connection.ID,
			&connection.RelationshipType,
			&connection.TargetEntityID,
			&connection.TargetName,
			&connection.TargetType,
			&connection.StartDate,
			&connection.EndDate,
		)

		if err != nil {
			return nil, err
		}

		connections = append(connections, connection)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return connections, nil
}
