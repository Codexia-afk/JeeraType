package config

import (
	"bufio"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// LoadThemeFromTOML parses a key-value style custom TOML theme file.
func LoadThemeFromTOML(filePath string) (Theme, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return ThemeAmber, err
	}
	defer file.Close()

	theme := ThemeAmber
	theme.Name = "custom"

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.ToLower(strings.TrimSpace(parts[0]))
		val := strings.Trim(strings.TrimSpace(parts[1]), "\"")

		switch key {
		case "name":
			theme.Name = val
		case "primary":
			theme.Primary = lipgloss.Color(val)
		case "secondary":
			theme.Secondary = lipgloss.Color(val)
		case "success":
			theme.Success = lipgloss.Color(val)
		case "error":
			theme.Error = lipgloss.Color(val)
		case "dim":
			theme.Dim = lipgloss.Color(val)
		case "subtle":
			theme.Subtle = lipgloss.Color(val)
		case "background":
			theme.Background = lipgloss.Color(val)
		case "highlight":
			theme.Highlight = lipgloss.Color(val)
		case "ghost_cursor":
			theme.GhostCursor = lipgloss.Color(val)
		}
	}

	return theme, scanner.Err()
}
