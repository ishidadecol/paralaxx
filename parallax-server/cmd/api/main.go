package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	"github.com/ishidadecol/parallax/internal/database"
	"github.com/ishidadecol/parallax/internal/graph"
	"github.com/ishidadecol/parallax/internal/people"
	"github.com/ishidadecol/parallax/internal/relationships"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	driver := database.NewDriver()

	//MARK: PERSON ROUTE
	personRepo := &people.Repository{
		Driver: driver,
	}

	personHandler := &people.Handler{
		Repository: personRepo,
	}

	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:3000",
		},
		AllowedMethods: []string{
			"GET",
			"POST",
		},
		AllowedHeaders: []string{
			"Accept",
			"Authorization",
			"Content-Type",
		},
	}))

	r.Post("/people", personHandler.Create)

	//MARK: RELATIONSHIP ROUTE
	relationshipRepo := &relationships.Repository{
		Driver: driver,
	}

	relationshipHandler := &relationships.Handler{
		Repository: relationshipRepo,
	}

	r.Post("/relationships", relationshipHandler.Create)

	//MARK: GRAPH ROUTE
	graphRepo := &graph.Repository{
		Driver: driver,
	}

	graphHandler := &graph.Handler{
		Repository: graphRepo,
	}

	r.Get("/graph", graphHandler.GetGraph)
	log.Println("API running on :8080")

	err = http.ListenAndServe(":8080", r)

	if err != nil {
		log.Fatal(err)
	}
}
