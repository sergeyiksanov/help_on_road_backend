package tg_client

import (
	"fmt"
	"log"
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

	if _, err := c.bot.Send(msg); err != nil {
		c.Alert(fmt.Sprintf("%s\n%s", err.Error(), msg.Text))
		log.Println("Не удалось отправки сообщение модерации: ", err)
	}
}

func formatUserInfo(user *models.User) string {
	return "Новый пользователь:\nID: " + strconv.FormatInt(user.Id, 10) + "\nИмя: " + user.LastName + " " + user.FirstName + " " + user.Surname + "\nНомер телефона: " + user.PhoneNumber + "\nМодель автомобиля: " + user.AutoModel + "\nГос номер автомобиля: " + user.AutoGosNumber + "\nVIN код: " + user.VinCode + "\nГод выпуска автомобиля: " + user.AutoYear
}
