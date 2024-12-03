package main

import (
	"log"

	"github.com/Kittonn/go-graphql/config"
	"github.com/Kittonn/go-graphql/internal/app"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	app := app.NewApp(config)
	if err := app.Run(); err != nil {
		log.Fatalf("Server failed to run: %v", err)
	}
}
