package db

import (
	"database/sql"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

// GetDBFilePath returns path to ~/.config/jeeratype/stats.db with permission fallback.
func GetDBFilePath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err == nil {
		appDir := filepath.Join(configDir, "jeeratype")
		if err := os.MkdirAll(appDir, 0755); err == nil {
			return filepath.Join(appDir, "stats.db"), nil
		}
	}

	if home := os.Getenv("HOME"); home != "" {
		appDir := filepath.Join(home, ".config", "jeeratype")
		if err := os.MkdirAll(appDir, 0755); err == nil {
			return filepath.Join(appDir, "stats.db"), nil
		}
	}

	appDir := filepath.Join(os.TempDir(), "jeeratype")
	_ = os.MkdirAll(appDir, 0755)
	return filepath.Join(appDir, "stats.db"), nil
}

// InitDB initializes SQLite connection and creates schema.
func InitDB() error {
	dbPath, err := GetDBFilePath()
	if err != nil {
		return err
	}

	database, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return err
	}

	DB = database

	schema := `
	CREATE TABLE IF NOT EXISTS test_runs (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		timestamp DATETIME,
		mode TEXT,
		wpm REAL,
		raw_wpm REAL,
		accuracy REAL,
		consistency REAL,
		duration_sec INTEGER
	);

	CREATE TABLE IF NOT EXISTS key_stats (
		key_char TEXT PRIMARY KEY,
		total_hits INTEGER DEFAULT 0,
		total_errors INTEGER DEFAULT 0,
		total_latency_ms INTEGER DEFAULT 0
	);

	CREATE TABLE IF NOT EXISTS bigram_stats (
		bigram TEXT PRIMARY KEY,
		count INTEGER DEFAULT 0,
		total_latency_ms INTEGER DEFAULT 0
	);

	CREATE TABLE IF NOT EXISTS file_offsets (
		file_path TEXT PRIMARY KEY,
		cursor_offset INTEGER DEFAULT 0,
		updated_at DATETIME
	);

	CREATE TABLE IF NOT EXISTS pb_samples (
		mode_key TEXT PRIMARY KEY,
		wpm REAL DEFAULT 0,
		samples_json TEXT
	);
	`

	_, err = DB.Exec(schema)
	return err
}

// CloseDB closes database connection.
func CloseDB() {
	if DB != nil {
		_ = DB.Close()
	}
}
