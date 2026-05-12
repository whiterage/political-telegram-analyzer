package validation

import (
	"fmt"

	"sofiasoft/internal/domain"
)

func ValidatePosts(posts []domain.Post) error {
	for _, post := range posts {
		if post.ID <= 0 {
			return fmt.Errorf("post %d: id must be greater than zero", post.ID)
		}

		if post.ChannelName == "" {
			return fmt.Errorf("post %d: channel name is required", post.ID)
		}

		if post.ChannelUsername == "" {
			return fmt.Errorf("post %d: channel username is required", post.ID)
		}

		if post.ChannelCategory == "" {
			return fmt.Errorf("post %d: channel category is required", post.ID)
		}

		if post.ChannelActorType == "" {
			return fmt.Errorf("post %d: channel actor type is required", post.ID)
		}

		if post.Text == "" {
			return fmt.Errorf("post %d: text is required", post.ID)
		}

		if post.PublishedAt == "" {
			return fmt.Errorf("post %d: published_at is required", post.ID)
		}

		if post.Views < 0 {
			return fmt.Errorf("post %d: views must be greater than or equal to zero", post.ID)
		}

		if post.Forwards < 0 {
			return fmt.Errorf("post %d: forwards must be greater than or equal to zero", post.ID)
		}

		if post.CommentsCount < 0 {
			return fmt.Errorf("post %d: comments_count must be greater than or equal to zero", post.ID)
		}

		for _, reaction := range post.Reactions {
			if reaction.Count < 0 {
				return fmt.Errorf("post %d: reaction count must be greater than or equal to zero", post.ID)
			}
		}

		if post.HasMedia && post.MediaType == "" {
			return fmt.Errorf("post %d: media_type is required when has_media is true", post.ID)
		}
	}

	return nil
}
