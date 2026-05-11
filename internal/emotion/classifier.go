package emotion

import "sofiasoft/internal/domain"

type Classifier interface {
	Classify(text string) domain.EmotionResult
}
