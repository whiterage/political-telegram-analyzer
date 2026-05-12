package filter

import (
	"reflect"
	"testing"

	"sofiasoft/internal/domain"
)

func TestByThresholds(t *testing.T) {
	t.Parallel()

	posts := []domain.Post{
		{
			ID:    1,
			Views: 1500,
			Reactions: []domain.Reaction{
				{Emoji: "🔥", Count: 5},
				{Emoji: "👍", Count: 7},
			},
		},
		{
			ID:    2,
			Views: 900,
			Reactions: []domain.Reaction{
				{Emoji: "🔥", Count: 20},
			},
		},
		{
			ID:    3,
			Views: 5000,
			Reactions: []domain.Reaction{
				{Emoji: "🔥", Count: 9},
			},
		},
	}

	got := ByThresholds(posts, 1000, 10)
	want := []domain.Post{
		posts[0],
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("ByThresholds() = %#v, want %#v", got, want)
	}
}

func TestByThresholdsInclusiveBoundary(t *testing.T) {
	t.Parallel()

	posts := []domain.Post{
		{
			ID:    1,
			Views: 1000,
			Reactions: []domain.Reaction{
				{Emoji: "🔥", Count: 10},
			},
		},
	}

	got := ByThresholds(posts, 1000, 10)
	want := []domain.Post{
		posts[0],
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("ByThresholds() = %#v, want %#v", got, want)
	}
}
