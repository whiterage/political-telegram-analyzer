package telegram

import (
	"fmt"

	"sofiasoft/internal/config"
	"sofiasoft/internal/domain"
)

type Source struct {
	telegram config.TelegramConfig
	channels []config.ChannelConfig
}

func New(telegram config.TelegramConfig, channels []config.ChannelConfig) *Source {
	return &Source{
		telegram: telegram,
		channels: channels,
	}
}

func (s *Source) LoadPosts() ([]domain.Post, error) {
	return nil, fmt.Errorf("telegram source is not implemented yet")
}
