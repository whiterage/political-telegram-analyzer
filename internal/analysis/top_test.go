package analysis

import (
	"reflect"
	"testing"

	"sofiasoft/internal/domain"
)

func TestTopPosts(t *testing.T) {
	t.Parallel()

	posts := []domain.AnalyzedPost{
		{Post: domain.Post{ID: 1}},
		{Post: domain.Post{ID: 2}},
		{Post: domain.Post{ID: 3}},
	}

	got := TopPosts(posts, 2)
	want := []domain.AnalyzedPost{
		{Post: domain.Post{ID: 1}},
		{Post: domain.Post{ID: 2}},
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("TopPosts() = %#v, want %#v", got, want)
	}
}

func TestTopPostsLimitGreaterThanLength(t *testing.T) {
	t.Parallel()

	posts := []domain.AnalyzedPost{
		{Post: domain.Post{ID: 1}},
		{Post: domain.Post{ID: 2}},
	}

	got := TopPosts(posts, 10)
	if !reflect.DeepEqual(got, posts) {
		t.Fatalf("TopPosts() = %#v, want %#v", got, posts)
	}
}

func TestTopPostsNonPositiveLimit(t *testing.T) {
	t.Parallel()

	posts := []domain.AnalyzedPost{
		{Post: domain.Post{ID: 1}},
	}

	got := TopPosts(posts, 0)
	if got != nil {
		t.Fatalf("TopPosts() = %#v, want nil", got)
	}
}
