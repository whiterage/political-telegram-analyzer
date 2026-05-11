package analysis

import (
	"sofiasoft/src/internal/domain"
	"sort"
)

func SortByERR(posts []domain.AnalyzedPost) {
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].ERR > posts[j].ERR
	})
}
