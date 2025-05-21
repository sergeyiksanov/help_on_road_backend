package service_provider

import "github.com/sergeyiksanov/help-on-road/internal/controllers/help_controller"

func (s *ServiceProvider) HelpController() *help_controller.HelpController {
	if s.helpController == nil {
		s.helpController = help_controller.NewHelpController(s.HelpService())
	}

	return s.helpController
}
