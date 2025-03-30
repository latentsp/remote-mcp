# Remote MCP Server

A Model Context Protocol (MCP) server implementation in Go. This server provides tools and resources for AI agents to interact with your development environment.

## Features

- MCP server implementation following the [Model Context Protocol](https://modelcontextprotocol.io/) specifications
- Cross-platform support (Linux, macOS, Windows)
- Automated releases for multiple platforms

## Getting Started

### Pre-built Binaries

You can download pre-built binaries from the [GitHub Releases](https://github.com/yourusername/remote-mcp/releases) page. The releases include binaries for:
- Linux (amd64)
- macOS (amd64)
- Windows (amd64)

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

### Release Process

This project uses GitHub Actions for automated releases. The release process is triggered by:
- Pushing to the `main` branch
- Creating and pushing a version tag (e.g., `v1.0.0`)

The workflow will:
1. Build binaries for all supported platforms
2. Create a GitHub release with the version tag
3. Upload the built binaries as release assets

To create a new release:
1. Update the version in your code
2. Create and push a new tag:
   ```bash
   git tag v1.0.0
   git push origin v1.0.0
   ```

## License

[Add your license information here]

## Contributing

[Add contribution guidelines here]