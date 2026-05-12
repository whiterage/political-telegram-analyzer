package config

import (
	"fmt"
	"os"
	"strconv"
)

func applyEnvOverrides(cfg *Config) error {
	if value := os.Getenv("TELEGRAM_API_ID"); value != "" {
		id, err := strconv.Atoi(value)
		if err != nil {
			return fmt.Errorf("parse TELEGRAM_API_ID: %w", err)
		}

		cfg.Telegram.APIID = id
	}

	if value := os.Getenv("TELEGRAM_API_HASH"); value != "" {
		cfg.Telegram.APIHash = value
	}

	if value := os.Getenv("TELEGRAM_SESSION_FILE"); value != "" {
		cfg.Telegram.SessionFile = value
	}

	return nil
}
