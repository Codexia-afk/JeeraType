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

## Usage

```bash
# Launch default 30s typing test
jeeratype

# Pipe text directly into JeeraType
cat essay.txt | jeeratype

# Open long file with auto-resume reading progress
jeeratype /path/to/book.txt

# View session history and key heatmap
jeeratype stats
jeeratype stats --heatmap
```

---

## License

Distributed under the MIT License. See [`LICENSE`](LICENSE) for details.
