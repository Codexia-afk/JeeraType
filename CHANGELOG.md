# Changelog — JeeraType

All notable changes to **JeeraType** will be documented in this file.

---

## [v2.0.0] — Master Feature Expansion (Major Release)

### 🚀 Tier 1 (Core Gameplay & UX Polish)
- **Punctuation & Numbers Toggles (`--punctuation` / `-p`, `--numbers` / `-n`)**: Injects capitalization, commas, periods, semicolons, question marks, and digit tokens into typing text streams. Menu toggles `p` and `n`.
- **Zen Mode (`--zen` / `-z`)**: Infinite typing mode with no timer or word cap. Text streams infinitely until `Esc` is pressed.
- **Countdown Before Start**: Centered `3…2…1…` countdown animation rendered before input starts (omitted in Zen Mode).
- **Persistent Runtime Footer**: Persistent footer hints bar `Tab: restart   Esc: quit   Ctrl+C: force quit` rendered at the bottom of the screen.
- **Personal Best (PB) Celebration**: Tracks and compares PB scores per mode/duration/punctuation/numbers combo and displays `🎉 NEW BEST!` banners.

### 📊 Tier 2 (Analytics, Streaks, Themes & Custom Lists)
- **Session History Command (`jeeratype stats`)**: Subcommand rendering a table of your last 20 typing sessions with trend indicators (`↑` / `↓` / `–`) versus prior runs.
- **Streak Counter**: Consecutive calendar day streak counter (`🔥 5 day streak`).
- **Custom Word Lists (`--wordlist /path/to/file.txt`)**: Load custom word lists (validates file exists and contains ≥ 50 words).
- **Lipgloss Themes Engine (`--theme <name>`)**: Themes (`amber`, `dracula`, `nord`, `solarized`, `catppuccin`, `gruvbox`, `matrix`) with JSON configuration persistence in `~/.config/jeeratype/config.json`.

### 🎮 Tier 3 (Advanced Hardcore & Multi-User Features)
- **Language Code Mode (`--mode code --lang python|js|go`)**: Snippet banks for Python, JavaScript, and Go.
- **Death Mode (`--death` / `-d`)**: Any typo immediately resets the test session.
- **Shaded Keyboard Heatmap (`jeeratype stats --heatmap`)**: ASCII QWERTY keyboard with 5 block shading intensity levels (`░`, `▒`, `▓`, `█`).
- **Multi-User Leaderboard (`--profile <name>`, `jeeratype stats --leaderboard`)**: Scopes history to profiles and displays top WPM per profile.
- **Replay Export & Race (`jeeratype export-replay`, `jeeratype race`)**: Export keystroke timelines to JSON and race against previous replays.
- **Sound Feedback (`--sound`)**: Optional non-blocking terminal bell (`\a`) audio feedback.

---

## [v1.0.2] — Master Architecture Initial Release
- Pure Go SQLite Engine (`modernc.org/sqlite`).
- UNIX STDIN Pipe support (`cat file.txt | jeeratype`).
- File Reader with Auto-Resume progress offset (`jeeratype /path/to/book.txt`).
- Hardcore modes (`--stop-on-error`, `--sudden-death`).
- Visual Keyboard Overlay (`-showkeys`).
- Scriptable UNIX Exporters (`-csv`, `-json`).
