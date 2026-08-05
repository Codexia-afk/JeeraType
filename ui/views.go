package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/guptarohit/asciigraph"
	"jeeratype/config"
	"jeeratype/stats"
)

type TestMode int

const (
	ModeParagraphs TestMode = iota
	ModeCode
	ModeAdaptive
)

func (m TestMode) String() string {
	switch m {
	case ModeParagraphs:
		return "Paragraphs"
	case ModeCode:
		return "Code Mode"
	case ModeAdaptive:
		return "Adaptive Weakness"
	default:
		return "Paragraphs"
	}
}

// RenderMenuView renders initial options: mode, duration, theme, ghost pacer, hotkeys.
func RenderMenuView(
	timeModes []int,
	selectedModeIdx int,
	currentMode TestMode,
	currentTheme config.Theme,
	ghostWPM float64,
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
	modes := []TestMode{ModeParagraphs, ModeCode, ModeAdaptive}
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

	// Duration Selector Pills
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
	b.WriteString("\n\n")

	// Ghost Pacer & Theme Status Line
	ghostStr := "Disabled"
	if ghostWPM > 0 {
		ghostStr = fmt.Sprintf("%.0f WPM 👻", ghostWPM)
	}
	infoLine := fmt.Sprintf("Theme: [%s]    Ghost Pacer: [%s]", strings.ToUpper(currentTheme.Name), ghostStr)
	b.WriteString(s.StyleSubtext.Render(infoLine))
	b.WriteString("\n\n")

	// Controls Guide
	helpText := "Controls:\n" +
		"  [m] Toggle Mode (Paragraph / Code / Adaptive)    [t] Cycle Theme\n" +
		"  [g] Toggle Ghost Pacer (Off/60/80/100/120 WPM)   [h] Keyboard Heatmap\n" +
		"  [← / → / 1-5] Select Duration                   [Enter / Space] Start Test"
	b.WriteString(s.StyleHelp.Render(helpText))

	content := b.String()
	if termWidth > 0 {
		return lipgloss.PlaceHorizontal(termWidth, lipgloss.Center, content)
	}
	return content
}

// RenderTestView renders live test with Ghost Pacer and target word wrap.
func RenderTestView(t *stats.Tracker, theme config.Theme, termWidth int) string {
	s := NewDynamicStyles(theme)
	var b strings.Builder

	// Header Bar: Time left, Live WPM, Accuracy, Ghost Pacer
	remSec := t.RemainingSeconds()
	liveWPM := t.CalculateWPM()
	liveAcc := t.CalculateAccuracy()

	timeBadge := s.StyleActive.Render(fmt.Sprintf(" ⏱ %ds ", remSec))
	wpmBadge := s.StyleCard.Render(fmt.Sprintf("WPM: %.0f", liveWPM))
	accBadge := s.StyleCard.Render(fmt.Sprintf("ACC: %.0f%%", liveAcc))

	headerItems := []string{timeBadge, "   ", wpmBadge, "   ", accBadge}
	if t.TargetWPM > 0 {
		ghostBadge := s.StyleCard.Render(fmt.Sprintf("GHOST: %.0f WPM", t.TargetWPM))
		headerItems = append(headerItems, "   ", ghostBadge)
	}

	b.WriteString(lipgloss.JoinHorizontal(lipgloss.Center, headerItems...))
	b.WriteString("\n\n")

	// Target text container calculation
	maxWidth := termWidth - 10
	if maxWidth < 40 {
		maxWidth = 40
	}

	var textBuilder strings.Builder
	for i, r := range t.TargetRunes {
		charStr := string(r)

		// Ghost Pacer cursor check
		isGhost := (t.TargetWPM > 0 && i == t.GhostCursorIdx && i != t.CursorIdx)

		if i == t.CursorIdx {
			if r == ' ' {
				textBuilder.WriteString(s.StyleCursor.Render(" "))
			} else {
				textBuilder.WriteString(s.StyleCursor.Render(charStr))
			}
		} else if isGhost {
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

	b.WriteString(s.StyleSubtext.Render("[Tab] Restart Test    [Esc] Return to Menu"))

	content := b.String()
	if termWidth > 0 {
		return lipgloss.PlaceHorizontal(termWidth, lipgloss.Center, content)
	}
	return content
}

// RenderResultsView renders post-game stats, asciigraph WPM timeline, and slowest key transitions breakdown.
func RenderResultsView(t *stats.Tracker, theme config.Theme, termWidth int) string {
	s := NewDynamicStyles(theme)
	var b strings.Builder

	// Logo Header
	b.WriteString(RenderLogo(theme))
	b.WriteString("\n\n")

	// Stat Cards
	netWPM := t.CalculateWPM()
	rawWPM := t.CalculateRawWPM()
	acc := t.CalculateAccuracy()
	consistency := t.CalculateConsistency()

	cardWPM := lipgloss.JoinVertical(lipgloss.Center,
		lipgloss.NewStyle().Foreground(theme.Subtle).Bold(true).Render("NET WPM"),
		s.StyleCardVal.Render(fmt.Sprintf("%.1f", netWPM)),
	)
	cardRaw := lipgloss.JoinVertical(lipgloss.Center,
		lipgloss.NewStyle().Foreground(theme.Subtle).Bold(true).Render("RAW WPM"),
		s.StyleCardVal.Render(fmt.Sprintf("%.1f", rawWPM)),
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
	box2 := s.StyleCard.Render(cardRaw)
	box3 := s.StyleCard.Render(cardAcc)
	box4 := s.StyleCard.Render(cardCons)

	statRow := lipgloss.JoinHorizontal(lipgloss.Center, box1, box2, box3, box4)
	b.WriteString(statRow)
	b.WriteString("\n\n")

	// Keystrokes & Slowest Keys Breakdown
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
