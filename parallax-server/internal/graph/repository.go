package graph

import (
	"context"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Repository struct {
	Driver neo4j.DriverWithContext
}

func (r *Repository) GetGraph() (Response, error) {
	session := r.Driver.NewSession(
		context.Background(),
		neo4j.SessionConfig{},
	)

	defer session.Close(context.Background())

	result, err := session.Run(
		context.Background(),
		`
		MATCH (a)-[r]->(b)
		RETURN a, r, b
		`,
		nil,
	)

	if err != nil {
		return Response{}, err
	}

	nodeMap := make(map[string]Node)
	edges := []Edge{}

	for result.Next(context.Background()) {
		record := result.Record()

		startNode, _ := record.Get("a")
		endNode, _ := record.Get("b")
		relationship, _ := record.Get("r")

		a := startNode.(neo4j.Node)
		b := endNode.(neo4j.Node)
		r := relationship.(neo4j.Relationship)

		sourceID := a.Props["id"].(string)
		targetID := b.Props["id"].(string)

		nodeMap[sourceID] = Node{
			ID:      sourceID,
			Type:    "person",
			Name:    a.Props["name"].(string),
			Surname: a.Props["surname"].(string),
		}

		nodeMap[targetID] = Node{
			ID:      targetID,
			Type:    "person",
			Name:    b.Props["name"].(string),
			Surname: b.Props["surname"].(string),
		}

		edges = append(edges, Edge{
			ID:     elementID(r),
			Source: sourceID,
			Target: targetID,
			Type:   r.Type,
		})
	}

	nodes := []Node{}

	for _, node := range nodeMap {
		nodes = append(nodes, node)
	}

	return Response{
		Nodes: nodes,
		Edges: edges,
	}, nil
}

func elementID(r neo4j.Relationship) string {
	return r.ElementId
}
