package service_provider

import "github.com/sergeyiksanov/help-on-road/internal/models"

func (s *ServiceProvider) UsersForModerationChannel() chan *models.User {
	if s.usersForModerationChannel == nil {
		s.usersForModerationChannel = make(chan *models.User, 50)
	}

	return s.usersForModerationChannel
}
