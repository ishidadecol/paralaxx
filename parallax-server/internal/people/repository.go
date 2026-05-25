package people

import (
	"context"

	"github.com/google/uuid"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Repository struct {
	Driver neo4j.DriverWithContext
}

func (r *Repository) Create(person Person) (Person, error) {
	session := r.Driver.NewSession(
		context.Background(),
		neo4j.SessionConfig{},
	)

	defer session.Close(context.Background())

	person.ID = uuid.NewString()

	_, err := session.Run(
		context.Background(),
		`
		CREATE (p:Person {
			id: $id,
			name : $name,
			surname: $surname,
			occupation: $occupation,
			city: $city
		})
		`,

		map[string]any{
			"id":         person.ID,
			"name":       person.Name,
			"surname":    person.Surname,
			"occupation": person.Occupation,
			"city":       person.City,
		},
	)

	return person, err
}
