package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/guptarohit/asciigraph"
	"github.com/Codexia-afk/JeeraType/config"
	"github.com/Codexia-afk/JeeraType/stats"
	"github.com/Codexia-afk/JeeraType/storage"
)

type TestMode int

const (
	ModeParagraphs TestMode = iota
	ModeCode
	ModeAdaptive
	ModeQuotes
	ModeStdin
	ModeFile
)

func (m TestMode) String() string {
	switch m {
	case ModeParagraphs:
		return "Paragraphs"
	case ModeCode:
		return "Code Mode"
	case ModeAdaptive:
		return "Adaptive Weakness"
	case ModeQuotes:
		return "Quotes"
	case ModeStdin:
		return "STDIN Pipe"
	case ModeFile:
		return "File Reader"
	default:
		return "Paragraphs"
	}
}

// RenderKeyboardHeatmap renders Keyboard Heatmap screen.
func RenderKeyboardHeatmap(theme config.Theme, termWidth int) string {
	return RenderShadedKeyboardHeatmap(theme, termWidth)
}

// RenderCountdownView renders centered 3...2...1... countdown animation.
func RenderCountdownView(count int, theme config.Theme, termWidth int) string {
	var b strings.Builder
	b.WriteString(RenderLogo(theme))
	b.WriteString("\n\n\n")

	bigNumStyle := lipgloss.NewStyle().
		Foreground(theme.Primary).
		Bold(true).
		Padding(1, 4).
		Border(lipgloss.DoubleBorder()).
		BorderForeground(theme.Secondary)

	b.WriteString(bigNumStyle.Render(fmt.Sprintf("   %d   ", count)))
	b.WriteString("\n\n")
	b.WriteString(lipgloss.NewStyle().Foreground(theme.Subtle).Italic(true).Render("Get ready to type..."))

	content := b.String()
	if termWidth > 0 {
		return lipgloss.PlaceHorizontal(termWidth, lipgloss.Center, content)
	}
	return content
}

// RenderMenuView renders initial options.
func RenderMenuView(
	timeModes []int,
	selectedModeIdx int,
	currentMode TestMode,
	currentTheme config.Theme,
	ghostWPM float64,
	showKeys bool,
	punctuation bool,
	numbers bool,
	isZen bool,
	termWidth int,
) string {
	s := NewDynamicStyles(currentTheme)
	var b strings.Builder

	// Top Logo
	b.WriteString(RenderLogo(currentTheme))
	b.WriteString("\n\n")

	// Category Mode Selector Pills
	b.WriteString(s.StyleSubtext.Render("Typing Mode:"))
	b.WriteString("\n")
	var modePills []string
	modes := []TestMode{ModeParagraphs, ModeCode, ModeAdaptive, ModeQuotes}
	for _, m := range modes {
		label := m.String()
		if m == currentMode {
			modePills = append(modePills, s.StyleActive.Render(" ✦ "+label+" "))
		} else {
			modePills = append(modePills, s.StyleInactive.Render("   "+label+" "))
		}
	}
	b.WriteString(lipgloss.JoinHorizontal(lipgloss.Center, modePills...))
	b.WriteString("\n\n")

	// Duration Selector Pills (or Zen mode indicator)
	if !isZen {
		b.WriteString(s.StyleSubtext.Render("Test Duration:"))
		b.WriteString("\n")
		var pills []string
		for i, duration := range timeModes {
			label := fmt.Sprintf("%ds", duration)
			if i == selectedModeIdx {
				pills = append(pills, s.StyleActive.Render(" > "+label+" < "))
			} else {
				pills = append(pills, s.StyleInactive.Render("   "+label+"   "))
			}
		}
		b.WriteString(lipgloss.JoinHorizontal(lipgloss.Center, pills...))
	} else {
		b.WriteString(s.StyleActive.Render(" 🧘 ZEN MODE (Infinite Text Stream) "))
	}
	b.WriteString("\n\n")

	// Status & Toggles Line
	puncStr := "OFF"
	if punctuation {
		puncStr = "ON"
	}
	numStr := "OFF"
	if numbers {
		numStr = "ON"
	}
	ghostStr := "Disabled"
	if ghostWPM > 0 {
		ghostStr = fmt.Sprintf("%.0f WPM 👻", ghostWPM)
	}

	infoLine := fmt.Sprintf("Theme: [%s] | Punctuation: [%s] | Numbers: [%s] | Ghost: [%s]",
		strings.ToUpper(currentTheme.Name), puncStr, numStr, ghostStr)
	b.WriteString(s.StyleSubtext.Render(infoLine))
	b.WriteString("\n\n")

	// Controls Guide
	helpText := "Controls:\n" +
		"  [p] Toggle Punctuation    [n] Toggle Numbers    [z] Toggle Zen Mode\n" +
		"  [m] Toggle Typing Mode    [t] Cycle Theme       [g] Ghost Pacer\n" +
		"  [v] Visual Keyboard       [1-5] Duration        [Enter / Space] Start"
	b.WriteString(s.StyleHelp.Render(helpText))

	content := b.String()
	if termWidth > 0 {
		return lipgloss.PlaceHorizontal(termWidth, lipgloss.Center, content)
	}
	return content
}

