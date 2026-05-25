package database

import (
	"context"
	"log"
	"os"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func NewDriver() neo4j.DriverWithContext {
	driver, err := neo4j.NewDriverWithContext(
		os.Getenv("NEO4J_URI"),
		neo4j.BasicAuth(
			os.Getenv("NEO4J_USERNAME"),
			os.Getenv("NEO4J_PASSWORD"),
			"",
		),
	)

	if err != nil {
		log.Fatal(err)
	}

	err = driver.VerifyConnectivity(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to Neo4j")

	return driver
}

