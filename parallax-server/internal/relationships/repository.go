package relationships

import (
	"context"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Repository struct {
	Driver neo4j.DriverWithContext
}

func (r *Repository) Create(rel Relationship) error {
	session := r.Driver.NewSession(
		context.Background(),
		neo4j.SessionConfig{},
	)

	defer session.Close(context.Background())

	query := fmt.Sprintf(`
		MATCH (a:Person {id:$sourcePersonId})
		MATCH (b:Person {id:$targetPersonId})

		CREATE (a)-[:%s]->(b)
	`, rel.Type)

	_, err := session.Run(
		context.Background(),
		query,
		map[string]any{
			"sourcePersonId": rel.SourcePersonID,
			"targetPersonId": rel.TargetPersonID,
		},
	)

	return err
}
