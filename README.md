# 🌶️ JeeraType (`jeeratype`)

> **The 100% Offline, High-Performance, Terminal-Based Typing Speed Tester & MonkeyType Clone.**

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

[![Go Version](https://img.shields.io/badge/Go-1.21%2B-00ADD8?style=flat-square&logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-amber?style=flat-square)](LICENSE)
[![Platform](https://img.shields.io/badge/Platform-macOS%20%7C%20Linux%20%7C%20Windows-blue?style=flat-square)](https://github.com/Codexia-afk/JeeraType/releases)
[![CGO](https://img.shields.io/badge/CGO-Disabled-success?style=flat-square)]()

---

## ✨ Features

- ⚡ **100% Offline & Zero Runtime Dependencies**: Embeds all text datasets, themes, and code snippets via `go:embed`.
- 📊 **Pure Go SQLite Analytics**: Uses `modernc.org/sqlite` (no CGO/GCC needed) to track historical performance and keystroke metrics in `~/.config/jeeratype/stats.db`.
- 🧠 **Adaptive Weakness Engine**: Measures millisecond latency between key transitions (bigrams/n-grams) and error rates, dynamically generating custom practice text targeting your weakest keys.
- 👻 **The Ghost Pacer**: Features a customizable pacemaker cursor (60, 80, 100, 120 WPM) moving through the text to challenge your target speed.
- ⌨️ **ASCII QWERTY Keyboard Heatmap**: Renders an interactive color-coded keyboard mapping keys from Green (fast/accurate) to Red (slow/high latency).
- 💻 **Developer Code Mode**: Practice typing real-world code blocks across Go, Rust, JavaScript, and Python with code-friendly indentation.
- 🎨 **Lipgloss Color Themes**: Switch seamlessly between **Amber**, **Catppuccin**, **Nord**, **Dracula**, and **Matrix**.
- 📈 **WPM Timeline Graphing**: Visualizes WPM performance over time using Unicode sparkline graphing (`asciigraph`).

---

## 📦 Installation

### Option 1: Quick Shell Install (macOS / Linux)
```bash
curl -sSL https://raw.githubusercontent.com/Codexia-afk/JeeraType/main/install.sh | sh
```

### Option 2: Go Install (Cross-Platform)
If you have Go 1.21+ installed:
```bash
go install github.com/Codexia-afk/JeeraType@latest
```

### Option 3: Download Pre-Built Binary
Download standalone executable archives for **macOS (Apple Silicon & Intel)**, **Linux**, or **Windows** directly from the [Releases Page](https://github.com/Codexia-afk/JeeraType/releases).

---

## 🎮 Usage & Controls

Simply launch `jeeratype` in your terminal:

```bash
jeeratype
```

### ⌨️ Hotkeys & Keybindings

#### 🏠 Main Menu
| Key | Action |
| :--- | :--- |
| `1` – `5` / `←` `→` | Change test duration (15s, 30s, 45s, 60s, 120s) |
| `m` | Toggle Mode (**Paragraphs** ➔ **Code Mode** ➔ **Adaptive Weakness**) |
| `t` | Cycle Theme (**Amber** ➔ **Catppuccin** ➔ **Nord** ➔ **Dracula** ➔ **Matrix**) |
| `g` | Cycle Ghost Pacer speed (**Off** ➔ **60** ➔ **80** ➔ **100** ➔ **120 WPM**) |
| `k` | Open ASCII Keyboard Heatmap |
| `Enter` / `Space` | Start typing test |
| `Esc` / `q` | Quit application |

#### ⚡ Active Typing Test
| Key | Action |
| :--- | :--- |
| `Character Keys` | Type target text |
| `Backspace` | Delete previous character |
| `Tab` | Instantly restart test with new text |
| `Esc` | Abort test and return to menu |

#### 📊 Results Screen
| Key | Action |
| :--- | :--- |
| `Tab` / `Enter` | Start a new test |
| `k` | View ASCII Keyboard Heatmap |
| `Esc` | Return to menu |
| `q` | Quit |

---

## 🛠️ CLI Options & Flags

Launch directly into specific modes or output stats directly from the command line:

```bash
# Render ASCII Keyboard Heatmap & analytics directly in terminal
jeeratype --stats

# Start with Catppuccin theme and 80 WPM Ghost Pacer
jeeratype --theme catppuccin --ghost 80

# Start in Developer Code Mode
jeeratype --mode code
```

---

## 🏗️ Building from Source

```bash
# Clone repository
git clone https://github.com/Codexia-afk/JeeraType.git
cd JeeraType

# Build standalone binary
go build -o jeeratype .

# Run binary
./jeeratype
```

---

## 📄 License

Distributed under the MIT License. See `LICENSE` for more information.
