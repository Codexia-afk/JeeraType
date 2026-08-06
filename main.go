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
	"github.com/Codexia-afk/JeeraType/export"
	"github.com/Codexia-afk/JeeraType/generator"
	"github.com/Codexia-afk/JeeraType/storage"
	"github.com/Codexia-afk/JeeraType/ui"
)

func main() {
	// Subcommand: `jeeratype theme` (list | preview <name>)
	if len(os.Args) > 1 && os.Args[1] == "theme" {
		if len(os.Args) > 2 && os.Args[2] == "preview" {
			themeName := "amber"
			if len(os.Args) > 3 {
				themeName = os.Args[3]
			}
			fmt.Println(cmd.RenderThemePreview(themeName))
		} else {
			fmt.Println(cmd.RenderThemeList())
		}
		os.Exit(0)
	}

	// Subcommand: `jeeratype stats` (--heatmap, --leaderboard)
	if len(os.Args) > 1 && os.Args[1] == "stats" {
		_ = db.InitDB()
		defer db.CloseDB()

		themeFlag := "amber"
		showHeatmap := false
		showLeaderboard := false

		for i, arg := range os.Args {
			if (arg == "--theme" || arg == "-t") && i+1 < len(os.Args) {
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

	showStats := flag.Bool("stats", false, "Render ASCII Keyboard Heatmap & historical stats in terminal")
	flag.BoolVar(showStats, "s", false, "Render ASCII Keyboard Heatmap (shorthand)")

	punctuation := flag.Bool("punctuation", false, "Inject punctuation into generated text stream")
	flag.BoolVar(punctuation, "p", false, "Inject punctuation (shorthand)")

	numbers := flag.Bool("numbers", false, "Inject numbers into generated text stream")
	flag.BoolVar(numbers, "n", false, "Inject numbers (shorthand)")

	isZen := flag.Bool("zen", false, "Enable Zen mode (infinite typing stream, no timer)")
	flag.BoolVar(isZen, "z", false, "Enable Zen mode (shorthand)")

	codeLang := flag.String("lang", "go", "Set programming language for Code mode (python, js, go)")

	deathMode := flag.Bool("death", false, "Enable Death mode (typo resets test immediately)")
	flag.BoolVar(deathMode, "d", false, "Enable Death mode (shorthand)")

	profile := flag.String("profile", "default", "Scope stats/history to a local user profile")

	sound := flag.Bool("sound", false, "Enable audio/terminal bell feedback on keypress/error")

	wordlistPath := flag.String("wordlist", "", "Path to custom wordlist text file (must contain at least 50 words)")

	themeName := flag.String("theme", "amber", "Set UI theme (amber, jewel, sunset, forest, neon, vintage, mono, cyberpunk, tokyonight, monokai, rose-pine, synthwave, dracula, nord, solarized, catppuccin, matrix, gruvbox)")
	flag.StringVar(themeName, "t", "amber", "Set UI theme (shorthand)")

	tomlThemePath := flag.String("config-theme", "", "Path to custom theme file")

	ghostWPM := flag.Float64("ghost", 0, "Set Ghost Pacer target WPM (e.g. 60, 80, 100)")
	flag.Float64Var(ghostWPM, "g", 0, "Set Ghost Pacer target WPM (shorthand)")

	modeName := flag.String("mode", "paragraphs", "Set initial mode (paragraphs, code, adaptive, quotes)")
	flag.StringVar(modeName, "m", "paragraphs", "Set initial mode (shorthand)")

	quotesMode := flag.Bool("quotes", false, "Launch directly in quotes mode")

	showKeys := flag.Bool("showkeys", false, "Enable live visual keyboard overlay")
	flag.BoolVar(showKeys, "sk", false, "Enable live visual keyboard overlay (shorthand)")

	stopOnError := flag.Bool("stop-on-error", false, "Enable Stop-On-Error mode")
	flag.BoolVar(stopOnError, "soe", false, "Enable Stop-On-Error mode (shorthand)")

	suddenDeath := flag.Bool("sudden-death", false, "Enable Sudden Death mode")
	flag.BoolVar(suddenDeath, "sd", false, "Enable Sudden Death mode (shorthand)")

	exportCSV := flag.Bool("csv", false, "Export history to CSV and exit")
	exportJSON := flag.Bool("json", false, "Export history to JSON and exit")

	flag.Parse()

	// Initialize SQLite database
	if err := db.InitDB(); err != nil {
		fmt.Fprintf(os.Stderr, "Warning: Database initialization error: %v\n", err)
	}
	defer db.CloseDB()

	// Handle CSV / JSON Exporters
	if *exportCSV || *exportJSON {
		records, err := storage.LoadHistory()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error loading history: %v\n", err)
			os.Exit(1)
		}
		if *exportCSV {
			_ = export.ExportCSV(records, os.Stdout)
		} else if *exportJSON {
			_ = export.ExportJSON(records, os.Stdout)
		}
		os.Exit(0)
	}

	// Handle --stats CLI flag mode directly
	if *showStats {
		th := config.GetThemeByName(*themeName)
		if *tomlThemePath != "" {
			if t, err := config.LoadThemeFromTOML(*tomlThemePath); err == nil {
				th = t
			}
		}
		fmt.Println(ui.RenderKeyboardHeatmap(th, 80))
		os.Exit(0)
	}

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
	if *themeName == "amber" && appCfg.Theme != "" {
		*themeName = appCfg.Theme
	}

	model := engine.NewModel()
	model.SetTheme(*themeName)
	if *tomlThemePath != "" {
		if t, err := config.LoadThemeFromTOML(*tomlThemePath); err == nil {
			model.SetCustomTheme(t)
		}
	}

	if *ghostWPM > 0 {
		model.SetGhostWPM(*ghostWPM)
	}
	model.SetPunctuation(*punctuation)
	model.SetNumbers(*numbers)
	model.SetZen(*isZen)
	model.SetCodeLang(*codeLang)
	model.SetProfile(*profile)
	model.SetSound(*sound)

	if *wordlistPath != "" {
		model.SetWordlistPath(*wordlistPath)
	}
	model.SetShowKeys(*showKeys)
	model.SetStopOnError(*stopOnError)

	if *deathMode || *suddenDeath {
		model.SetSuddenDeath(true)
	}

	if *quotesMode {
		model.SetMode(ui.ModeQuotes)
	} else {
		mName := strings.ToLower(*modeName)
		switch mName {
		case "code":
			model.SetMode(ui.ModeCode)
		case "adaptive":
			model.SetMode(ui.ModeAdaptive)
		case "quotes":
			model.SetMode(ui.ModeQuotes)
		default:
			model.SetMode(ui.ModeParagraphs)
		}
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
