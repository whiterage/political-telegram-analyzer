package analysis

import "sofiasoft/internal/domain"

func TopPostsByChannel(posts []domain.AnalyzedPost, limit int) []domain.AnalyzedPost {
	if limit <= 0 {
		return []domain.AnalyzedPost{}
	}

	grouped := make(map[string][]domain.AnalyzedPost)

	for _, post := range posts {
		channelName := post.Post.ChannelName
		grouped[channelName] = append(grouped[channelName], post)
	}

	result := make([]domain.AnalyzedPost, 0, len(posts))

	for _, channelPosts := range grouped {
		SortByERR(channelPosts)

		topPosts := TopPosts(channelPosts, limit)

		result = append(result, topPosts...)
	}

	SortByERR(result)

	return result
}
