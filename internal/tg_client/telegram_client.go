package tg_client

import (
	"context"
	"fmt"
	"log"
	"time"

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

func RecoverWithAlert(alert func(string)) func() {
	return func() {
		if r := recover(); r != nil {
			alert(fmt.Sprintf("PANIC: %v", r))
		}
	}
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
	var offset int

	log.Print("start telegram polling")

	go c.handleUserModeration(ctx)
	go c.handleHelpRequests(ctx)

	for {
		select {
		case <-ctx.Done():
			log.Print("telegram polling stopped")
			return ctx.Err()
		default:
		}

		updates, err := c.bot.GetUpdates(tgbotapi.UpdateConfig{
			Offset:  offset,
			Timeout: 60,
		})
		if err != nil {
			log.Printf("getUpdates error: %v", err)
			c.Alert("Failed get updates: " + err.Error())
			time.Sleep(3 * time.Second)
			continue
		}

		for _, update := range updates {
			offset = update.UpdateID + 1
			c.processUpdate(update)
		}
	}
}

func (c *TelegramClient) handleUserModeration(ctx context.Context) {
	defer RecoverWithAlert(c.Alert)()

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
	defer RecoverWithAlert(c.Alert)()

	for {
		select {
		case <-ctx.Done():
			return
		case call := <-c.helpChannel:
			c.sendHelpCallMessage(call)
		}
	}
}
