package analysis

import (
	"reflect"
	"testing"

	"sofiasoft/internal/domain"
)

func TestTopPostsByChannel(t *testing.T) {
	t.Parallel()

	posts := []domain.AnalyzedPost{
		{Post: domain.Post{ID: 200, ChannelName: "beta"}, ERR: 10},
		{Post: domain.Post{ID: 102, ChannelName: "alpha"}, ERR: 10},
		{Post: domain.Post{ID: 101, ChannelName: "alpha"}, ERR: 10},
		{Post: domain.Post{ID: 201, ChannelName: "beta"}, ERR: 20},
		{Post: domain.Post{ID: 301, ChannelName: "gamma"}, ERR: 10},
		{Post: domain.Post{ID: 302, ChannelName: "gamma"}, ERR: 3},
	}

	got := TopPostsByChannel(posts, 2)
	want := []domain.AnalyzedPost{
		{Post: domain.Post{ID: 201, ChannelName: "beta"}, ERR: 20},
		{Post: domain.Post{ID: 101, ChannelName: "alpha"}, ERR: 10},
		{Post: domain.Post{ID: 102, ChannelName: "alpha"}, ERR: 10},
		{Post: domain.Post{ID: 200, ChannelName: "beta"}, ERR: 10},
		{Post: domain.Post{ID: 301, ChannelName: "gamma"}, ERR: 10},
		{Post: domain.Post{ID: 302, ChannelName: "gamma"}, ERR: 3},
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("TopPostsByChannel() = %#v, want %#v", got, want)
	}
}

func TestTopPostsByChannelNonPositiveLimit(t *testing.T) {
	t.Parallel()

	got := TopPostsByChannel(nil, 0)
	if len(got) != 0 {
		t.Fatalf("TopPostsByChannel() len = %d, want 0", len(got))
	}
}
