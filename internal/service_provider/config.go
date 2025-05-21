package service_provider

import "github.com/sergeyiksanov/help-on-road/internal/config"

func (s *ServiceProvider) Config() *config.Config {
	if s.config == nil {
		s.config = config.LoadConfig()
	}

	return s.config
}
