package service_provider

import "github.com/sergeyiksanov/help-on-road/internal/services/help_service"

func (s *ServiceProvider) HelpService() *help_service.HelpService {
	if s.helpService == nil {
		s.helpService = help_service.NewHelpService(s.TokensRepository(), s.UserRepository(), s.HelpRepository(), s.HelpChannel())
	}

	return s.helpService
}
