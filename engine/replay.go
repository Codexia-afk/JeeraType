package engine

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Codexia-afk/JeeraType/storage"
)

type ReplayKeystroke struct {
	TimestampMs int64 `json:"timestamp_ms"`
	Char        rune  `json:"char"`
	IsCorrect   bool  `json:"is_correct"`
}

type ReplayRun struct {
	RunID       string            `json:"run_id"`
	Mode        string            `json:"mode"`
	DurationSec int               `json:"duration_sec"`
	WPM         float64           `json:"wpm"`
	WPMSamples  []float64         `json:"wpm_samples"`
	Keystrokes  []ReplayKeystroke `json:"keystrokes"`
}

// ExportReplay exports last/specified run timeline to JSON file.
func ExportReplay(targetPath string) error {
	records, err := storage.LoadHistory()
	if err != nil || len(records) == 0 {
		return fmt.Errorf("no history runs available to export")
	}

	lastRun := records[len(records)-1]
	replay := ReplayRun{
		RunID:       fmt.Sprintf("run_%d", lastRun.Timestamp.Unix()),
		Mode:        "Paragraphs",
		DurationSec: lastRun.DurationSec,
		WPM:         lastRun.WPM,
		WPMSamples:  []float64{lastRun.WPM},
		Keystrokes:  make([]ReplayKeystroke, 0),
	}

	data, err := json.MarshalIndent(replay, "", "  ")
	if err != nil {
		return err
	}

	if targetPath == "" {
		targetPath = fmt.Sprintf("replay_%d.json", lastRun.Timestamp.Unix())
	}

	return os.WriteFile(targetPath, data, 0644)
}

// LoadReplay loads ReplayRun from JSON file.
func LoadReplay(filePath string) (*ReplayRun, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read replay file: %v", err)
	}

	var replay ReplayRun
	if err := json.Unmarshal(data, &replay); err != nil {
		return nil, fmt.Errorf("failed to parse replay JSON: %v", err)
	}

	return &replay, nil
}
