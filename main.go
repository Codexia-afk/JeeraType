package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"golang.org/x/term"
	"github.com/Codexia-afk/JeeraType/cmd"
	"github.com/Codexia-afk/JeeraType/config"
	"github.com/Codexia-afk/JeeraType/db"
	"github.com/Codexia-afk/JeeraType/engine"
	"github.com/Codexia-afk/JeeraType/generator"
	"github.com/Codexia-afk/JeeraType/ui"
)

var Version = "v2.5.0"

func main() {
	// Subcommand: `jeeratype theme list` & `jeeratype theme preview`
	if len(os.Args) > 1 {
		firstArg := strings.ToLower(os.Args[1])
		if firstArg == "theme" || firstArg == "themes" {
			if len(os.Args) > 2 && strings.ToLower(os.Args[2]) == "preview" {
				themeName := "default"
				if len(os.Args) > 3 {
					themeName = os.Args[3]
				}
				fmt.Println(cmd.RenderThemePreview(themeName))
			} else {
				fmt.Println(cmd.RenderThemeList())
			}
			os.Exit(0)
		}
	}

	// Subcommand: `jeeratype stats` (--heatmap, --leaderboard)
	if len(os.Args) > 1 && os.Args[1] == "stats" {
		_ = db.InitDB()
		defer db.CloseDB()

		themeFlag := "default"
		showHeatmap := false
		showLeaderboard := false

		for i, arg := range os.Args {
			if arg == "--theme" && i+1 < len(os.Args) {
				themeFlag = os.Args[i+1]
			}
			if arg == "--heatmap" {
				showHeatmap = true
			}
			if arg == "--leaderboard" {
				showLeaderboard = true
			}
		}
		cmd.RunStatsSubcommand(themeFlag, showHeatmap, showLeaderboard)
		os.Exit(0)
	}

	// Subcommand: `jeeratype export-replay <run-id>`
	if len(os.Args) > 1 && os.Args[1] == "export-replay" {
		targetPath := ""
		if len(os.Args) > 2 {
			targetPath = os.Args[2]
		}
		if err := engine.ExportReplay(targetPath); err != nil {
			fmt.Fprintf(os.Stderr, "Error exporting replay: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("✅ Replay exported successfully!")
		os.Exit(0)
	}

	// Subcommand: `jeeratype race replay.json`
	if len(os.Args) > 1 && os.Args[1] == "race" {
		if len(os.Args) < 3 {
			fmt.Fprintf(os.Stderr, "Usage: jeeratype race <replay.json>\n")
			os.Exit(1)
		}
		replay, err := engine.LoadReplay(os.Args[2])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error loading replay: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("🏎️  Racing against recorded replay (WPM: %.1f)...\n", replay.WPM)
		os.Exit(0)
	}

	// Exact Approved Flag Definitions
	punctuation := flag.Bool("punctuation", false, "realistic capitalization, commas, periods")
	numbers := flag.Bool("numbers", false, "inject digit tokens")
	zen := flag.Bool("zen", false, "infinite mode, no timer/cap, quit with Esc")
	death := flag.Bool("death", false, "sudden-death: one typo resets the test")
	wordlistPath := flag.String("wordlist", "", "custom word list, one word per line, ≥50 words")
	themeName := flag.String("theme", "default", "default | dracula | nord | solarized | jewel | sunset | forest | neon | vintage | mono")
	profile := flag.String("profile", "default", "scopes history/streak/PB to a named local profile")
	sound := flag.Bool("sound", false, "optional terminal bell/click feedback")
	modeName := flag.String("mode", "time", "base test type (time|code)")
	codeLang := flag.String("lang", "go", "only used with --mode code (python|js|go)")
	showVersion := flag.Bool("version", false, "print installed version")

	flag.Parse()

	if *showVersion {
		fmt.Printf("jeeratype version %s\n", Version)
		os.Exit(0)
	}

	// Initialize SQLite database
	if err := db.InitDB(); err != nil {
		fmt.Fprintf(os.Stderr, "Warning: Database initialization error: %v\n", err)
	}
	defer db.CloseDB()

	// Validate Custom Wordlist if specified
	if *wordlistPath != "" {
		_, err := generator.LoadCustomWordlist(*wordlistPath, 30)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	}

	// Load JSON app configuration
	appCfg := config.LoadConfig()
	if *themeName == "default" && appCfg.Theme != "" {
		*themeName = appCfg.Theme
	}

	model := engine.NewModel()
	model.SetTheme(*themeName)
	model.SetPunctuation(*punctuation)
	model.SetNumbers(*numbers)
	model.SetZen(*zen)
	model.SetCodeLang(*codeLang)
	model.SetProfile(*profile)
	model.SetSound(*sound)

	if *wordlistPath != "" {
		model.SetWordlistPath(*wordlistPath)
	}

	if *death {
		model.SetSuddenDeath(true)
	}

	mName := strings.ToLower(*modeName)
	switch mName {
	case "code":
		model.SetMode(ui.ModeCode)
	default:
		model.SetMode(ui.ModeParagraphs)
	}

	// Save active settings to config.json
	appCfg.Theme = *themeName
	appCfg.Punctuation = *punctuation
	appCfg.Numbers = *numbers
	appCfg.Sound = *sound
	appCfg.DefaultProfile = *profile
	_ = config.SaveConfig(appCfg)

	// Check STDIN Pipe Input (e.g. cat file.txt | jeeratype)
	if !term.IsTerminal(int(os.Stdin.Fd())) {
		pipedBytes, err := io.ReadAll(os.Stdin)
		if err == nil && len(pipedBytes) > 0 {
			pipedStr := strings.TrimSpace(string(pipedBytes))
			if len(pipedStr) > 0 {
				model.SetStdinText(pipedStr)
			}
		}
	} else if flag.NArg() > 0 {
		// Positional argument file path (e.g. jeeratype /path/to/book.txt)
		filePath := flag.Arg(0)
		model.SetFilePath(filePath)
	}

	p := tea.NewProgram(model, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error running JeeraType: %v\n", err)
		os.Exit(1)
	}
}
