package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"time"
)

type Config struct {
	Environment string
	APIBaseURL  string
	Timeout     time.Duration
}

func fromEnv() (*Config, error) {
	env := os.Getenv("SCREENER_ENV")
	if env == "" {
		return nil, fmt.Errorf("SCREENER_ENV is not set")
	}
	url := os.Getenv("BYBIT_URL")
	if url == "" {
		return nil, fmt.Errorf("BYBIT_URL is not set")
	}
	to := os.Getenv("TIMEOUT")
	if to == "" {
		to = "10"
	}

	timeout, err := strconv.Atoi(to)
	if err != nil {
		return nil, fmt.Errorf("Ошибка преобразования TIMEOUT: %w", err)
	}
	return &Config{
		Environment: env,
		APIBaseURL:  url,
		Timeout:     time.Duration(timeout) * time.Second,
	}, nil
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("Ошибка загрузки .env файла: %w", err)
	}

	return fromEnv()
}
