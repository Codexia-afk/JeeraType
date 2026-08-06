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

echo "🚀 Installing / Updating JeeraType for ${OS}/${ARCH}..."

# Fetch latest version tag from GitHub API
LATEST_TAG=$(curl -s https://api.github.com/repos/${REPO}/tags | grep '"name":' | head -n 1 | sed -E 's/.*"([^"]+)".*/\1/')

if [ -z "$LATEST_TAG" ]; then
  LATEST_TAG=$(curl -s https://api.github.com/repos/${REPO}/releases/latest | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
fi

if [ -z "$LATEST_TAG" ]; then
  LATEST_TAG="v2.4.0"
fi

VERSION_NUM="${LATEST_TAG#v}"
TAR_FILE="jeeratype_${VERSION_NUM}_${OS}_${ARCH}.tar.gz"

TMP_DIR=$(mktemp -d)
cd "$TMP_DIR"

URL_LOWER="https://github.com/${REPO}/releases/download/${LATEST_TAG}/jeeratype_${VERSION_NUM}_${OS}_${ARCH}.tar.gz"
URL_UPPER="https://github.com/${REPO}/releases/download/${LATEST_TAG}/JeeraType_${VERSION_NUM}_${OS}_${ARCH}.tar.gz"

echo "📥 Downloading JeeraType ${LATEST_TAG}..."
DOWNLOAD_SUCCESS=0

if curl -sSLf "$URL_LOWER" -o "$TAR_FILE" 2>/dev/null; then
  DOWNLOAD_SUCCESS=1
elif curl -sSLf "$URL_UPPER" -o "$TAR_FILE" 2>/dev/null; then
  DOWNLOAD_SUCCESS=1
fi

if [ "$DOWNLOAD_SUCCESS" -eq 1 ]; then
  tar -xzf "$TAR_FILE"
  FOUND_BIN=$(find . -maxdepth 2 -type f \( -name "jeeratype" -o -name "JeeraType" \) | head -n 1)
  if [ -n "$FOUND_BIN" ]; then
    chmod +x "$FOUND_BIN"
  fi
fi

# Fallback: if release tarball not found, compile directly from main branch via Go
if [ -z "$FOUND_BIN" ] || [ ! -f "$FOUND_BIN" ]; then
  if command -v go >/dev/null 2>&1; then
    echo "⚙️ Building latest JeeraType binary directly from GitHub main branch..."
    go install "github.com/${REPO}@main"
    GO_BIN="$(go env GOPATH)/bin/jeeratype"
    if [ -f "$GO_BIN" ]; then
      FOUND_BIN="$GO_BIN"
    fi
  fi
fi

if [ -z "$FOUND_BIN" ] || [ ! -f "$FOUND_BIN" ]; then
  echo "❌ Error: Could not download or locate jeeratype binary."
  echo "Please check available releases at: https://github.com/${REPO}/releases"
  exit 1
fi

# Global destination directory: /usr/local/bin
GLOBAL_DIR="/usr/local/bin"

if [ -w "$GLOBAL_DIR" ]; then
  cp -f "$FOUND_BIN" "$GLOBAL_DIR/$BINARY_NAME"
  chmod +x "$GLOBAL_DIR/$BINARY_NAME"
  echo ""
  echo "✅ JeeraType successfully installed to ${GLOBAL_DIR}/${BINARY_NAME}!"
  echo "Type 'jeeratype' anywhere in your terminal to start!"
  exit 0
fi

# Elevate with sudo for global installation in /usr/local/bin
if command -v sudo >/dev/null 2>&1; then
  echo "⚙️ Installing globally to ${GLOBAL_DIR}..."
  if sudo cp -f "$FOUND_BIN" "$GLOBAL_DIR/$BINARY_NAME" && sudo chmod +x "$GLOBAL_DIR/$BINARY_NAME"; then
    echo ""
    echo "✅ JeeraType successfully installed to ${GLOBAL_DIR}/${BINARY_NAME}!"
    echo "Type 'jeeratype' anywhere in your terminal to start!"
    exit 0
  fi
fi

# User-level fallback: $HOME/.local/bin
FALLBACK_DIR="$HOME/.local/bin"
mkdir -p "$FALLBACK_DIR"
cp -f "$FOUND_BIN" "$FALLBACK_DIR/$BINARY_NAME"
chmod +x "$FALLBACK_DIR/$BINARY_NAME"

PROFILES="$HOME/.zshrc $HOME/.bashrc $HOME/.profile $HOME/.bash_profile $HOME/.config/fish/config.fish"
for PROFILE in $PROFILES; do
  if [ -f "$PROFILE" ] && [ -w "$PROFILE" ]; then
    if ! grep -q "$FALLBACK_DIR" "$PROFILE"; then
      echo 'export PATH="$HOME/.local/bin:$PATH"' >> "$PROFILE"
    fi
  fi
done

export PATH="$FALLBACK_DIR:$PATH"

echo ""
echo "✅ JeeraType installed to ${FALLBACK_DIR}/${BINARY_NAME}!"
echo "Type 'jeeratype' to start!"
