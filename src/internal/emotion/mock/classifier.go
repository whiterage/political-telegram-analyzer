package mock

import "sofiasoft/src/internal/domain"

type Classifier struct{}

func NewClassifier() *Classifier {
	return &Classifier{}
}

func (c *Classifier) Classify(text string) domain.EmotionResult {
	return domain.EmotionResult{
		Frame:      "anger_accusation",
		Emotion:    "anger",
		Confidence: 0.90,
		Method:     "mock",
		Reason:     "Mock classifier returns fixed emotion result for architecture testing.",
		Markers:    []string{"mock"},
	}
}
