package config

import "testing"

func TestApplyEnvOverrides(t *testing.T) {
	t.Setenv("TELEGRAM_API_ID", "123")
	t.Setenv("TELEGRAM_API_HASH", "hash")
	t.Setenv("TELEGRAM_SESSION_FILE", "override.session")

	cfg := Config{
		Telegram: TelegramConfig{
			APIID:       1,
			APIHash:     "old",
			SessionFile: "old.session",
		},
	}

	if err := applyEnvOverrides(&cfg); err != nil {
		t.Fatalf("applyEnvOverrides() error = %v", err)
	}

	if cfg.Telegram.APIID != 123 {
		t.Fatalf("APIID = %d, want 123", cfg.Telegram.APIID)
	}

	if cfg.Telegram.APIHash != "hash" {
		t.Fatalf("APIHash = %q, want %q", cfg.Telegram.APIHash, "hash")
	}

	if cfg.Telegram.SessionFile != "override.session" {
		t.Fatalf(
			"SessionFile = %q, want %q",
			cfg.Telegram.SessionFile,
			"override.session",
		)
	}
}

func TestApplyEnvOverridesInvalidAPIID(t *testing.T) {
	t.Setenv("TELEGRAM_API_ID", "invalid")

	cfg := Config{}

	if err := applyEnvOverrides(&cfg); err == nil {
		t.Fatal("applyEnvOverrides() error = nil, want non-nil")
	}
}
