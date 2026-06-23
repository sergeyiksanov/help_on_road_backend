package executor_service

import (
	"context"

	"github.com/sergeyiksanov/help-on-road/internal/models"
)

func (s *ExecutorService) UpdateCoordinates(ctx context.Context, token string, coordinates *models.Point) error {
	return nil // TODO: mock update coordinates
}
