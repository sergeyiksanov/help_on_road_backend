package help_service

import (
	"context"

	"github.com/sergeyiksanov/help-on-road/internal/models"
	"github.com/sergeyiksanov/help-on-road/internal/services"
)

func (hs *HelpService) GetServices(ctx context.Context, token string) ([]models.Service, error) {
	res, err := hs.servicesRepository.GetAll(ctx)
	if err != nil {
		return []models.Service{}, services.InternalServerError
	}

	return res, nil
}
