package ui

import (
	"github.com/charmbracelet/lipgloss"
	"jeeratype/config"
)

// DynamicStyles constructs Lipgloss styles dynamically based on selected Theme.
type DynamicStyles struct {
	Theme          config.Theme
	StyleUntyped   lipgloss.Style
	StyleCorrect   lipgloss.Style
	StyleIncorrect lipgloss.Style
	StyleCursor    lipgloss.Style
	StyleGhost     lipgloss.Style
	StyleHeader    lipgloss.Style
	StyleSubtext   lipgloss.Style
	StyleActive    lipgloss.Style
	StyleInactive  lipgloss.Style
	StyleCard      lipgloss.Style
	StyleCardVal   lipgloss.Style
	StyleHelp      lipgloss.Style
}

// NewDynamicStyles creates themed styles.
func NewDynamicStyles(theme config.Theme) DynamicStyles {
	return DynamicStyles{
		Theme: theme,
		StyleUntyped: lipgloss.NewStyle().
			Foreground(theme.Dim),
		StyleCorrect: lipgloss.NewStyle().
			Foreground(theme.Success).
			Bold(true),
		StyleIncorrect: lipgloss.NewStyle().
			Foreground(theme.Error).
			Underline(true).
			Bold(true),
		StyleCursor: lipgloss.NewStyle().
			Background(theme.Secondary).
			Foreground(theme.Background).
			Bold(true),
		StyleGhost: lipgloss.NewStyle().
			Foreground(theme.GhostCursor).
			Underline(true),
		StyleHeader: lipgloss.NewStyle().
			Foreground(theme.Primary).
			Bold(true),
		StyleSubtext: lipgloss.NewStyle().
			Foreground(theme.Subtle).
			Italic(true),
		StyleActive: lipgloss.NewStyle().
			Background(theme.Primary).
			Foreground(theme.Background).
			Bold(true).
			Padding(0, 1),
		StyleInactive: lipgloss.NewStyle().
			Foreground(theme.Subtle).
			Padding(0, 1),
		StyleCard: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(theme.Primary).
			Padding(0, 2).
			Margin(0, 1),
		StyleCardVal: lipgloss.NewStyle().
			Foreground(theme.Secondary).
			Bold(true),
		StyleHelp: lipgloss.NewStyle().
			Foreground(theme.Dim).
			MarginTop(1),
	}
}
