package tg_client

import (
	"strconv"

	"github.com/sergeyiksanov/help-on-road/internal/models"
	tgbotapi "github.com/sergeyiksanov/telegram-bot-api"
)

func (c *TelegramClient) sendHelpCallMessage(call *models.HelpCall) {
	text := formatHelpInfo(call)

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Мастер отправлен", "accept_call_"+strconv.FormatInt(call.Caller.Id, 10)),
			tgbotapi.NewInlineKeyboardButtonData("Отклонить вызов", "reject_call_"+strconv.FormatInt(call.Caller.Id, 10)),
		),
	)

	msg := tgbotapi.NewMessage(c.mainChatId, text)
	msg.TopicID = int(c.helpChatThread)
	msg.ReplyMarkup = keyboard
	loc := tgbotapi.NewLocation(c.mainChatId, call.Latitude, call.Longitude)
	loc.TopicID = int(c.helpChatThread)

	c.bot.Send(msg)
	c.bot.Send(loc)
}

func formatHelpInfo(call *models.HelpCall) string {
	return "Новый вызов:\nОписание: " + call.Description + "\nУслуга: " + call.Service + "\nСпособ оплаты: " + call.PayType + "\n\nПользователь\nID: " + strconv.FormatInt(call.Caller.Id, 10) + "\nИмя: " + call.Caller.LastName + " " + call.Caller.FirstName + " " + call.Caller.Surname + "\nНомер телефона: " + call.Caller.PhoneNumber + "\nМодель автомобиля: " + call.Caller.AutoModel + "\nГос номер автомобиля: " + call.Caller.AutoGosNumber + "\nVIN код: " + call.Caller.VinCode
}
