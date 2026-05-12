package rulebased

import (
	"reflect"
	"testing"
)

func TestClassifierClassify(t *testing.T) {
	t.Parallel()

	classifier := NewClassifier()

	tests := []struct {
		name        string
		text        string
		wantEmotion string
		wantFrame   string
		wantMarkers []string
	}{
		{
			name:        "detects anger accusation",
			text:        "Это ложь и позор.",
			wantEmotion: "anger",
			wantFrame:   "anger_accusation",
			wantMarkers: []string{"ложь", "позор"},
		},
		{
			name:        "detects fear threat",
			text:        "Это реальная угроза и опасность.",
			wantEmotion: "fear",
			wantFrame:   "fear_threat",
			wantMarkers: []string{"угроза", "опасность"},
		},
		{
			name:        "detects pride victory",
			text:        "Это победа, герои добились своего.",
			wantEmotion: "pride",
			wantFrame:   "pride_victory",
			wantMarkers: []string{"герои", "победа", "добились"},
		},
		{
			name:        "detects sadness empathy",
			text:        "Мы скорбим, дети и жертвы вызывают боль.",
			wantEmotion: "sadness",
			wantFrame:   "sadness_empathy",
			wantMarkers: []string{"скорбим", "жертвы", "боль", "дети"},
		},
		{
			name:        "returns neutral when no markers found",
			text:        "Информационное сообщение без выраженной окраски.",
			wantEmotion: "neutral",
			wantFrame:   "neutral",
			wantMarkers: nil,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := classifier.Classify(tt.text)

			if got.Emotion != tt.wantEmotion {
				t.Fatalf("Emotion = %q, want %q", got.Emotion, tt.wantEmotion)
			}

			if got.Frame != tt.wantFrame {
				t.Fatalf("Frame = %q, want %q", got.Frame, tt.wantFrame)
			}

			if got.Method != "rulebased" {
				t.Fatalf("Method = %q, want %q", got.Method, "rulebased")
			}

			if !reflect.DeepEqual(got.Markers, tt.wantMarkers) {
				t.Fatalf("Markers = %#v, want %#v", got.Markers, tt.wantMarkers)
			}
		})
	}
}
