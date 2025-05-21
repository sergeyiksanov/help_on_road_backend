package user_controller

import (
	"encoding/json"
	"net/http"

	"github.com/sergeyiksanov/help-on-road/internal/models"
)

type (
	signUpRequest struct {
		PhoneNumber   string `json:"phone_number,omitempty"`
		Password      string `json:"password,omitempty"`
		FirstName     string `json:"first_name,omitempty"`
		LastName      string `json:"last_name,omitempty"`
		Surname       string `json:"surname,omitempty"`
		AutoModel     string `json:"auto_model,omitempty"`
		AutoGosNumber string `json:"auto_gos_number,omitempty"`
		VinCode       string `json:"vin_code,omitempty"`
	}
)

func (uc *UserController) SignUp(w http.ResponseWriter, r *http.Request) {
	var req signUpRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid fields", http.StatusBadRequest)
		return
	}

	if req.PhoneNumber == "" || req.Password == "" || req.FirstName == "" || req.LastName == "" {
		http.Error(w, "Invalid fields", http.StatusBadRequest)
		return
	}

	userModel := models.User{
		PhoneNumber:   req.PhoneNumber,
		Password:      req.Password,
		FirstName:     req.FirstName,
		LastName:      req.LastName,
		Surname:       req.Surname,
		AutoModel:     req.AutoModel,
		AutoGosNumber: req.AutoGosNumber,
		VinCode:       req.VinCode,
	}

	err := uc.userService.SignUp(r.Context(), &userModel)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"success": true,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
