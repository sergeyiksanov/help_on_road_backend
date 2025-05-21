package service_provider

import "github.com/sergeyiksanov/help-on-road/internal/controllers/user_controller"

func (s *ServiceProvider) UserController() *user_controller.UserController {
	if s.userController == nil {
		s.userController = user_controller.NewUserController(s.UserService())
	}

	return s.userController
}
