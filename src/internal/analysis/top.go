package analysis

import "sofiasoft/src/internal/domain"

func TopPosts(posts []domain.AnalyzedPost, limit int) []domain.AnalyzedPost {
	if limit <= 0 {
		return nil
	}

	if limit > len(posts) {
		limit = len(posts)
	}

	return posts[:limit]
}