// RenderTestView renders live test screen.
func RenderTestView(t *stats.Tracker, theme config.Theme, showKeys bool, activeRune rune, termWidth int) string {
	s := NewDynamicStyles(theme)
	var b strings.Builder

	// Header Bar: Time left, Live WPM, Accuracy, Ghost Pacer, Hardcore Mode Badges
	liveWPM := t.CalculateWPM()
	liveAcc := t.CalculateAccuracy()

	var timeBadge string
	if t.IsZen {
		timeBadge = s.StyleActive.Render(" 🧘 ZEN MODE ")
	} else {
		remSec := t.RemainingSeconds()
		timeBadge = s.StyleActive.Render(fmt.Sprintf(" ⏱ %ds ", remSec))
	}

	wpmBadge := s.StyleCard.Render(fmt.Sprintf("WPM: %.0f", liveWPM))
	accBadge := s.StyleCard.Render(fmt.Sprintf("ACC: %.0f%%", liveAcc))

	headerItems := []string{timeBadge, "   ", wpmBadge, "   ", accBadge}
	if t.TargetWPM > 0 {
		ghostBadge := s.StyleCard.Render(fmt.Sprintf("GHOST: %.0f WPM 👻", t.TargetWPM))
		headerItems = append(headerItems, "   ", ghostBadge)
	}
	if t.PBWPM > 0 {
		pbBadge := s.StyleCard.Render(fmt.Sprintf("PB: %.0f WPM 🏆", t.PBWPM))
		headerItems = append(headerItems, "   ", pbBadge)
	}

	b.WriteString(lipgloss.JoinHorizontal(lipgloss.Center, headerItems...))
	b.WriteString("\n\n")

	maxWidth := termWidth - 10
	if maxWidth < 40 {
		maxWidth = 40
	}

	var textBuilder strings.Builder
	for i, r := range t.TargetRunes {
		charStr := string(r)

		isTargetGhost := (t.TargetWPM > 0 && i == t.GhostCursorIdx && i != t.CursorIdx)
		isPBGhost := (t.PBWPM > 0 && i == t.PBGhostCursorIdx && i != t.CursorIdx)

		if i == t.CursorIdx {
			if r == ' ' {
				textBuilder.WriteString(s.StyleCursor.Render(" "))
			} else {
				textBuilder.WriteString(s.StyleCursor.Render(charStr))
			}
		} else if isPBGhost {
			textBuilder.WriteString(lipgloss.NewStyle().Foreground(theme.Highlight).Underline(true).Render("🏆"))
		} else if isTargetGhost {
			textBuilder.WriteString(s.StyleGhost.Render(charStr))
		} else if i < t.CursorIdx {
			if t.CharStates[i] == stats.StatusCorrect {
				textBuilder.WriteString(s.StyleCorrect.Render(charStr))
			} else {
				if r == ' ' {
					textBuilder.WriteString(s.StyleIncorrect.Render("_"))
				} else {
					textBuilder.WriteString(s.StyleIncorrect.Render(charStr))
				}
			}
		} else {
			textBuilder.WriteString(s.StyleUntyped.Render(charStr))
		}
	}

	textBoxStyle := lipgloss.NewStyle().
		Width(maxWidth).
		Padding(1, 2).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(theme.Dim)

	b.WriteString(textBoxStyle.Render(textBuilder.String()))
	b.WriteString("\n\n")

	// Render Visual Keyboard Overlay if active
	if showKeys {
		b.WriteString(RenderVisualKeyboardOverlay(activeRune, theme, termWidth))
		b.WriteString("\n\n")
	}

	// Persistent Footer Hints Bar
	footerBar := lipgloss.NewStyle().
		Foreground(theme.Subtle).
		Background(theme.Dim).
		Padding(0, 2).
		Render("Tab: restart   Esc: quit   Ctrl+C: force quit")

	b.WriteString(footerBar)

	content := b.String()
	if termWidth > 0 {
		return lipgloss.PlaceHorizontal(termWidth, lipgloss.Center, content)
	}
	return content
}

