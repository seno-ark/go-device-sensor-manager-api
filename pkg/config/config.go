package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

type Config struct {
	Port string

	DatabaseHost string
	DatabasePort string
	DatabaseUser string
	DatabasePass string
	DatabaseName string
}

func GetConfig() *Config {
	return &Config{
		Port: os.Getenv("PORT"),

		DatabaseHost: os.Getenv("DB_HOST"),
		DatabasePort: os.Getenv("DB_PORT"),
		DatabaseUser: os.Getenv("DB_USER"),
		DatabasePass: os.Getenv("DB_PASS"),
		DatabaseName: os.Getenv("DB_NAME"),
	}
}
