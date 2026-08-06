package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/Codexia-afk/JeeraType/config"
	"github.com/Codexia-afk/JeeraType/db"
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

// RenderShadedKeyboardHeatmap renders ASCII QWERTY keyboard with 5 block shading levels.
func RenderShadedKeyboardHeatmap(theme config.Theme, termWidth int) string {
	row1 := []rune{'q', 'w', 'e', 'r', 't', 'y', 'u', 'i', 'o', 'p'}
	row2 := []rune{'a', 's', 'd', 'f', 'g', 'h', 'j', 'k', 'l'}
	row3 := []rune{'z', 'x', 'c', 'v', 'b', 'n', 'm'}

	// Fetch error stats from SQLite
	weakKeys := db.GetTopWeakKeys(26)
	errorMap := make(map[rune]int)
	maxErr := 1
	for idx, r := range weakKeys {
		cnt := len(weakKeys) - idx
		errorMap[r] = cnt
		if cnt > maxErr {
			maxErr = cnt
		}
	}

	var b strings.Builder
	b.WriteString(lipgloss.NewStyle().Foreground(theme.Primary).Bold(true).Render("[Heatmap] Cumulative Key Error Heatmap"))
	b.WriteString("\n\n")

	b.WriteString(renderShadedRow(row1, errorMap, maxErr, theme, ""))
	b.WriteString("\n")
	b.WriteString(renderShadedRow(row2, errorMap, maxErr, theme, "  "))
	b.WriteString("\n")
	b.WriteString(renderShadedRow(row3, errorMap, maxErr, theme, "    "))
	b.WriteString("\n\n")

	legend := fmt.Sprintf("Legend: %s Low  %s Mid  %s High  %s Critical",
		lipgloss.NewStyle().Foreground(theme.Dim).Render("░"),
		lipgloss.NewStyle().Foreground(theme.Secondary).Render("▒"),
		lipgloss.NewStyle().Foreground(theme.Highlight).Render("▓"),
		lipgloss.NewStyle().Foreground(theme.Error).Bold(true).Render("█"),
	)
	b.WriteString(legend)

	content := b.String()
	if termWidth > 0 {
		return lipgloss.PlaceHorizontal(termWidth, lipgloss.Center, content)
	}
	return content
}

func renderShadedRow(row []rune, errorMap map[rune]int, maxErr int, theme config.Theme, indent string) string {
	var keyBlocks []string

	for _, r := range row {
		charUpper := strings.ToUpper(string(r))
		errCount := errorMap[r]

		var shadeGlyph string
		var fgColor lipgloss.Color

		if errCount == 0 {
			shadeGlyph = " "
			fgColor = theme.Subtle
		} else {
			ratio := float64(errCount) / float64(maxErr)
			if ratio < 0.25 {
				shadeGlyph = "░"
				fgColor = theme.Dim
			} else if ratio < 0.50 {
				shadeGlyph = "▒"
				fgColor = theme.Secondary
			} else if ratio < 0.75 {
				shadeGlyph = "▓"
				fgColor = theme.Highlight
			} else {
				shadeGlyph = "█"
				fgColor = theme.Error
			}
		}

		keyStyle := lipgloss.NewStyle().
			Foreground(fgColor).
			Border(lipgloss.NormalBorder()).
			BorderForeground(theme.Dim).
			Padding(0, 1)

		keyBlocks = append(keyBlocks, keyStyle.Render(fmt.Sprintf("%s%s", charUpper, shadeGlyph)))
	}

	return indent + lipgloss.JoinHorizontal(lipgloss.Center, keyBlocks...)
}
