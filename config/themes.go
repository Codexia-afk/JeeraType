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

	ThemeCyberpunk = Theme{
		Name:        "cyberpunk",
		Primary:     lipgloss.Color("#00F0FF"), // Electric Neon Cyan
		Secondary:   lipgloss.Color("#FFE600"), // Neon Yellow
		Success:     lipgloss.Color("#00FF66"), // Neon Mint
		Error:       lipgloss.Color("#FF007F"), // Hot Pink
		Dim:         lipgloss.Color("#3A1C71"), // Dark Purple Void
		Subtle:      lipgloss.Color("#E0C3FC"), // Soft Violet
		Background:  lipgloss.Color("#0D0221"),
		Highlight:   lipgloss.Color("#FF007F"),
		GhostCursor: lipgloss.Color("#4F1787"),
	}

	ThemeTokyoNight = Theme{
		Name:        "tokyonight",
		Primary:     lipgloss.Color("#7AA2F7"), // Tokyo Blue
		Secondary:   lipgloss.Color("#E0AF68"), // Soft Gold
		Success:     lipgloss.Color("#9ECE6A"), // Leaf Green
		Error:       lipgloss.Color("#F7768E"), // Red Pink
		Dim:         lipgloss.Color("#414868"), // Dark Slate
		Subtle:      lipgloss.Color("#C0CAF5"), // Pale Blue
		Background:  lipgloss.Color("#1A1B26"),
		Highlight:   lipgloss.Color("#7DCFFF"), // Cyan
		GhostCursor: lipgloss.Color("#24283B"),
	}

	ThemeMonokai = Theme{
		Name:        "monokai",
		Primary:     lipgloss.Color("#FFD866"), // Warm Gold
		Secondary:   lipgloss.Color("#FC9867"), // Sunset Orange
		Success:     lipgloss.Color("#A9DC76"), // Mint Green
		Error:       lipgloss.Color("#FF6188"), // Crimson Pink
		Dim:         lipgloss.Color("#5B585C"), // Medium Charcoal
		Subtle:      lipgloss.Color("#FCE5C0"), // Cream
		Background:  lipgloss.Color("#2D2A2E"), // Dark Obsidian
		Highlight:   lipgloss.Color("#78DCE8"), // Cyan
		GhostCursor: lipgloss.Color("#403E41"),
	}

	ThemeRosePine = Theme{
		Name:        "rose-pine",
		Primary:     lipgloss.Color("#EBBCBA"), // Rose Gold
		Secondary:   lipgloss.Color("#F6C177"), // Gold
		Success:     lipgloss.Color("#31748F"), // Pine Cyan
		Error:       lipgloss.Color("#EB6F92"), // Love Red
		Dim:         lipgloss.Color("#403C58"), // Dark Dusk
		Subtle:      lipgloss.Color("#E0DEF4"), // Subtext
		Background:  lipgloss.Color("#191724"),
		Highlight:   lipgloss.Color("#C4A7E7"), // Iris Purple
		GhostCursor: lipgloss.Color("#26233A"),
	}

	ThemeSynthwave = Theme{
		Name:        "synthwave",
		Primary:     lipgloss.Color("#FF7ED4"), // Retro Pink
		Secondary:   lipgloss.Color("#FEDE5D"), // Neon Yellow
		Success:     lipgloss.Color("#72F1B8"), // Mint Green
		Error:       lipgloss.Color("#FE4450"), // Bright Red
		Dim:         lipgloss.Color("#493967"), // Arcade Dark
		Subtle:      lipgloss.Color("#F9F8F6"), // Light Chalk
		Background:  lipgloss.Color("#241B2F"),
		Highlight:   lipgloss.Color("#36F9F6"), // Neon Cyan
		GhostCursor: lipgloss.Color("#342948"),
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

	ThemeSolarized = Theme{
		Name:        "solarized",
		Primary:     lipgloss.Color("#268BD2"), // Blue
		Secondary:   lipgloss.Color("#B58900"), // Yellow
		Success:     lipgloss.Color("#859900"), // Green
		Error:       lipgloss.Color("#DC322F"), // Red
		Dim:         lipgloss.Color("#586E75"), // Base01
		Subtle:      lipgloss.Color("#93A1A1"), // Base1
		Background:  lipgloss.Color("#002B36"), // Base03
		Highlight:   lipgloss.Color("#2AA198"), // Cyan
		GhostCursor: lipgloss.Color("#073642"),
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

	ThemeGruvbox = Theme{
		Name:        "gruvbox",
		Primary:     lipgloss.Color("#FE8019"), // Orange
		Secondary:   lipgloss.Color("#FABD2F"), // Yellow
		Success:     lipgloss.Color("#B8BB26"), // Green
		Error:       lipgloss.Color("#FB4934"), // Red
		Dim:         lipgloss.Color("#665C54"), // Gray
		Subtle:      lipgloss.Color("#EBDBB2"), // Light
		Background:  lipgloss.Color("#282828"),
		Highlight:   lipgloss.Color("#83A598"), // Blue
		GhostCursor: lipgloss.Color("#504945"),
	}

	AvailableThemes = []Theme{
		ThemeAmber,
		ThemeCyberpunk,
		ThemeTokyoNight,
		ThemeMonokai,
		ThemeRosePine,
		ThemeSynthwave,
		ThemeDracula,
		ThemeNord,
		ThemeSolarized,
		ThemeCatppuccin,
		ThemeGruvbox,
		ThemeMatrix,
	}
)

// GetThemeByName returns matching Theme or default Amber.
func GetThemeByName(name string) Theme {
	n := strings.ToLower(strings.TrimSpace(name))
	if n == "default" {
		n = "amber"
	}
	for _, t := range AvailableThemes {
		if t.Name == n {
			return t
		}
	}
	return ThemeAmber
}
