package stats

import (
	"testing"
	"time"

	"github.com/Codexia-afk/JeeraType/storage"
)

func TestTrackerMath(t *testing.T) {
	text := "The quick brown fox jumps over the lazy dog and types very fast"
	tracker := NewTracker(text, 15, 0, false, false, false)

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

func TestPersonalBestScoping(t *testing.T) {
	// Mode + duration + punctuation + numbers combo 1: "Paragraphs_15s_punc_false_num_false"
	key15 := storage.GetPBKey("Paragraphs", 15, false, false)
	key30 := storage.GetPBKey("Paragraphs", 30, false, false)

	if key15 == key30 {
		t.Fatalf("expected different PB storage keys for 15s vs 30s, got identical: %s", key15)
	}

	// Simulate sequence of runs on 15s mode
	pbMap := make(map[string]float64)

	// Run 1: 50 WPM
	run1WPM := 50.0
	prevPB1, isNewPB1 := storage.CheckAndUpdatePBInMap(pbMap, key15, run1WPM)
	if !isNewPB1 || prevPB1 != 0 {
		t.Errorf("Run 1 should be new PB (prev: 0), got isNewPB=%v, prevPB=%.1f", isNewPB1, prevPB1)
	}

	// Run 2: 40 WPM (lower than 50 WPM)
	run2WPM := 40.0
	prevPB2, isNewPB2 := storage.CheckAndUpdatePBInMap(pbMap, key15, run2WPM)
	if isNewPB2 || prevPB2 != 50.0 {
		t.Errorf("Run 2 should NOT be new PB, got isNewPB=%v, prevPB=%.1f", isNewPB2, prevPB2)
	}

	// Run 3: 72 WPM (beats 50 WPM)
	run3WPM := 72.0
	prevPB3, isNewPB3 := storage.CheckAndUpdatePBInMap(pbMap, key15, run3WPM)
	if !isNewPB3 || prevPB3 != 50.0 {
		t.Errorf("Run 3 should be new PB (prev: 50.0), got isNewPB=%v, prevPB=%.1f", isNewPB3, prevPB3)
	}

	// Run 4: 30s mode run with 80 WPM
	run4WPM := 80.0
	prevPB30, isNewPB30 := storage.CheckAndUpdatePBInMap(pbMap, key30, run4WPM)
	if !isNewPB30 || prevPB30 != 0 {
		t.Errorf("30s Run should be independent PB (prev: 0), got isNewPB=%v, prevPB=%.1f", isNewPB30, prevPB30)
	}

	// Verify 15s PB remains unchanged at 72.0
	if pbMap[key15] != 72.0 {
		t.Errorf("expected 15s PB to remain 72.0, got %.1f", pbMap[key15])
	}
	if pbMap[key30] != 80.0 {
		t.Errorf("expected 30s PB to be 80.0, got %.1f", pbMap[key30])
	}
}

func TestStreakCalculationDates(t *testing.T) {
	now := time.Now()

	// 1. Consecutive 3 days (Today, Yesterday, 2 days ago)
	records3Days := []storage.HistoryRecord{
		{Timestamp: now.AddDate(0, 0, -2)},
		{Timestamp: now.AddDate(0, 0, -1)},
		{Timestamp: now},
	}
	streak3 := CalculateStreak(records3Days)
	if streak3 != 3 {
		t.Errorf("expected streak of 3 consecutive days, got %d", streak3)
	}

	// 2. Skipped day (Today, 2 days ago - yesterday skipped)
	recordsSkipped := []storage.HistoryRecord{
		{Timestamp: now.AddDate(0, 0, -2)},
		{Timestamp: now},
	}
	streakSkipped := CalculateStreak(recordsSkipped)
	if streakSkipped != 1 {
		t.Errorf("expected streak of 1 day due to skipped yesterday, got %d", streakSkipped)
	}
}
