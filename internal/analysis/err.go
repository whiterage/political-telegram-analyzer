package analysis

func CalculateERR(reactions, comments, forwards, views int) float64 {
	if views == 0 {
		return 0
	}

	rate := float64(reactions+comments+forwards) / float64(views) * 100

	return rate
}
