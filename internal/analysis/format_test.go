package analysis

import (
	"strings"
	"testing"
)

func TestDetectFormat(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		text string
		want string
	}{
		{
			name: "short on upper boundary",
			text: strings.Repeat("я", 300),
			want: "short",
		},
		{
			name: "medium after short boundary",
			text: strings.Repeat("я", 301),
			want: "medium",
		},
		{
			name: "medium on upper boundary",
			text: strings.Repeat("я", 1000),
			want: "medium",
		},
		{
			name: "long after medium boundary",
			text: strings.Repeat("я", 1001),
			want: "long",
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := DetectFormat(tt.text)
			if got != tt.want {
				t.Fatalf("DetectFormat() = %q, want %q", got, tt.want)
			}
		})
	}
}
