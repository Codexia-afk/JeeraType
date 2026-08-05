package ui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/Codexia-afk/JeeraType/config"
)

// RenderVisualKeyboardOverlay renders live keyboard highlighting pressed keys in real-time.
func RenderVisualKeyboardOverlay(activeRune rune, theme config.Theme, termWidth int) string {
	row1 := []rune{'q', 'w', 'e', 'r', 't', 'y', 'u', 'i', 'o', 'p'}
	row2 := []rune{'a', 's', 'd', 'f', 'g', 'h', 'j', 'k', 'l'}
	row3 := []rune{'z', 'x', 'c', 'v', 'b', 'n', 'm'}

	var b strings.Builder

	b.WriteString(renderOverlayRow(row1, activeRune, theme, ""))
	b.WriteString("\n")
	b.WriteString(renderOverlayRow(row2, activeRune, theme, "  "))
	b.WriteString("\n")
	b.WriteString(renderOverlayRow(row3, activeRune, theme, "    "))

	content := b.String()
	if termWidth > 0 {
		return lipgloss.PlaceHorizontal(termWidth, lipgloss.Center, content)
	}
	return content
}

func renderOverlayRow(row []rune, activeRune rune, theme config.Theme, indent string) string {
	var keyBlocks []string
	activeLower := strings.ToLower(string(activeRune))

	for _, r := range row {
		charStr := strings.ToLower(string(r))
		charUpper := strings.ToUpper(string(r))

		var keyStyle lipgloss.Style
		if charStr == activeLower {
			keyStyle = lipgloss.NewStyle().
				Background(theme.Primary).
				Foreground(theme.Background).
				Bold(true).
				Border(lipgloss.RoundedBorder()).
				BorderForeground(theme.Secondary).
				Padding(0, 1)
		} else {
			keyStyle = lipgloss.NewStyle().
				Foreground(theme.Subtle).
				Border(lipgloss.NormalBorder()).
				BorderForeground(theme.Dim).
				Padding(0, 1)
		}

		keyBlocks = append(keyBlocks, keyStyle.Render(charUpper))
	}

	return indent + lipgloss.JoinHorizontal(lipgloss.Center, keyBlocks...)
}
