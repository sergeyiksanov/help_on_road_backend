package service_provider

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func (s *ServiceProvider) DB() *gorm.DB {
	if s.db == nil {
		db, err := gorm.Open(postgres.Open(s.Config().GetPostgresConnStr()), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		s.db = db
	}

	return s.db
}
