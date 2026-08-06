package engine

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/Codexia-afk/JeeraType/config"
	"github.com/Codexia-afk/JeeraType/db"
	"github.com/Codexia-afk/JeeraType/generator"
	"github.com/Codexia-afk/JeeraType/stats"
	"github.com/Codexia-afk/JeeraType/storage"
	"github.com/Codexia-afk/JeeraType/ui"
)

// Model represents top-level Elm architecture model for JeeraType.
type Model struct {
	state           AppState
	timeModes       []int
	selectedModeIdx int
	currentMode     ui.TestMode
	currentTheme    config.Theme
	ghostWPM        float64
	punctuation     bool
	numbers         bool
	isZen           bool
	codeLang        string
	profile         string
	sound           bool
	countdownCount  int // 3...2...1...0 (0 = start input)
	wordlistPath    string
	showKeys        bool
	stopOnError     bool
	suddenDeath     bool
	stdinText       string
	filePath        string
	fileOffset      int
	lastPressedRune rune
	tracker         *stats.Tracker
	width           int
	height          int
	historySaved    bool
}

// NewModel constructs an initialized Model.
func NewModel() *Model {
	return &Model{
		state:           StateMenu,
		timeModes:       []int{15, 30, 45, 60, 120},
		selectedModeIdx: 0,
		currentMode:     ui.ModeParagraphs,
		currentTheme:    config.ThemeAmber,
		ghostWPM:        0,
		punctuation:     false,
		numbers:         false,
		isZen:           false,
		codeLang:        "go",
		profile:         "default",
		sound:           false,
		countdownCount:  0,
		showKeys:        false,
		stopOnError:     false,
		suddenDeath:     false,
		width:           80,
		height:          24,
	}
}

func (m *Model) SetTheme(themeName string) {
	m.currentTheme = config.GetThemeByName(themeName)
}

func (m *Model) SetCustomTheme(theme config.Theme) {
	m.currentTheme = theme
}

func (m *Model) SetGhostWPM(wpm float64) {
	m.ghostWPM = wpm
}

func (m *Model) SetPunctuation(punc bool) {
	m.punctuation = punc
}

func (m *Model) SetNumbers(num bool) {
	m.numbers = num
}

func (m *Model) SetZen(zen bool) {
	m.isZen = zen
}

func (m *Model) SetCodeLang(lang string) {
	m.codeLang = lang
}

func (m *Model) SetProfile(prof string) {
	m.profile = prof
}

func (m *Model) SetSound(snd bool) {
	m.sound = snd
}

func (m *Model) SetWordlistPath(path string) {
	m.wordlistPath = path
}

func (m *Model) SetShowKeys(show bool) {
	m.showKeys = show
}

func (m *Model) SetStopOnError(soe bool) {
	m.stopOnError = soe
}

func (m *Model) SetSuddenDeath(sd bool) {
	m.suddenDeath = sd
}

func (m *Model) SetMode(mode ui.TestMode) {
	m.currentMode = mode
}

func (m *Model) SetStdinText(text string) {
	m.stdinText = text
	m.currentMode = ui.ModeStdin
}

func (m *Model) SetFilePath(path string) {
	m.filePath = path
	m.currentMode = ui.ModeFile
	m.fileOffset = db.GetFileOffset(path)
}

func (m *Model) SetState(st AppState) {
	m.state = st
}

// Init starts tick loops.
func (m *Model) Init() tea.Cmd {
	return tea.Batch(
		CmdTick100ms(),
		CmdTick250ms(),
		CmdTick1s(),
	)
}

