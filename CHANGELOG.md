# Changelog — JeeraType

## v2.0.0 — Major Update Release Notes

Welcome to JeeraType v2.0.0! This release introduces powerful new practice modes, rich stats tracking, custom themes, and under-the-hood polish.

### New ways to practice

- **Punctuation & numbers toggle**: Practice realistic text containing capital letters, commas, periods, and numbers instead of plain lowercase words.
  *Command / Key:* Press `p` for punctuation or `n` for numbers on the menu, or run `jeeratype --punctuation --numbers`.

- **Zen mode**: Type freely without any timer or word limits. Whenever you want to stop, press `Esc` to view your final typing speed and accuracy.
  *Command / Key:* Press `z` on the menu or run `jeeratype --zen`.

- **Code mode**: Practice typing real computer code snippets with correct indentation for Python, JavaScript, and Go.
  *Command / Key:* Run `jeeratype --mode code --lang python` (or `js` / `go`).

- **Death mode**: Test your ultimate accuracy by restarting the test immediately if you make a single typo.
  *Command / Key:* Run `jeeratype --death`.

- **Countdown timer**: A clear 3…2…1… countdown appears before each timed test starts so you can prepare your hands.
  *Command / Key:* Starts automatically whenever you launch any timed typing test.

- **Custom word lists**: Load your own custom text file to practice specific words or vocabulary lists.
  *Command / Key:* Run `jeeratype --wordlist /path/to/mywords.txt`.

### New stats & tracking

- **Personal best tracking**: Automatically saves your highest typing speed for each mode and shows a celebration banner when you set a new record.
  *Command / Key:* Appears on the results screen automatically whenever you beat your previous best.

- **Session history**: View a clean summary table of your last 20 typing tests with trend arrows showing if your speed is improving over time.
  *Command / Key:* Run `jeeratype stats`.

- **Streak counter**: Tracks how many consecutive days you have practiced typing to keep you motivated.
  *Command / Key:* Shown automatically on your test results screen (e.g., `🔥 5 day streak`).

- **Key heatmap**: Displays a visual keyboard shaded by error frequency to help you identify which keys give you trouble.
  *Command / Key:* Run `jeeratype stats --heatmap`.

- **Local leaderboard**: Allows multiple people using the same computer to track separate practice profiles and compare top scores.
  *Command / Key:* Run `jeeratype --profile Alex` and view rankings with `jeeratype stats --leaderboard`.

- **Replay & race**: Save a past typing run to a file and re-type the passage to race live against your own past performance.
  *Command / Key:* Run `jeeratype export-replay` and `jeeratype race replay.json`.

### Customization & Theme System

- **6 New Premium Themes**:
  - `jewel` — deep saturated emerald, sapphire, ruby, and gold on a near-black background
  - `sunset` — warm coral, tangerine, and magenta on a deep plum background
  - `forest` — muted earthy moss green, bark brown, and cream text for low-contrast relaxed typing
  - `neon` — ultra-saturated electric cyan, hot pink, and acid green on pure black
  - `vintage` — desaturated mustard, rust, sage, and paper cream retro print poster palette
  - `mono` — high-contrast stark white, gold, and crimson two-tone for minimal distraction
- **Theme Preview Commands**:
  - `jeeratype theme list` displays color swatches `██` for all 18 registered themes.
  - `jeeratype theme preview <name>` renders a live UI preview box for any theme before applying it.

- **Optional sound feedback**: Hear a subtle audio click or bell sound every time you press a key or make a typo.
  *Command / Key:* Run `jeeratype --sound`.

- **Always-visible footer hints**: Important hotkeys are displayed at the bottom of the screen during tests so you never forget controls.
  *Command / Key:* Shown at the bottom of the screen (`Tab: restart   Esc: quit   Ctrl+C: force quit`).

### Under the hood improvements

- **100% offline**: Everything is stored directly on your computer, requiring zero internet connection.
- **Single executable file**: Runs as a single file without needing extra software installations or dependencies.
- **Cross-platform**: Behaves identically across Windows, macOS, and Linux terminal applications.

---

### How to update / install

Simply download the latest single binary for your operating system (Windows, macOS, or Linux) from the GitHub Releases page and run it directly in your terminal — no installation steps required!
