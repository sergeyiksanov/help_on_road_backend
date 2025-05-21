package tg_client

import (
	"strconv"

	"github.com/sergeyiksanov/help-on-road/internal/models"
	tgbotapi "github.com/sergeyiksanov/telegram-bot-api"
)

func (c *TelegramClient) sendModerationMessage(user *models.User) {
	text := formatUserInfo(user)

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Принять", "accept_user_"+strconv.FormatInt(user.Id, 10)),
			tgbotapi.NewInlineKeyboardButtonData("Отклонить", "reject_user_"+strconv.FormatInt(user.Id, 10)),
		),
	)

	msg := tgbotapi.NewMessage(c.mainChatId, text)
	msg.TopicID = int(c.moderationChatThread)
	msg.ReplyMarkup = keyboard

	c.bot.Send(msg)
}

func formatUserInfo(user *models.User) string {
	return "Новый пользователь:\nID: " + strconv.FormatInt(user.Id, 10) + "\nИмя: " + user.LastName + " " + user.FirstName + " " + user.Surname + "\nНомер телефона: " + user.PhoneNumber + "\nМодель автомобиля: " + user.AutoModel + "\nГос номер автомобиля: " + user.AutoGosNumber + "\nVIN код: " + user.VinCode
}
