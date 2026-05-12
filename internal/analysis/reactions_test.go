package analysis

import (
	"testing"

	"sofiasoft/internal/domain"
)

func TestTotalReactions(t *testing.T) {
	t.Parallel()

	reactions := []domain.Reaction{
		{Emoji: "🔥", Count: 3},
		{Emoji: "❤️", Count: 5},
		{Emoji: "🤡", Count: 2},
	}

	got := TotalReactions(reactions)
	if got != 10 {
		t.Fatalf("TotalReactions() = %d, want 10", got)
	}
}

func TestTotalReactionsEmpty(t *testing.T) {
	t.Parallel()

	got := TotalReactions(nil)
	if got != 0 {
		t.Fatalf("TotalReactions() = %d, want 0", got)
	}
}
