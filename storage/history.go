package storage

import (
	"bufio"
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

// HistoryRecord represents a saved test result entry.
type HistoryRecord struct {
	Timestamp    time.Time `json:"timestamp"`
	DurationSec  int       `json:"duration_sec"`
	WPM          float64   `json:"wpm"`
	RawWPM       float64   `json:"raw_wpm"`
	Accuracy     float64   `json:"accuracy"`
	Consistency  float64   `json:"consistency"`
	TotalChars   int       `json:"total_chars"`
	CorrectChars int       `json:"correct_chars"`
	ErrorCount   int       `json:"error_count"`
}

// GetHistoryFilePath resolves standard OS user config directory ~/.config/jeeratype/history.jsonl
// with fallback to working directory / temp dir if permissions are restricted.
func GetHistoryFilePath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err == nil {
		appDir := filepath.Join(configDir, "jeeratype")
		if err := os.MkdirAll(appDir, 0755); err == nil {
			return filepath.Join(appDir, "history.jsonl"), nil
		}
	}

	// Fallback 1: ~/.config/jeeratype
	if home := os.Getenv("HOME"); home != "" {
		appDir := filepath.Join(home, ".config", "jeeratype")
		if err := os.MkdirAll(appDir, 0755); err == nil {
			return filepath.Join(appDir, "history.jsonl"), nil
		}
	}

	// Fallback 2: Local temp dir
	appDir := filepath.Join(os.TempDir(), "jeeratype")
	_ = os.MkdirAll(appDir, 0755)
	return filepath.Join(appDir, "history.jsonl"), nil
}

// SaveRecord appends a single completed test result to history.jsonl
func SaveRecord(rec HistoryRecord) error {
	filePath, err := GetHistoryFilePath()
	if err != nil {
		return err
	}

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := json.Marshal(rec)
	if err != nil {
		return err
	}

	_, err = file.Write(append(data, '\n'))
	return err
}

// LoadHistory reads all past records from history.jsonl
func LoadHistory() ([]HistoryRecord, error) {
	filePath, err := GetHistoryFilePath()
	if err != nil {
		return nil, err
	}

	file, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	defer file.Close()

	var records []HistoryRecord
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 {
			continue
		}
		var rec HistoryRecord
		if err := json.Unmarshal(line, &rec); err == nil {
			records = append(records, rec)
		}
	}

	return records, scanner.Err()
}
