package export

import (
	"encoding/csv"
	"fmt"
	"io"

	"github.com/Codexia-afk/JeeraType/storage"
)

// ExportCSV formats history records into CSV output.
func ExportCSV(records []storage.HistoryRecord, w io.Writer) error {
	writer := csv.NewWriter(w)
	defer writer.Flush()

	// Header row
	header := []string{"timestamp", "mode", "duration_sec", "wpm", "raw_wpm", "accuracy", "consistency", "total_chars", "correct_chars", "error_count"}
	if err := writer.Write(header); err != nil {
		return err
	}

	for _, r := range records {
		row := []string{
			r.Timestamp.Format("2006-01-02 15:04:05"),
			fmt.Sprintf("%ds", r.DurationSec),
			fmt.Sprintf("%d", r.DurationSec),
			fmt.Sprintf("%.2f", r.WPM),
			fmt.Sprintf("%.2f", r.RawWPM),
			fmt.Sprintf("%.2f", r.Accuracy),
			fmt.Sprintf("%.2f", r.Consistency),
			fmt.Sprintf("%d", r.TotalChars),
			fmt.Sprintf("%d", r.CorrectChars),
			fmt.Sprintf("%d", r.ErrorCount),
		}
		if err := writer.Write(row); err != nil {
			return err
		}
	}
	return nil
}
