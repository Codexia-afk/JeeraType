#!/bin/sh
set -e

# JeeraType Universal Installer Script
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
  LATEST_TAG="v1.0.0"
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

# Find binary inside extracted directory (case insensitive search)
FOUND_BIN=$(find . -maxdepth 2 -type f \( -name "jeeratype" -o -name "JeeraType" \) | head -n 1)

if [ -z "$FOUND_BIN" ]; then
  echo "❌ Error: Could not locate jeeratype binary inside extracted package."
  exit 1
fi

INSTALL_DIR="/usr/local/bin"
if [ ! -w "$INSTALL_DIR" ]; then
  INSTALL_DIR="$HOME/.local/bin"
  mkdir -p "$INSTALL_DIR"
fi

mv "$FOUND_BIN" "$INSTALL_DIR/$BINARY_NAME"
chmod +x "$INSTALL_DIR/$BINARY_NAME"

echo "✅ JeeraType installed successfully to ${INSTALL_DIR}/${BINARY_NAME}!"
echo "Run 'jeeratype' in your terminal to start typing."
