package help_controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/sergeyiksanov/help-on-road/internal/models"
)

type (
	helpRequest struct {
		Service     string  `json:"service,omitempty"`
		Latitude    float64 `json:"latitude,omitempty"`
		Longitude   float64 `json:"longitude,omitempty"`
		Description string  `json:"description,omitempty"`
		PayType     string  `json:"pay_type,omitempty"`
	}
)

func (hc *HelpController) Help(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, "Invalid fields", http.StatusBadRequest)
		return
	}

	var req helpRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Print(err)
		http.Error(w, "Invalid fields", http.StatusBadRequest)
		return
	}

	if req.Service == "" {
		log.Print("null fields")
		http.Error(w, "Invalid fields", http.StatusBadRequest)
		return
	}

	if err := hc.helpService.HelpCall(r.Context(), token, &models.HelpCall{
		Service:     req.Service,
		Latitude:    req.Latitude,
		Longitude:   req.Longitude,
		Description: req.Description,
		Status:      models.Pending,
		PayType:     req.PayType,
	}); err != nil {
		log.Print(err)
		http.Error(w, "Invalid fields", http.StatusBadRequest)
		return
	}

	response := map[string]interface{}{
		"success": true,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
