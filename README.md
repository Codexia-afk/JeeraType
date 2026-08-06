# JeeraType

A 100% offline, cross-platform terminal typing test for macOS, Windows, and Linux that runs instantly in any shell without network access or web browser tabs.

[![Go Version](https://img.shields.io/github/go-mod/go-version/Codexia-afk/JeeraType?style=flat-square)](https://go.dev)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg?style=flat-square)](https://opensource.org/licenses/MIT)
[![Release](https://img.shields.io/github/v/release/Codexia-afk/JeeraType?style=flat-square)](https://github.com/Codexia-afk/JeeraType/releases)
[![Platform](https://img.shields.io/badge/platform-macOS%20%7C%20Linux%20%7C%20Windows-blue?style=flat-square)](#installation)

<!-- TODO: add demo.gif -->

Unlike browser-based typing tests that require web tabs and accounts, JeeraType runs fully offline directly inside your terminal window as a standalone tool.

---

## Why JeeraType?

- **100% Offline & Private**: Zero network calls, telemetry, or external web requests — word banks, code snippets, and themes are completely self-contained.
- **Single Executable File**: Ships as a single standalone executable per operating system with zero external dependencies.
- **Low Latency Input**: Sub-16ms keystroke rendering engine designed for fluid, responsive typing.
- **Multi-Platform Support**: Identical user interface and behavior across macOS Terminal, iTerm2, Linux terminals, and Windows Terminal.

---

## Features

- **Punctuation & Numbers**: Practice realistic text with capital letters, punctuation, and digits (`jeeratype --punctuation --numbers`).
- **Zen Mode**: Type infinitely without timers or word caps (`jeeratype --zen`).
- **Code Mode**: Practice real code snippets for Python, JavaScript, and Go (`jeeratype --mode code --lang python`).
- **Death Mode**: Restart instantly upon a single typo for strict accuracy training (`jeeratype --death`).
- **Custom Word Lists**: Load your own vocabulary files (`jeeratype --wordlist /path/to/words.txt`).
- **Personal Best Tracking**: Automatically alerts you when you set a new WPM speed record.
- **Session History**: Table of your last 20 typing runs with speed trend indicators (`jeeratype stats`).
- **Key Error Heatmap**: Visual ASCII keyboard shaded by error frequency (`jeeratype stats --heatmap`).
- **Local Leaderboard**: Track separate profiles for multiple users (`jeeratype --profile Alex` and `jeeratype stats --leaderboard`).
- **Replay & Race**: Export keystroke timelines to JSON and race against past runs (`jeeratype export-replay` & `jeeratype race replay.json`).
- **Color Themes**: Built-in schemes like Dracula, Nord, Solarized, Amber, Catppuccin, Gruvbox, and Matrix (`jeeratype --theme dracula`).
- **Optional Audio**: Subtle terminal bell sound feedback on keypresses and typos (`jeeratype --sound`).

---

## Installation

### 🍎 macOS & 🐧 Linux
Run in your terminal to install globally:
```bash
curl -sSL https://raw.githubusercontent.com/Codexia-afk/JeeraType/main/install.sh | sh
```
*Or with `sudo`:*
```bash
curl -sSL https://raw.githubusercontent.com/Codexia-afk/JeeraType/main/install.sh | sudo sh
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

## 🗑️ Uninstalling

If you wish to remove JeeraType in one command:

### 🍎 macOS & 🐧 Linux
```bash
curl -sSL https://raw.githubusercontent.com/Codexia-afk/JeeraType/main/uninstall.sh | sh
```
*Or if installed with `sudo`:*
```bash
curl -sSL https://raw.githubusercontent.com/Codexia-afk/JeeraType/main/uninstall.sh | sudo sh
```

### 🪟 Windows (PowerShell & Command Prompt)
In **PowerShell**, run:
```powershell
irm https://raw.githubusercontent.com/Codexia-afk/JeeraType/main/uninstall.ps1 | iex
```
*Or in Command Prompt (`cmd.exe`):*
```cmd
powershell -c "irm https://raw.githubusercontent.com/Codexia-afk/JeeraType/main/uninstall.ps1 | iex"
```

---

## Usage & CLI Reference

```bash
# 🚀 Launch Default Interactive Session
jeeratype

# 📊 Stats, Analytics & Leaderboards
jeeratype stats               # View history table of last 20 runs + trend arrows (↑/↓/–)
jeeratype stats --heatmap     # Render 5-level shaded ASCII key error heatmap
jeeratype stats --leaderboard # View multi-profile leaderboard rankings

# 🧘 Practice Modes
jeeratype --zen               # Infinite typing stream (no timer, press Esc to end)
jeeratype --punctuation -p    # Enable realistic capitalization & punctuation
jeeratype --numbers -n        # Enable numeric digit tokens
jeeratype -p -n               # Enable both punctuation and numbers
jeeratype --quotes            # Practice structured English quotes

# 💻 Code Mode
jeeratype --mode code --lang python # Practice Python code snippets
jeeratype --mode code --lang js     # Practice JavaScript code snippets
jeeratype --mode code --lang go     # Practice Go code snippets

# 🎯 Target Drills & Hardcore Modes
jeeratype --mode adaptive     # Target passage targeting your weakest keys & bigrams
jeeratype --death -d          # Single typo immediately resets test session
jeeratype --stop-on-error -soe# Forces error correction before advancing
jeeratype --ghost 80 -g 80    # Set Target Ghost Pacer to 80 WPM

# 🎨 Customization & Audio
jeeratype --theme dracula -t dracula # Options: amber, dracula, nord, solarized, catppuccin, gruvbox, matrix
jeeratype --sound            # Enable terminal bell / audio click feedback
jeeratype --showkeys -sk      # Enable live visual QWERTY keyboard overlay at bottom

# 👤 Multi-User Profiles & Custom Word Lists
jeeratype --profile Alex      # Scope test session & stats to a specific profile
jeeratype --wordlist /path/to/mywords.txt # Practice custom vocabulary (≥50 words required)

# 🏎️ Replays & UNIX Pipelines
jeeratype export-replay       # Export last run timeline to JSON (e.g. replay_12345.json)
jeeratype race replay.json    # Race live against a past recorded replay
cat essay.txt | jeeratype     # Pipe text directly into JeeraType
jeeratype /path/to/book.txt   # Read file with auto-resume progress offset
jeeratype --csv               # Export raw test history to CSV format
jeeratype --json              # Export raw test history to JSON format
```

---

## License

Distributed under the MIT License. See [`LICENSE`](LICENSE) for details.
