package service_provider

import "github.com/sergeyiksanov/help-on-road/internal/repositories"

func (s *ServiceProvider) ServicesRepository() *repositories.ServiceRepository {
	if s.servicesRepository == nil {
		s.servicesRepository = repositories.NewServiceRepository(s.DB())
	}

	return s.servicesRepository
}
