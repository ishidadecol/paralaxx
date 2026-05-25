package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"

	"github.com/ishidadecol/parallax/internal/database"
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

	r.Post("/people", personHandler.Create)

	//MARK: RELATIONSHIP ROUTE
	relationshipRepo := &relationships.Repository{
		Driver: driver,
	}

	relationshipHandler := &relationships.Handler{
		Repository: relationshipRepo,
	}

	r.Post("/relationships", relationshipHandler.Create)
	log.Println("API running on :8080")

	http.ListenAndServe(":8080", r)
}
