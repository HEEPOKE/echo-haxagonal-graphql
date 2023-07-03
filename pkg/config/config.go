package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	Cfg *Config
)

type Config struct {
	PORT      string
	MONGO_URL string
	DB_NAME   string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return nil, err
	}

	config := &Config{
		PORT:      os.Getenv("PORT"),
		MONGO_URL: os.Getenv("MONGO_URL"),
		DB_NAME:   os.Getenv("DB_NAME"),
	}

	Cfg = config

	return config, nil
}
