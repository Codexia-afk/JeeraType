package stats

import (
	"testing"
	"time"
)

func TestTrackerMath(t *testing.T) {
	text := "The quick brown fox jumps over the lazy dog and types very fast"
	tracker := NewTracker(text, 15, 0)

	if tracker.RemainingSeconds() != 15 {
		t.Errorf("expected 15 remaining seconds, got %d", tracker.RemainingSeconds())
	}

	tracker.Start()

	// Type "The " correctly
	for _, r := range "The " {
		tracker.RecordRune(r)
	}

	if tracker.CursorIdx != 4 {
		t.Errorf("expected cursor index 4, got %d", tracker.CursorIdx)
	}

	if tracker.CountCorrectChars() != 4 {
		t.Errorf("expected 4 correct chars, got %d", tracker.CountCorrectChars())
	}

	if tracker.CalculateAccuracy() != 100.0 {
		t.Errorf("expected 100%% accuracy, got %.2f", tracker.CalculateAccuracy())
	}

	// Type incorrect char 'x' instead of 'q'
	tracker.RecordRune('x')
	if tracker.CharStates[4] != StatusIncorrect {
		t.Errorf("expected char index 4 to be StatusIncorrect")
	}

	// Backspace
	tracker.Backspace()
	if tracker.CursorIdx != 4 {
		t.Errorf("expected cursor index 4 after backspace, got %d", tracker.CursorIdx)
	}

	// Fixed start and end time for WPM testing (exactly 60 seconds)
	now := time.Now()
	tracker.StartTime = now.Add(-60 * time.Second)
	tracker.EndTime = now
	tracker.IsFinished = true

	// Set 50 correct characters within slice bounds
	tracker.CursorIdx = 50
	tracker.CorrectKeystrokes = 50
	tracker.TotalKeystrokes = 50
	for i := 0; i < 50; i++ {
		tracker.CharStates[i] = StatusCorrect
	}

	// (50 / 5) / 1 min = 10 WPM
	wpm := tracker.CalculateWPM()
	if wpm != 10.0 {
		t.Errorf("expected 10 WPM, got %f", wpm)
	}
}
