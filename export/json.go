package export

import (
	"encoding/json"
	"io"

	"github.com/Codexia-afk/JeeraType/storage"
)

// ExportJSON formats history records into JSON output.
func ExportJSON(records []storage.HistoryRecord, w io.Writer) error {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	return encoder.Encode(records)
}
