package connection

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

// MARK: CREATE CONNECTION
func (h *Handler) CreateConnection(w http.ResponseWriter, r *http.Request) {
	var request CreateEntityConnectionRequest

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)

		return
	}

	connection, err := h.service.CreateConnection(r.Context(), request)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	json.NewEncoder(w).Encode(connection)
}

// MARK: GET ALL CONNECTIONS
func (h *Handler) GetAllConnections(w http.ResponseWriter, r *http.Request) {

	connections, err := h.service.GetConnections(r.Context())

	if err != nil {
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

	json.NewEncoder(w).Encode(connections)
}

// MARK: GET CONNECTIONS OF AN ENTITY
func (h *Handler) GetConnectionsOfAnEntity(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	typeFilter := r.URL.Query()["type"]

	input := GetEntityConnectionsRequest{
		EntityID:   id,
		TypeFilter: typeFilter,
	}

	connections, err := h.service.GetConnectionsForEntity(
		r.Context(),
		input,
	)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(connections)
}
