package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Env   string
	Port  string
	DBUrl string
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("Error loading .env file: %v", err)
	}

	dbUrl := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	return &Config{
		Env:   os.Getenv("ENV"),
		Port:  os.Getenv("PORT"),
		DBUrl: dbUrl,
	}, nil
}
