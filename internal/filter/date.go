package filter

import (
	"fmt"
	"time"

	"sofiasoft/internal/domain"
)

func ByDateRange(posts []domain.Post, fromRaw, toRaw string) ([]domain.Post, error) {
	from, err := time.Parse(time.RFC3339, fromRaw)
	if err != nil {
		return nil, fmt.Errorf("parse date_from: %w", err)
	}

	to, err := time.Parse(time.RFC3339, toRaw)
	if err != nil {
		return nil, fmt.Errorf("parse date_to: %w", err)
	}

	filtered := make([]domain.Post, 0, len(posts))

	for _, post := range posts {
		if post.PublishedAt == "" {
			return nil, fmt.Errorf("post %d: published_at is required", post.ID)
		}

		publishedAt, err := time.Parse(time.RFC3339, post.PublishedAt)
		if err != nil {
			return nil, fmt.Errorf("post %d: parse published_at: %w", post.ID, err)
		}

		if publishedAt.Before(from) || publishedAt.After(to) {
			continue
		}

		filtered = append(filtered, post)
	}

	return filtered, nil
}
