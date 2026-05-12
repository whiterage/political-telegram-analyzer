package telegram

import (
	"fmt"

	"sofiasoft/internal/config"
	"sofiasoft/internal/domain"
)

type Source struct {
	cfg config.TelegramConfig
}

func New(cfg config.TelegramConfig) *Source {
	return &Source{
		cfg: cfg,
	}
}

func (s *Source) LoadPosts() ([]domain.Post, error) {
	return nil, fmt.Errorf("telegram source is not implemented yet")
}
