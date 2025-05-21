package help_service

import (
	"context"

	"github.com/sergeyiksanov/help-on-road/internal/models"
	"github.com/sergeyiksanov/help-on-road/internal/services"
)

func (hs *HelpService) HelpCall(ctx context.Context, token string, helpCall *models.HelpCall) error {
	userId, err := hs.tokenRepository.GetUserIDByToken(ctx, token)
	if err != nil {
		return services.InternalServerError
	}

	user, err := hs.userRepository.GetById(userId)
	if err != nil {
		return services.InternalServerError
	}
	if !user.IsValid {
		return services.AccessDenied
	}

	helpCall.Caller = user
	helpCall.Status = models.Pending

	l, err := hs.helpRepository.GetByUserId(ctx, userId)
	if err != nil {
		return err
	}
	for _, el := range l {
		if el.Status == models.Pending || el.Status == models.Helping {
			return services.AlreadyExistsError
		}
	}
	// if len(l) != 0 {
	// return services.AlreadyExistsError
	// }
	if err := hs.helpRepository.Add(ctx, helpCall); err != nil {
		return err
	}
	hs.helpChannel <- helpCall

	return nil
}
