package analysis

import "sofiasoft/internal/domain"

func TotalReactions(reactions []domain.Reaction) int {
	total := 0

	for _, r := range reactions {
		total += r.Count
	}

	return total
}
