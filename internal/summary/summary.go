package summary

import "sofiasoft/internal/domain"

type EmotionSummary struct {
	Type       string
	Name       string
	Count      int
	AverageERR float64
}

type accumulator struct {
	Count    int
	TotalERR float64
}

func BuildEmotionSummary(posts []domain.AnalyzedPost) []EmotionSummary {
	stats := make(map[string]accumulator)

	for _, post := range posts {
		add(stats, "frame", post.Emotion.Frame, post.ERR)
		add(stats, "reaction_emotion", post.ReactionEmotion.DominantEmotion, post.ERR)
	}

	result := make([]EmotionSummary, 0, len(stats))

	for key, value := range stats {
		summaryType, name := splitKey(key)

		result = append(result, EmotionSummary{
			Type:       summaryType,
			Name:       name,
			Count:      value.Count,
			AverageERR: value.TotalERR / float64(value.Count),
		})
	}

	return result
}

func add(stats map[string]accumulator, summaryType, name string, err float64) {
	key := summaryType + ":" + name

	current := stats[key]
	current.Count++
	current.TotalERR += err

	stats[key] = current
}

func splitKey(key string) (string, string) {
	for i, char := range key {
		if char == ':' {
			return key[:i], key[i+1:]
		}
	}

	return "unknown", key
}
