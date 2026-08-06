#!/bin/sh
set -e

# JeeraType Universal Uninstaller Script for macOS & Linux
BINARY_NAME="jeeratype"

echo "🗑️ Uninstalling JeeraType from your system..."

# Remove from /usr/local/bin
if [ -f "/usr/local/bin/$BINARY_NAME" ]; then
  if [ -w "/usr/local/bin/$BINARY_NAME" ]; then
    rm -f "/usr/local/bin/$BINARY_NAME"
    echo "  - Removed /usr/local/bin/$BINARY_NAME"
  elif command -v sudo >/dev/null 2>&1; then
    sudo rm -f "/usr/local/bin/$BINARY_NAME" 2>/dev/null || true
    echo "  - Removed /usr/local/bin/$BINARY_NAME"
  fi
fi

# Remove from $HOME/.local/bin
if [ -f "$HOME/.local/bin/$BINARY_NAME" ]; then
  rm -f "$HOME/.local/bin/$BINARY_NAME"
  echo "  - Removed $HOME/.local/bin/$BINARY_NAME"
fi

# Remove config / user stats directory
if [ -d "$HOME/.config/jeeratype" ]; then
  rm -rf "$HOME/.config/jeeratype"
  echo "  - Removed user config and data directory ($HOME/.config/jeeratype)"
fi

echo ""
echo "✅ JeeraType has been completely uninstalled from your system!"
