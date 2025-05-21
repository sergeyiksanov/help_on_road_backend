package service_provider

import "github.com/sergeyiksanov/help-on-road/internal/repositories"

func (s *ServiceProvider) TransactionManager() *repositories.GormTransactionManager {
	if s.transactionManager == nil {
		s.transactionManager = repositories.NewGormTransactionManager(s.DB())
	}

	return s.transactionManager
}
