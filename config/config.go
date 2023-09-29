package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	HTTPPort string

	PostgresPort     int
	PostgresHost     string
	PostgresUser     string
	PostgresDB       string
	PostgresPassword string
}

func Load() Config {
	var config Config

	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	v := viper.New()
	v.AutomaticEnv()

	v.SetDefault("HTTP_PORT", ":8080")

	config.HTTPPort = v.GetString("HTTP_PORT")
	config.PostgresDB = v.GetString("POSTGRES_DB")
	config.PostgresHost = v.GetString("POSTGRES_HOST")
	config.PostgresUser = v.GetString("POSTGRES_USER")
	config.PostgresPort = v.GetInt("POSTGRES_PORT")
	config.PostgresPassword = v.GetString("POSTGRES_PASSWORD")

	return config
}
