#!/bin/sh
set -e

# JeeraType Installer Script — Prebuilt Binary Installation Only
REPO="Codexia-afk/JeeraType"
BINARY_NAME="jeeratype"

# 1. Detect OS and Architecture
OS_RAW="$(uname -s)"
ARCH_RAW="$(uname -m)"

case "$OS_RAW" in
  Darwin)  OS="darwin" ;;
  Linux)   OS="linux" ;;
  MINGW*|MSYS*|CYGWIN*) OS="windows" ;;
  *) echo "❌ Error: Unsupported Operating System: $OS_RAW"; exit 1 ;;
esac

case "$ARCH_RAW" in
  x86_64|amd64)  ARCH="amd64" ;;
  arm64|aarch64) ARCH="arm64" ;;
  *) echo "❌ Error: Unsupported Architecture: $ARCH_RAW"; exit 1 ;;
esac

echo "🚀 Installing JeeraType for ${OS}/${ARCH}..."

# 2. Query GitHub Releases API for the latest release tag
RELEASE_JSON=$(curl -sSL "https://api.github.com/repos/${REPO}/releases/latest" || true)

LATEST_TAG=$(echo "$RELEASE_JSON" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')

if [ -z "$LATEST_TAG" ]; then
  echo "❌ Error: Unable to fetch latest release from GitHub API (https://api.github.com/repos/${REPO}/releases/latest)."
  exit 1
fi

VERSION_NUM="${LATEST_TAG#v}"

# Determine file extension based on OS (.zip for windows, .tar.gz for darwin/linux)
if [ "$OS" = "windows" ]; then
  ASSET_NAME="jeeratype_${VERSION_NUM}_${OS}_${ARCH}.zip"
else
  ASSET_NAME="jeeratype_${VERSION_NUM}_${OS}_${ARCH}.tar.gz"
fi

DOWNLOAD_URL="https://github.com/${REPO}/releases/download/${LATEST_TAG}/${ASSET_NAME}"

TMP_DIR=$(mktemp -d)
trap 'rm -rf "$TMP_DIR"' EXIT
cd "$TMP_DIR"

# 3. Download the matching prebuilt binary asset directly from the release
echo "📥 Downloading prebuilt binary asset ${ASSET_NAME} (${LATEST_TAG})..."

if ! curl -sSLf "$DOWNLOAD_URL" -o "$ASSET_NAME" 2>/dev/null; then
  echo "❌ Error: No matching prebuilt binary asset found for OS: ${OS}, Arch: ${ARCH} in release ${LATEST_TAG}."
  echo "Download URL attempted: ${DOWNLOAD_URL}"
  echo "Please check available release assets at https://github.com/${REPO}/releases/tag/${LATEST_TAG}"
  exit 1
fi

# Extract prebuilt archive
if [ "$OS" = "windows" ]; then
  unzip -q "$ASSET_NAME"
else
  tar -xzf "$ASSET_NAME"
fi

FOUND_BIN=$(find . -type f \( -name "jeeratype" -o -name "jeeratype.exe" \) | head -n 1)

if [ -z "$FOUND_BIN" ] || [ ! -f "$FOUND_BIN" ]; then
  echo "❌ Error: Binary '${BINARY_NAME}' not found inside downloaded archive ${ASSET_NAME}."
  exit 1
fi

chmod +x "$FOUND_BIN"

# 4. Install binary to PATH (/usr/local/bin or $HOME/.local/bin)
GLOBAL_DIR="/usr/local/bin"
FALLBACK_DIR="$HOME/.local/bin"
INSTALLED_PATH=""

if [ -w "$GLOBAL_DIR" ]; then
  cp -f "$FOUND_BIN" "$GLOBAL_DIR/$BINARY_NAME"
  chmod +x "$GLOBAL_DIR/$BINARY_NAME"
  INSTALLED_PATH="$GLOBAL_DIR/$BINARY_NAME"
elif command -v sudo >/dev/null 2>&1 && [ -t 0 ]; then
  echo "⚙️  Installing to ${GLOBAL_DIR}/${BINARY_NAME} (requires sudo)..."
  if sudo cp -f "$FOUND_BIN" "$GLOBAL_DIR/$BINARY_NAME" && sudo chmod +x "$GLOBAL_DIR/$BINARY_NAME"; then
    INSTALLED_PATH="$GLOBAL_DIR/$BINARY_NAME"
  fi
fi

if [ -z "$INSTALLED_PATH" ]; then
  mkdir -p "$FALLBACK_DIR"
  cp -f "$FOUND_BIN" "$FALLBACK_DIR/$BINARY_NAME"
  chmod +x "$FALLBACK_DIR/$BINARY_NAME"
  INSTALLED_PATH="$FALLBACK_DIR/$BINARY_NAME"
fi

echo ""
echo "✅ JeeraType successfully installed to ${INSTALLED_PATH}!"

# 5. Print installed version at end using jeeratype --version
if command -v "$INSTALLED_PATH" >/dev/null 2>&1; then
  "$INSTALLED_PATH" --version || true
elif [ -x "$INSTALLED_PATH" ]; then
  "$INSTALLED_PATH" --version || true
fi
