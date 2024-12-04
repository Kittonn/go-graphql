package main

import (
	"log"

	"github.com/Kittonn/go-graphql/config"
	"github.com/Kittonn/go-graphql/internal/app"
	"github.com/Kittonn/go-graphql/pkg/database/redis"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	redisClient, err := redis.NewRedisClient(config)
	if err != nil {
		log.Fatalf("Failed to connect to redis: %v", err)
	}

	defer redisClient.Client.Close()

	app := app.NewApp(config)
	if err := app.Run(); err != nil {
		log.Fatalf("Server failed to run: %v", err)
	}
}
