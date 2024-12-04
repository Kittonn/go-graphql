package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/Kittonn/go-graphql/config"
	"github.com/redis/go-redis/v9"
)

const ctxTimeOut = 10

type redisManager struct {
	Client *redis.Client
}

func NewRedisClient(config *config.Config) (*redisManager, error) {
	addr := fmt.Sprintf("%v:%v", config.RedisHost, config.RedisPort)

	client := redis.NewClient(&redis.Options{
		Password: config.RedisPassword,
		DB:       config.RedisDB,
		Addr:     addr,
	})

	ctx, cancel := context.WithTimeout(context.Background(), ctxTimeOut*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return &redisManager{
		Client: client,
	}, nil
}
