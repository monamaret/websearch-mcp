# Multi-Platform Support Update

This document summarizes the multi-platform build support added to the WebSearch MCP Server project.

## 📋 Summary

The WebSearch MCP Server now supports building and distributing binaries for **6 different platform/architecture combinations**:

### Supported Platforms

1. **macOS Intel (x86_64/amd64)** - Intel-based Macs
2. **macOS Apple Silicon (ARM64)** - M1/M2/M3 Macs
3. **Windows Intel (x86_64/amd64)** - Standard Windows PCs
4. **Windows ARM64** - Windows on ARM devices
5. **Linux Intel (x86_64/amd64)** - Standard Linux servers/desktops
6. **Linux ARM64** - ARM-based Linux systems, Raspberry Pi 4+

## 🔧 What Changed

### 1. Makefile Updates

**File:** `Makefile`

**Changes:**
- Updated `build-all-platforms` target to include Windows ARM64
- Improved build messages with clearer platform/architecture labels
- Enhanced help text with platform details and documentation references

**New capabilities:**
```bash
make build-all-platforms
```
Now builds 6 binaries instead of 5, including Windows ARM64 support.

### 2. GitHub Workflows

#### Build Workflow

**File:** `.github/workflows/build.yml`

**Major changes:**
- Converted from single-build to **matrix strategy** for parallel builds
- Builds run on native runners for each OS:
  - `macos-latest` for macOS builds
  - `windows-latest` for Windows builds
  - `ubuntu-latest` for Linux builds
- Cross-compilation for different architectures (amd64 and arm64)
- Platform-specific archive creation:
  - `.tar.gz` for Unix-like systems
  - `.zip` for Windows
- Enhanced snapshot releases with detailed platform instructions
- Combined checksums file for all platforms

**Benefits:**
- Faster builds through parallelization
- Platform-native build environment
- Comprehensive testing across all platforms

#### Release Workflow

**File:** `.github/workflows/release.yml`

**Major changes:**
- Implemented matrix strategy matching build workflow
- Creates 6 release archives per release
- Enhanced release notes with:
  - Platform-specific installation instructions
  - Download links for each platform
  - Verification instructions
- Combined checksums for all platforms

**Release assets now include:**
- 6 binary archives (3 for macOS/Linux as `.tar.gz`, 2 for Windows as `.zip`)
- Single `checksums.txt` with all SHA256 hashes

### 3. Documentation

#### New Documentation Files

1. **`docs/BUILDING.md`** - Comprehensive build guide
   - Prerequisites for each platform
   - Platform-specific build instructions
   - Build configuration and customization
   - Makefile targets reference
   - Verification and testing procedures
   - Troubleshooting guide
   - Advanced build topics

2. **`docs/PLATFORM_SUPPORT.md`** - Platform compatibility guide
   - Detailed platform requirements
   - Installation instructions per platform
   - Platform-specific issues and solutions
   - Running as a service (systemd, launchd, Windows Service)
   - Platform detection scripts
   - Upgrade procedures

3. **`docs/RELEASE_GUIDE.md`** - Release management guide
   - Pre-release checklist
   - Step-by-step release creation
   - Release types (stable, pre-release, snapshot)
   - Post-release tasks
   - Rollback procedures
   - Best practices

4. **`docs/MULTI_PLATFORM_UPDATE.md`** - This file
   - Summary of multi-platform support
   - Migration guide
   - Quick reference

#### Updated Documentation Files

1. **`README.md`**
   - Added Quick Start section with platform-specific download instructions
   - Added Supported Platforms table
   - Updated configuration examples with platform-specific binary names
   - Added links to new documentation
   - Enhanced troubleshooting section

2. **`docs/WORKFLOWS.md`**
   - Completely rewritten for matrix builds
   - Added platform-specific build details
   - Enhanced with troubleshooting section
   - Added metrics and best practices
   - Included instructions for adding new platforms

## 📦 Binary Naming Convention

