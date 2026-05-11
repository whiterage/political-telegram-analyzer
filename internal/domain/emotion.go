package domain

type EmotionResult struct {
	Emotion    string
	Frame      string
	Confidence float64
	Method     string
	Reason     string
	Markers    []string
}
