package telegram

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/gotd/td/session"
	tdtelegram "github.com/gotd/td/telegram"
	"github.com/gotd/td/tg"

	"sofiasoft/internal/config"
	"sofiasoft/internal/domain"
)

type Source struct {
	cfg      config.TelegramConfig
	channels []config.ChannelConfig
	logger   *slog.Logger
}

func New(cfg config.TelegramConfig, channels []config.ChannelConfig, logger *slog.Logger) *Source {
	if logger == nil {
		logger = slog.Default()
	}

	return &Source{
		cfg:      cfg,
		channels: channels,
		logger:   logger,
	}
}

func (s *Source) LoadPosts() ([]domain.Post, error) {
	ctx := context.Background()

	storage := &session.FileStorage{
		Path: s.cfg.SessionFile,
	}

	client := tdtelegram.NewClient(s.cfg.APIID, s.cfg.APIHash, tdtelegram.Options{
		SessionStorage: storage,
	})

	err := client.Run(ctx, func(ctx context.Context) error {
		if err := client.Auth().IfNecessary(ctx, newAuthFlow()); err != nil {
			return fmt.Errorf("auth: %w", err)
		}

		api := tg.NewClient(client)

		if len(s.channels) == 0 {
			return fmt.Errorf("telegram channels are required")
		}

		if err := s.fetchRecentMessages(ctx, api, s.channels[0]); err != nil {
			return err
		}

		return fmt.Errorf("telegram message fetch ok, mapping to domain.Post is not implemented yet")
	})
	if err != nil {
		return nil, err
	}

	return nil, fmt.Errorf("telegram source returned no posts")
}

func (s *Source) fetchRecentMessages(ctx context.Context, api *tg.Client, channel config.ChannelConfig) error {
	resolved, err := api.ContactsResolveUsername(ctx, &tg.ContactsResolveUsernameRequest{
		Username: channel.Username,
	})
	if err != nil {
		return fmt.Errorf("resolve username %s: %w", channel.Username, err)
	}

	if len(resolved.Chats) == 0 {
		return fmt.Errorf("resolve username %s: no chats found", channel.Username)
	}

	inputPeer, err := inputPeerFromChat(resolved.Chats[0])
	if err != nil {
		return fmt.Errorf("build input peer for %s: %w", channel.Username, err)
	}

	history, err := api.MessagesGetHistory(ctx, &tg.MessagesGetHistoryRequest{
		Peer:  inputPeer,
		Limit: s.cfg.FetchLimit,
	})
	if err != nil {
		return fmt.Errorf("get history for %s: %w", channel.Username, err)
	}

	messages, ok := history.(*tg.MessagesChannelMessages)
	if !ok {
		return fmt.Errorf("unexpected history type %T", history)
	}

	for _, msgClass := range messages.Messages {
		msg, ok := msgClass.(*tg.Message)
		if !ok {
			continue
		}

		s.logger.Info(
			"telegram message fetched",
			"channel_username",
			channel.Username,
			"id",
			msg.ID,
			"date",
			time.Unix(int64(msg.Date), 0).Format(time.RFC3339),
			"text",
			msg.Message,
			"views",
			msg.Views,
			"forwards",
			msg.Forwards,
		)
	}

	return nil
}

func inputPeerFromChat(chat tg.ChatClass) (tg.InputPeerClass, error) {
	switch c := chat.(type) {
	case *tg.Channel:
		return &tg.InputPeerChannel{
			ChannelID:  c.ID,
			AccessHash: c.AccessHash,
		}, nil

	case *tg.Chat:
		return &tg.InputPeerChat{
			ChatID: c.ID,
		}, nil

	default:
		return nil, fmt.Errorf("unsupported chat type %T", chat)
	}
}
