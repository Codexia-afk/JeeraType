#!/bin/sh
set -e

# JeeraType Universal Global Installer Script
REPO="Codexia-afk/JeeraType"
BINARY_NAME="jeeratype"

OS="$(uname -s | tr '[:upper:]' '[:lower:]')"
ARCH="$(uname -m)"

case "$ARCH" in
  x86_64|amd64) ARCH="amd64" ;;
  arm64|aarch64) ARCH="arm64" ;;
  *) echo "Unsupported architecture: $ARCH"; exit 1 ;;
esac

echo "🚀 Installing JeeraType for ${OS}/${ARCH}..."

# Fetch latest version tag from GitHub API
LATEST_TAG=$(curl -s https://api.github.com/repos/${REPO}/releases/latest | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')

if [ -z "$LATEST_TAG" ]; then
  LATEST_TAG="v1.0.2"
fi

VERSION_NUM="${LATEST_TAG#v}"
TAR_FILE="jeeratype_${VERSION_NUM}_${OS}_${ARCH}.tar.gz"

TMP_DIR=$(mktemp -d)
cd "$TMP_DIR"

URL_LOWER="https://github.com/${REPO}/releases/download/${LATEST_TAG}/jeeratype_${VERSION_NUM}_${OS}_${ARCH}.tar.gz"
URL_UPPER="https://github.com/${REPO}/releases/download/${LATEST_TAG}/JeeraType_${VERSION_NUM}_${OS}_${ARCH}.tar.gz"

echo "📥 Downloading JeeraType ${LATEST_TAG}..."
if ! curl -sSLf "$URL_LOWER" -o "$TAR_FILE" 2>/dev/null; then
  if ! curl -sSLf "$URL_UPPER" -o "$TAR_FILE" 2>/dev/null; then
    echo "❌ Failed to download release archive from GitHub."
    echo "Please check available releases at: https://github.com/${REPO}/releases"
    exit 1
  fi
fi

tar -xzf "$TAR_FILE"

# Find binary inside extracted directory
FOUND_BIN=$(find . -maxdepth 2 -type f \( -name "jeeratype" -o -name "JeeraType" \) | head -n 1)

if [ -z "$FOUND_BIN" ]; then
  echo "❌ Error: Could not locate jeeratype binary inside extracted package."
  exit 1
fi

chmod +x "$FOUND_BIN"

# Global binary destination
GLOBAL_DIR="/usr/local/bin"

if [ -w "$GLOBAL_DIR" ]; then
  mv "$FOUND_BIN" "$GLOBAL_DIR/$BINARY_NAME"
  chmod +x "$GLOBAL_DIR/$BINARY_NAME"
  echo ""
  echo "✅ JeeraType installed globally to ${GLOBAL_DIR}/${BINARY_NAME}!"
  echo "Type 'jeeratype' anywhere in your terminal to start!"
  exit 0
fi

# Try elevating with sudo if available
if command -v sudo >/dev/null 2>&1; then
  echo "⚙️ Installing to global system directory ${GLOBAL_DIR}..."
  if sudo mv "$FOUND_BIN" "$GLOBAL_DIR/$BINARY_NAME" 2>/dev/null && sudo chmod +x "$GLOBAL_DIR/$BINARY_NAME" 2>/dev/null; then
    echo ""
    echo "✅ JeeraType installed globally to ${GLOBAL_DIR}/${BINARY_NAME}!"
    echo "Type 'jeeratype' anywhere in your terminal to start!"
    exit 0
  fi
fi

# User-level fallback: $HOME/.local/bin
FALLBACK_DIR="$HOME/.local/bin"
mkdir -p "$FALLBACK_DIR"
mv "$FOUND_BIN" "$FALLBACK_DIR/$BINARY_NAME"
chmod +x "$FALLBACK_DIR/$BINARY_NAME"

# Automatically append to PATH in all common shell profiles
PROFILES="$HOME/.zshrc $HOME/.bashrc $HOME/.profile $HOME/.bash_profile $HOME/.config/fish/config.fish"
for PROFILE in $PROFILES; do
  if [ -f "$PROFILE" ] && [ -w "$PROFILE" ]; then
    if ! grep -q "$FALLBACK_DIR" "$PROFILE"; then
      echo 'export PATH="$HOME/.local/bin:$PATH"' >> "$PROFILE"
    fi
  fi
done

# Instantly export PATH for current shell subshell
export PATH="$FALLBACK_DIR:$PATH"

echo ""
echo "✅ JeeraType installed to ${FALLBACK_DIR}/${BINARY_NAME}!"
echo "Type 'jeeratype' to start!"
