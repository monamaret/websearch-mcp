# Building WebSearch MCP Server

This guide covers building the WebSearch MCP Server from source for multiple platforms and architectures.

## 📋 Table of Contents

- [Prerequisites](#prerequisites)
- [Quick Start](#quick-start)
- [Platform-Specific Builds](#platform-specific-builds)
- [Build for All Platforms](#build-for-all-platforms)
- [Build Configuration](#build-configuration)
- [Makefile Targets](#makefile-targets)
- [Verifying Builds](#verifying-builds)
- [Troubleshooting](#troubleshooting)

## Prerequisites

### Required

- **Go 1.21 or later**: [Download Go](https://golang.org/dl/)
- **Git**: For cloning the repository
- **Make**: For using the Makefile (optional but recommended)

### Optional

- **zip**: For creating Windows archives (usually pre-installed)
- **tar**: For creating Unix archives (usually pre-installed)

## Quick Start

Build for your current platform:

```bash
# Clone the repository
git clone https://github.com/yourusername/websearch-mcp.git
cd websearch-mcp

# Download dependencies
go mod download

# Build
make build

# Or without Make
go build -o websearch-mcp .
```

## Platform-Specific Builds

### macOS

#### Apple Silicon (M1/M2/M3)

```bash
# Using Make
make build

# Or manually
GOOS=darwin GOARCH=arm64 go build -o websearch-mcp-darwin-arm64 .
```

#### Intel

```bash
# Manually (cross-compile on Apple Silicon)
GOOS=darwin GOARCH=amd64 go build -o websearch-mcp-darwin-amd64 .
```

### Windows

#### Intel/AMD64

```bash
# Using Make (on macOS/Linux)
GOOS=windows GOARCH=amd64 go build -o websearch-mcp-windows-amd64.exe .
```

#### ARM64

```bash
# Using Make (on macOS/Linux)
GOOS=windows GOARCH=arm64 go build -o websearch-mcp-windows-arm64.exe .
```

**Note for Windows users**: If building on Windows, use PowerShell or Command Prompt:

```powershell
# Build for Windows
go build -o websearch-mcp.exe .

# Cross-compile for ARM64 (from Windows)
$env:GOOS="windows"; $env:GOARCH="arm64"; go build -o websearch-mcp-arm64.exe .
```

### Linux

#### Intel/AMD64

```bash
# Using Make
GOOS=linux GOARCH=amd64 go build -o websearch-mcp-linux-amd64 .
```

#### ARM64

```bash
# Using Make
GOOS=linux GOARCH=arm64 go build -o websearch-mcp-linux-arm64 .
```

## Build for All Platforms

Build binaries for all supported platforms at once:

```bash
make build-all-platforms
```

This will create:
- Binaries in `dist/` directory
- Compressed archives (.tar.gz for Unix, .zip for Windows)
- SHA256 checksums in `dist/checksums.txt`

### Supported Platforms

| Platform | Architecture | Binary Name | Archive Format |
|----------|--------------|-------------|----------------|
| macOS | Intel (amd64) | `websearch-mcp-darwin-amd64` | `.tar.gz` |
| macOS | Apple Silicon (arm64) | `websearch-mcp-darwin-arm64` | `.tar.gz` |
| Windows | Intel (amd64) | `websearch-mcp-windows-amd64.exe` | `.zip` |
| Windows | ARM64 | `websearch-mcp-windows-arm64.exe` | `.zip` |
| Linux | Intel (amd64) | `websearch-mcp-linux-amd64` | `.tar.gz` |
| Linux | ARM64 | `websearch-mcp-linux-arm64` | `.tar.gz` |

## Build Configuration

### Version Information

The build process embeds version information into the binary:

```bash
# Automatic version detection (from git tags)
make build

# Custom version
VERSION=v1.2.3 make build
```

Version information is set via ldflags:
- `version`: Git tag or "dev"
- `buildTime`: RFC3339 timestamp
- `gitCommit`: Short git commit hash

### Build Flags

The default build uses these flags for optimization:

```bash
-ldflags="-w -s -X main.version=... -X main.buildTime=... -X main.gitCommit=..."
-a -installsuffix cgo
```

Explanation:
- `-w`: Omit DWARF symbol table
- `-s`: Omit symbol table and debug info
- `-X`: Set version variables
- `-a`: Force rebuilding of packages
- `-installsuffix cgo`: Add suffix to package installation directory
- `CGO_ENABLED=0`: Disable CGO for static binary

### Custom Build

For a custom build with different flags:

```bash
go build \
  -ldflags="-X main.version=custom" \
  -o my-custom-binary \
  .
```

## Makefile Targets

The Makefile provides several convenient targets:

### Basic Targets

```bash
# Build for current platform
make build

# Build optimized release binary
make build-release

# Build for all platforms
make build-all-platforms

# Run the server
make run

# Clean build artifacts
make clean
```

### Development Targets

```bash
# Download dependencies
make deps

# Run tests
make test

# Format code
make fmt

# Run go vet
make vet

# Run linter (requires golangci-lint)
make lint

# Development mode with hot reload
make dev
```

### Information Targets

```bash
# Show version information
make version

# Show help
make help
```

## Verifying Builds

### Check Binary

After building, verify the binary:

```bash
# macOS/Linux
file websearch-mcp-darwin-arm64
ls -lh websearch-mcp-darwin-arm64

# Check it runs
./websearch-mcp-darwin-arm64 --help
```

### Verify Checksums

When building all platforms:

```bash
# Check checksums
cd dist
sha256sum -c checksums.txt

# Or on macOS
shasum -a 256 -c checksums.txt
```

### Test Binary

Run a quick test:

```bash
# Start server
./websearch-mcp-darwin-arm64 &
SERVER_PID=$!

# Wait for startup
sleep 2

# Test health endpoint
curl http://localhost:8080/health

# Stop server
kill $SERVER_PID
```

## Troubleshooting

### Common Build Issues

#### 1. Go Version Too Old

**Error**: `go: go.mod requires go >= 1.21`

**Solution**: Update Go to version 1.21 or later
```bash
# Check version
go version

# Download from https://golang.org/dl/
```

#### 2. Missing Dependencies

**Error**: `cannot find package...`

**Solution**: Download dependencies
```bash
go mod download
go mod tidy
```

#### 3. Cross-Compilation Fails

**Error**: Build fails when cross-compiling

**Solution**: Ensure CGO is disabled
```bash
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build .
```

#### 4. Archive Creation Fails

**Error**: `tar: command not found` or `zip: command not found`

**Solution**: 
- On macOS: These are pre-installed
- On Linux: Install via package manager
  ```bash
  # Debian/Ubuntu
  sudo apt-get install tar zip
  
  # CentOS/RHEL
  sudo yum install tar zip
  ```
- On Windows: Use Git Bash or WSL

#### 5. Permission Denied

**Error**: Cannot execute binary

**Solution**: Make it executable
```bash
chmod +x websearch-mcp-darwin-arm64
```

### Platform-Specific Issues

#### macOS

**Issue**: "Cannot open because the developer cannot be verified"

**Solution**: Remove quarantine attribute
```bash
xattr -d com.apple.quarantine websearch-mcp-darwin-arm64
```

Or allow in System Preferences → Security & Privacy

#### Windows

**Issue**: Windows Defender blocks execution

**Solution**: Add exception in Windows Security settings

#### Linux

**Issue**: Executable format error on ARM

**Solution**: Ensure you're using the correct architecture binary
```bash
# Check your architecture
uname -m

# arm64, aarch64 → use arm64 binary
# x86_64 → use amd64 binary
```

### Build Performance

#### Slow Builds

Enable module caching:
```bash
# Set cache directory
export GOMODCACHE=$HOME/go/pkg/mod

# Use build cache
go env -w GOCACHE=$HOME/.cache/go-build
```

#### Parallel Builds

When building multiple platforms manually:
```bash
# Build in parallel (bash)
for platform in darwin/amd64 darwin/arm64 linux/amd64 linux/arm64 windows/amd64 windows/arm64; do
  GOOS=${platform%/*} GOARCH=${platform#*/} go build -o dist/websearch-mcp-${platform%/*}-${platform#*/} . &
done
wait
```

## Advanced Topics

### Static vs Dynamic Linking

Our builds use static linking (`CGO_ENABLED=0`) for portability. For dynamic linking:

```bash
# Enable CGO (may require C toolchain)
CGO_ENABLED=1 go build .
```

### Optimization Levels

Different optimization approaches:

```bash
# Default (fastest build)
go build .

# Optimized size (our default)
go build -ldflags="-w -s" .

# Optimized speed
go build -gcflags="-N -l" .

# Profile-guided optimization
go build -pgo=auto .
```

### Custom Target OS/Arch

View all supported platforms:
```bash
go tool dist list
```

Build for other platforms:
```bash
# FreeBSD
GOOS=freebsd GOARCH=amd64 go build .

# OpenBSD
GOOS=openbsd GOARCH=amd64 go build .

# Raspberry Pi
GOOS=linux GOARCH=arm GOARM=7 go build .
```

## Next Steps

After building:

1. **Test the binary**: Run basic functionality tests
2. **Configure MCP**: Set up `.mcp_servers` configuration
3. **Integration**: Connect with Tabnine or other MCP clients
4. **Deployment**: Deploy to your target environment

See also:
- [USAGE.md](USAGE.md) - Using the server
- [TABNINE_SETUP.md](TABNINE_SETUP.md) - Tabnine integration
- [WORKFLOWS.md](WORKFLOWS.md) - CI/CD workflows
