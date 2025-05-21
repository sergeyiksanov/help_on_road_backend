package user_service

import (
	"context"
	"errors"

	"github.com/sergeyiksanov/help-on-road/internal/models"
	"github.com/sergeyiksanov/help-on-road/internal/utils"
)

func (us *UserService) SignUp(ctx context.Context, user *models.User) error {
	cnt, err := us.userRepository.GetCountByNumber(user.PhoneNumber)
	if err != nil {
		return errors.New("Internal server error")
	}
	if cnt > 0 {
		return errors.New("User already exist")
	}

	user.IsValid = false
	user.IsModerate = false

	hash, err := utils.HashPassword(user.Password)
	if err != nil {
		return errors.New("Internal server error")
	}

	user.Password = hash

	_, err = us.userRepository.Create(nil, user)
	if err != nil {
		return err
	}

	us.channelUsersForModeration <- user

	return nil
}
