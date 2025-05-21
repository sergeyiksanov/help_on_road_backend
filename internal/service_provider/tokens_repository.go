package service_provider

import "github.com/sergeyiksanov/help-on-road/internal/repositories"

func (s *ServiceProvider) TokensRepository() *repositories.TokensRepository {
	if s.tokensRepository == nil {
		s.tokensRepository = repositories.NewTokensRepository(s.RedisClient(), s.Config().JWT.SecretKey)
	}

	return s.tokensRepository
}
