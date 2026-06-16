package entity

import (
	"encoding/json"
	"log"
	"net/http"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

// MARK: Get all entities names
func (h *Handler) GetAllEntitiesNames(w http.ResponseWriter, r *http.Request) {

	typeFilter := r.URL.Query()["filter"]
	log.Printf("filters: %+v", typeFilter)

	entities, err := h.service.GetAllEntitiesNames(r.Context(), typeFilter)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(entities)
}
