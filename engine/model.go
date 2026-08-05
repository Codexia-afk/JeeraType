package engine

import (
	"time"

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
		width:           80,
		height:          24,
	}
}

func (m *Model) SetTheme(themeName string) {
	m.currentTheme = config.GetThemeByName(themeName)
}

func (m *Model) SetGhostWPM(wpm float64) {
	m.ghostWPM = wpm
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
	var targetText string

	switch m.currentMode {
	case ui.ModeCode:
		targetText = generator.GenerateCodeText(duration)
	case ui.ModeAdaptive:
		weakKeys := db.GetTopWeakKeys(5)
		weakBigrams := db.GetTopSlowestBigrams(5)
		targetText = generator.GenerateAdaptiveText(weakKeys, weakBigrams, duration)
	default:
		targetText = generator.GenerateText(duration)
	}

	m.tracker = stats.NewTracker(targetText, duration, m.ghostWPM)
	m.state = StateTest
	m.historySaved = false
}

// FinishTest concludes test session and records metrics in JSONL and SQLite database.
func (m *Model) FinishTest() {
	if m.tracker == nil || m.tracker.IsFinished {
		return
	}
	m.tracker.Finish()
	m.state = StateResults

	if !m.historySaved {
		m.historySaved = true

		netWPM := m.tracker.CalculateWPM()
		rawWPM := m.tracker.CalculateRawWPM()
		acc := m.tracker.CalculateAccuracy()
		consistency := m.tracker.CalculateConsistency()

		// Save JSONL record
		rec := storage.HistoryRecord{
			Timestamp:    time.Now(),
			DurationSec:  m.tracker.DurationSec,
			WPM:          netWPM,
			RawWPM:       rawWPM,
			Accuracy:     acc,
			Consistency:  consistency,
			TotalChars:   m.tracker.CursorIdx,
			CorrectChars: m.tracker.CountCorrectChars(),
			ErrorCount:   m.tracker.ErrorKeystrokes,
		}
		_ = storage.SaveRecord(rec)

		// Save SQLite run record
		_ = db.SaveTestRun(m.currentMode.String(), netWPM, rawWPM, acc, consistency, m.tracker.DurationSec)
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
		if m.state == StateTest && m.tracker != nil && m.tracker.IsStarted && !m.tracker.IsFinished {
			if m.tracker.RemainingSeconds() <= 0 {
				m.FinishTest()
			}
		}
		return m, CmdTick250ms()

	case Tick1sMsg:
		if m.state == StateTest && m.tracker != nil && m.tracker.IsStarted && !m.tracker.IsFinished {
			m.tracker.SampleWPM()
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
	case "m":
		// Cycle typing mode
		switch m.currentMode {
		case ui.ModeParagraphs:
			m.currentMode = ui.ModeCode
		case ui.ModeCode:
			m.currentMode = ui.ModeAdaptive
		default:
			m.currentMode = ui.ModeParagraphs
		}
	case "t":
		// Cycle themes
		themes := config.AvailableThemes
		for i, th := range themes {
			if th.Name == m.currentTheme.Name {
				m.currentTheme = themes[(i+1)%len(themes)]
				break
			}
		}
	case "g":
		// Cycle ghost pacer speed
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
		m.state = StateMenu
		return m, nil
	case "tab":
		m.StartTest()
		return m, nil
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
		for _, r := range runes {
			m.tracker.RecordRune(r)
			if m.tracker.IsFinished {
				m.FinishTest()
				break
			}
		}
	} else if msg.String() == "space" {
		m.tracker.RecordRune(' ')
		if m.tracker.IsFinished {
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

// View renders the active screen state.
func (m *Model) View() string {
	switch m.state {
	case StateMenu:
		return ui.RenderMenuView(m.timeModes, m.selectedModeIdx, m.currentMode, m.currentTheme, m.ghostWPM, m.width)
	case StateTest:
		return ui.RenderTestView(m.tracker, m.currentTheme, m.width)
	case StateResults:
		return ui.RenderResultsView(m.tracker, m.currentTheme, m.width)
	case StateHeatmap:
		return ui.RenderKeyboardHeatmap(m.currentTheme, m.width)
	default:
		return "Unknown State"
	}
}
