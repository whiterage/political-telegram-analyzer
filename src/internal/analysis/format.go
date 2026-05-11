package analysis

func DetectFormat(text string) string {
	lenT := len([]rune(text))

	if lenT <= 300 {
		return "short"
	}

	if lenT <= 1000 {
		return "medium"
	}

	return "long"
}
