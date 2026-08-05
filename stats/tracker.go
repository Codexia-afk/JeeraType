package stats

import (
	"math"
	"time"

	"jeeratype/db"
)

type CharStatus int

const (
	StatusUntyped CharStatus = iota
	StatusCorrect
	StatusIncorrect
)

type BigramLatency struct {
	Bigram    string
	LatencyMs int64
}

// Tracker handles real-time typing metrics, latency tracking, and Ghost Pacer calculations.
type Tracker struct {
	TargetText        string
	TargetRunes       []rune
	CharStates        []CharStatus
	CursorIdx         int
	StartTime         time.Time
	EndTime           time.Time
	LastKeyTime       time.Time
	LastKeyRune       rune
	IsStarted         bool
	IsFinished        bool
	TotalKeystrokes   int
	CorrectKeystrokes int
	ErrorKeystrokes   int
	BackspaceCount    int
	WPMSamples        []float64
	DurationSec       int
	TargetWPM         float64 // Ghost Pacer Target (0 = disabled)
	GhostCursorIdx    int     // Position of Ghost Pacer
	RunBigramLatencies []BigramLatency
	MissedKeysMap     map[rune]int
}

// NewTracker creates a new typing tracker initialized with target text, duration, and target WPM.
func NewTracker(targetText string, durationSec int, targetWPM float64) *Tracker {
	runes := []rune(targetText)
	states := make([]CharStatus, len(runes))
	for i := range states {
		states[i] = StatusUntyped
	}
	return &Tracker{
		TargetText:        targetText,
		TargetRunes:       runes,
		CharStates:        states,
		CursorIdx:         0,
		DurationSec:       durationSec,
		TargetWPM:         targetWPM,
		WPMSamples:        make([]float64, 0, durationSec),
		RunBigramLatencies: make([]BigramLatency, 0),
		MissedKeysMap:     make(map[rune]int),
	}
}

// Start begins tracking time on the first keystroke.
func (t *Tracker) Start() {
	if !t.IsStarted {
		now := time.Now()
		t.IsStarted = true
		t.StartTime = now
		t.LastKeyTime = now
	}
}

// RecordRune processes typed character input and records latency.
func (t *Tracker) RecordRune(r rune) {
	if t.IsFinished || t.CursorIdx >= len(t.TargetRunes) {
		return
	}
	now := time.Now()
	if !t.IsStarted {
		t.Start()
		now = t.StartTime
	}

	t.TotalKeystrokes++

	// Latency calculation
	var latencyMs int64
	if !t.LastKeyTime.IsZero() {
		latencyMs = now.Sub(t.LastKeyTime).Milliseconds()
		if latencyMs > 2000 { // Cap idle latency at 2s
			latencyMs = 2000
		}
	}
	t.LastKeyTime = now

	expected := t.TargetRunes[t.CursorIdx]
	isError := (r != expected)

	if !isError {
		t.CharStates[t.CursorIdx] = StatusCorrect
		t.CorrectKeystrokes++
	} else {
		t.CharStates[t.CursorIdx] = StatusIncorrect
		t.ErrorKeystrokes++
		t.MissedKeysMap[expected]++
	}

	// Save SQLite key and bigram metrics
	keyStr := string(expected)
	_ = db.RecordKeyHit(keyStr, isError, latencyMs)

	if t.LastKeyRune != 0 {
		bg := string([]rune{t.LastKeyRune, expected})
		_ = db.RecordBigramHit(bg, latencyMs)
		t.RunBigramLatencies = append(t.RunBigramLatencies, BigramLatency{
			Bigram:    bg,
			LatencyMs: latencyMs,
		})
	}
	t.LastKeyRune = expected

	t.CursorIdx++
	if t.CursorIdx >= len(t.TargetRunes) {
		t.Finish()
	}
}

