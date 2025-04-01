# PowerShell script to install remote-mcp

# Determine architecture
$arch = "amd64"
if ($env:PROCESSOR_ARCHITECTURE -eq "ARM64") {
    $arch = "arm64"
}

# Get latest release
Write-Host "Fetching the latest release..."
$latestRelease = Invoke-RestMethod -Uri "https://api.github.com/repos/latentsp/remote-mcp/releases/latest"
$version = $latestRelease.tag_name

# Download URL
$downloadUrl = "https://github.com/latentsp/remote-mcp/releases/download/$version/remote-mcp-windows-$arch.exe"
$outputPath = "$env:USERPROFILE\remote-mcp.exe"

Write-Host "Downloading remote-mcp $version for Windows-$arch..."
Invoke-WebRequest -Uri $downloadUrl -OutFile $outputPath

# Add to PATH
$binPath = "C:\Windows\System32"
Copy-Item -Path $outputPath -Destination "$binPath\remote-mcp.exe" -Force
Write-Host "remote-mcp successfully installed to $binPath"
Write-Host "Run 'remote-mcp --help' to get started"