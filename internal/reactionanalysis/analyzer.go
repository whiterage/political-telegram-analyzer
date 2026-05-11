package reactionanalysis

import "sofiasoft/internal/domain"

func AnalyzeReactions(reactions []domain.Reaction) domain.ReactionEmotionResult {
	distribution := make(map[string]int)

	total := 0
	dominantEmoji := ""
	dominantCount := 0

	for _, reaction := range reactions {
		distribution[reaction.Emoji] += reaction.Count
		total += reaction.Count

		if reaction.Count > dominantCount {
			dominantEmoji = reaction.Emoji
			dominantCount = reaction.Count
		}
	}

	return domain.ReactionEmotionResult{
		DominantEmoji:   dominantEmoji,
		DominantEmotion: detectReactionEmotion(dominantEmoji),
		TotalReactions:  total,
		Distribution:    distribution,
	}
}

func detectReactionEmotion(emoji string) string {
	switch emoji {
	case "🔥", "👍", "❤️", "👏":
		return "approval"
	case "😁", "😂":
		return "positive_laughter"
	case "🤡":
		return "ridicule"
	case "😡", "🤬":
		return "anger"
	case "😢", "😭":
		return "sadness"
	case "😱", "🤯":
		return "shock"
	case "💩":
		return "contempt"
	case "🤮":
		return "disgust"
	default:
		return "unknown"
	}
}
