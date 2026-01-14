package tg_client

import tgbotapi "github.com/sergeyiksanov/telegram-bot-api"

func (c *TelegramClient) Alert(errMsg string) {
	msg := tgbotapi.NewMessage(c.mainChatId, "⚠️ "+errMsg)
	msg.TopicID = int(c.moderationChatThread)

	c.bot.Send(msg)
}