// RenderResultsView renders post-game stats.
func RenderResultsView(t *stats.Tracker, theme config.Theme, termWidth int) string {
	s := NewDynamicStyles(theme)
	var b strings.Builder

	// Logo Header
	b.WriteString(RenderLogo(theme))
	b.WriteString("\n\n")

	// Streak Counter Badge
	records, _ := storage.LoadHistory()
	streak := stats.CalculateStreak(records)
	streakBadge := lipgloss.NewStyle().
		Foreground(theme.Primary).
		Bold(true).
		Render(stats.FormatStreakBadge(streak))
	b.WriteString(lipgloss.PlaceHorizontal(termWidth, lipgloss.Center, streakBadge))
	b.WriteString("\n\n")

	// Personal Best Banner if beaten
	netWPM := t.CalculateWPM()
	if t.IsNewPB {
		pbBanner := lipgloss.NewStyle().
			Foreground(theme.Background).
			Background(theme.Secondary).
			Bold(true).
			Padding(0, 2).
			Render(fmt.Sprintf("🎉 NEW BEST! %.1f WPM (prev: %.1f WPM)", netWPM, t.PreviousPBWPM))
		b.WriteString(pbBanner)
		b.WriteString("\n\n")
	}

	// Stat Cards
	grossWPM := t.CalculateGrossWPM()
	acc := t.CalculateAccuracy()
	consistency := t.CalculateConsistency()

	cardWPM := lipgloss.JoinVertical(lipgloss.Center,
		lipgloss.NewStyle().Foreground(theme.Subtle).Bold(true).Render("NET WPM"),
		s.StyleCardVal.Render(fmt.Sprintf("%.1f", netWPM)),
	)
	cardGross := lipgloss.JoinVertical(lipgloss.Center,
		lipgloss.NewStyle().Foreground(theme.Subtle).Bold(true).Render("GROSS WPM"),
		s.StyleCardVal.Render(fmt.Sprintf("%.1f", grossWPM)),
	)
	cardAcc := lipgloss.JoinVertical(lipgloss.Center,
		lipgloss.NewStyle().Foreground(theme.Subtle).Bold(true).Render("ACCURACY"),
		s.StyleCardVal.Render(fmt.Sprintf("%.1f%%", acc)),
	)
	cardCons := lipgloss.JoinVertical(lipgloss.Center,
		lipgloss.NewStyle().Foreground(theme.Subtle).Bold(true).Render("CONSISTENCY"),
		s.StyleCardVal.Render(fmt.Sprintf("%.1f%%", consistency)),
	)

	box1 := s.StyleCard.Render(cardWPM)
	box2 := s.StyleCard.Render(cardGross)
	box3 := s.StyleCard.Render(cardAcc)
	box4 := s.StyleCard.Render(cardCons)

	statRow := lipgloss.JoinHorizontal(lipgloss.Center, box1, box2, box3, box4)
	b.WriteString(statRow)
	b.WriteString("\n\n")

	// Keystrokes Detail Line
	b.WriteString(s.StyleSubtext.Render(fmt.Sprintf(
		"Keystrokes: Correct %d | Errors %d | Backspaces %d | Total %d",
		t.CorrectKeystrokes, t.ErrorKeystrokes, t.BackspaceCount, t.TotalKeystrokes,
	)))
	b.WriteString("\n\n")

	// WPM Timeline Graph
	if len(t.WPMSamples) > 1 {
		graphWidth := termWidth - 20
		if graphWidth < 30 {
			graphWidth = 30
		}
		if graphWidth > 70 {
			graphWidth = 70
		}

		graphPlot := asciigraph.Plot(
			t.WPMSamples,
			asciigraph.Height(6),
			asciigraph.Width(graphWidth),
			asciigraph.Caption("WPM Speed Timeline (Seconds)"),
		)
		graphStyle := lipgloss.NewStyle().
			Foreground(theme.Primary).
			Border(lipgloss.NormalBorder()).
			BorderForeground(theme.Dim).
			Padding(0, 1)

		b.WriteString(graphStyle.Render(graphPlot))
		b.WriteString("\n\n")
	}

	// Action Controls Footer
	footer := "[Tab / Enter] Test Again    [h] Keyboard Heatmap    [Esc] Menu    [q] Quit"
	b.WriteString(s.StyleActive.Render(footer))

	content := b.String()
	if termWidth > 0 {
		return lipgloss.PlaceHorizontal(termWidth, lipgloss.Center, content)
	}
	return content
}
