package cmd

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/Codexia-afk/JeeraType/config"
)

// RenderThemeList renders color swatches for all registered themes.
func RenderThemeList() string {
	var b strings.Builder

	titleStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#F59E0B")).
		Bold(true).
		Padding(0, 1)

	b.WriteString(titleStyle.Render("🎨 JeeraType — Registered Color Themes"))
	b.WriteString("\n\n")

	for _, th := range config.AvailableThemes {
		swatchPrimary := lipgloss.NewStyle().Foreground(th.Primary).Render("██")
		swatchSecondary := lipgloss.NewStyle().Foreground(th.Secondary).Render("██")
		swatchSuccess := lipgloss.NewStyle().Foreground(th.Success).Render("██")
		swatchError := lipgloss.NewStyle().Foreground(th.Error).Render("██")
		swatchDim := lipgloss.NewStyle().Foreground(th.Dim).Render("██")
		swatchSubtle := lipgloss.NewStyle().Foreground(th.Subtle).Render("██")

		swatchRow := fmt.Sprintf("%s %s %s %s %s %s",
			swatchPrimary, swatchSecondary, swatchSuccess, swatchError, swatchDim, swatchSubtle)

		nameStyle := lipgloss.NewStyle().
			Foreground(th.Primary).
			Bold(true).
			Width(16)

		b.WriteString(fmt.Sprintf("%s  %s\n", nameStyle.Render(th.Name), swatchRow))
	}

	b.WriteString("\nUsage: jeeratype --theme <name>  OR  jeeratype theme preview <name>\n")
	return b.String()
}

// RenderThemePreview renders a live preview of a specific theme.
func RenderThemePreview(name string) string {
	th := config.GetThemeByName(name)
	var b strings.Builder

	boxStyle := lipgloss.NewStyle().
		Background(th.Background).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(th.Primary).
		Padding(1, 3).
		Width(65)

	header := lipgloss.NewStyle().
		Foreground(th.Primary).
		Bold(true).
		Render(fmt.Sprintf("🎨 Theme Preview: %s", strings.ToUpper(th.Name)))

	b.WriteString(header)
	b.WriteString("\n\n")

	// Sample typing passage render
	correctText := lipgloss.NewStyle().Foreground(th.Success).Render("The quick brown fox ")
	cursorText := lipgloss.NewStyle().Background(th.Primary).Foreground(th.Background).Bold(true).Render("j")
	untypedText := lipgloss.NewStyle().Foreground(th.Dim).Render("umps over the lazy dog.")
	errorText := lipgloss.NewStyle().Foreground(th.Error).Underline(true).Render(" typo")

	b.WriteString("Sample Text: ")
	b.WriteString(correctText)
	b.WriteString(cursorText)
	b.WriteString(untypedText)
	b.WriteString(errorText)
	b.WriteString("\n\n")

	// Swatch Palette Breakdown
	p := lipgloss.NewStyle().Foreground(th.Primary).Render("██ Primary")
	sec := lipgloss.NewStyle().Foreground(th.Secondary).Render("██ Secondary")
	suc := lipgloss.NewStyle().Foreground(th.Success).Render("██ Success")
	err := lipgloss.NewStyle().Foreground(th.Error).Render("██ Error")
	dim := lipgloss.NewStyle().Foreground(th.Dim).Render("██ Dim")

	b.WriteString(fmt.Sprintf("Palette: %s  %s  %s  %s  %s\n", p, sec, suc, err, dim))

	return boxStyle.Render(b.String())
}
