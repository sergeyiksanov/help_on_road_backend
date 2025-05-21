package service_provider

import "github.com/redis/go-redis/v9"

func (s *ServiceProvider) RedisClient() *redis.Client {
	if s.redisClient == nil {
		s.redisClient = redis.NewClient(&redis.Options{
			Addr:     s.Config().GetRedisAddr(),
			Password: s.Config().Redis.Password,
			DB:       s.Config().Redis.DB,
		})
	}

	return s.redisClient
}
