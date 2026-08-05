# JeeraType (`jeeratype`)

JeeraType is a 100% offline, cross-platform terminal-based typing speed tester written in Go using Charm's Bubble Tea and Lipgloss libraries.

```text
     ___                     _____                 
    |_  |                   |_   _|                
      | | ___  ___ _ __ __ _  | |   _   _ _ __   ___ 
      | |/ _ \/ _ \ '__/ _` || |  | |_| | | '_ \ / _ \
  /\__/ /  __/  __/ | | (_| || |  |  _  | |_) |  __/
  \____/ \___|\___|_|  \__,_||_|   \__, | .__/ \___|
                                   __/ | |          
                                  |___/|_|          
  [ The Offline Terminal Typing Tester ]
```

---

## Key Features

- **100% Offline & Single Binary**: Uses Go's native `//go:embed` for paragraphs, quotes, and code datasets. Zero runtime network calls.
- **Pure Go SQLite Engine**: Uses `modernc.org/sqlite` (zero CGO/GCC requirement) to persist completed test runs, key transition latencies, and file reading progress offsets in `~/.config/jeeratype/stats.db`.
- **UNIX STDIN & Pipe Support**: Accepts text via standard input (e.g. `cat file.txt | jeeratype`).
- **File Reader with Auto-Resume**: Opening a text file (`jeeratype /path/to/book.txt`) saves cursor progress in SQLite and automatically resumes where you left off.
- **Industrial Math Engine**: Computes Gross WPM `((total_keystrokes/5) / elapsed_minutes)`, Net WPM, Accuracy %, and WPM consistency.
- **Adaptive Weakness Engine**: Measures millisecond latency between key transitions (bigrams) and error rates, generating custom drill passages targeting your weakest keys.
- **Ghost Pacers & PB Replay**: Includes Target Ghost Pacer (60, 80, 100, 120 WPM) and Personal Best (PB) Ghost Replay `🏆` stored in SQLite.
- **Hardcore Modes**:
  - `--stop-on-error` / `-soe`: Forces error correction before advancing.
  - `--sudden-death` / `-sd`: Single typo ends the test immediately.
- **Visual Keyboard Overlay (`-showkeys` / `-sk`)**: Displays a live QWERTY keyboard at the bottom of the screen highlighting key presses in real-time.
- **Scriptable UNIX Exporters (`-csv` / `-json`)**: Outputs historical typing data directly in structured CSV or JSON for shell scripts and spreadsheets.
- **Color Themes & TOML Loader**: Built-in themes (`Amber`, `Catppuccin`, `Nord`, `Dracula`, `Matrix`, `Gruvbox`) plus custom `.toml` theme loading (`--config-theme`).

---

## Installation

### 🍎 macOS & 🐧 Linux (Global One-Line Install)
```bash
curl -sSL https://raw.githubusercontent.com/Codexia-afk/JeeraType/main/install.sh | sudo sh
```
*Or without `sudo`:*
```bash
curl -sSL https://raw.githubusercontent.com/Codexia-afk/JeeraType/main/install.sh | sh
```

### 🪟 Windows (One-Line PowerShell Install)
Open **PowerShell** and run:
```powershell
irm https://raw.githubusercontent.com/Codexia-afk/JeeraType/main/install.ps1 | iex
```

### ⚡ Go Install (Cross-Platform)
```bash
go install github.com/Codexia-afk/JeeraType@latest
```

---

## Usage & CLI Flags

```bash
# Launch interactive TUI
jeeratype

# Pipe UNIX text into JeeraType
cat essay.txt | jeeratype

# Open file with auto-resume progress offset
jeeratype /path/to/book.txt

# Hardcore Modes
jeeratype --stop-on-error
jeeratype --sudden-death

# Visual Keyboard Overlay & Custom Theme
jeeratype --showkeys --theme catppuccin --ghost 80

# Output Keyboard Heatmap
jeeratype --stats

# Export UNIX CSV / JSON data
jeeratype --csv
jeeratype --json
```

---

## License

Distributed under the MIT License. See `LICENSE` for details.
