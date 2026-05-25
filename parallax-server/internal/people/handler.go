package people

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	Repository *Repository
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var person Person

	err := json.NewDecoder(r.Body).Decode(&person)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	created, err := h.Repository.Create(person)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(created)
}
