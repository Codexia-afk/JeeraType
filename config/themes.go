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

	ThemeCyberpunk = Theme{
		Name:        "cyberpunk",
		Primary:     lipgloss.Color("#00F0FF"),
		Secondary:   lipgloss.Color("#FFE600"),
		Success:     lipgloss.Color("#00FF66"),
		Error:       lipgloss.Color("#FF007F"),
		Dim:         lipgloss.Color("#3A1C71"),
		Subtle:      lipgloss.Color("#E0C3FC"),
		Background:  lipgloss.Color("#0D0221"),
		Highlight:   lipgloss.Color("#FF007F"),
		GhostCursor: lipgloss.Color("#4F1787"),
	}

	ThemeTokyoNight = Theme{
		Name:        "tokyonight",
		Primary:     lipgloss.Color("#7AA2F7"),
		Secondary:   lipgloss.Color("#E0AF68"),
		Success:     lipgloss.Color("#9ECE6A"),
		Error:       lipgloss.Color("#F7768E"),
		Dim:         lipgloss.Color("#414868"),
		Subtle:      lipgloss.Color("#C0CAF5"),
		Background:  lipgloss.Color("#1A1B26"),
		Highlight:   lipgloss.Color("#7DCFFF"),
		GhostCursor: lipgloss.Color("#24283B"),
	}

	ThemeMonokai = Theme{
		Name:        "monokai",
		Primary:     lipgloss.Color("#FFD866"),
		Secondary:   lipgloss.Color("#FC9867"),
		Success:     lipgloss.Color("#A9DC76"),
		Error:       lipgloss.Color("#FF6188"),
		Dim:         lipgloss.Color("#5B585C"),
		Subtle:      lipgloss.Color("#FCE5C0"),
		Background:  lipgloss.Color("#2D2A2E"),
		Highlight:   lipgloss.Color("#78DCE8"),
		GhostCursor: lipgloss.Color("#403E41"),
	}

	ThemeRosePine = Theme{
		Name:        "rose-pine",
		Primary:     lipgloss.Color("#EBBCBA"),
		Secondary:   lipgloss.Color("#F6C177"),
		Success:     lipgloss.Color("#31748F"),
		Error:       lipgloss.Color("#EB6F92"),
		Dim:         lipgloss.Color("#403C58"),
		Subtle:      lipgloss.Color("#E0DEF4"),
		Background:  lipgloss.Color("#191724"),
		Highlight:   lipgloss.Color("#C4A7E7"),
		GhostCursor: lipgloss.Color("#26233A"),
	}

	ThemeSynthwave = Theme{
		Name:        "synthwave",
		Primary:     lipgloss.Color("#FF7ED4"),
		Secondary:   lipgloss.Color("#FEDE5D"),
		Success:     lipgloss.Color("#72F1B8"),
		Error:       lipgloss.Color("#FE4450"),
		Dim:         lipgloss.Color("#493967"),
		Subtle:      lipgloss.Color("#F9F8F6"),
		Background:  lipgloss.Color("#241B2F"),
		Highlight:   lipgloss.Color("#36F9F6"),
		GhostCursor: lipgloss.Color("#342948"),
	}

	ThemeCatppuccin = Theme{
		Name:        "catppuccin",
		Primary:     lipgloss.Color("#CBA6F7"),
		Secondary:   lipgloss.Color("#F9E2AF"),
		Success:     lipgloss.Color("#A6E3A1"),
		Error:       lipgloss.Color("#F38BA8"),
		Dim:         lipgloss.Color("#585B70"),
		Subtle:      lipgloss.Color("#BAC2DE"),
		Background:  lipgloss.Color("#1E1E2E"),
		Highlight:   lipgloss.Color("#89B4FA"),
		GhostCursor: lipgloss.Color("#6C7086"),
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

	ThemeMatrix = Theme{
		Name:        "matrix",
		Primary:     lipgloss.Color("#00FF66"),
		Secondary:   lipgloss.Color("#CCFF00"),
		Success:     lipgloss.Color("#00FF99"),
		Error:       lipgloss.Color("#FF0055"),
		Dim:         lipgloss.Color("#004D1A"),
		Subtle:      lipgloss.Color("#66FF99"),
		Background:  lipgloss.Color("#051A05"),
		Highlight:   lipgloss.Color("#00FFCC"),
		GhostCursor: lipgloss.Color("#006622"),
	}

	ThemeGruvbox = Theme{
		Name:        "gruvbox",
		Primary:     lipgloss.Color("#FE8019"),
		Secondary:   lipgloss.Color("#FABD2F"),
		Success:     lipgloss.Color("#B8BB26"),
		Error:       lipgloss.Color("#FB4934"),
		Dim:         lipgloss.Color("#665C54"),
		Subtle:      lipgloss.Color("#EBDBB2"),
		Background:  lipgloss.Color("#282828"),
		Highlight:   lipgloss.Color("#83A598"),
		GhostCursor: lipgloss.Color("#504945"),
	}

	AvailableThemes = []Theme{
		ThemeAmber,
		ThemeJewel,
		ThemeSunset,
		ThemeForest,
		ThemeNeon,
		ThemeVintage,
		ThemeMono,
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
