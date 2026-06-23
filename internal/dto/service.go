package dto

type Service struct {
	ID          int64  `gorm:"column:id;primaryKey;autoIncrement"`
	Name        string `gorm:"column:name"`
	Description string `gorm:"column:description"`
	Price       int64  `gorm:"column:price"`
}

func (Service) TableName() string {
	return "services"
}
