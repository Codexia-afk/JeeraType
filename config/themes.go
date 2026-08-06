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
	ThemeDefault = Theme{
		Name:        "default",
		Primary:     lipgloss.Color("#3B82F6"), // Vibrant Blue
		Secondary:   lipgloss.Color("#60A5FA"), // Sky Blue
		Success:     lipgloss.Color("#10B981"), // Emerald Green
		Error:       lipgloss.Color("#EF4444"), // Crimson Red
		Dim:         lipgloss.Color("#4B5563"), // Dark Slate
		Subtle:      lipgloss.Color("#9CA3AF"), // Light Gray
		Background:  lipgloss.Color("#0F172A"), // Deep Navy/Slate
		Highlight:   lipgloss.Color("#F59E0B"), // Amber Highlight
		GhostCursor: lipgloss.Color("#6B7280"),
	}

	ThemeDracula = Theme{
		Name:        "dracula",
		Primary:     lipgloss.Color("#BD93F9"),
		Secondary:   lipgloss.Color("#F1FA8C"),
		Success:     lipgloss.Color("#50FA7B"),
		Error:       lipgloss.Color("#FF5555"),
		Dim:         lipgloss.Color("#6272A4"),
		Subtle:      lipgloss.Color("#F8F8F2"),
		Background:  lipgloss.Color("#282A36"),
		Highlight:   lipgloss.Color("#8BE9FD"),
		GhostCursor: lipgloss.Color("#44475A"),
	}

	ThemeNord = Theme{
		Name:        "nord",
		Primary:     lipgloss.Color("#88C0D0"),
		Secondary:   lipgloss.Color("#EBCB8B"),
		Success:     lipgloss.Color("#A3BE8C"),
		Error:       lipgloss.Color("#BF616A"),
		Dim:         lipgloss.Color("#4C566A"),
		Subtle:      lipgloss.Color("#D8DEE9"),
		Background:  lipgloss.Color("#2E3440"),
		Highlight:   lipgloss.Color("#81A1C1"),
		GhostCursor: lipgloss.Color("#434C5E"),
	}

	ThemeSolarized = Theme{
		Name:        "solarized",
		Primary:     lipgloss.Color("#268BD2"),
		Secondary:   lipgloss.Color("#B58900"),
		Success:     lipgloss.Color("#859900"),
		Error:       lipgloss.Color("#DC322F"),
		Dim:         lipgloss.Color("#586E75"),
		Subtle:      lipgloss.Color("#93A1A1"),
		Background:  lipgloss.Color("#002B36"),
		Highlight:   lipgloss.Color("#2AA198"),
		GhostCursor: lipgloss.Color("#073642"),
	}

	ThemeJewel = Theme{
		Name:        "jewel",
		Primary:     lipgloss.Color("#50C878"), // Emerald
		Secondary:   lipgloss.Color("#FFD700"), // Gold
		Success:     lipgloss.Color("#0F52BA"), // Sapphire
		Error:       lipgloss.Color("#E0115F"), // Ruby
		Dim:         lipgloss.Color("#3A2E56"), // Deep Amethyst
		Subtle:      lipgloss.Color("#B39DDB"), // Light Lavender
		Background:  lipgloss.Color("#0A0915"),
		Highlight:   lipgloss.Color("#FFD700"),
		GhostCursor: lipgloss.Color("#281E3D"),
	}

	ThemeSunset = Theme{
		Name:        "sunset",
		Primary:     lipgloss.Color("#FF7F50"), // Coral
		Secondary:   lipgloss.Color("#FFA07A"), // Light Salmon
		Success:     lipgloss.Color("#FF6347"), // Tomato
		Error:       lipgloss.Color("#E63946"), // Crimson
		Dim:         lipgloss.Color("#5C3A6B"), // Plum Slate
		Subtle:      lipgloss.Color("#E2B6CF"), // Soft Mauve
		Background:  lipgloss.Color("#1E0F24"),
		Highlight:   lipgloss.Color("#DA70D6"), // Orchid
		GhostCursor: lipgloss.Color("#3B1E47"),
	}

	ThemeForest = Theme{
		Name:        "forest",
		Primary:     lipgloss.Color("#8FBC8F"), // Moss Green
		Secondary:   lipgloss.Color("#EEDC82"), // Bark Gold
		Success:     lipgloss.Color("#708090"), // Slate
		Error:       lipgloss.Color("#D9534F"), // Autumn Red
		Dim:         lipgloss.Color("#3A4F3D"), // Pine Shadow
		Subtle:      lipgloss.Color("#C2D4C3"), // Sage Cream
		Background:  lipgloss.Color("#1B261D"),
		Highlight:   lipgloss.Color("#F5F5DC"), // Cream
		GhostCursor: lipgloss.Color("#28382A"),
	}

	ThemeNeon = Theme{
		Name:        "neon",
		Primary:     lipgloss.Color("#00F5FF"), // Electric Cyan
		Secondary:   lipgloss.Color("#FFE600"), // Acid Yellow
		Success:     lipgloss.Color("#39FF14"), // Neon Green
		Error:       lipgloss.Color("#FF007F"), // Hot Pink
		Dim:         lipgloss.Color("#333333"), // Dark Charcoal
		Subtle:      lipgloss.Color("#CCCCCC"), // Light Gray
		Background:  lipgloss.Color("#000000"),
		Highlight:   lipgloss.Color("#FF007F"),
		GhostCursor: lipgloss.Color("#222222"),
	}

	ThemeVintage = Theme{
		Name:        "vintage",
		Primary:     lipgloss.Color("#E5C07B"), // Mustard
		Secondary:   lipgloss.Color("#D19A66"), // Rust Orange
		Success:     lipgloss.Color("#98C379"), // Sage Green
		Error:       lipgloss.Color("#E06C75"), // Muted Red
		Dim:         lipgloss.Color("#4B5263"), // Dark Slate
		Subtle:      lipgloss.Color("#ABB2BF"), // Soft Gray
		Background:  lipgloss.Color("#21252B"),
		Highlight:   lipgloss.Color("#F0E6D2"), // Paper Cream
		GhostCursor: lipgloss.Color("#2C313A"),
	}

	ThemeMono = Theme{
		Name:        "mono",
		Primary:     lipgloss.Color("#FFFFFF"), // Stark White
		Secondary:   lipgloss.Color("#FFD700"), // Gold
		Success:     lipgloss.Color("#CCCCCC"), // Light Gray
		Error:       lipgloss.Color("#FF3333"), // Red
		Dim:         lipgloss.Color("#444444"), // Dark Gray
		Subtle:      lipgloss.Color("#888888"), // Mid Gray
		Background:  lipgloss.Color("#101010"),
		Highlight:   lipgloss.Color("#FFD700"),
		GhostCursor: lipgloss.Color("#222222"),
	}

	AvailableThemes = []Theme{
		ThemeDefault,
		ThemeDracula,
		ThemeNord,
		ThemeSolarized,
		ThemeJewel,
		ThemeSunset,
		ThemeForest,
		ThemeNeon,
		ThemeVintage,
		ThemeMono,
	}
)

// GetThemeByName returns matching Theme or default Theme.
func GetThemeByName(name string) Theme {
	n := strings.ToLower(strings.TrimSpace(name))
	for _, t := range AvailableThemes {
		if t.Name == n {
			return t
		}
	}
	return ThemeDefault
}
