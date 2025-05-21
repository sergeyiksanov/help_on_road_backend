package user_service

import (
	"context"

	"github.com/sergeyiksanov/help-on-road/internal/models"
	"github.com/sergeyiksanov/help-on-road/internal/services"
)

func (us *UserService) GetByToken(ctx context.Context, token string) (*models.User, error) {
	userId, err := us.tokenRepository.GetUserIDByToken(ctx, token)
	if err != nil {
		return nil, services.InternalServerError
	}

	user, err := us.userRepository.GetById(userId)
	if err != nil {
		return nil, err
	}

	return user, nil
}
