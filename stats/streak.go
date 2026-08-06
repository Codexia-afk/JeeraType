package stats

import (
	"fmt"
	"time"

	"github.com/Codexia-afk/JeeraType/storage"
)

// CalculateStreak returns consecutive calendar days with at least 1 completed test.
func CalculateStreak(records []storage.HistoryRecord) int {
	if len(records) == 0 {
		return 0
	}

	// Map unique dates YYYY-MM-DD
	dayMap := make(map[string]bool)
	for _, r := range records {
		dayStr := r.Timestamp.Format("2006-01-02")
		dayMap[dayStr] = true
	}

	today := time.Now()
	streak := 0

	// Check if tested today
	currentDate := today
	currentDayStr := currentDate.Format("2006-01-02")

	if !dayMap[currentDayStr] {
		// Check yesterday if haven't tested today yet
		currentDate = currentDate.AddDate(0, 0, -1)
		currentDayStr = currentDate.Format("2006-01-02")
		if !dayMap[currentDayStr] {
			return 0
		}
	}

	// Count backwards consecutive days
	for {
		dStr := currentDate.Format("2006-01-02")
		if dayMap[dStr] {
			streak++
			currentDate = currentDate.AddDate(0, 0, -1)
		} else {
			break
		}
	}

	return streak
}

// FormatStreakBadge returns formatted streak text.
func FormatStreakBadge(streak int) string {
	if streak == 0 {
		return "[0 day streak]"
	}
	return fmt.Sprintf("🔥 %d day streak", streak)
}
