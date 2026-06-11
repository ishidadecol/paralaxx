package company

import (
	"encoding/json"
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

// MARK: GET ALL COMPANIES
func (h *Handler) GetAllCompanies(w http.ResponseWriter, r *http.Request) {
	companies, err := h.service.GetCompanies(r.Context())

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	json.NewEncoder(w).Encode(companies)
}
