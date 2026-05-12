package summary

import (
	"reflect"
	"testing"

	"sofiasoft/internal/domain"
)

func TestBuildChannelGroupSummary(t *testing.T) {
	t.Parallel()

	posts := []domain.AnalyzedPost{
		{
			Post: domain.Post{
				ChannelCategory:  "news",
				ChannelActorType: "media",
				Views:            1000,
			},
			TotalReactions: 100,
			ERR:            10,
		},
		{
			Post: domain.Post{
				ChannelCategory:  "news",
				ChannelActorType: "media",
				Views:            3000,
			},
			TotalReactions: 300,
			ERR:            20,
		},
		{
			Post: domain.Post{
				ChannelCategory:  "personal",
				ChannelActorType: "political_actor",
				Views:            2000,
			},
			TotalReactions: 250,
			ERR:            12.5,
		},
	}

	got := BuildChannelGroupSummary(posts)
	want := []ChannelGroupSummary{
		{
			GroupType:        "channel_category",
			Name:             "news",
			Count:            2,
			AverageViews:     2000,
			AverageReactions: 200,
			AverageERR:       15,
		},
		{
			GroupType:        "channel_category",
			Name:             "personal",
			Count:            1,
			AverageViews:     2000,
			AverageReactions: 250,
			AverageERR:       12.5,
		},
		{
			GroupType:        "channel_actor_type",
			Name:             "media",
			Count:            2,
			AverageViews:     2000,
			AverageReactions: 200,
			AverageERR:       15,
		},
		{
			GroupType:        "channel_actor_type",
			Name:             "political_actor",
			Count:            1,
			AverageViews:     2000,
			AverageReactions: 250,
			AverageERR:       12.5,
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("BuildChannelGroupSummary() = %#v, want %#v", got, want)
	}
}
