package help_controller

import (
	"encoding/json"
	"net/http"
)

func (hc *HelpController) Get(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, "Invalid fields", http.StatusBadRequest)
		return
	}

	helps, err := hc.helpService.GetByToken(r.Context(), token)
	if err != nil {
		http.Error(w, "Invalid fields", http.StatusBadRequest)
		return
	}

	response := map[string]interface{}{
		"helps": helps,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
