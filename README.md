# JeeraType — #1 Terminal Typing Test ⚡

> **The #1 100% Offline, Cross-Platform CLI Typing Speed & Accuracy Trainer for macOS, Windows, and Linux.**

[![Release](https://img.shields.io/github/v/release/Codexia-afk/JeeraType?style=flat-square&color=blue)](https://github.com/Codexia-afk/JeeraType/releases)
[![Go Version](https://img.shields.io/github/go-mod/go-version/Codexia-afk/JeeraType?style=flat-square&color=00ADD8)](https://go.dev)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg?style=flat-square)](https://opensource.org/licenses/MIT)
[![Platform](https://img.shields.io/badge/platform-macOS%20%7C%20Linux%20%7C%20Windows-blue?style=flat-square)](#-installation)
[![Support on Ko-fi](https://img.shields.io/badge/Support-Ko--fi-FF5E5B?style=flat-square&logo=ko-fi&logoColor=white)](https://ko-fi.com/srinjoypramanick)

---

## ☕ Support & Sponsor

If **JeeraType** helps you boost your typing speed, consider supporting the project! Your sponsorship keeps JeeraType 100% free, offline, and actively maintained.

<p align="center">
  <a href="https://ko-fi.com/srinjoypramanick" target="_blank">
    <img src="https://storage.ko-fi.com/cdn/kofi2.png?v=3" alt="Buy Me a Coffee at ko-fi.com" height="48">
  </a>
</p>

<p align="center">
  👉 <b><a href="https://ko-fi.com/srinjoypramanick">Sponsor JeeraType on Ko-fi (ko-fi.com/srinjoypramanick)</a></b> 👈
</p>

---

## 📌 Table of Contents

- [Why JeeraType?](#-why-jeeratype)
- [Features](#-features)
- [Installation](#-installation)
  - [macOS & Linux](#-macos--linux)
  - [Windows](#-windows-powershell--command-prompt)
- [Usage & CLI Reference](#-usage--cli-reference)
- [Themes & Customization](#-themes--customization)
- [Uninstalling](#-uninstalling)
- [Support & Sponsor](#-support--sponsor-1)
- [License](#-license)

---

## ⚡ Why JeeraType?

Unlike browser-based typing tests that require web tabs, accounts, and internet access, **JeeraType** is a standalone, lightweight, ultra-low latency CLI & TUI terminal typing test designed for developers, sysadmins, and keyboard enthusiasts.

- 🔒 **100% Offline & Private**: Zero network calls, telemetry, or external web requests. Word banks, code snippets, and themes are 100% self-contained.
- 🚀 **Single Standalone Binary**: Ships as a single executable with zero external dependencies.
- ⚡ **Sub-16ms Keystroke Latency**: Fluid, responsive rendering powered by Go and Bubble Tea.
- 🖥️ **Cross-Platform Compatibility**: Identical experience across macOS Terminal, iTerm2, Linux terminals, and Windows Terminal.
- 🎯 **Offline MonkeyType Alternative**: Built-in WPM analytics, streak tracking, error heatmaps, and customizable themes directly inside your terminal.

---

## ✨ Features

- **Punctuation & Numbers**: Practice realistic text with capital letters, commas, periods, and numeric digits (`jeeratype --punctuation --numbers`).
- **Zen Mode**: Type infinitely without timers or word caps for relaxed flow practice (`jeeratype --zen`).
- **Code Mode**: Practice real code snippets for Python, JavaScript, and Go (`jeeratype --mode code --lang python`).
- **Death Mode**: Sudden-death accuracy drill — a single typo instantly resets the test (`jeeratype --death`).
- **Custom Word Lists**: Load your own custom vocabulary files (`jeeratype --wordlist /path/to/words.txt`).
- **Personal Best Tracking**: Automatic WPM record detection with celebratory banners.
- **Session History & Analytics**: SQLite-backed history table of your last 20 typing runs with speed trend indicators (`jeeratype stats`).
- **Key Error Heatmap**: Visual 5-level shaded ASCII QWERTY keyboard map highlighting error frequencies (`jeeratype stats --heatmap`).
- **Local Leaderboard**: Multi-profile support for tracking multiple users or practice routines (`jeeratype --profile Alex`).
- **Replays & Racing**: Export keystroke timelines to JSON and race live against past runs (`jeeratype export-replay` & `jeeratype race replay.json`).
- **18 Custom Color Themes**: Themes like Dracula, Nord, Solarized, Catppuccin, Gruvbox, Jewel, Neon, and Matrix (`jeeratype --theme dracula`).
- **Audio Feedback**: Optional terminal bell sound feedback on keypresses and typos (`jeeratype --sound`).

---

## 📦 Installation

### 🍎 macOS & 🐧 Linux

Run in your terminal to install globally:

```bash
curl -sSL https://raw.githubusercontent.com/Codexia-afk/JeeraType/main/install.sh | sh
```

*Or with `sudo`:*

```bash
curl -sSL https://raw.githubusercontent.com/Codexia-afk/JeeraType/main/install.sh | sudo sh
```

---

### 🪟 Windows (PowerShell & Command Prompt)

In **PowerShell**, run:

```powershell
irm https://raw.githubusercontent.com/Codexia-afk/JeeraType/main/install.ps1 | iex
```

*Or in Command Prompt (`cmd.exe`):*

```cmd
powershell -c "irm https://raw.githubusercontent.com/Codexia-afk/JeeraType/main/install.ps1 | iex"
```

> 💡 **Note on Updating**: Running the install command again automatically upgrades your binary to the latest version while preserving all your saved stats, history, and themes!

---

## 💻 Usage & CLI Reference

```bash
# 🚀 Launch Default Interactive Session
jeeratype

# 📊 Stats, Analytics & Leaderboards
jeeratype stats               # History table of last 20 runs + trend arrows (↑/↓/–)
jeeratype stats --heatmap     # Visual 5-level shaded ASCII key error heatmap
jeeratype stats --leaderboard # Multi-profile leaderboard rankings

# 🧘 Practice Modes
jeeratype --zen               # Infinite typing stream (no timer, press Esc to end)
jeeratype --punctuation -p    # Enable capitalization & punctuation
jeeratype --numbers -n        # Enable numeric digit tokens
jeeratype -p -n               # Enable both punctuation and numbers
jeeratype --quotes            # Practice English quotes

# 💻 Code Mode
jeeratype --mode code --lang python # Practice Python code snippets
jeeratype --mode code --lang js     # Practice JavaScript code snippets
jeeratype --mode code --lang go     # Practice Go code snippets

# 🎯 Hardcore & Target Drills
jeeratype --death -d          # Single typo immediately resets test session
jeeratype --wordlist /path/to/words.txt # Practice custom vocabulary list

# 🎨 Customization & Theme Preview
jeeratype theme list          # Display color swatches for all 18 themes
jeeratype theme preview jewel # Render live UI preview for a specific theme
jeeratype --theme dracula     # Options: dracula, nord, solarized, catppuccin, gruvbox, jewel, sunset, forest, neon, vintage, mono, matrix, amber, cyberpunk, tokyonight, monokai, rose-pine, synthwave
jeeratype --sound            # Enable terminal bell / audio click feedback

# 👤 Multi-User Profiles & UNIX Pipelines
jeeratype --profile Alex      # Scope test session & stats to a specific profile
cat essay.txt | jeeratype     # Pipe text directly into JeeraType
jeeratype /path/to/book.txt   # Read file with progress offset
```

---

## 🎨 Themes & Customization

JeeraType ships with **18 built-in color themes**:

| Theme | Command | Description |
| :--- | :--- | :--- |
| **Dracula** | `jeeratype --theme dracula` | Dark violet & pastel accent theme |
| **Nord** | `jeeratype --theme nord` | Arctic icy blue theme |
| **Solarized** | `jeeratype --theme solarized` | Precision dark teal theme |
| **Catppuccin** | `jeeratype --theme catppuccin` | Warm pastel theme |
| **Gruvbox** | `jeeratype --theme gruvbox` | Retro groove earth-tone palette |
| **Jewel** | `jeeratype --theme jewel` | Saturated emerald, sapphire, & ruby |
| **Sunset** | `jeeratype --theme sunset` | Warm coral, tangerine, & plum |
| **Forest** | `jeeratype --theme forest` | Muted moss green & bark brown |
| **Neon** | `jeeratype --theme neon` | Electric cyan, hot pink, & acid green |
| **Matrix** | `jeeratype --theme matrix` | Classic digital rain green on black |

List all available themes and preview swatches:

```bash
jeeratype theme list
jeeratype theme preview jewel
```

---

## ☕ Support & Sponsor

If you love **JeeraType**, please consider buying a coffee on Ko-fi to support open-source development!

[![Ko-fi Sponsor](https://img.shields.io/badge/Sponsor%20on-Ko--fi-ff5e5b?style=for-the-badge&logo=ko-fi&logoColor=white)](https://ko-fi.com/srinjoypramanick)

- ☕ **Ko-fi Page**: [ko-fi.com/srinjoypramanick](https://ko-fi.com/srinjoypramanick)
- ⭐️ **Star the Repo**: Give [JeeraType a Star on GitHub](https://github.com/Codexia-afk/JeeraType) to help others discover it!

---

## 🗑️ Uninstalling

### 🍎 macOS & 🐧 Linux
```bash
curl -sSL https://raw.githubusercontent.com/Codexia-afk/JeeraType/main/uninstall.sh | sh
```

### 🪟 Windows (PowerShell & Command Prompt)
```powershell
irm https://raw.githubusercontent.com/Codexia-afk/JeeraType/main/uninstall.ps1 | iex
```

---

## 📜 License

Distributed under the **MIT License**. See [`LICENSE`](LICENSE) for details.
