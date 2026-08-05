package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"jeeratype/config"
	"jeeratype/db"
)

// RenderKeyboardHeatmap renders a full color-coded ASCII QWERTY keyboard heatmap.
func RenderKeyboardHeatmap(theme config.Theme, termWidth int) string {
	metrics := db.GetKeyMetrics()

	row1 := []rune{'q', 'w', 'e', 'r', 't', 'y', 'u', 'i', 'o', 'p'}
	row2 := []rune{'a', 's', 'd', 'f', 'g', 'h', 'j', 'k', 'l'}
	row3 := []rune{'z', 'x', 'c', 'v', 'b', 'n', 'm'}

	var b strings.Builder

	b.WriteString(lipgloss.NewStyle().Foreground(theme.Primary).Bold(true).Render("⌨  ASCII KEYBOARD HEATMAP (Historical Performance)"))
	b.WriteString("\n\n")

	b.WriteString(renderKeyRow(row1, metrics, theme, ""))
	b.WriteString("\n")
	b.WriteString(renderKeyRow(row2, metrics, theme, "  "))
	b.WriteString("\n")
	b.WriteString(renderKeyRow(row3, metrics, theme, "    "))
	b.WriteString("\n\n")

	// Heatmap Legend
	legendGreen := lipgloss.NewStyle().Foreground(theme.Success).Render("■ Excellent (Fast & Accurate)")
	legendAmber := lipgloss.NewStyle().Foreground(theme.Secondary).Render("■ Moderate")
	legendRed := lipgloss.NewStyle().Foreground(theme.Error).Render("■ Weak / High Latency")
	legendDim := lipgloss.NewStyle().Foreground(theme.Dim).Render("■ Untested")

	legend := fmt.Sprintf("Legend: %s   %s   %s   %s", legendGreen, legendAmber, legendRed, legendDim)
	b.WriteString(legend)

	content := b.String()
	if termWidth > 0 {
		return lipgloss.PlaceHorizontal(termWidth, lipgloss.Center, content)
	}
	return content
}

func renderKeyRow(row []rune, metrics map[rune]db.KeyMetric, theme config.Theme, indent string) string {
	var keyBlocks []string
	for _, r := range row {
		m, exists := metrics[r]
		charUpper := strings.ToUpper(string(r))

		var keyStyle lipgloss.Style
		if !exists || m.TotalHits == 0 {
			keyStyle = lipgloss.NewStyle().
				Foreground(theme.Dim).
				Border(lipgloss.NormalBorder()).
				BorderForeground(theme.Dim).
				Padding(0, 1)
		} else if m.Accuracy >= 95.0 && m.AvgLatency < 300 {
			keyStyle = lipgloss.NewStyle().
				Foreground(theme.Success).
				Bold(true).
				Border(lipgloss.RoundedBorder()).
				BorderForeground(theme.Success).
				Padding(0, 1)
		} else if m.Accuracy >= 85.0 {
			keyStyle = lipgloss.NewStyle().
				Foreground(theme.Secondary).
				Bold(true).
				Border(lipgloss.RoundedBorder()).
				BorderForeground(theme.Secondary).
				Padding(0, 1)
		} else {
			keyStyle = lipgloss.NewStyle().
				Foreground(theme.Error).
				Bold(true).
				Border(lipgloss.RoundedBorder()).
				BorderForeground(theme.Error).
				Padding(0, 1)
		}

		keyBlocks = append(keyBlocks, keyStyle.Render(charUpper))
	}

	return indent + lipgloss.JoinHorizontal(lipgloss.Center, keyBlocks...)
}
