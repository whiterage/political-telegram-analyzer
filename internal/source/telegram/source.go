package telegram

import (
	"context"
	"fmt"

	"sofiasoft/internal/config"
	"sofiasoft/internal/domain"

	"github.com/gotd/td/session"
	tdtelegram "github.com/gotd/td/telegram"
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

	ctx := context.Background()
	storage := &session.FileStorage{
		Path: s.telegram.SessionFile,
	}
	client := tdtelegram.NewClient(s.telegram.APIID, s.telegram.APIHash, tdtelegram.Options{
		SessionStorage: storage,
	})
	return nil, client.Run(ctx, func(ctx context.Context) error {
		if err := client.Auth().IfNecessary(ctx, newAuthFlow()); err != nil {
			return fmt.Errorf("auth: %w", err)
		}
		self, err := client.Self(ctx)
		if err != nil {
			return fmt.Errorf("get self: %w", err)
		}
		_ = s.channels
		return fmt.Errorf(
			"telegram auth ok: user_id=%d username=%s, history collection is not implemented yet",
			self.ID,
			self.Username,
		)
	})

}
