package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// AppConfig defines user settings saved to ~/.config/jeeratype/config.json
type AppConfig struct {
	Theme          string `json:"theme"`
	DefaultProfile string `json:"default_profile"`
	Punctuation    bool   `json:"punctuation"`
	Numbers        bool   `json:"numbers"`
	Sound          bool   `json:"sound"`
	LastMode       string `json:"last_mode"`
	LastDuration   int    `json:"last_duration"`
}

func DefaultConfig() AppConfig {
	return AppConfig{
		Theme:          "amber",
		DefaultProfile: "default",
		Punctuation:    false,
		Numbers:        false,
		Sound:          false,
		LastMode:       "paragraphs",
		LastDuration:   30,
	}
}

// GetConfigPath returns path to ~/.config/jeeratype/config.json
func GetConfigPath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		home := os.Getenv("HOME")
		if home != "" {
			configDir = filepath.Join(home, ".config")
		} else {
			configDir = os.TempDir()
		}
	}
	appDir := filepath.Join(configDir, "jeeratype")
	_ = os.MkdirAll(appDir, 0755)
	return filepath.Join(appDir, "config.json"), nil
}

// LoadConfig reads config.json or returns sensible defaults if missing.
func LoadConfig() AppConfig {
	cfg := DefaultConfig()
	path, err := GetConfigPath()
	if err != nil {
		return cfg
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return cfg
	}

	_ = json.Unmarshal(data, &cfg)
	if cfg.Theme == "" {
		cfg.Theme = "amber"
	}
	if cfg.DefaultProfile == "" {
		cfg.DefaultProfile = "default"
	}
	return cfg
}

// SaveConfig persists AppConfig to config.json
func SaveConfig(cfg AppConfig) error {
	path, err := GetConfigPath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}
