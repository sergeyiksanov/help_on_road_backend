package help_service

import (
	"context"

	"github.com/sergeyiksanov/help-on-road/internal/models"
	"github.com/sergeyiksanov/help-on-road/internal/services"
)

func (hs *HelpService) GetByToken(ctx context.Context, token string) ([]*models.HelpCall, error) {
	userId, err := hs.tokenRepository.GetUserIDByToken(ctx, token)
	if err != nil {
		return []*models.HelpCall{}, services.InternalServerError
	}

	return hs.helpRepository.GetByUserId(ctx, userId)
}
