package domain

type ReactionEmotionResult struct {
	DominantEmoji   string
	DominantEmotion string
	TotalReactions  int
	Distribution    map[string]int
}
