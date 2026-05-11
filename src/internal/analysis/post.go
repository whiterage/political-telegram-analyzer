package analysis

import (
	"sofiasoft/src/internal/domain"
	"sofiasoft/src/internal/emotion"
)

func AnalyzePost(post domain.Post, classifier emotion.Classifier) domain.AnalyzedPost {
	totalReactions := TotalReactions(post.Reactions)
	err := CalculateERR(
		totalReactions,
		post.CommentsCount,
		post.Forwards,
		post.Views,
	)
	textLength := len([]rune(post.Text))
	formatType := DetectFormat(post.Text)
	emotionResult := classifier.Classify(post.Text)

	return domain.AnalyzedPost{
		Post:           post,
		TotalReactions: totalReactions,
		ERR:            err,
		TextLength:     textLength,
		FormatType:     formatType,
		Emotion:        emotionResult,
	}
}

func AnalyzePosts(posts []domain.Post, classifier emotion.Classifier) []domain.AnalyzedPost {
	analyzedPosts := make([]domain.AnalyzedPost, len(posts))

	for i, post := range posts {
		analyzedPosts[i] = AnalyzePost(post, classifier)
	}

	return analyzedPosts
}
