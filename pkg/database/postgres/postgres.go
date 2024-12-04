package postgres

import (
	"fmt"

	"github.com/Kittonn/go-graphql/config"
	postgresDriver "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	DB *gorm.DB
}

func NewPostgres(config *config.Config) (*Postgres, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Bangkok",
		config.PostgresHost,
		config.PostgresUsername,
		config.PostgresPassword,
		config.PostgresDBName,
		config.PostgresPort)

	db, err := gorm.Open(postgresDriver.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return &Postgres{
		DB: db,
	}, nil
}
