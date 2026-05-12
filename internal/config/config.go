package config

import (
	"errors"
	"fmt"
	"time"
)

type Config struct {
	InputFile                     string          `yaml:"input_file"`
	OutputFile                    string          `yaml:"output_file"`
	SummaryOutputFile             string          `yaml:"summary_output_file"`
	TopLimit                      int             `yaml:"top_limit"`
	ChannelSummaryOutputFile      string          `yaml:"channel_summary_output_file"`
	ChannelGroupSummaryOutputFile string          `yaml:"channel_group_summary_output_file"`
	DateFrom                      string          `yaml:"date_from"`
	DateTo                        string          `yaml:"date_to"`
	Source                        string          `yaml:"source"`
	Channels                      []ChannelConfig `yaml:"channels"`
	Telegram                      TelegramConfig  `yaml:"telegram"`
}

type TelegramConfig struct {
	APIID       int    `yaml:"api_id"`
	APIHash     string `yaml:"api_hash"`
	SessionFile string `yaml:"session_file"`
}

type ChannelConfig struct {
	Username  string `yaml:"username"`
	Name      string `yaml:"name"`
	Category  string `yaml:"category"`
	ActorType string `yaml:"actor_type"`
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
	if c.ChannelGroupSummaryOutputFile == "" {
		return errors.New("channel_group_summary_output_file is required")
	}
	if c.DateFrom == "" {
		return errors.New("date_from is required")
	}
	if c.DateTo == "" {
		return errors.New("date_to is required")
	}
	if c.Source == "" {
		return errors.New("source is required")
	}
	if c.Source != "json" && c.Source != "telegram" {
		return errors.New("source must be either json or telegram")
	}

	from, err := time.Parse(time.RFC3339, c.DateFrom)
	if err != nil {
		return fmt.Errorf("date_from must be RFC3339: %w", err)
	}

	to, err := time.Parse(time.RFC3339, c.DateTo)
	if err != nil {
		return fmt.Errorf("date_to must be RFC3339: %w", err)
	}

	if !from.Before(to) {
		return errors.New("date_from must be before date_to")
	}

	if len(c.Channels) == 0 {
		return errors.New("channels are required")
	}

	for _, channel := range c.Channels {
		if channel.Username == "" {
			return errors.New("channel username is required")
		}

		if channel.Name == "" {
			return errors.New("channel name is required")
		}

		if channel.Category == "" {
			return errors.New("channel category is required")
		}

		if channel.ActorType == "" {
			return errors.New("channel actor_type is required")
		}
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
	}

	return nil
}
