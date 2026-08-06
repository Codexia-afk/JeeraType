package storage

import (
	"bufio"
	"encoding/json"
	"fmt"
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
	Profile      string    `json:"profile,omitempty"`
}

// GetHistoryFilePath resolves standard OS user config directory ~/.config/jeeratype/history.jsonl
func GetHistoryFilePath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err == nil {
		appDir := filepath.Join(configDir, "jeeratype")
		if err := os.MkdirAll(appDir, 0755); err == nil {
			return filepath.Join(appDir, "history.jsonl"), nil
		}
	}

	if home := os.Getenv("HOME"); home != "" {
		appDir := filepath.Join(home, ".config", "jeeratype")
		if err := os.MkdirAll(appDir, 0755); err == nil {
			return filepath.Join(appDir, "history.jsonl"), nil
		}
	}

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
			if rec.Profile == "" {
				rec.Profile = "default"
			}
			records = append(records, rec)
		}
	}

	return records, scanner.Err()
}

// GetPBKey builds a unique key string for scoping Personal Bests per mode+duration+punctuation+numbers combo.
func GetPBKey(mode string, durationSec int, punctuation bool, numbers bool) string {
	puncStr := "no_punc"
	if punctuation {
		puncStr = "punc"
	}
	numStr := "no_num"
	if numbers {
		numStr = "num"
	}
	return fmt.Sprintf("%s_%ds_%s_%s", mode, durationSec, puncStr, numStr)
}

// CheckAndUpdatePBInMap checks if a WPM beats previous PB for a given modeKey.
func CheckAndUpdatePBInMap(pbMap map[string]float64, modeKey string, wpm float64) (float64, bool) {
	prevPB := pbMap[modeKey]
	if wpm > prevPB && wpm > 0 {
		pbMap[modeKey] = wpm
		return prevPB, true
	}
	return prevPB, false
}
