# JeeraType (`jeeratype`) v2.0.0

JeeraType is a 100% offline, cross-platform terminal-based typing speed tester written in Go using Charm's Bubble Tea and Lipgloss libraries.

```text
     _                _____                 
    | | ___  ___ _ __/ /_ /___ _   _  ___   
 _  | |/ _ \/ _ \ '__/ /_ / __| | | |/ _ \  
| |_| |  __/  __/ | / / / (__| |_| |  __/  
 \___/ \___|\___|_|/_/  \___|\__, |\___|   
                              |___/        
   [ terminal typing speed test — 100% offline ]
```

---

## Key Features in v2.0.0

- **100% Offline & Pure Go**: Zero CGO dependencies (`modernc.org/sqlite`). All paragraphs, quotes, themes, and code snippets are embedded via `go:embed`.
- **Punctuation & Numbers Toggles (`--punctuation`, `--numbers`)**: Injects capitalization, punctuation (`.`, `,`, `?`, `;`), and numbers (`0`-`9`) into text streams. Menu toggles `p` and `n`.
- **Zen Mode (`--zen`)**: Infinite typing stream with no timer or word limit until `Esc` is pressed.
- **Countdown Before Start**: Centered `3…2…1…` countdown animation before input begins.
- **Persistent Runtime Footer**: Persistent hints bar (`Tab: restart   Esc: quit   Ctrl+C: force quit`) at the bottom of the active screen.
- **Personal Best (PB) Tracking**: Automatically tracks PB scores per mode/duration/punctuation/numbers combo and displays `🎉 NEW BEST!` banners.
- **Session History Subcommand (`jeeratype stats`)**: Formatted table of the last 20 runs with performance trend indicators (`↑` / `↓` / `–`).
- **Streak Counter**: Tracks consecutive calendar days with at least 1 completed test (`🔥 5 day streak`).
- **Shaded Keyboard Heatmap (`jeeratype stats --heatmap`)**: Renders an ASCII QWERTY keyboard with 5 block shading intensity levels (`░`, `▒`, `▓`, `█`).
- **Multi-User Leaderboard (`--profile <name>`, `jeeratype stats --leaderboard`)**: Profile-scoped history and leaderboard rankings.
- **Code Mode Expansion (`--mode code --lang python|js|go`)**: Snippets for Python, JavaScript, and Go.
- **Death Mode (`--death`)**: Any typo immediately resets the test session.
- **Custom Word Lists (`--wordlist /path/to/file.txt`)**: Validated custom wordlist loader (requires ≥ 50 words).
- **Replay Export & Race (`jeeratype export-replay`, `jeeratype race`)**: Export keystroke timelines to JSON and race against past replays.
- **Themes Engine (`--theme <name>`)**: Themes (`amber`, `dracula`, `nord`, `solarized`, `catppuccin`, `gruvbox`, `matrix`) saved in `~/.config/jeeratype/config.json`.

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

### 🪟 Windows (PowerShell & Command Prompt)
In **PowerShell**, run:
```powershell
irm https://raw.githubusercontent.com/Codexia-afk/JeeraType/main/install.ps1 | iex
```
*Or in Command Prompt (`cmd.exe`):*
```cmd
powershell -c "irm https://raw.githubusercontent.com/Codexia-afk/JeeraType/main/install.ps1 | iex"
```

---

## Usage & Subcommands

```bash
# Launch interactive TUI
jeeratype

# Subcommands
jeeratype stats               # Session history table (last 20 runs + trend)
jeeratype stats --heatmap     # Shaded ASCII QWERTY key error heatmap
jeeratype stats --leaderboard # Local multi-user profile leaderboard
jeeratype export-replay       # Export keystroke timeline to JSON

# CLI Flags
jeeratype --punctuation --numbers
jeeratype --zen
jeeratype --theme dracula
jeeratype --mode code --lang python
jeeratype --death
jeeratype --profile srinjoy
jeeratype --wordlist /path/to/mywords.txt
```

---

## 🗑️ Uninstalling

### 🍎 macOS & 🐧 Linux
```bash
curl -sSL https://raw.githubusercontent.com/Codexia-afk/JeeraType/main/uninstall.sh | sh
```

### 🪟 Windows
```powershell
irm https://raw.githubusercontent.com/Codexia-afk/JeeraType/main/uninstall.ps1 | iex
```

---

## License

Distributed under the MIT License. See `LICENSE` for details.
