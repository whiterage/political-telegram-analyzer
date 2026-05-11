package emotion

import "sofiasoft/src/internal/domain"

type Classifier interface {
	Classify(text string) domain.EmotionResult
}
