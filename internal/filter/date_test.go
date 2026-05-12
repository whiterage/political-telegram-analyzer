package filter

import (
	"testing"

	"sofiasoft/internal/domain"
)

func TestByDateRange(t *testing.T) {
	t.Parallel()

	posts := []domain.Post{
		{
			ID:          1,
			PublishedAt: "2026-05-10T12:00:00+03:00",
		},
		{
			ID:          2,
			PublishedAt: "2026-04-30T23:59:59+03:00",
		},
		{
			ID:          3,
			PublishedAt: "2026-06-01T00:00:00+03:00",
		},
	}

	got, err := ByDateRange(
		posts,
		"2026-05-01T00:00:00+03:00",
		"2026-05-31T23:59:59+03:00",
	)
	if err != nil {
		t.Fatalf("ByDateRange() error = %v", err)
	}

	if len(got) != 1 {
		t.Fatalf("ByDateRange() len = %d, want 1", len(got))
	}

	if got[0].ID != 1 {
		t.Fatalf("ByDateRange() first ID = %d, want 1", got[0].ID)
	}
}

func TestByDateRangeInvalidFrom(t *testing.T) {
	t.Parallel()

	_, err := ByDateRange(nil, "invalid", "2026-05-31T23:59:59+03:00")
	if err == nil {
		t.Fatal("ByDateRange() error = nil, want non-nil")
	}
}

func TestByDateRangeInvalidTo(t *testing.T) {
	t.Parallel()

	_, err := ByDateRange(nil, "2026-05-01T00:00:00+03:00", "invalid")
	if err == nil {
		t.Fatal("ByDateRange() error = nil, want non-nil")
	}
}

func TestByDateRangeInvalidPostDate(t *testing.T) {
	t.Parallel()

	posts := []domain.Post{
		{
			ID:          1,
			PublishedAt: "invalid",
		},
	}

	_, err := ByDateRange(
		posts,
		"2026-05-01T00:00:00+03:00",
		"2026-05-31T23:59:59+03:00",
	)
	if err == nil {
		t.Fatal("ByDateRange() error = nil, want non-nil")
	}
}

func TestByDateRangeMissingPostDate(t *testing.T) {
	t.Parallel()

	posts := []domain.Post{
		{
			ID: 1,
		},
	}

	_, err := ByDateRange(
		posts,
		"2026-05-01T00:00:00+03:00",
		"2026-05-31T23:59:59+03:00",
	)
	if err == nil {
		t.Fatal("ByDateRange() error = nil, want non-nil")
	}
}
