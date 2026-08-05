package db

import (
	"encoding/json"
	"strings"
	"time"
)

type KeyMetric struct {
	Char       rune
	TotalHits  int
	ErrorCount int
	AvgLatency float64
	Accuracy   float64
}

// SaveTestRun records a test session in test_runs table.
func SaveTestRun(mode string, wpm, rawWPM, accuracy, consistency float64, durationSec int) error {
	if DB == nil {
		return nil
	}
	query := `INSERT INTO test_runs (timestamp, mode, wpm, raw_wpm, accuracy, consistency, duration_sec)
	          VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err := DB.Exec(query, time.Now(), mode, wpm, rawWPM, accuracy, consistency, durationSec)
	return err
}

// SaveFileOffset records or updates file reading cursor offset.
func SaveFileOffset(filePath string, offset int) error {
	if DB == nil || filePath == "" {
		return nil
	}
	query := `
	INSERT INTO file_offsets (file_path, cursor_offset, updated_at)
	VALUES (?, ?, ?)
	ON CONFLICT(file_path) DO UPDATE SET
		cursor_offset = ?,
		updated_at = ?;
	`
	now := time.Now()
	_, err := DB.Exec(query, filePath, offset, now, offset, now)
	return err
}

// GetFileOffset retrieves saved reading offset for a file.
func GetFileOffset(filePath string) int {
	if DB == nil || filePath == "" {
		return 0
	}
	query := `SELECT cursor_offset FROM file_offsets WHERE file_path = ?`
	var offset int
	err := DB.QueryRow(query, filePath).Scan(&offset)
	if err != nil {
		return 0
	}
	return offset
}

// SavePBRace records personal best WPM and 1-second sample timeline.
func SavePBRace(modeKey string, wpm float64, samples []float64) error {
	if DB == nil || modeKey == "" {
		return nil
	}
	data, err := json.Marshal(samples)
	if err != nil {
		return err
	}

	query := `
	INSERT INTO pb_samples (mode_key, wpm, samples_json)
	VALUES (?, ?, ?)
	ON CONFLICT(mode_key) DO UPDATE SET
		wpm = ?,
		samples_json = ?
	WHERE ? > wpm;
	`
	_, err = DB.Exec(query, modeKey, wpm, string(data), wpm, string(data), wpm)
	return err
}

// GetPBRace retrieves personal best WPM and sample timeline.
func GetPBRace(modeKey string) (float64, []float64) {
	if DB == nil || modeKey == "" {
		return 0, nil
	}
	query := `SELECT wpm, samples_json FROM pb_samples WHERE mode_key = ?`
	var wpm float64
	var jsonStr string
	err := DB.QueryRow(query, modeKey).Scan(&wpm, &jsonStr)
	if err != nil {
		return 0, nil
	}
	var samples []float64
	_ = json.Unmarshal([]byte(jsonStr), &samples)
	return wpm, samples
}

// RecordKeyHit updates hits, errors, and latency for a single key.
func RecordKeyHit(keyChar string, isError bool, latencyMs int64) error {
	if DB == nil || keyChar == "" {
		return nil
	}
	keyChar = strings.ToLower(keyChar)
	errInc := 0
	if isError {
		errInc = 1
	}

	query := `
	INSERT INTO key_stats (key_char, total_hits, total_errors, total_latency_ms)
	VALUES (?, 1, ?, ?)
	ON CONFLICT(key_char) DO UPDATE SET
		total_hits = total_hits + 1,
		total_errors = total_errors + ?,
		total_latency_ms = total_latency_ms + ?;
	`
	_, err := DB.Exec(query, keyChar, errInc, latencyMs, errInc, latencyMs)
	return err
}

// RecordBigramHit updates count and latency for a key-to-key transition.
func RecordBigramHit(bigram string, latencyMs int64) error {
	if DB == nil || len(bigram) < 2 {
		return nil
	}
	bigram = strings.ToLower(bigram)
	query := `
	INSERT INTO bigram_stats (bigram, count, total_latency_ms)
	VALUES (?, 1, ?)
	ON CONFLICT(bigram) DO UPDATE SET
		count = count + 1,
		total_latency_ms = total_latency_ms + ?;
	`
	_, err := DB.Exec(query, bigram, latencyMs, latencyMs)
	return err
}

// GetTopWeakKeys returns the most missed or slowest keys.
func GetTopWeakKeys(limit int) []rune {
	if DB == nil {
		return []rune{'e', 't', 'a', 'o', 'i'}
	}
	query := `
	SELECT key_char, (CAST(total_errors AS REAL) / MAX(total_hits, 1)) + (CAST(total_latency_ms AS REAL) / MAX(total_hits, 1) / 1000.0) AS score
	FROM key_stats
	WHERE total_hits > 3
	ORDER BY score DESC
	LIMIT ?;
	`
	rows, err := DB.Query(query, limit)
	if err != nil {
		return []rune{'e', 't', 'a', 'o', 'i'}
	}
	defer rows.Close()

	var result []rune
	for rows.Next() {
		var kStr string
		var score float64
		if err := rows.Scan(&kStr, &score); err == nil && len(kStr) > 0 {
			result = append(result, []rune(kStr)[0])
		}
	}
	if len(result) == 0 {
		return []rune{'e', 't', 'a', 'o', 'i'}
	}
	return result
}

// GetTopSlowestBigrams returns top slowest key transition pairs.
func GetTopSlowestBigrams(limit int) []string {
	if DB == nil {
		return []string{"th", "he", "in", "er", "an"}
	}
	query := `
	SELECT bigram, (CAST(total_latency_ms AS REAL) / count) AS avg_latency
	FROM bigram_stats
	WHERE count > 2
	ORDER BY avg_latency DESC
	LIMIT ?;
	`
	rows, err := DB.Query(query, limit)
	if err != nil {
		return []string{"th", "he", "in", "er", "an"}
	}
	defer rows.Close()

	var result []string
	for rows.Next() {
		var bg string
		var avg float64
		if err := rows.Scan(&bg, &avg); err == nil && bg != "" {
			result = append(result, bg)
		}
	}
	if len(result) == 0 {
		return []string{"th", "he", "in", "er", "an"}
	}
	return result
}

// GetKeyMetrics returns map of key stats for ASCII Keyboard Heatmap rendering.
func GetKeyMetrics() map[rune]KeyMetric {
	res := make(map[rune]KeyMetric)
	if DB == nil {
		return res
	}
	query := `SELECT key_char, total_hits, total_errors, total_latency_ms FROM key_stats`
	rows, err := DB.Query(query)
	if err != nil {
		return res
	}
	defer rows.Close()

	for rows.Next() {
		var kStr string
		var hits, errors, latency int
		if err := rows.Scan(&kStr, &hits, &errors, &latency); err == nil && len(kStr) > 0 {
			r := []rune(kStr)[0]
			avgLat := 0.0
			acc := 100.0
			if hits > 0 {
				avgLat = float64(latency) / float64(hits)
				acc = ((float64(hits) - float64(errors)) / float64(hits)) * 100.0
			}
			res[r] = KeyMetric{
				Char:       r,
				TotalHits:  hits,
				ErrorCount: errors,
				AvgLatency: avgLat,
				Accuracy:   acc,
			}
		}
	}
	return res
}
