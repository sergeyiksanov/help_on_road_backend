package executor_service

import (
	"context"

	"github.com/sergeyiksanov/help-on-road/internal/models"
)

func (s *ExecutorService) GetTask(ctx context.Context, token string, userID int64) (*models.HelpCall, error) {
	return nil, nil // TODO: mock get task for executor
}
