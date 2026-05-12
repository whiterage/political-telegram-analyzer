package summary

import (
	"sort"

	"sofiasoft/internal/domain"
)

type ChannelSummary struct {
	ChannelName      string
	Count            int
	AverageViews     float64
	AverageReactions float64
	AverageERR       float64
}

type channelAccumulator struct {
	Count          int
	TotalViews     int
	TotalReactions int
	TotalERR       float64
}

func BuildChannelSummary(posts []domain.AnalyzedPost) []ChannelSummary {
	stats := make(map[string]channelAccumulator)

	for _, post := range posts {
		channelName := post.Post.ChannelName

		current := stats[channelName]
		current.Count++
		current.TotalViews += post.Post.Views
		current.TotalReactions += post.TotalReactions
		current.TotalERR += post.ERR

		stats[channelName] = current
	}

	result := make([]ChannelSummary, 0, len(stats))

	for channelName, value := range stats {
		result = append(result, ChannelSummary{
			ChannelName:      channelName,
			Count:            value.Count,
			AverageViews:     float64(value.TotalViews) / float64(value.Count),
			AverageReactions: float64(value.TotalReactions) / float64(value.Count),
			AverageERR:       value.TotalERR / float64(value.Count),
		})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].AverageERR > result[j].AverageERR
	})

	return result
}
