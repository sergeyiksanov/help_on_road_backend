package service_provider

import "github.com/sergeyiksanov/help-on-road/internal/services/user_service"

func (s *ServiceProvider) UserService() *user_service.UserService {
	if s.userService == nil {
		s.userService = user_service.NewUserService(s.UserRepository(), s.TokensRepository(), s.UsersForModerationChannel(), s.TransactionManager())
	}

	return s.userService
}
