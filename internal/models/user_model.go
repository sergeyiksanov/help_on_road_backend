package models

type User struct {
	Id            int64  `json:"id,omitempty"`
	PhoneNumber   string `json:"phone_number,omitempty"`
	FirstName     string `json:"first_name,omitempty"`
	LastName      string `json:"last_name,omitempty"`
	Surname       string `json:"surname,omitempty"`
	Password      string `json:"password,omitempty"`
	AutoModel     string `json:"auto_model,omitempty"`
	AutoGosNumber string `json:"auto_gos_number,omitempty"`
	VinCode       string `json:"vin_code,omitempty"`
	IsValid       bool   `json:"is_valid,omitempty"`
	IsModerate    bool   `json:"is_moderate,omitempty"`
}
