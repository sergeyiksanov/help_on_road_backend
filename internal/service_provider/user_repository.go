package service_provider

import "github.com/sergeyiksanov/help-on-road/internal/repositories"

func (s *ServiceProvider) UserRepository() *repositories.UserRepository {
	if s.userRepository == nil {
		s.userRepository = repositories.NewUserRepository(s.DB())
	}

	return s.userRepository
}
