package main

import (
	"log"
	"net/http"

	"github.com/ishidadecol/parallax/internal/database"
	"github.com/ishidadecol/parallax/internal/person"

	"github.com/go-chi/chi/v5"
)

func main() {
	// Initialize database connection pool
	db := database.NewPostgresPool()

	//MARK: Person repository, service, and handler
	personRepository :=
		person.NewRepository(db)

	personService :=
		person.NewService(personRepository)

	personHandler :=
		person.NewHandler(personService)

	router := chi.NewRouter()

	// Define API routes
	router.Get(
		"/people",
		personHandler.GetPeople,
	)

	log.Println("API running on :8080")

	http.ListenAndServe(
		":8080",
		router,
	)
}
