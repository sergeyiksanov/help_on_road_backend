package user_controller

import (
	"encoding/json"
	"net/http"
)

type (
	signInRequest struct {
		PhoneNumber string `json:"phone_number,omitempty"`
		Password    string `json:"password,omitempty"`
	}

	signInResponse struct {
		Token string `json:"token,omitempty"`
	}
)

func (uc *UserController) SignIn(w http.ResponseWriter, r *http.Request) {
	var req signInRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid fields", http.StatusBadRequest)
		return
	}

	if req.PhoneNumber == "" || req.Password == "" {
		http.Error(w, "Invalid fields", http.StatusBadRequest)
		return
	}

	token, err := uc.userService.SignIn(r.Context(), req.PhoneNumber, req.Password)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(signInResponse{
		Token: token,
	})
}
