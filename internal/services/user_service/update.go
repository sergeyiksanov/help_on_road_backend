package user_service

import (
	"context"

	"github.com/sergeyiksanov/help-on-road/internal/models"
	"github.com/sergeyiksanov/help-on-road/internal/services"
)

func (s *UserService) Update(ctx context.Context, token string, user *models.User) error {
	userId, err := s.tokenRepository.GetUserIDByToken(ctx, token)
	if err != nil {
		return services.InternalServerError
	}

	user.Id = userId
	user.IsModerate = false
	user.IsValid = false
	if err := s.userRepository.Update(ctx, user); err != nil {
		return err
	}

	s.channelUsersForModeration <- user
	return nil
}
