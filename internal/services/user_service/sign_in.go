package user_service

import (
	"context"

	"github.com/sergeyiksanov/help-on-road/internal/utils"
)

func (us *UserService) SignIn(ctx context.Context, phoneNumber, password string) (string, error) {
	user, err := us.userRepository.GetByPhoneNumber(phoneNumber)
	if err != nil {
		return "", err
	}

	if !utils.ValidatePassword(password, user.Password) {
		return "", err
	}

	token, err := us.tokenRepository.CreateToken(ctx, user.Id)
	if err != nil {
		return "", err
	}

	return token, nil
}
