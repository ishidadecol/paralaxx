package relationships

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	Repository *Repository
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var rel Relationship

	err := json.NewDecoder(r.Body).Decode(&rel)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.Repository.Create(rel)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
