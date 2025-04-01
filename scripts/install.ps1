# PowerShell script to install remote-mcp

# Determine architecture
$ARCH = if ($env:PROCESSOR_ARCHITECTURE -eq "ARM64") { "arm64" } else { "amd64" }

# Determine OS
$OS = "windows"

# Get latest release
Write-Host "Fetching the latest release..."
$LATEST_RELEASE = Invoke-RestMethod -Uri "https://api.github.com/repos/latentsp/remote-mcp/releases/latest"
$VERSION = $LATEST_RELEASE.tag_name

# Download URL
$DOWNLOAD_URL = "https://github.com/latentsp/remote-mcp/releases/download/$VERSION/remote-mcp-$OS-$ARCH.exe"
Write-Host "Downloading remote-mcp $VERSION for $OS-$ARCH..."

# Create temp directory
$TEMP_DIR = Join-Path $env:TEMP "remote-mcp-install"
New-Item -ItemType Directory -Force -Path $TEMP_DIR | Out-Null

# Download the binary
$BINARY_PATH = Join-Path $TEMP_DIR "remote-mcp.exe"
Invoke-WebRequest -Uri $DOWNLOAD_URL -OutFile $BINARY_PATH

# Move to Program Files
$INSTALL_DIR = "C:\Program Files\remote-mcp"
New-Item -ItemType Directory -Force -Path $INSTALL_DIR | Out-Null
Move-Item -Force $BINARY_PATH (Join-Path $INSTALL_DIR "remote-mcp.exe")

# Add to PATH if not already present
$PATH = [Environment]::GetEnvironmentVariable("Path", "Machine")
if ($PATH -notlike "*$INSTALL_DIR*") {
    [Environment]::SetEnvironmentVariable("Path", $PATH + ";$INSTALL_DIR", "Machine")
}

Write-Host "remote-mcp successfully installed to $INSTALL_DIR"
Write-Host "Run 'remote-mcp --help' to get started"