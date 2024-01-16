package redis

import (
	"context"
	"github.com/blazee5/auth-microservice/internal/config"
	"github.com/redis/go-redis/v9"
)

func NewRedisClient(cfg *config.Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Host,
		Password: cfg.Redis.Password,
		DB:       0,
	})

	err := client.Ping(context.Background()).Err()

	if err != nil {
		panic(err)
	}

	return client
}
