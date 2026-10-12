package config

import (
	"testing"
	"time"
)

func TestFromEnv(t *testing.T) {
	t.Setenv("SCREENER_ENV", "development")
	t.Setenv("BYBIT_URL", "https://api.bybit.com")
	t.Setenv("TIMEOUT", "10")

	cfg, err := fromEnv()
	if err != nil {
		t.Fatalf("fromEnv() returned an unexpected error: %v", err)
	}

	if cfg.Environment != "development" {
		t.Errorf("Environment = %q, want %q", cfg.Environment, "development")
	}
	if cfg.APIBaseURL != "https://api.bybit.com" {
		t.Errorf("APIBaseURL = %q, want %q", cfg.APIBaseURL, "https://api.bybit.com")
	}
	if cfg.Timeout != 10*time.Second {
		t.Errorf("Timeout = %v, want %v", cfg.Timeout, 10*time.Second)
	}
}

func TestFromEnvInvalidTimeout(t *testing.T) {
	t.Setenv("SCREENER_ENV", "development")
	t.Setenv("BYBIT_URL", "https://api.bybit.com")
	t.Setenv("TIMEOUT", "not-a-number")

	_, err := fromEnv()
	if err == nil {
		t.Fatal("fromEnv() returned nil error, want an error")
	}
}
