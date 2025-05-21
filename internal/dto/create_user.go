package dto

type User struct {
	ID            int64  `gorm:"column:id;primaryKey;autoIncrement"`
	PhoneNumber   string `gorm:"column:phone_number;uniqueIndex"`
	FirstName     string `gorm:"column:first_name"`
	LastName      string `gorm:"column:last_name"`
	Surname       string `gorm:"column:surname"`
	Password      string `gorm:"column:password"`
	AutoModel     string `gorm:"column:auto_model"`
	AutoGosNumber string `gorm:"column:auto_gos_number"`
	VinCode       string `gorm:"column:vin_code"`
	IsValid       bool   `gorm:"column:is_valid"`
	IsModerate    bool   `gorm:"column:is_moderate"`
}

func (User) TableName() string {
	return "users"
}