### Format
```
websearch-mcp-{platform}-{architecture}{extension}
```

### Examples
- `websearch-mcp-darwin-amd64` (macOS Intel, no extension)
- `websearch-mcp-darwin-arm64` (macOS Apple Silicon, no extension)
- `websearch-mcp-windows-amd64.exe` (Windows Intel)
- `websearch-mcp-windows-arm64.exe` (Windows ARM)
- `websearch-mcp-linux-amd64` (Linux Intel, no extension)
- `websearch-mcp-linux-arm64` (Linux ARM, no extension)

### Archive Naming
```
websearch-mcp-{version}-{platform}-{architecture}.{tar.gz|zip}
```

Examples:
- `websearch-mcp-v1.2.3-darwin-arm64.tar.gz`
- `websearch-mcp-v1.2.3-windows-amd64.zip`

## 🚀 Usage Changes

### Before (Single Binary)

```bash
# Build
make build

# Output
websearch-mcp

# Run
./websearch-mcp
```

### After (Platform-Specific Binaries)

```bash
# Build for current platform
make build

# Output (example on macOS ARM)
websearch-mcp

# Build for all platforms
make build-all-platforms

# Output in dist/
websearch-mcp-darwin-amd64
websearch-mcp-darwin-arm64
websearch-mcp-windows-amd64.exe
websearch-mcp-windows-arm64.exe
websearch-mcp-linux-amd64
websearch-mcp-linux-arm64

# Run (using full platform name)
./websearch-mcp-darwin-arm64
```

## 🔄 Migration Guide

### For Users

#### Updating Configuration Files

**Old `.mcp_servers` configuration:**
```json
{
  "mcpServers": {
    "websearch": {
      "command": "./websearch-mcp",
      "args": [],
      "env": {
        "PORT": "8080"
      }
    }
  }
}
```

**New `.mcp_servers` configuration:**
```json
{
  "mcpServers": {
    "websearch": {
      "command": "./websearch-mcp-darwin-arm64",
      "args": [],
      "env": {
        "PORT": "8080"
      }
    }
  }
}
```

**Platform-specific command values:**
- macOS Apple Silicon: `./websearch-mcp-darwin-arm64`
- macOS Intel: `./websearch-mcp-darwin-amd64`
- Windows Intel: `websearch-mcp-windows-amd64.exe`
- Windows ARM: `websearch-mcp-windows-arm64.exe`
- Linux Intel: `./websearch-mcp-linux-amd64`
- Linux ARM: `./websearch-mcp-linux-arm64`

### For Developers

#### Building

**Old method:**
```bash
make build
# or
go build -o websearch-mcp .
```

**New method (same, but output differs):**
```bash
# Build for current platform
make build
# Output: websearch-mcp (local builds keep simple name)

# Build for all platforms
make build-all-platforms
# Output: 6 platform-specific binaries in dist/
```

#### Release Process

**Old process:**
- Created 1 binary archive per release

**New process:**
- Creates 6 binary archives per release
- Each archive optimized for its platform
- Combined checksums file
- Enhanced release notes

**No changes required** - automated via GitHub Actions when you push a tag.

## 📊 Benefits

### For Users

1. **Native Performance**: Binaries optimized for each platform
2. **No Cross-Platform Issues**: Platform-specific builds avoid compatibility problems
3. **Easy Selection**: Clear naming makes it obvious which binary to download
4. **Better ARM Support**: Native ARM builds for Apple Silicon and Windows ARM
5. **Verification**: SHA256 checksums for all platforms

### For Developers

1. **Parallel Builds**: Faster CI/CD with matrix strategy
2. **Platform Testing**: Each build runs on its native OS
3. **Easier Debugging**: Platform-specific issues isolated
4. **Future-Proof**: Easy to add new platforms
5. **Automated**: No manual intervention needed

### For Maintainers

1. **Clear Documentation**: Comprehensive guides for all aspects
2. **Consistent Naming**: Standard naming across all platforms
3. **Version Tracking**: Version embedded in each binary
4. **Release Automation**: Fully automated release process
5. **Quality Assurance**: Each platform built and tested independently

