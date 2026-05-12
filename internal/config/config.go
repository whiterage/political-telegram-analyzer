package config

import "errors"

type Config struct {
	InputFile                string         `yaml:"input_file"`
	OutputFile               string         `yaml:"output_file"`
	SummaryOutputFile        string         `yaml:"summary_output_file"`
	TopLimit                 int            `yaml:"top_limit"`
	ChannelSummaryOutputFile string         `yaml:"channel_summary_output_file"`
	Source                   string         `yaml:"source"`
	Telegram                 TelegramConfig `yaml:"telegram"`
}

type TelegramConfig struct {
	APIID       int             `yaml:"api_id"`
	APIHash     string          `yaml:"api_hash"`
	SessionFile string          `yaml:"session_file"`
	Channels    []ChannelConfig `yaml:"channels"`
}

type ChannelConfig struct {
	Username string `yaml:"username"`
	Name     string `yaml:"name"`
}

func (c Config) Validate() error {
	if c.InputFile == "" {
		return errors.New("input_file is required")
	}
	if c.OutputFile == "" {
		return errors.New("output_file is required")
	}
	if c.TopLimit <= 0 {
		return errors.New("top_limit must be greater than zero")
	}
	if c.SummaryOutputFile == "" {
		return errors.New("summary_output_file is required")
	}
	if c.ChannelSummaryOutputFile == "" {
		return errors.New("channel_summary_output_file is required")
	}
	if c.Source == "" {
		return errors.New("source is required")
	}
	if c.Source != "json" && c.Source != "telegram" {
		return errors.New("source must be either json or telegram")
	}
	if c.Source == "telegram" {
		if c.Telegram.APIID == 0 {
			return errors.New("telegram.api_id is required")
		}

		if c.Telegram.APIHash == "" {
			return errors.New("telegram.api_hash is required")
		}

		if c.Telegram.SessionFile == "" {
			return errors.New("telegram.session_file is required")
		}

		if len(c.Telegram.Channels) == 0 {
			return errors.New("telegram.channels is required")
		}

		for _, channel := range c.Telegram.Channels {
			if channel.Username == "" {
				return errors.New("telegram channel username is required")
			}
		}
	}

	return nil
}
