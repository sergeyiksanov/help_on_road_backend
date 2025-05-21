package service_provider

import "github.com/sergeyiksanov/help-on-road/internal/repositories"

func (s *ServiceProvider) HelpRepository() *repositories.HelpRepository {
	if s.heloRepository == nil {
		s.heloRepository = repositories.NewHelpRepository(s.RedisClient())
	}

	return s.heloRepository
}
