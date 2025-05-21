package help_service

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
	}

	IHelpRepository interface {
		GetAll(ctx context.Context) ([]*models.HelpCall, error)
		GetByUserId(ctx context.Context, id int64) ([]*models.HelpCall, error)
		Add(ctx context.Context, m *models.HelpCall) error
		Delete(ctx context.Context, id int64) error
		Update(ctx context.Context, userID int64, index int, updatedCall *models.HelpCall) error
	}

	HelpService struct {
		tokenRepository ITokenRepository
		userRepository  IUserRepository
		helpRepository  IHelpRepository
		helpChannel     chan<- *models.HelpCall
	}
)

func NewHelpService(tr ITokenRepository, ur IUserRepository, hr IHelpRepository, hc chan *models.HelpCall) *HelpService {
	hs := HelpService{
		tokenRepository: tr,
		userRepository:  ur,
		helpRepository:  hr,
		helpChannel:     hc,
	}

	calls, err := hs.helpRepository.GetAll(context.Background())
	if err != nil {
		panic(err)
	}

	for _, call := range calls {
		if call.Status == models.Pending || call.Status == models.Helping {
			hs.helpChannel <- call
		}
	}

	return &hs
}
