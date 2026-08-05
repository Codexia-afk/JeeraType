package engine

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// Tick100msMsg triggers smooth Ghost Pacer cursor updates.
type Tick100msMsg time.Time

// Tick250msMsg triggers UI re-renders and timer countdown updates.
type Tick250msMsg time.Time

// Tick1sMsg triggers sampling WPM for timeline graphing.
type Tick1sMsg time.Time

// CmdTick100ms schedules a 100ms tick message.
func CmdTick100ms() tea.Cmd {
	return tea.Tick(100*time.Millisecond, func(t time.Time) tea.Msg {
		return Tick100msMsg(t)
	})
}

// CmdTick250ms schedules a 250ms tick message.
func CmdTick250ms() tea.Cmd {
	return tea.Tick(250*time.Millisecond, func(t time.Time) tea.Msg {
		return Tick250msMsg(t)
	})
}

// CmdTick1s schedules a 1s tick message.
func CmdTick1s() tea.Cmd {
	return tea.Tick(1*time.Second, func(t time.Time) tea.Msg {
		return Tick1sMsg(t)
	})
}
