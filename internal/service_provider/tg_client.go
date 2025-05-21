package service_provider

import "github.com/sergeyiksanov/help-on-road/internal/tg_client"

func (s *ServiceProvider) TgClient() *tg_client.TelegramClient {
	if s.tgClient == nil {
		cl, err := tg_client.NewTelegramClient(s.Config().Tg.Token, s.UserService(), s.HelpService(), s.UsersForModerationChannel(), s.HelpChannel(), s.Config().Tg.MainCharId, s.Config().Tg.ModerationChatId, s.Config().Tg.HelpChatId)
		if err != nil {
			panic(err)
		}

		s.tgClient = cl
	}

	return s.tgClient
}