// UpdateGhostPacer updates position of the Ghost Pacer cursor.
func (t *Tracker) UpdateGhostPacer() {
	if t.TargetWPM <= 0 || !t.IsStarted || t.IsFinished {
		return
	}
	elapsed := t.ElapsedSeconds()
	// Target WPM * 5 chars/word / 60 seconds
	charsPerSec := (t.TargetWPM * 5.0) / 60.0
	ghostIdx := int(elapsed * charsPerSec)
	if ghostIdx > len(t.TargetRunes) {
		ghostIdx = len(t.TargetRunes)
	}
	t.GhostCursorIdx = ghostIdx
}

// Backspace handles deleting the previously typed character.
func (t *Tracker) Backspace() {
	if t.IsFinished || t.CursorIdx == 0 {
		return
	}
	t.BackspaceCount++
	t.CursorIdx--
	t.CharStates[t.CursorIdx] = StatusUntyped
}

// Finish concludes the typing test session.
func (t *Tracker) Finish() {
	if !t.IsFinished {
		t.IsFinished = true
		t.EndTime = time.Now()
	}
}

// ElapsedSeconds calculates time spent typing.
func (t *Tracker) ElapsedSeconds() float64 {
	if !t.IsStarted {
		return 0
	}
	if t.IsFinished {
		return t.EndTime.Sub(t.StartTime).Seconds()
	}
	return time.Since(t.StartTime).Seconds()
}

// RemainingSeconds returns countdown timer seconds left.
func (t *Tracker) RemainingSeconds() int {
	if !t.IsStarted {
		return t.DurationSec
	}
	elapsed := int(t.ElapsedSeconds())
	rem := t.DurationSec - elapsed
	if rem < 0 {
		return 0
	}
	return rem
}

// CountCorrectChars returns total correctly typed characters currently on the board.
func (t *Tracker) CountCorrectChars() int {
	count := 0
	for i := 0; i < t.CursorIdx && i < len(t.CharStates); i++ {
		if t.CharStates[i] == StatusCorrect {
			count++
		}
	}
	return count
}

// CalculateWPM returns Net WPM: (correct_chars / 5) / elapsed_minutes.
func (t *Tracker) CalculateWPM() float64 {
	elapsedMin := t.ElapsedSeconds() / 60.0
	if elapsedMin <= 0 {
		return 0
	}
	netChars := float64(t.CountCorrectChars())
	wpm := (netChars / 5.0) / elapsedMin
	if wpm < 0 {
		return 0
	}
	return wpm
}

// CalculateRawWPM returns Raw WPM: (total_typed_chars / 5) / elapsed_minutes.
func (t *Tracker) CalculateRawWPM() float64 {
	elapsedMin := t.ElapsedSeconds() / 60.0
	if elapsedMin <= 0 {
		return 0
	}
	rawChars := float64(t.CursorIdx)
	return (rawChars / 5.0) / elapsedMin
}

// CalculateAccuracy returns Accuracy percentage.
func (t *Tracker) CalculateAccuracy() float64 {
	if t.TotalKeystrokes == 0 {
		return 100.0
	}
	acc := (float64(t.CorrectKeystrokes) / float64(t.TotalKeystrokes)) * 100.0
	if acc < 0 {
		return 0
	}
	if acc > 100 {
		return 100
	}
	return acc
}

// CalculateConsistency calculates WPM consistency percentage across samples.
func (t *Tracker) CalculateConsistency() float64 {
	if len(t.WPMSamples) < 2 {
		return 100.0
	}
	var sum float64
	for _, val := range t.WPMSamples {
		sum += val
	}
	mean := sum / float64(len(t.WPMSamples))
	if mean <= 0 {
		return 100.0
	}

	var varianceSum float64
	for _, val := range t.WPMSamples {
		diff := val - mean
		varianceSum += diff * diff
	}
	stdDev := math.Sqrt(varianceSum / float64(len(t.WPMSamples)))

	cv := stdDev / mean
	consistency := (1.0 - cv) * 100.0
	if consistency < 0 {
		consistency = 0
	}
	if consistency > 100 {
		consistency = 100
	}
	return consistency
}

// SampleWPM appends current Net WPM to samples slice for graph plotting.
func (t *Tracker) SampleWPM() {
	if t.IsStarted && !t.IsFinished {
		wpm := t.CalculateWPM()
		t.WPMSamples = append(t.WPMSamples, wpm)
	}
}
