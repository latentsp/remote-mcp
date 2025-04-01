# Remote MCP Client

A command-line client for interacting with remote MCP servers.

## Overview

This client establishes a connection to a remote MCP server, relays commands from standard input, and forwards responses back to standard output. It uses Server-Sent Events (SSE) to maintain a persistent connection with the server and discover the endpoint URL.

## Installation

### Unix-like systems (Linux, macOS)

```bash
curl -s https://raw.githubusercontent.com/latentsp/remote-mcp/main/scripts/install.sh | bash
```

### Windows (PowerShell)

```powershell
Invoke-WebRequest -Uri https://raw.githubusercontent.com/latentsp/remote-mcp/main/scripts/install.ps1 -OutFile install.ps1; .\install.ps1
```

The installation script will automatically:
- Detect your system architecture (amd64 or arm64)
- Download the appropriate binary
- Install it to the correct location
- Add it to your system PATH

### Manual Installation

If you prefer to install manually, you can:
1. Visit the [Releases](https://github.com/latentsp/remote-mcp/releases) page
2. Download the appropriate binary for your platform and architecture
3. Make the file executable (Linux/macOS):
   ```bash
   chmod +x remote-mcp
   ```
4. Move it to a directory in your PATH:
   - Linux/macOS: `/usr/local/bin/`
   - Windows: `C:\Program Files\remote-mcp\`

## Usage

```sh
remote-mcp --endpoint-url=<remote-mcp-server-url> [--verbose]
```

## Usage inside applications (Cursor, Claude, etc.)

Configure your application
```json
{
  "mcpServers": {
      "command": "remote-mcp",
      "args": [
         "--endpoint-url",
         "<your remote mcp server>"
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