// StartTest initializes a new typing test session.
func (m *Model) StartTest() {
	duration := m.timeModes[m.selectedModeIdx]
	if m.isZen {
		duration = 0
	}
	var targetText string

	if m.wordlistPath != "" {
		customText, err := generator.LoadCustomWordlist(m.wordlistPath, duration)
		if err == nil && len(customText) > 0 {
			targetText = customText
		} else {
			targetText = generator.GenerateText(duration, m.punctuation, m.numbers)
		}
	} else {
		switch m.currentMode {
		case ui.ModeCode:
			targetText = generator.GenerateLanguageCodeText(m.codeLang, duration)
		case ui.ModeAdaptive:
			weakKeys := db.GetTopWeakKeys(5)
			weakBigrams := db.GetTopSlowestBigrams(5)
			targetText = generator.GenerateAdaptiveText(weakKeys, weakBigrams, duration, m.punctuation, m.numbers)
		case ui.ModeQuotes:
			targetText = generator.GenerateQuoteText()
		case ui.ModeStdin:
			if m.stdinText != "" {
				targetText = m.stdinText
			} else {
				targetText = generator.GenerateText(duration, m.punctuation, m.numbers)
			}
		case ui.ModeFile:
			if m.filePath != "" {
				content, err := os.ReadFile(m.filePath)
				if err == nil && len(content) > 0 {
					fullText := string(content)
					if m.fileOffset < len(fullText) {
						targetText = fullText[m.fileOffset:]
					} else {
						targetText = fullText
						m.fileOffset = 0
					}
				} else {
					targetText = generator.GenerateText(duration, m.punctuation, m.numbers)
				}
			} else {
				targetText = generator.GenerateText(duration, m.punctuation, m.numbers)
			}
		default:
			targetText = generator.GenerateText(duration, m.punctuation, m.numbers)
		}
	}

	m.tracker = stats.NewTracker(targetText, duration, m.ghostWPM, m.stopOnError, m.suddenDeath, m.isZen)

	// Combo key for Personal Best checking
	puncStr := "no_punc"
	if m.punctuation {
		puncStr = "punc"
	}
	numStr := "no_num"
	if m.numbers {
		numStr = "num"
	}
	modeKey := fmt.Sprintf("%s_%ds_%s_%s", m.currentMode.String(), duration, puncStr, numStr)
	pbWPM, pbSamples := db.GetPBRace(modeKey)
	m.tracker.PreviousPBWPM = pbWPM
	m.tracker.PBWPM = pbWPM
	m.tracker.PBSamples = pbSamples

	if !m.isZen {
		m.countdownCount = 3
	} else {
		m.countdownCount = 0
	}

	m.state = StateTest
	m.historySaved = false
}

// FinishTest concludes test session.
func (m *Model) FinishTest() {
	if m.tracker == nil || m.tracker.IsFinished {
		return
	}
	m.tracker.Finish()
	m.state = StateResults

	if !m.historySaved {
		m.historySaved = true

		netWPM := m.tracker.CalculateWPM()
		grossWPM := m.tracker.CalculateGrossWPM()
		acc := m.tracker.CalculateAccuracy()
		consistency := m.tracker.CalculateConsistency()
		duration := m.tracker.DurationSec

		// Check if new Personal Best
		if netWPM > m.tracker.PreviousPBWPM && netWPM > 0 {
			m.tracker.IsNewPB = true
		}

		// Save JSONL record with Profile scope
		rec := storage.HistoryRecord{
			Timestamp:    m.tracker.EndTime,
			DurationSec:  duration,
			WPM:          netWPM,
			RawWPM:       grossWPM,
			Accuracy:     acc,
			Consistency:  consistency,
			TotalChars:   m.tracker.CursorIdx,
			CorrectChars: m.tracker.CountCorrectChars(),
			ErrorCount:   m.tracker.ErrorKeystrokes,
			Profile:      m.profile,
		}
		_ = storage.SaveRecord(rec)

		// Save SQLite run record
		modeStr := m.currentMode.String()
		_ = db.SaveTestRun(modeStr, netWPM, grossWPM, acc, consistency, duration)

		// Save PB Race Keyframes if new PB
		puncStr := "no_punc"
		if m.punctuation {
			puncStr = "punc"
		}
		numStr := "no_num"
		if m.numbers {
			numStr = "num"
		}
		modeKey := fmt.Sprintf("%s_%ds_%s_%s", modeStr, duration, puncStr, numStr)
		_ = db.SavePBRace(modeKey, netWPM, m.tracker.WPMSamples)

		// Save File Offset Progress
		if m.currentMode == ui.ModeFile && m.filePath != "" {
			newOffset := m.fileOffset + m.tracker.CursorIdx
			_ = db.SaveFileOffset(m.filePath, newOffset)
		}
	}
}

// Update handles state transitions and input events.
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case Tick100msMsg:
		if m.state == StateTest && m.tracker != nil && m.tracker.IsStarted && !m.tracker.IsFinished {
			m.tracker.UpdateGhostPacer()
		}
		return m, CmdTick100ms()

	case Tick250msMsg:
		if m.state == StateTest && m.tracker != nil && m.tracker.IsStarted && !m.tracker.IsFinished && !m.isZen {
			if m.tracker.RemainingSeconds() <= 0 {
				m.FinishTest()
			}
		}
		return m, CmdTick250ms()

	case Tick1sMsg:
		if m.state == StateTest {
			if m.countdownCount > 0 {
				m.countdownCount--
			} else if m.tracker != nil && m.tracker.IsStarted && !m.tracker.IsFinished {
				m.tracker.SampleWPM()
			}
		}
		return m, CmdTick1s()

	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}

		switch m.state {
		case StateMenu:
			return m.updateMenu(msg)
		case StateTest:
			return m.updateTest(msg)
		case StateResults:
			return m.updateResults(msg)
		case StateHeatmap:
			return m.updateHeatmap(msg)
		}
	}

	return m, nil
}

