package tg_client

import (
	"context"
	"log"

	"github.com/sergeyiksanov/help-on-road/internal/models"
	"github.com/sergeyiksanov/help-on-road/internal/services/help_service"
	"github.com/sergeyiksanov/help-on-road/internal/services/user_service"
	tgbotapi "github.com/sergeyiksanov/telegram-bot-api"
)

type TelegramClient struct {
	bot                  *tgbotapi.BotAPI
	userService          *user_service.UserService
	helpService          *help_service.HelpService
	userChannel          <-chan *models.User
	helpChannel          <-chan *models.HelpCall
	mainChatId           int64
	moderationChatThread int64
	helpChatThread       int64
}

func NewTelegramClient(token string, userService *user_service.UserService, helpService *help_service.HelpService, userChannel chan *models.User, helpChannel chan *models.HelpCall, mainChatId, moderationChatThread, helpChatThread int64) (*TelegramClient, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	log.Print("init bot api")

	return &TelegramClient{
		bot:                  bot,
		userService:          userService,
		helpService:          helpService,
		userChannel:          userChannel,
		helpChannel:          helpChannel,
		mainChatId:           mainChatId,
		moderationChatThread: moderationChatThread,
		helpChatThread:       helpChatThread,
	}, nil
}

func (c *TelegramClient) Start(ctx context.Context) error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := c.bot.GetUpdatesChan(u)
	log.Print("init upd chan")

	go c.handleUserModeration(ctx)
	log.Print("handle user modertaion msgs")
	go c.handleHelpRequests(ctx)

	for update := range updates {
		go c.processUpdate(update)
	}

	return nil
}

func (c *TelegramClient) handleUserModeration(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case user := <-c.userChannel:
			c.sendModerationMessage(user)
		}
	}
}

func (c *TelegramClient) handleHelpRequests(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case call := <-c.helpChannel:
			c.sendHelpCallMessage(call)
		}
	}
}
