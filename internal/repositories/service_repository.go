package repositories

import (
	"context"
	"errors"

	"github.com/sergeyiksanov/help-on-road/internal/dto"
	"github.com/sergeyiksanov/help-on-road/internal/models"
	"gorm.io/gorm"
)

type ServiceRepository struct {
	db *gorm.DB
}

func NewServiceRepository(db *gorm.DB) *ServiceRepository {
	return &ServiceRepository{
		db: db,
	}
}

func (sr *ServiceRepository) GetAll(tx context.Context) ([]models.Service, error) {
	var servicesDto []dto.Service

	ctx := sr.db
	if tx != nil {
		var ok bool
		ctx, ok = tx.(*gorm.DB)
		if !ok {
			return nil, errors.New("Invalid ctx")
		}
	}

	if err := ctx.Find(&servicesDto).Error; err != nil {
		return nil, err
	}

	services := make([]models.Service, 0, len(servicesDto))
	for _, s := range servicesDto {
		services = append(services, models.Service{
			ID:          s.ID,
			Name:        s.Name,
			Description: s.Description,
			Price:       s.Price,
		})
	}

	return services, nil
}
