package generator

import (
	"strings"
	"testing"
)

func TestGenerateText(t *testing.T) {
	durations := []int{15, 30, 45, 60, 120}
	for _, d := range durations {
		text := GenerateText(d)
		if len(text) == 0 {
			t.Errorf("expected non-empty text for duration %ds", d)
		}
		words := strings.Fields(text)
		if len(words) < 30 {
			t.Errorf("expected at least 30 words for duration %ds, got %d", d, len(words))
		}
	}
}
