package config

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// Theme defines color tokens for JeeraType UI.
type Theme struct {
	Name        string
	Primary     lipgloss.Color
	Secondary   lipgloss.Color
	Success     lipgloss.Color
	Error       lipgloss.Color
	Dim         lipgloss.Color
	Subtle      lipgloss.Color
	Background  lipgloss.Color
	Highlight   lipgloss.Color
	GhostCursor lipgloss.Color
}

var (
	ThemeAmber = Theme{
		Name:        "amber",
		Primary:     lipgloss.Color("#D97706"),
		Secondary:   lipgloss.Color("#F59E0B"),
		Success:     lipgloss.Color("#10B981"),
		Error:       lipgloss.Color("#EF4444"),
		Dim:         lipgloss.Color("#4B5563"),
		Subtle:      lipgloss.Color("#9CA3AF"),
		Background:  lipgloss.Color("#111827"),
		Highlight:   lipgloss.Color("#FBBF24"),
		GhostCursor: lipgloss.Color("#6B7280"),
	}

	ThemeCatppuccin = Theme{
		Name:        "catppuccin",
		Primary:     lipgloss.Color("#CBA6F7"), // Mauve
		Secondary:   lipgloss.Color("#F9E2AF"), // Yellow
		Success:     lipgloss.Color("#A6E3A1"), // Green
		Error:       lipgloss.Color("#F38BA8"), // Red
		Dim:         lipgloss.Color("#585B70"), // Surface 2
		Subtle:      lipgloss.Color("#BAC2DE"), // Subtext 1
		Background:  lipgloss.Color("#1E1E2E"), // Base
		Highlight:   lipgloss.Color("#89B4FA"), // Blue
		GhostCursor: lipgloss.Color("#6C7086"), // Overlay 0
	}

	ThemeNord = Theme{
		Name:        "nord",
		Primary:     lipgloss.Color("#88C0D0"), // Frost Cyan
		Secondary:   lipgloss.Color("#EBCB8B"), // Yellow
		Success:     lipgloss.Color("#A3BE8C"), // Green
		Error:       lipgloss.Color("#BF616A"), // Red
		Dim:         lipgloss.Color("#4C566A"), // Polar Night
		Subtle:      lipgloss.Color("#D8DEE9"), // Snow Storm
		Background:  lipgloss.Color("#2E3440"),
		Highlight:   lipgloss.Color("#81A1C1"),
		GhostCursor: lipgloss.Color("#434C5E"),
	}

	ThemeDracula = Theme{
		Name:        "dracula",
		Primary:     lipgloss.Color("#BD93F9"), // Purple
		Secondary:   lipgloss.Color("#F1FA8C"), // Yellow
		Success:     lipgloss.Color("#50FA7B"), // Green
		Error:       lipgloss.Color("#FF5555"), // Red
		Dim:         lipgloss.Color("#6272A4"), // Comment Gray
		Subtle:      lipgloss.Color("#F8F8F2"), // Foreground
		Background:  lipgloss.Color("#282A36"),
		Highlight:   lipgloss.Color("#8BE9FD"), // Cyan
		GhostCursor: lipgloss.Color("#44475A"),
	}

	ThemeMatrix = Theme{
		Name:        "matrix",
		Primary:     lipgloss.Color("#00FF66"), // Bright Neon Green
		Secondary:   lipgloss.Color("#CCFF00"), // Lime Yellow
		Success:     lipgloss.Color("#00FF99"), // Mint Green
		Error:       lipgloss.Color("#FF0055"), // Red
		Dim:         lipgloss.Color("#004D1A"), // Dark Green
		Subtle:      lipgloss.Color("#66FF99"), // Pale Green
		Background:  lipgloss.Color("#051A05"),
		Highlight:   lipgloss.Color("#00FFCC"),
		GhostCursor: lipgloss.Color("#006622"),
	}

	AvailableThemes = []Theme{
		ThemeAmber,
		ThemeCatppuccin,
		ThemeNord,
		ThemeDracula,
		ThemeMatrix,
	}
)

// GetThemeByName returns matching Theme or default Amber.
func GetThemeByName(name string) Theme {
	n := strings.ToLower(strings.TrimSpace(name))
	for _, t := range AvailableThemes {
		if t.Name == n {
			return t
		}
	}
	return ThemeAmber
}
