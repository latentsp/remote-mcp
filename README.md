# Remote MCP Client

A command-line client for interacting with remote MCP servers.

## Overview

This client establishes a connection to a remote MCP server, relays commands from standard input, and forwards responses back to standard output. It uses Server-Sent Events (SSE) to maintain a persistent connection with the server and discover the endpoint URL.

## Usage

```
./remote-mcp --endpoint-url=<remote-mcp-server-url> [--verbose]
```

## Usage inside applications (Cursor, Claude, etc.)

Configure your application
```json
{
  "mcpServers": {
      "command": "remote-mcp",
      "args": [
         "--endpoint-url",
         "<your remote mcp server"
      ]
  }
}
```

### Command-line Arguments

- `--endpoint-url`: (Required) The base URL of the remote MCP server
- `--verbose`: (Optional) Enable verbose logging

## How It Works

1. The client connects to the SSE endpoint of the remote server (`/sse`).
2. It waits to receive the actual endpoint URL through the SSE connection.
3. Once the endpoint URL is received, it reads JSON-RPC commands from standard input.
4. Each command is forwarded to the remote MCP server.
5. All interactions are logged to `~/remote-mcp/server.log`.

## Log Location

Logs are stored in `~/remote-mcp/server.log`

## Building from source

To build the Remote MCP Client from source:

1. Ensure you have Go 1.18+ installed on your system.
   ```
   go version
   ```

2. Clone the repository:
   ```
   git clone https://github.com/yourusername/remote-mcp.git
   cd remote-mcp
   ```

3. Build the executable:
   ```
   go build -o remote-mcp .
   ```

4. The binary will be created in the current directory.

## Downloading from releases

Pre-built binaries are available for the following platforms:
- Linux (amd64, arm64)
- macOS (amd64, arm64)
- Windows (amd64)

To download and use a pre-built binary:

1. Visit the [Releases](https://github.com/yourusername/remote-mcp/releases) page.
2. Download the appropriate binary for your platform.
3. Make the file executable (Linux/macOS):
   ```
   chmod +x remote-mcp
   ```
4. Run the client with the required parameters.

## Development

### Prerequisites

- Go 1.18+
- Git

### Development Workflow

1. Fork the repository.
2. Create a feature branch:
   ```
   git checkout -b feature/your-feature-name
   ```
3. Make your changes and ensure the code compiles without errors.
4. Run tests:
   ```
   go test ./...
   ```
5. Submit a pull request with a detailed description of the changes.

## Easy Installation

### Option 1: Automatic Install Script

#### Linux and macOS

```bash
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
LATEST_RELEASE=$(curl -s https://api.github.com/latentsp/remote-mcp/remote-mcp/releases/latest)
VERSION=$(echo $LATEST_RELEASE | grep -o '"tag_name": "[^"]*' | cut -d'"' -f4)

# Download URL
DOWNLOAD_URL="https://github.com/yourusername/remote-mcp/releases/download/$VERSION/remote-mcp-$OS-$ARCH"
echo "Downloading remote-mcp $VERSION for $OS-$ARCH..."
curl -L -o remote-mcp "$DOWNLOAD_URL"

# Make executable and move to bin directory
chmod +x remote-mcp
sudo mv remote-mcp /usr/local/bin/
echo "remote-mcp successfully installed to /usr/local/bin/"
echo "Run 'remote-mcp --help' to get started"
```

Save this script to a file (e.g., `install.sh`), make it executable with `chmod +x install.sh`, then run it with `./install.sh`.

#### Windows (PowerShell)

```powershell
# Determine architecture
$arch = "amd64"
if ($env:PROCESSOR_ARCHITECTURE -eq "ARM64") {
    $arch = "arm64"
}

# Get latest release
Write-Host "Fetching the latest release..."
$latestRelease = Invoke-RestMethod -Uri "https://api.github.com/repos/yourusername/remote-mcp/releases/latest"
$version = $latestRelease.tag_name

# Download URL
$downloadUrl = "https://github.com/yourusername/remote-mcp/releases/download/$version/remote-mcp-windows-$arch.exe"
$outputPath = "$env:USERPROFILE\remote-mcp.exe"

Write-Host "Downloading remote-mcp $version for Windows-$arch..."
Invoke-WebRequest -Uri $downloadUrl -OutFile $outputPath

# Add to PATH
$binPath = "C:\Windows\System32"
Copy-Item -Path $outputPath -Destination "$binPath\remote-mcp.exe" -Force
Write-Host "remote-mcp successfully installed to $binPath"
Write-Host "Run 'remote-mcp --help' to get started"
```

Save this script to a file (e.g., `install.ps1`), then run it in PowerShell with administrator privileges.

### Option 2: Manual Installation

You can also manually download the latest version from the [Releases](https://github.com/yourusername/remote-mcp/releases) page and:

#### Linux and macOS:
```bash
chmod +x remote-mcp
sudo mv remote-mcp /usr/local/bin/
```

#### Windows:
Move the downloaded `.exe` file to a directory in your PATH.
