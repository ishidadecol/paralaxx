package person

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetPeople(w http.ResponseWriter, r *http.Request) {

	people, err :=
		h.service.GetPeople(r.Context())

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	json.NewEncoder(w).Encode(people)
}

// MARK: GET PERSON BY ID
func (h *Handler) GetPersonById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	person, err :=
		h.service.GetPersonById(
			r.Context(),
			id,
		)

	if err != nil {

		if errors.Is(err, pgx.ErrNoRows) {
			http.Error(
				w,
				"person not found",
				http.StatusNotFound,
			)

			return
		}

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(person)
}

// MARK: CREATE NEW PERSON
func (h *Handler) CreatePerson(w http.ResponseWriter, r *http.Request) {

	var request CreatePersonRequest

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)

		return
	}

	person, err := h.service.Create(r.Context(), request)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	json.NewEncoder(w).Encode(person)
}
