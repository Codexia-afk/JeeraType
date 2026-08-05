package ui

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/Codexia-afk/JeeraType/config"
)

const RawLogo = `     ___                     _____                 
    |_  |                   |_   _|                
      | | ___  ___ _ __ __ _  | |   _   _ _ __   ___ 
      | |/ _ \/ _ \ '__/ _` + "`" + ` || |  | | | | '_ \ / _ \
  /\__/ /  __/  __/ | | (_| || |  | |_| | |_) |  __/
  \____/ \___|\___|_|  \__,_||_|   \__, | .__/ \___|
                                   __/ | |          
                                  |___/|_|          
  [ The Offline Terminal Typing Tester ]`

// RenderLogo renders the branded JeeraType logo themed dynamically.
func RenderLogo(theme config.Theme) string {
	logoStyle := lipgloss.NewStyle().
		Foreground(theme.Primary).
		Bold(true)
	return logoStyle.Render(RawLogo)
}
