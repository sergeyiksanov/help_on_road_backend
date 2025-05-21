package service_provider

import (
	"github.com/redis/go-redis/v9"
	"github.com/sergeyiksanov/help-on-road/internal/config"
	"github.com/sergeyiksanov/help-on-road/internal/controllers/help_controller"
	"github.com/sergeyiksanov/help-on-road/internal/controllers/user_controller"
	"github.com/sergeyiksanov/help-on-road/internal/models"
	"github.com/sergeyiksanov/help-on-road/internal/repositories"
	"github.com/sergeyiksanov/help-on-road/internal/services/help_service"
	"github.com/sergeyiksanov/help-on-road/internal/services/user_service"
	"github.com/sergeyiksanov/help-on-road/internal/tg_client"
	"gorm.io/gorm"
)

type ServiceProvider struct {
	db                        *gorm.DB
	usersForModerationChannel chan *models.User
	helpChannel               chan *models.HelpCall
	redisClient               *redis.Client
	config                    *config.Config

	transactionManager *repositories.GormTransactionManager
	userRepository     *repositories.UserRepository
	tokensRepository   *repositories.TokensRepository
	heloRepository     *repositories.HelpRepository

	userService *user_service.UserService
	helpService *help_service.HelpService

	userController *user_controller.UserController
	helpController *help_controller.HelpController

	tgClient *tg_client.TelegramClient
}