## 🎯 Quick Reference

### Download Pre-built Binary

Visit: `https://github.com/youruser/websearch-mcp/releases/latest`

Choose your platform:
- macOS Intel → `*-darwin-amd64.tar.gz`
- macOS Apple Silicon → `*-darwin-arm64.tar.gz`
- Windows Intel → `*-windows-amd64.zip`
- Windows ARM → `*-windows-arm64.zip`
- Linux Intel → `*-linux-amd64.tar.gz`
- Linux ARM → `*-linux-arm64.tar.gz`

### Build from Source

```bash
# Clone
git clone <repository-url>
cd websearch-mcp

# Build for current platform
make build

# Build for all platforms
make build-all-platforms

# Output in dist/
ls -lh dist/
```

### Create a Release

```bash
# 1. Update CHANGES.md
vim docs/CHANGES.md

# 2. Commit and push
git add docs/CHANGES.md
git commit -m "Update CHANGES for v1.2.3"
git push

# 3. Create and push tag
git tag -a v1.2.3 -m "Release v1.2.3"
git push origin v1.2.3

# 4. GitHub Actions handles the rest!
```

## 📚 Documentation Index

All documentation is in the `docs/` folder:

- **[BUILDING.md](BUILDING.md)** - How to build from source
- **[PLATFORM_SUPPORT.md](PLATFORM_SUPPORT.md)** - Platform compatibility and installation
- **[RELEASE_GUIDE.md](RELEASE_GUIDE.md)** - Creating and managing releases
- **[WORKFLOWS.md](WORKFLOWS.md)** - CI/CD workflows documentation
- **[USAGE.md](USAGE.md)** - How to use the server
- **[TABNINE_SETUP.md](TABNINE_SETUP.md)** - Tabnine integration guide

## ❓ FAQ

**Q: Why platform-specific binaries instead of one universal binary?**

A: Platform-specific binaries provide:
- Better performance (native compilation)
- Smaller file sizes
- No cross-platform compatibility issues
- Clear identification of supported platforms

**Q: Can I still build a simple binary for local development?**

A: Yes! `make build` creates a simple `websearch-mcp` binary for quick local development.

**Q: Do I need to update my configuration files?**

A: Only if you're using absolute paths. Update the binary name to match your platform (e.g., `websearch-mcp-darwin-arm64`).

**Q: What if my platform isn't supported?**

A: You can build from source! Go supports many platforms. See [BUILDING.md](BUILDING.md) for instructions.

**Q: How do I verify my download?**

A: Use the SHA256 checksums in `checksums.txt`:
```bash
# Unix
sha256sum -c checksums.txt

# Windows
Get-FileHash websearch-mcp-*.exe -Algorithm SHA256
```

**Q: Can I request support for a new platform?**

A: Yes! Open an issue on GitHub describing the platform and your use case.

## 🔜 Future Enhancements

Potential future additions:

1. **Additional Platforms**
   - FreeBSD (amd64, arm64)
   - OpenBSD (amd64)
   - 32-bit architectures (if needed)

2. **Distribution Methods**
   - Homebrew formula (macOS)
   - Chocolatey package (Windows)
   - APT/YUM repositories (Linux)
   - Docker images (multi-arch)

3. **Build Improvements**
   - Signed binaries
   - Notarized macOS binaries
   - Windows installer
   - Universal macOS binary (combined Intel + ARM)

## 📞 Support

For questions or issues:

1. Check documentation in `docs/` folder
2. Search existing GitHub issues
3. Open a new issue with:
   - Your platform and architecture
   - Steps to reproduce
   - Error messages or logs

## 🙏 Credits

Multi-platform support implemented using:
- Go's cross-compilation capabilities
- GitHub Actions matrix strategy
- Best practices from the Go community

---

**Last Updated:** 2024-01-15  
**Version:** 1.0.0  
**Status:** Complete
