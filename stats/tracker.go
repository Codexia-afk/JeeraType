package stats

import (
	"math"
	"time"

	"github.com/Codexia-afk/JeeraType/db"
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

// Tracker handles real-time typing metrics, latency tracking, Ghost Pacer, and PB replay.
type Tracker struct {
	TargetText         string
	TargetRunes        []rune
	CharStates         []CharStatus
	CursorIdx          int
	StartTime          time.Time
	EndTime            time.Time
	LastKeyTime        time.Time
	LastKeyRune        rune
	IsStarted          bool
	IsFinished         bool
	TotalKeystrokes    int
	CorrectKeystrokes  int
	ErrorKeystrokes    int
	BackspaceCount     int
	WPMSamples         []float64
	DurationSec        int
	IsZen              bool    // Zen Mode (infinite, no timer)
	TargetWPM          float64 // Ghost Pacer Target (0 = disabled)
	GhostCursorIdx     int     // Position of Ghost Pacer
	PBWPM              float64 // Personal Best WPM
	PreviousPBWPM      float64 // Previously stored PB WPM
	IsNewPB            bool    // Set to true if current run beat previous PB
	PBSamples          []float64
	PBGhostCursorIdx   int  // Position of Personal Best Ghost
	StopOnError        bool // Stop-on-Error mode
	SuddenDeath        bool // Sudden Death mode
	RunBigramLatencies []BigramLatency
	MissedKeysMap      map[rune]int
	OnNeedMoreText     func()
}

// AppendText appends more text to the active typing stream (used in Zen Mode).
func (t *Tracker) AppendText(moreText string) {
	if len(moreText) == 0 {
		return
	}
	prefix := ""
	if len(t.TargetRunes) > 0 && t.TargetRunes[len(t.TargetRunes)-1] != ' ' {
		prefix = " "
	}
	added := prefix + moreText
	addedRunes := []rune(added)
	t.TargetText += added
	t.TargetRunes = append(t.TargetRunes, addedRunes...)
	newStates := make([]CharStatus, len(addedRunes))
	for i := range newStates {
		newStates[i] = StatusUntyped
	}
	t.CharStates = append(t.CharStates, newStates...)
}

// NewTracker creates a new typing tracker.
func NewTracker(targetText string, durationSec int, targetWPM float64, stopOnError bool, suddenDeath bool, isZen bool) *Tracker {
	runes := []rune(targetText)
	states := make([]CharStatus, len(runes))
	for i := range states {
		states[i] = StatusUntyped
	}
	return &Tracker{
		TargetText:         targetText,
		TargetRunes:        runes,
		CharStates:         states,
		CursorIdx:          0,
		DurationSec:        durationSec,
		IsZen:              isZen,
		TargetWPM:          targetWPM,
		StopOnError:        stopOnError,
		SuddenDeath:        suddenDeath,
		WPMSamples:         make([]float64, 0, durationSec),
		RunBigramLatencies: make([]BigramLatency, 0),
		MissedKeysMap:      make(map[rune]int),
	}
}

// Start begins tracking time on the first keypress.
func (t *Tracker) Start() {
	if !t.IsStarted {
		now := time.Now()
		t.IsStarted = true
		t.StartTime = now
		t.LastKeyTime = now
	}
}

// RecordRune processes typed character input.
func (t *Tracker) RecordRune(r rune) {
	if t.IsZen && t.CursorIdx >= len(t.TargetRunes)-20 {
		if t.OnNeedMoreText != nil {
			t.OnNeedMoreText()
		} else {
			// Fallback auto-extend by repeating TargetText
			t.AppendText(t.TargetText)
		}
	}

	if t.IsFinished || t.CursorIdx >= len(t.TargetRunes) {
		return
	}

	// Stop-on-Error check: if previous char was wrong, force backspace
	if t.StopOnError && t.CursorIdx > 0 && t.CharStates[t.CursorIdx-1] == StatusIncorrect {
		return
	}

	now := time.Now()
	if !t.IsStarted {
		t.Start()
		now = t.StartTime
	}

	t.TotalKeystrokes++

	var latencyMs int64
	if !t.LastKeyTime.IsZero() {
		latencyMs = now.Sub(t.LastKeyTime).Milliseconds()
		if latencyMs > 2000 {
			latencyMs = 2000
		}
	}
	t.LastKeyTime = now

	expected := t.TargetRunes[t.CursorIdx]
	isError := (r != expected)

	// Sudden Death check: single typo ends test
	if t.SuddenDeath && isError {
		t.CharStates[t.CursorIdx] = StatusIncorrect
		t.ErrorKeystrokes++
		t.Finish()
		return
	}

	if !isError {
		t.CharStates[t.CursorIdx] = StatusCorrect
		t.CorrectKeystrokes++
	} else {
		t.CharStates[t.CursorIdx] = StatusIncorrect
		t.ErrorKeystrokes++
		t.MissedKeysMap[expected]++
	}

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
		if !t.IsZen {
			t.Finish()
		}
	}
}

// UpdateGhostPacer updates position of Target Ghost & PB Ghost.
func (t *Tracker) UpdateGhostPacer() {
	if !t.IsStarted || t.IsFinished {
		return
	}
	elapsed := t.ElapsedSeconds()

	// Target Ghost Pacer
	if t.TargetWPM > 0 {
		charsPerSec := (t.TargetWPM * 5.0) / 60.0
		ghostIdx := int(elapsed * charsPerSec)
		if ghostIdx > len(t.TargetRunes) {
			ghostIdx = len(t.TargetRunes)
		}
		t.GhostCursorIdx = ghostIdx
	}

	// PB Ghost Replay
	if t.PBWPM > 0 {
		pbCharsPerSec := (t.PBWPM * 5.0) / 60.0
		pbIdx := int(elapsed * pbCharsPerSec)
		if pbIdx > len(t.TargetRunes) {
			pbIdx = len(t.TargetRunes)
		}
		t.PBGhostCursorIdx = pbIdx
	}
}

// Backspace handles deleting character.
func (t *Tracker) Backspace() {
	if t.IsFinished || t.CursorIdx == 0 {
		return
	}
	t.BackspaceCount++
	t.CursorIdx--
	t.CharStates[t.CursorIdx] = StatusUntyped
}

// Finish concludes session.
func (t *Tracker) Finish() {
	if !t.IsFinished {
		t.IsFinished = true
		t.EndTime = time.Now()
	}
}

// ElapsedSeconds returns time spent typing.
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
	if t.IsZen {
		return 0
	}
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

// CountCorrectChars returns count of correct characters currently on board.
func (t *Tracker) CountCorrectChars() int {
	count := 0
	for i := 0; i < t.CursorIdx && i < len(t.CharStates); i++ {
		if t.CharStates[i] == StatusCorrect {
			count++
		}
	}
	return count
}

// GrossWPM: (Total Keystrokes / 5) / Elapsed Minutes
func (t *Tracker) CalculateGrossWPM() float64 {
	elapsedMin := t.ElapsedSeconds() / 60.0
	if elapsedMin <= 0 {
		return 0
	}
	return (float64(t.TotalKeystrokes) / 5.0) / elapsedMin
}

// NetWPM: Gross WPM - (Uncorrected Errors / Elapsed Minutes)
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

func (t *Tracker) CalculateRawWPM() float64 {
	return t.CalculateGrossWPM()
}

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

func (t *Tracker) SampleWPM() {
	if t.IsStarted && !t.IsFinished {
		wpm := t.CalculateWPM()
		t.WPMSamples = append(t.WPMSamples, wpm)
	}
}
