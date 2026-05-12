package reactionanalysis

import (
	"reflect"
	"testing"

	"sofiasoft/internal/domain"
)

func TestAnalyzeReactions(t *testing.T) {
	t.Parallel()

	reactions := []domain.Reaction{
		{Emoji: "🔥", Count: 2},
		{Emoji: "😂", Count: 3},
		{Emoji: "🔥", Count: 4},
	}

	got := AnalyzeReactions(reactions)

	if got.DominantEmoji != "🔥" {
		t.Fatalf("DominantEmoji = %q, want %q", got.DominantEmoji, "🔥")
	}

	if got.DominantEmotion != "approval" {
		t.Fatalf("DominantEmotion = %q, want %q", got.DominantEmotion, "approval")
	}

	if got.TotalReactions != 9 {
		t.Fatalf("TotalReactions = %d, want 9", got.TotalReactions)
	}

	wantDistribution := map[string]int{
		"🔥": 6,
		"😂": 3,
	}

	if !reflect.DeepEqual(got.Distribution, wantDistribution) {
		t.Fatalf("Distribution = %#v, want %#v", got.Distribution, wantDistribution)
	}
}

func TestAnalyzeReactionsEmpty(t *testing.T) {
	t.Parallel()

	got := AnalyzeReactions(nil)

	if got.DominantEmoji != "" {
		t.Fatalf("DominantEmoji = %q, want empty string", got.DominantEmoji)
	}

	if got.DominantEmotion != "unknown" {
		t.Fatalf("DominantEmotion = %q, want %q", got.DominantEmotion, "unknown")
	}

	if got.TotalReactions != 0 {
		t.Fatalf("TotalReactions = %d, want 0", got.TotalReactions)
	}

	if len(got.Distribution) != 0 {
		t.Fatalf("Distribution len = %d, want 0", len(got.Distribution))
	}
}
