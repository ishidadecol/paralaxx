package main

import (
	"log"
	"net/http"

	"github.com/ishidadecol/parallax/internal/database"
	"github.com/ishidadecol/parallax/internal/person"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main() {
	router := chi.NewRouter()

	// Add CORS middleware
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}).Handler)

	// Initialize database connection pool
	db := database.NewPostgresPool()

	//MARK: Person repository, service, and handler
	personRepository := person.NewRepository(db)

	personService := person.NewService(personRepository)

	personHandler := person.NewHandler(personService)

	// Define API routes
	router.Get(
		"/person",
		personHandler.GetPeople,
	)

	router.Post(
		"/person",
		personHandler.CreatePerson,
	)

	log.Println("API running on :8080")

	http.ListenAndServe(
		":8080",
		router,
	)
}
