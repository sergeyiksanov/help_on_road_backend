package user_service

import (
	"context"
)

func (us *UserService) ModerateUser(ctx context.Context, userId int64, result bool) error {
	user, err := us.userRepository.GetById(userId)
	if err != nil {
		return err
	}

	user.IsModerate = true
	user.IsValid = result

	return us.userRepository.Update(nil, user)
}
