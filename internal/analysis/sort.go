package analysis

import (
	"sofiasoft/internal/domain"
	"sort"
)

func SortByERR(posts []domain.AnalyzedPost) {
	sort.SliceStable(posts, func(i, j int) bool {
		if posts[i].ERR != posts[j].ERR {
			return posts[i].ERR > posts[j].ERR
		}

		if posts[i].Post.ChannelName != posts[j].Post.ChannelName {
			return posts[i].Post.ChannelName < posts[j].Post.ChannelName
		}

		return posts[i].Post.ID < posts[j].Post.ID
	})
}
