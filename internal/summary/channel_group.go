package summary

import (
	"sort"

	"sofiasoft/internal/domain"
)

type ChannelGroupSummary struct {
	GroupType        string
	Name             string
	Count            int
	AverageViews     float64
	AverageReactions float64
	AverageERR       float64
}

func BuildChannelGroupSummary(posts []domain.AnalyzedPost) []ChannelGroupSummary {
	stats := make(map[string]channelAccumulator)

	for _, post := range posts {
		addChannelGroup(stats, "channel_category", post.Post.ChannelCategory, post)
		addChannelGroup(stats, "channel_actor_type", post.Post.ChannelActorType, post)
	}

	result := make([]ChannelGroupSummary, 0, len(stats))

	for key, value := range stats {
		groupType, name := splitChannelGroupKey(key)

		result = append(result, ChannelGroupSummary{
			GroupType:        groupType,
			Name:             name,
			Count:            value.Count,
			AverageViews:     float64(value.TotalViews) / float64(value.Count),
			AverageReactions: float64(value.TotalReactions) / float64(value.Count),
			AverageERR:       value.TotalERR / float64(value.Count),
		})
	}

	groupOrder := map[string]int{
		"channel_category":   0,
		"channel_actor_type": 1,
	}

	sort.Slice(result, func(i, j int) bool {
		if result[i].GroupType != result[j].GroupType {
			return groupOrder[result[i].GroupType] < groupOrder[result[j].GroupType]
		}

		if result[i].AverageERR != result[j].AverageERR {
			return result[i].AverageERR > result[j].AverageERR
		}

		return result[i].Name < result[j].Name
	})

	return result
}

func addChannelGroup(
	stats map[string]channelAccumulator,
	groupType, name string,
	post domain.AnalyzedPost,
) {
	key := groupType + ":" + name

	current := stats[key]
	current.Count++
	current.TotalViews += post.Post.Views
	current.TotalReactions += post.TotalReactions
	current.TotalERR += post.ERR

	stats[key] = current
}

func splitChannelGroupKey(key string) (string, string) {
	for i, char := range key {
		if char == ':' {
			return key[:i], key[i+1:]
		}
	}

	return "unknown", key
}
