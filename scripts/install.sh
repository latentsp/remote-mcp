#!/bin/bash
set -e

# Determine architecture
ARCH=$(uname -m)
case $ARCH in
    x86_64)
        ARCH="amd64"
        ;;
    arm64|aarch64)
        ARCH="arm64"
        ;;
    *)
        echo "Unsupported architecture: $ARCH"
        exit 1
        ;;
esac

# Determine OS
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
case $OS in
    linux*)
        OS="linux"
        ;;
    darwin*)
        OS="darwin"
        ;;
    *)
        echo "Unsupported OS: $OS"
        exit 1
        ;;
esac

# Get latest release
echo "Fetching the latest release..."
LATEST_RELEASE=$(curl -s https://api.github.com/repos/latentsp/remote-mcp/releases/latest)
VERSION=$(echo $LATEST_RELEASE | grep -o '"tag_name": "[^"]*' | cut -d'"' -f4)

# Download URL
DOWNLOAD_URL="https://github.com/latentsp/remote-mcp/releases/download/$VERSION/remote-mcp-$OS-$ARCH"
echo "Downloading remote-mcp $VERSION for $OS-$ARCH..."

# Create temp directory
TEMP_DIR=$(mktemp -d)
trap 'rm -rf "$TEMP_DIR"' EXIT

# Download the binary
BINARY_PATH="$TEMP_DIR/remote-mcp"
curl -L -o "$BINARY_PATH" "$DOWNLOAD_URL"

# Make executable and move to bin directory
chmod +x "$BINARY_PATH"
sudo mv "$BINARY_PATH" /usr/local/bin/

echo "remote-mcp successfully installed to /usr/local/bin/"
echo "Run 'remote-mcp --help' to get started"