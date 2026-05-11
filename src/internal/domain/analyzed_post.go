package domain

type AnalyzedPost struct {
	Post           Post
	TotalReactions int
	ERR            float64
	TextLength     int
	FormatType     string
	Emotion        EmotionResult
}
