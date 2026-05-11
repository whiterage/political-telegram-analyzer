package rulebased

import (
	"strings"

	"sofiasoft/internal/domain"
)

type Classifier struct{}

func NewClassifier() *Classifier {
	return &Classifier{}
}

func (c *Classifier) Classify(text string) domain.EmotionResult {
	lowerText := strings.ToLower(text)

	if containsAny(lowerText, []string{"врут", "ложь", "предательство", "обман", "позор"}) {
		return domain.EmotionResult{
			Frame:      "anger_accusation",
			Emotion:    "anger",
			Confidence: 0.70,
			Method:     "rulebased",
			Reason:     "Текст содержит обвинительную лексику.",
			Markers:    findMarkers(lowerText, []string{"врут", "ложь", "предательство", "обман", "позор"}),
		}
	}

	if containsAny(lowerText, []string{"угроза", "опасность", "удар", "атака", "катастрофа"}) {
		return domain.EmotionResult{
			Frame:      "fear_threat",
			Emotion:    "fear",
			Confidence: 0.70,
			Method:     "rulebased",
			Reason:     "Текст содержит лексику угрозы или опасности.",
			Markers:    findMarkers(lowerText, []string{"угроза", "опасность", "удар", "атака", "катастрофа"}),
		}
	}

	if containsAny(lowerText, []string{"герои", "победа", "слава", "добились", "поддерживает"}) {
		return domain.EmotionResult{
			Frame:      "pride_victory",
			Emotion:    "pride",
			Confidence: 0.70,
			Method:     "rulebased",
			Reason:     "Текст содержит лексику победы, поддержки или героизации.",
			Markers:    findMarkers(lowerText, []string{"герои", "победа", "слава", "добились", "поддерживает"}),
		}
	}

	if containsAny(lowerText, []string{"погибли", "скорбим", "жертвы", "боль", "дети"}) {
		return domain.EmotionResult{
			Frame:      "sadness_empathy",
			Emotion:    "sadness",
			Confidence: 0.70,
			Method:     "rulebased",
			Reason:     "Текст содержит лексику утраты, сочувствия или эмпатии.",
			Markers:    findMarkers(lowerText, []string{"погибли", "скорбим", "жертвы", "боль", "дети"}),
		}
	}

	return domain.EmotionResult{
		Frame:      "neutral",
		Emotion:    "neutral",
		Confidence: 0.50,
		Method:     "rulebased",
		Reason:     "Явных эмоциональных маркеров по текущим правилам не найдено.",
		Markers:    nil,
	}
}

func containsAny(text string, markers []string) bool {
	for _, marker := range markers {
		if strings.Contains(text, marker) {
			return true
		}
	}

	return false
}

func findMarkers(text string, markers []string) []string {
	found := make([]string, 0)

	for _, marker := range markers {
		if strings.Contains(text, marker) {
			found = append(found, marker)
		}
	}

	return found
}
