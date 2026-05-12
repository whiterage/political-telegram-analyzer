package analysis

import (
	"math"
	"testing"
)

func TestCalculateERR(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		reactions int
		comments  int
		forwards  int
		views     int
		want      float64
	}{
		{
			name:      "calculates engagement rate",
			reactions: 40,
			comments:  10,
			forwards:  5,
			views:     200,
			want:      27.5,
		},
		{
			name:      "returns zero when views are zero",
			reactions: 10,
			comments:  5,
			forwards:  2,
			views:     0,
			want:      0,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := CalculateERR(tt.reactions, tt.comments, tt.forwards, tt.views)
			if math.Abs(got-tt.want) > 1e-9 {
				t.Fatalf("CalculateERR() = %v, want %v", got, tt.want)
			}
		})
	}
}
