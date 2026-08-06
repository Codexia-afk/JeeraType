package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/Codexia-afk/JeeraType/config"
	"github.com/Codexia-afk/JeeraType/storage"
	"github.com/Codexia-afk/JeeraType/ui"
)

// RenderStatsTable renders formatted table of last 20 runs with trend indicators.
func RenderStatsTable(records []storage.HistoryRecord, theme config.Theme) string {
	var b strings.Builder

	titleStyle := lipgloss.NewStyle().
		Foreground(theme.Primary).
		Bold(true).
		Padding(0, 1)

	b.WriteString(titleStyle.Render("📊 JeeraType — Session History (Last 20 Runs)"))
	b.WriteString("\n\n")

	if len(records) == 0 {
		b.WriteString(lipgloss.NewStyle().Foreground(theme.Subtle).Render("No test records found. Complete a test run first!"))
		return b.String()
	}

	n := len(records)
	startIdx := n - 20
	if startIdx < 0 {
		startIdx = 0
	}
	subset := records[startIdx:]

	rev := make([]storage.HistoryRecord, len(subset))
	for i, r := range subset {
		rev[len(subset)-1-i] = r
	}

	headerStyle := lipgloss.NewStyle().Foreground(theme.Secondary).Bold(true)
	b.WriteString(fmt.Sprintf("%-16s %-12s %-6s %-8s %-7s %-7s %-5s\n",
		headerStyle.Render("DATE"),
		headerStyle.Render("MODE"),
		headerStyle.Render("DUR"),
		headerStyle.Render("WPM"),
		headerStyle.Render("ACC"),
		headerStyle.Render("CONS"),
		headerStyle.Render("TREND"),
	))
	b.WriteString(lipgloss.NewStyle().Foreground(theme.Dim).Render(strings.Repeat("─", 72)) + "\n")

	rowStyle := lipgloss.NewStyle().Foreground(theme.Subtle)
	upTrend := lipgloss.NewStyle().Foreground(theme.Success).Bold(true).Render("↑")
	downTrend := lipgloss.NewStyle().Foreground(theme.Error).Bold(true).Render("↓")
	flatTrend := lipgloss.NewStyle().Foreground(theme.Dim).Render("–")

	for i, r := range rev {
		dateStr := r.Timestamp.Format("2006-01-02 15:04")
		durStr := fmt.Sprintf("%ds", r.DurationSec)
		wpmStr := fmt.Sprintf("%.1f", r.WPM)
		accStr := fmt.Sprintf("%.0f%%", r.Accuracy)
		consStr := fmt.Sprintf("%.0f%%", r.Consistency)

		trendStr := flatTrend
		for j := i + 1; j < len(rev); j++ {
			prev := rev[j]
			if prev.DurationSec == r.DurationSec {
				if r.WPM > prev.WPM+0.5 {
					trendStr = upTrend
				} else if r.WPM < prev.WPM-0.5 {
					trendStr = downTrend
				}
				break
			}
		}

		b.WriteString(rowStyle.Render(fmt.Sprintf("%-16s %-12s %-6s %-8s %-7s %-7s %-5s\n",
			dateStr, "Paragraphs", durStr, wpmStr, accStr, consStr, trendStr,
		)))
	}

	return b.String()
}

// RunStatsSubcommand executes `jeeratype stats` command.
func RunStatsSubcommand(themeName string, showHeatmap bool, showLeaderboard bool) {
	th := config.GetThemeByName(themeName)

	if showHeatmap {
		fmt.Println(ui.RenderShadedKeyboardHeatmap(th, 80))
		return
	}

	if showLeaderboard {
		leaderboard, _ := storage.GetLeaderboard()
		fmt.Println(storage.FormatLeaderboardTable(leaderboard))
		return
	}

	records, err := storage.LoadHistory()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading history: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(RenderStatsTable(records, th))
}
