package filter

import "sofiasoft/internal/domain"

func ByThresholds(posts []domain.Post, minViews, minTotalReactions int) []domain.Post {
	filtered := make([]domain.Post, 0, len(posts))

	for _, post := range posts {
		if post.Views < minViews {
			continue
		}

		if totalReactions(post.Reactions) < minTotalReactions {
			continue
		}

		filtered = append(filtered, post)
	}

	return filtered
}

func totalReactions(reactions []domain.Reaction) int {
	total := 0

	for _, reaction := range reactions {
		total += reaction.Count
	}

	return total
}
