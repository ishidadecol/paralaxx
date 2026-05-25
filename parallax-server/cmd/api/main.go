package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"

	"github.com/ishidadecol/parallax/internal/database"
	"github.com/ishidadecol/parallax/internal/people"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	driver := database.NewDriver()

	personRepo := &people.Repository{
		Driver: driver,
	}

	personHandler := &people.Handler{
		Repository: personRepo,
	}

	r := chi.NewRouter()

	r.Post("/people", personHandler.Create)

	log.Println("API running on :8080")

	http.ListenAndServe(":8080", r)
}

