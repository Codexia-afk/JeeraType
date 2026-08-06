package storage

import (
	"os"
	"testing"
	"time"
)

func TestStorageSaveAndLoad(t *testing.T) {
	tmpDir := t.TempDir()
	os.Setenv("HOME", tmpDir)

	rec := HistoryRecord{
		Timestamp:    time.Now().Truncate(time.Second),
		DurationSec:  30,
		WPM:          85.5,
		RawWPM:       90.0,
		Accuracy:     98.2,
		Consistency:  95.0,
		TotalChars:   250,
		CorrectChars: 245,
		ErrorCount:   5,
		Profile:      "default",
	}

	path, err := GetHistoryFilePath()
	if err != nil {
		t.Fatalf("failed to get history path: %v", err)
	}
	if path == "" {
		t.Fatalf("expected non-empty history path")
	}

	err = SaveRecord(rec)
	if err != nil {
		t.Fatalf("failed to save record: %v", err)
	}

	records, err := LoadHistory()
	if err != nil {
		t.Fatalf("failed to load history: %v", err)
	}

	if len(records) == 0 {
		t.Fatalf("expected at least 1 record in history, got 0")
	}
}