func (m *Model) updateMenu(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "left", "h":
		if m.selectedModeIdx > 0 {
			m.selectedModeIdx--
		}
	case "right", "l":
		if m.selectedModeIdx < len(m.timeModes)-1 {
			m.selectedModeIdx++
		}
	case "1":
		m.selectedModeIdx = 0
	case "2":
		m.selectedModeIdx = 1
	case "3":
		m.selectedModeIdx = 2
	case "4":
		m.selectedModeIdx = 3
	case "5":
		m.selectedModeIdx = 4
	case "p":
		m.punctuation = !m.punctuation
	case "n":
		m.numbers = !m.numbers
	case "z":
		m.isZen = !m.isZen
	case "m":
		switch m.currentMode {
		case ui.ModeParagraphs:
			m.currentMode = ui.ModeCode
		case ui.ModeCode:
			m.currentMode = ui.ModeAdaptive
		case ui.ModeAdaptive:
			m.currentMode = ui.ModeQuotes
		default:
			m.currentMode = ui.ModeParagraphs
		}
	case "t":
		themes := config.AvailableThemes
		for i, th := range themes {
			if th.Name == m.currentTheme.Name {
				m.currentTheme = themes[(i+1)%len(themes)]
				break
			}
		}
	case "g":
		switch m.ghostWPM {
		case 0:
			m.ghostWPM = 60
		case 60:
			m.ghostWPM = 80
		case 80:
			m.ghostWPM = 100
		case 100:
			m.ghostWPM = 120
		default:
			m.ghostWPM = 0
		}
	case "v":
		m.showKeys = !m.showKeys
	case "k":
		m.state = StateHeatmap
	case "enter", " ":
		m.StartTest()
	case "esc", "q":
		return m, tea.Quit
	}
	return m, nil
}

func (m *Model) updateTest(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "esc":
		if m.isZen || (m.tracker != nil && m.tracker.IsStarted) {
			m.FinishTest()
		} else {
			m.state = StateMenu
		}
		return m, nil
	case "tab", "ctrl+r":
		m.StartTest()
		return m, nil
	}

	// Block text input during 3..2..1.. countdown
	if m.countdownCount > 0 {
		return m, nil
	}

	switch msg.String() {
	case "backspace":
		m.tracker.Backspace()
		return m, nil
	case "enter":
		if m.currentMode == ui.ModeCode {
			m.tracker.RecordRune('\n')
			if m.tracker.IsFinished {
				m.FinishTest()
			}
			return m, nil
		}
	}

	runes := msg.Runes
	if len(runes) > 0 {
		m.lastPressedRune = runes[0]
		for _, r := range runes {
			currCursor := m.tracker.CursorIdx
			m.tracker.RecordRune(r)

			// Sound feedback on error or keypress
			if m.sound {
				fmt.Print("\a")
			}

			// Death Mode: any typo resets test immediately
			if m.suddenDeath && m.tracker.CursorIdx <= currCursor && m.tracker.ErrorKeystrokes > 0 {
				m.StartTest()
				return m, nil
			}

			if m.tracker.IsFinished && !m.isZen {
				m.FinishTest()
				break
			}
		}
	} else if msg.String() == "space" {
		m.lastPressedRune = ' '
		m.tracker.RecordRune(' ')
		if m.sound {
			fmt.Print("\a")
		}
		if m.tracker.IsFinished && !m.isZen {
			m.FinishTest()
		}
	}

	return m, nil
}

func (m *Model) updateResults(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "tab", "enter":
		m.StartTest()
	case "h":
		m.state = StateHeatmap
	case "esc":
		m.state = StateMenu
	case "q":
		return m, tea.Quit
	}
	return m, nil
}

func (m *Model) updateHeatmap(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "esc", "q", "enter", "backspace":
		m.state = StateMenu
	}
	return m, nil
}

// View renders active screen.
func (m *Model) View() string {
	switch m.state {
	case StateMenu:
		return ui.RenderMenuView(m.timeModes, m.selectedModeIdx, m.currentMode, m.currentTheme, m.ghostWPM, m.showKeys, m.punctuation, m.numbers, m.isZen, m.width)
	case StateTest:
		if m.countdownCount > 0 {
			return ui.RenderCountdownView(m.countdownCount, m.currentTheme, m.width)
		}
		return ui.RenderTestView(m.tracker, m.currentTheme, m.showKeys, m.lastPressedRune, m.width)
	case StateResults:
		return ui.RenderResultsView(m.tracker, m.currentTheme, m.width)
	case StateHeatmap:
		return ui.RenderKeyboardHeatmap(m.currentTheme, m.width)
	default:
		return "Unknown State"
	}
}
