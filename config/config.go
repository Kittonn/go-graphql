package config

import "github.com/spf13/viper"

type Config struct {
	Port             int    `mapstructure:"PORT"`
	RedisPort        int    `mapstructure:"REDIS_PORT"`
	RedisPassword    string `mapstructure:"REDIS_PASSWORD"`
	RedisHost        string `mapstructure:"REDIS_HOST"`
	RedisDB          int    `mapstructure:"REDIS_DB"`
	PostgresHost     string `mapstructure:"POSTGRES_HOST"`
	PostgresPort     int    `mapstructure:"POSTGRES_PORT"`
	PostgresUsername string `mapstructure:"POSTGRES_USERNAME"`
	PostgresPassword string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresDBName   string `mapstructure:"POSTGRES_DBNAME"`
}

func LoadConfig() (*Config, error) {
	viper.AddConfigPath("./")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
