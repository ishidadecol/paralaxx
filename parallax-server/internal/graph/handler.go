package graph

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	Repository *Repository
}

func (h *Handler) GetGraph(
	w http.ResponseWriter,
	r *http.Request,
) {
	graph, err := h.Repository.GetGraph()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(graph)
}
