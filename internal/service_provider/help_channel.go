package service_provider

import "github.com/sergeyiksanov/help-on-road/internal/models"

func (s *ServiceProvider) HelpChannel() chan *models.HelpCall {
	if s.helpChannel == nil {
		s.helpChannel = make(chan *models.HelpCall, 100)
	}

	return s.helpChannel
}
