package validation

import (
	"testing"

	"sofiasoft/internal/domain"
)

func TestValidatePosts(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		posts   []domain.Post
		wantErr bool
	}{
		{
			name: "valid posts pass",
			posts: []domain.Post{
				{
					ID:               1,
					ChannelName:      "Channel",
					ChannelUsername:  "channel",
					ChannelCategory:  "news",
					ChannelActorType: "media",
					Text:             "Post text",
					PublishedAt:      "2026-05-01T10:00:00+03:00",
					Views:            10,
					Forwards:         1,
					CommentsCount:    2,
					Reactions: []domain.Reaction{
						{Emoji: "🔥", Count: 3},
					},
				},
			},
		},
		{
			name: "missing text fails",
			posts: []domain.Post{
				{
					ID:               2,
					ChannelName:      "Channel",
					ChannelUsername:  "channel",
					ChannelCategory:  "news",
					ChannelActorType: "media",
					PublishedAt:      "2026-05-01T10:00:00+03:00",
				},
			},
			wantErr: true,
		},
		{
			name: "negative reaction count fails",
			posts: []domain.Post{
				{
					ID:               3,
					ChannelName:      "Channel",
					ChannelUsername:  "channel",
					ChannelCategory:  "news",
					ChannelActorType: "media",
					Text:             "Post text",
					PublishedAt:      "2026-05-01T10:00:00+03:00",
					Reactions: []domain.Reaction{
						{Emoji: "🔥", Count: -1},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "missing media type fails when has media",
			posts: []domain.Post{
				{
					ID:               4,
					ChannelName:      "Channel",
					ChannelUsername:  "channel",
					ChannelCategory:  "news",
					ChannelActorType: "media",
					Text:             "Post text",
					PublishedAt:      "2026-05-01T10:00:00+03:00",
					HasMedia:         true,
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := ValidatePosts(tt.posts)
			if tt.wantErr && err == nil {
				t.Fatal("ValidatePosts() error = nil, want non-nil")
			}

			if !tt.wantErr && err != nil {
				t.Fatalf("ValidatePosts() error = %v, want nil", err)
			}
		})
	}
}
