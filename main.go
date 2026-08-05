package main

import (
	"flag"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"jeeratype/config"
	"jeeratype/db"
	"jeeratype/engine"
	"jeeratype/ui"
)

func main() {
	showStats := flag.Bool("stats", false, "Render ASCII Keyboard Heatmap & historical stats in terminal")
	flag.BoolVar(showStats, "s", false, "Render ASCII Keyboard Heatmap & historical stats in terminal (shorthand)")
	themeName := flag.String("theme", "amber", "Set UI theme (amber, catppuccin, nord, dracula, matrix)")
	flag.StringVar(themeName, "t", "amber", "Set UI theme (shorthand)")
	ghostWPM := flag.Float64("ghost", 0, "Set Ghost Pacer target WPM (e.g. 60, 80, 100)")
	flag.Float64Var(ghostWPM, "g", 0, "Set Ghost Pacer target WPM (shorthand)")
	modeName := flag.String("mode", "paragraphs", "Set initial mode (paragraphs, code, adaptive)")
	flag.StringVar(modeName, "m", "paragraphs", "Set initial mode (shorthand)")

	flag.Parse()

	// Initialize SQLite database
	if err := db.InitDB(); err != nil {
		fmt.Printf("Warning: Database initialization error: %v\n", err)
	}
	defer db.CloseDB()

	// Handle --stats CLI flag mode directly
	if *showStats {
		th := config.GetThemeByName(*themeName)
		fmt.Println(ui.RenderKeyboardHeatmap(th, 80))
		os.Exit(0)
	}

	model := engine.NewModel()
	model.SetTheme(*themeName)
	if *ghostWPM > 0 {
		model.SetGhostWPM(*ghostWPM)
	}

	p := tea.NewProgram(model, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running JeeraType: %v\n", err)
		os.Exit(1)
	}
}
