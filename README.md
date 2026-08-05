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

- **100% Offline & Single Binary**: Uses Go's native `//go:embed` for paragraph and code datasets. No runtime API or network calls.
- **Pure Go SQLite Engine**: Uses `modernc.org/sqlite` (zero CGO/GCC requirement) to persist completed test runs and keystroke statistics in `~/.config/jeeratype/stats.db`.
- **Keystroke Latency & Error Tracking**: Records millisecond timing between consecutive keystrokes and calculates error counts per key.
- **Adaptive Weakness Mode**: Generates text passages prioritizing your top slowest and most missed keys.
- **Ghost Pacer Cursor**: Optional pacemaker cursor moving through the target text at a fixed speed (60, 80, 100, or 120 WPM).
- **ASCII QWERTY Keyboard Heatmap**: Displays an ASCII QWERTY keyboard color-coded by accuracy and latency.
- **Developer Code Mode**: Typing passages sourced from Go, Rust, JavaScript, and Python code blocks.
- **Color Themes**: Selectable themes including Amber, Catppuccin, Nord, Dracula, and Matrix.
- **WPM Timeline Graphing**: Renders WPM over time using `asciigraph`.

---

## Installation

### 🍎 macOS & 🐧 Linux (Global One-Line Install)
Open terminal and run:
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

### 📦 Pre-Built Binaries
Pre-compiled static binaries for macOS, Linux, and Windows are available on the [GitHub Releases](https://github.com/Codexia-afk/JeeraType/releases) page.

---

## Usage

Run `jeeratype` in your terminal:

```bash
jeeratype
```

### Keybindings

#### Main Menu
- `1` – `5` or `←` / `→` : Select test duration (15s, 30s, 45s, 60s, 120s)
- `m` : Cycle typing mode (**Paragraphs** ➔ **Code Mode** ➔ **Adaptive Weakness**)
- `t` : Cycle theme (**Amber** ➔ **Catppuccin** ➔ **Nord** ➔ **Dracula** ➔ **Matrix**)
- `g` : Cycle Ghost Pacer speed (**Off** ➔ **60** ➔ **80** ➔ **100** ➔ **120 WPM**)
- `k` : Open ASCII Keyboard Heatmap
- `Enter` / `Space` : Start test
- `Esc` / `q` : Quit

#### Active Test
- Type text as displayed
- `Backspace` : Delete last character
- `Tab` : Restart test with new text
- `Esc` : Return to main menu

#### Results Screen
- `Tab` / `Enter` : Start new test
- `k` : View ASCII Keyboard Heatmap
- `Esc` : Main menu
- `q` : Quit

---

## Command-Line Flags

```bash
# Output ASCII Keyboard Heatmap directly in terminal
jeeratype --stats

# Start with a specific theme and Ghost Pacer speed
jeeratype --theme catppuccin --ghost 80

# Start in Developer Code Mode
jeeratype --mode code
```

---

## Building From Source

```bash
git clone https://github.com/Codexia-afk/JeeraType.git
cd JeeraType
go build -o jeeratype .
./jeeratype
```

---

## License

Distributed under the MIT License. See `LICENSE` for details.
