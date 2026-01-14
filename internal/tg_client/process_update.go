package tg_client

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/sergeyiksanov/help-on-road/internal/models"
	tgbotapi "github.com/sergeyiksanov/telegram-bot-api"
)

func (c *TelegramClient) processUpdate(update tgbotapi.Update) {
	defer RecoverWithAlert(c.Alert)

	if update.CallbackQuery != nil {
		c.handleCallback(update.CallbackQuery)
		return
	}

	// if update.Message != nil && update.Message.NewChatMembers != nil {
	// 	for _, member := range update.Message.NewChatMembers {
	// 		if member.ID == c.bot.Self.ID {
	// 			c.handleBotAddedToChat(update.Message.Chat)
	// 		} else {
	// 			user := convertTelegramUserToModel(member)
	// 			c.userChannel <- user
	// 		}
	// 	}
	// }

	// if update.Message != nil && update.Message.IsCommand() {
	// c.handleCommand(update.Message)
	// }
}

func (c *TelegramClient) handleCallback(callback *tgbotapi.CallbackQuery) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var err error

	data := callback.Data
	log.Print(data)

	if len(data) > 11 && data[:11] == "accept_user" {
		userID := data[12:]
		var id int64
		id, err = strconv.ParseInt(userID, 10, 64)

		if err == nil {
			err = c.userService.ModerateUser(ctx, id, true)
		}
	} else if len(data) > 11 && data[:11] == "reject_user" {
		userID := data[12:]
		var id int64
		id, err = strconv.ParseInt(userID, 10, 64)

		if err == nil {
			err = c.userService.ModerateUser(ctx, id, false)
		}
	}

	isHelping := false
	if len(data) > 11 && data[:11] == "accept_call" {
		userID := data[12:]
		var id int64
		id, err = strconv.ParseInt(userID, 10, 64)

		if err == nil {
			err = c.helpService.CommitHelp(ctx, id, models.Helping)
			isHelping = true
		}
	} else if len(data) > 11 && data[:11] == "reject_call" {
		userID := data[12:]
		var id int64
		id, err = strconv.ParseInt(userID, 10, 64)

		if err == nil {
			err = c.helpService.CommitHelp(ctx, id, models.Rejected)
		}
	} else if len(data) > 11 && data[:11] == "complt_call" {
		userID := data[12:]
		var id int64
		id, err = strconv.ParseInt(userID, 10, 64)

		if err == nil {
			err = c.helpService.CommitHelp(ctx, id, models.Helped)
		}
	}

	if err == nil {
		// if isHelping {
		var keyb tgbotapi.InlineKeyboardMarkup
		if isHelping {
			completedButton := tgbotapi.NewInlineKeyboardButtonData("✅ Выполнено", "complt_call_"+data[12:])
			keyb = tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(completedButton),
			)
		} else {
			keyb = tgbotapi.InlineKeyboardMarkup{
				InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{},
			}
		}

		editMsg := tgbotapi.NewEditMessageReplyMarkup(
			callback.Message.Chat.ID,
			callback.Message.MessageID,
			keyb,
		)
		_, err = c.bot.Send(editMsg)

		if err != nil {
			errorMsg := tgbotapi.NewMessage(callback.Message.Chat.ID, "Ошибка при обновлении сообщения: "+err.Error())
			c.bot.Send(errorMsg)
		}
		// } else {
		// deleteMsg := tgbotapi.NewDeleteMessage(callback.Message.Chat.ID, callback.Message.MessageID)
		// c.bot.Send(deleteMsg)
		// }
	} else {
		errorMsg := tgbotapi.NewMessage(callback.Message.Chat.ID, "Ошибка при обработке запроса: "+err.Error())
		c.bot.Send(errorMsg)
	}
}
