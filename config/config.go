package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config holds database connection information
type Config struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
}

// LoadConfig reads environment variables and returns a Config struct
func LoadConfig() (*Config, error) {
	cfg := &Config{}

	// Load environment variables from .env file
	envLoadErr := godotenv.Load()
	if envLoadErr != nil {
		return nil, fmt.Errorf("error loading .env file: %w", envLoadErr)
	}

	var err error
	cfg.Host = os.Getenv("DB_HOST")
	cfg.Port, err = strconv.Atoi(os.Getenv("DB_PORT"))

	if err != nil {
		return nil, fmt.Errorf("error parsing DB_PORT: %w", err)
	}
	cfg.Username = os.Getenv("DB_USERNAME")
	cfg.Password = os.Getenv("DB_PASSWORD")
	cfg.DBName = os.Getenv("DB_NAME")

	return cfg, nil
}
