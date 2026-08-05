package engine

// AppState represents the active screen in the Bubble Tea state machine.
type AppState int

const (
	StateMenu AppState = iota
	StateTest
	StateResults
	StateHeatmap
)
