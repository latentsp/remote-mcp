# Remote MCP Server

A Model Context Protocol (MCP) server implementation in Go. This server provides tools and resources for AI agents to interact with your development environment.

## Features

- MCP server implementation following the [Model Context Protocol](https://modelcontextprotocol.io/) specifications
- Cross-platform support (Linux, macOS, Windows)
- Automated builds for multiple platforms

### Usage

1. Download the relevant binary
2. Run `remote-mcp --endpoint-url=<URL OF SSE MCP SERVER>`

## Getting Started

### Pre-built Binaries

You can download pre-built binaries from the [GitHub Releases](https://github.com/yourusername/remote-mcp/releases) page. The releases include binaries for:
- Linux (amd64)
- macOS (amd64)
- Windows (amd64)

Each release is tagged with the commit SHA from the main branch (e.g., `main-abc1234`).

### Building from Source

To build the project from source:

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/remote-mcp.git
   cd remote-mcp
   ```

2. Make sure you have Go 1.22 or later installed.

3. Build the project:
   ```bash
   go build -o mcp-server ./...
   ```

## Development

### Automated Builds

This project uses GitHub Actions for automated builds. Every push to the `main` branch will:
1. Build binaries for all supported platforms
2. Create a new release tagged with the commit SHA
3. Upload the built binaries as release assets

The releases are marked as pre-releases to indicate they are development builds from the main branch.

## License

[Add your license information here]

## Contributing

[Add contribution guidelines here]