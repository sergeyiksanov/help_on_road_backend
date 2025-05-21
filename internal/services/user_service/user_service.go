package user_service

import (
	"context"

	"github.com/sergeyiksanov/help-on-road/internal/models"
	"github.com/sergeyiksanov/help-on-road/internal/services"
)

type (
	ITokenRepository interface {
		CreateToken(ctx context.Context, userID int64) (string, error)
		GetUserTokens(ctx context.Context, userID int64) ([]string, error)
		DeleteToken(ctx context.Context, token string) error
		GetUserIDByToken(ctx context.Context, token string) (int64, error)
	}

	IUserRepository interface {
		Create(tx services.TransactionContext, model *models.User) (int64, error)
		GetById(id int64) (*models.User, error)
		Update(tx services.TransactionContext, model *models.User) error
		Delete(tx services.TransactionContext, id int64) error
		GetCountByNumber(number string) (int64, error)
		GetByPhoneNumber(phoneNumber string) (*models.User, error)
		GetNotValidUsers() ([]*models.User, error)
	}

	UserService struct {
		userRepository            IUserRepository
		tokenRepository           ITokenRepository
		channelUsersForModeration chan<- *models.User
		txManager                 services.ITransactionManager
	}
)

func NewUserService(ur IUserRepository, tr ITokenRepository, ch chan<- *models.User, txManager services.ITransactionManager) *UserService {
	usersForModeration, err := ur.GetNotValidUsers()
	if err != nil {
		panic(err)
	}
	for _, user := range usersForModeration {
		ch <- user
	}

	return &UserService{
		userRepository:            ur,
		tokenRepository:           tr,
		channelUsersForModeration: ch,
		txManager:                 txManager,
	}
}
