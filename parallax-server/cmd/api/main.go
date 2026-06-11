package main

import (
	"log"
	"net/http"

	"github.com/ishidadecol/parallax/internal/company"
	"github.com/ishidadecol/parallax/internal/connection"
	"github.com/ishidadecol/parallax/internal/database"
	"github.com/ishidadecol/parallax/internal/entity"
	"github.com/ishidadecol/parallax/internal/person"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main() {
	router := chi.NewRouter()
	//TODO: ORGANIZE ROUTING
	//TODO: add logging middleware
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

	// Enity connection repository, service and handler
	entityConnectionRepository := connection.NewRepository(db)
	entityConnectionService := connection.NewService(entityConnectionRepository)
	entityConnectionHandler := connection.NewHandler(entityConnectionService)

	//MARK: Entity service and repository
	entityRepository := entity.NewRepository(db)

	//MARK: Person repository, service, and handler
	personRepository := person.NewRepository(db)
	personService := person.NewService(personRepository, entityRepository)
	personHandler := person.NewHandler(personService)

	//MARK: Company repository and handler
	companyRepository := company.NewRepository(db)
	companyService := company.NewService(companyRepository)
	companyHandler := company.NewHandler(companyService)

	// Define API routes

	//MARK: PERSON ROUTES
	router.Get(
		"/person",
		personHandler.GetPeople,
	)
	router.Get(
		"/person/{id}",
		personHandler.GetPersonById,
	)
	router.Get(
		"/person/{id}/connections",
		personHandler.GetConnectionsForPerson,
	)
	router.Post(
		"/person",
		personHandler.CreatePerson,
	)

	//MARK: ENTITY CONNECTION ROUTES
	router.Get(
		"/connection",
		entityConnectionHandler.GetAllConnections,
	)

	router.Get(
		"/connection/entity/{id}",
		entityConnectionHandler.GetConnectionsOfAnEntity,
	)

	router.Post(
		"/connection",
		entityConnectionHandler.CreateConnection,
	)

	//MARK: COMPANY ROUTES
	router.Get(
		"/company",
		companyHandler.GetAllCompanies,
	)

	log.Println("API running on :8080")

	http.ListenAndServe(
		":8080",
		router,
	)
}
