package user_controller

import (
	"encoding/json"
	"net/http"
)

type userResp struct {
	Id            int64  `json:"id,omitempty"`
	PhoneNumber   string `json:"phone_number,omitempty"`
	FirstName     string `json:"first_name,omitempty"`
	LastName      string `json:"last_name,omitempty"`
	Surname       string `json:"surname,omitempty"`
	AutoModel     string `json:"auto_model,omitempty"`
	AutoGosNumber string `json:"auto_gos_number,omitempty"`
	VinCode       string `json:"vin_code,omitempty"`
	IsValid       bool   `json:"is_valid,omitempty"`
	IsModerate    bool   `json:"is_moderate,omitempty"`
	AutoYear      string `json:"auto_year"`
}

func (c *UserController) GetMe(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, "Invalid fields", http.StatusBadRequest)
		return
	}

	user, err := c.userService.GetByToken(r.Context(), token)
	if err != nil {
		http.Error(w, "Invalid fields", http.StatusBadRequest)
		return
	}

	resp := userResp{
		Id:            user.Id,
		PhoneNumber:   user.PhoneNumber,
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		Surname:       user.Surname,
		AutoModel:     user.AutoModel,
		AutoGosNumber: user.AutoGosNumber,
		VinCode:       user.VinCode,
		IsValid:       user.IsValid,
		IsModerate:    user.IsModerate,
		AutoYear:      user.AutoYear,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
