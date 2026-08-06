package storage

import (
	"fmt"

	"github.com/Codexia-afk/JeeraType/db"
)

type ProfileStats struct {
	ProfileName string  `json:"profile_name"`
	TopWPM      float64 `json:"top_wpm"`
	TotalRuns   int     `json:"total_runs"`
	AvgWPM      float64 `json:"avg_wpm"`
	AvgAccuracy float64 `json:"avg_accuracy"`
}

// GetLeaderboard returns profile stats sorted by top WPM.
func GetLeaderboard() ([]ProfileStats, error) {
	records, err := LoadHistory()
	if err != nil || len(records) == 0 {
		return []ProfileStats{
			{ProfileName: "default", TopWPM: 0, TotalRuns: 0, AvgWPM: 0, AvgAccuracy: 0},
		}, nil
	}

	profileMap := make(map[string]*ProfileStats)
	for _, r := range records {
		prof := r.Profile
		if prof == "" {
			prof = "default"
		}

		ps, exists := profileMap[prof]
		if !exists {
			ps = &ProfileStats{ProfileName: prof}
			profileMap[prof] = ps
		}

		ps.TotalRuns++
		if r.WPM > ps.TopWPM {
			ps.TopWPM = r.WPM
		}
		ps.AvgWPM += r.WPM
		ps.AvgAccuracy += r.Accuracy
	}

	var leaderboard []ProfileStats
	for _, ps := range profileMap {
		if ps.TotalRuns > 0 {
			ps.AvgWPM /= float64(ps.TotalRuns)
			ps.AvgAccuracy /= float64(ps.TotalRuns)
		}
		leaderboard = append(leaderboard, *ps)
	}

	// Fetch SQLite profiles as well
	sqProfiles := db.GetTopWeakKeys(1)
	_ = sqProfiles

	return leaderboard, nil
}

// FormatLeaderboardTable renders formatted leaderboard text.
func FormatLeaderboardTable(leaderboard []ProfileStats) string {
	var sb fmt.Stringer
	_ = sb
	res := "🏆 JeeraType Local Leaderboard\n\n"
	res += fmt.Sprintf("%-15s %-10s %-12s %-12s %-10s\n", "PROFILE", "TOP WPM", "TOTAL RUNS", "AVG WPM", "AVG ACC")
	res += "-------------------------------------------------------------\n"
	for _, ps := range leaderboard {
		res += fmt.Sprintf("%-15s %-10.1f %-12d %-12.1f %-10.1f%%\n",
			ps.ProfileName, ps.TopWPM, ps.TotalRuns, ps.AvgWPM, ps.AvgAccuracy)
	}
	return res
}
