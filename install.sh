#!/bin/sh
set -e

# JeeraType Universal Installer Script
REPO="yourusername/jeeratype"
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

TAR_FILE="${BINARY_NAME}_${LATEST_TAG#v}_${OS}_${ARCH}.tar.gz"
DOWNLOAD_URL="https://github.com/${REPO}/releases/download/${LATEST_TAG}/${TAR_FILE}"

TMP_DIR=$(mktemp -d)
cd "$TMP_DIR"

echo "📥 Downloading from ${DOWNLOAD_URL}..."
curl -sSL "$DOWNLOAD_URL" -o "$TAR_FILE"
tar -xzf "$TAR_FILE"

INSTALL_DIR="/usr/local/bin"
if [ ! -w "$INSTALL_DIR" ]; then
  INSTALL_DIR="$HOME/.local/bin"
  mkdir -p "$INSTALL_DIR"
fi

mv "$BINARY_NAME" "$INSTALL_DIR/$BINARY_NAME"
chmod +x "$INSTALL_DIR/$BINARY_NAME"

echo "✅ JeeraType installed successfully to ${INSTALL_DIR}/${BINARY_NAME}!"
echo "Run 'jeeratype' in your terminal to start typing."
