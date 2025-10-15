# Platform Support

This document describes the platforms and architectures supported by WebSearch MCP Server, along with installation and compatibility notes.

## 📊 Supported Platforms

WebSearch MCP Server is built and tested on the following platforms:

### ✅ Fully Supported

| Platform | Architecture | Binary Available | Tested | Notes |
|----------|--------------|------------------|--------|-------|
| macOS 11+ | Apple Silicon (ARM64) | ✅ | ✅ | M1, M2, M3 chips |
| macOS 11+ | Intel (x86_64) | ✅ | ✅ | Intel-based Macs |
| Windows 10/11 | Intel/AMD (x86_64) | ✅ | ✅ | 64-bit Windows |
| Windows 10/11 | ARM64 | ✅ | ⚠️ | Windows on ARM devices |
| Linux | Intel/AMD (x86_64) | ✅ | ✅ | Most distributions |
| Linux | ARM64 | ✅ | ✅ | ARM servers, Raspberry Pi 4+ |

### Legend
- ✅ Fully tested and supported
- ⚠️ Built but limited testing
- ❌ Not supported

## 🖥️ Platform-Specific Details

### macOS

#### System Requirements
- **macOS Version**: 11.0 (Big Sur) or later
- **Architectures**: 
  - Apple Silicon (ARM64): M1, M2, M3 chips
  - Intel (x86_64): 64-bit Intel processors
- **Permissions**: May require security approval on first run

#### Installation

##### Apple Silicon (ARM64)
```bash
# Download the ARM64 binary
curl -LO https://github.com/youruser/websearch-mcp/releases/latest/download/websearch-mcp-VERSION-darwin-arm64.tar.gz

# Extract
tar -xzf websearch-mcp-VERSION-darwin-arm64.tar.gz

# Make executable
chmod +x websearch-mcp-darwin-arm64

# Run
./websearch-mcp-darwin-arm64
```

##### Intel (x86_64)
```bash
# Download the Intel binary
curl -LO https://github.com/youruser/websearch-mcp/releases/latest/download/websearch-mcp-VERSION-darwin-amd64.tar.gz

# Extract
tar -xzf websearch-mcp-VERSION-darwin-amd64.tar.gz

# Make executable
chmod +x websearch-mcp-darwin-amd64

# Run
./websearch-mcp-darwin-amd64
```

#### Common Issues

**"Cannot open because the developer cannot be verified"**

Solution 1: Remove quarantine attribute
```bash
xattr -d com.apple.quarantine websearch-mcp-darwin-arm64
```

Solution 2: System Preferences
1. Open System Preferences → Security & Privacy
2. Click "Open Anyway" after the first blocked attempt

**Rosetta 2 Compatibility**

Apple Silicon Macs can run Intel binaries through Rosetta 2:
```bash
# Install Rosetta 2 (if needed)
softwareupdate --install-rosetta

# Run Intel binary on Apple Silicon
./websearch-mcp-darwin-amd64
```

However, use the native ARM64 binary for best performance.

### Windows

#### System Requirements
- **Windows Version**: Windows 10 (version 1809+) or Windows 11
- **Architectures**:
  - Intel/AMD (x86_64): 64-bit processors
  - ARM64: Windows on ARM devices (Surface Pro X, etc.)
- **Runtime**: No additional runtime required (statically compiled)

#### Installation

##### Intel/AMD (x86_64)

**PowerShell:**
```powershell
# Download (replace VERSION)
Invoke-WebRequest -Uri "https://github.com/youruser/websearch-mcp/releases/latest/download/websearch-mcp-VERSION-windows-amd64.zip" -OutFile "websearch-mcp.zip"

# Extract
Expand-Archive -Path websearch-mcp.zip -DestinationPath .

# Run
.\websearch-mcp-windows-amd64.exe
```

**Command Prompt:**
```cmd
REM Download manually from GitHub releases

REM Extract using File Explorer or:
tar -xf websearch-mcp-VERSION-windows-amd64.zip

REM Run
websearch-mcp-windows-amd64.exe
```

##### ARM64

```powershell
# Download ARM64 version
Invoke-WebRequest -Uri "https://github.com/youruser/websearch-mcp/releases/latest/download/websearch-mcp-VERSION-windows-arm64.zip" -OutFile "websearch-mcp-arm64.zip"

# Extract and run
Expand-Archive -Path websearch-mcp-arm64.zip -DestinationPath .
.\websearch-mcp-windows-arm64.exe
```

#### Common Issues

**Windows Defender SmartScreen Warning**

Solution 1: Click "More info" → "Run anyway"

Solution 2: Add to exclusions
```powershell
# Add folder to exclusions (requires admin)
Add-MpPreference -ExclusionPath "C:\path\to\websearch-mcp"
```

**Windows Firewall Blocking**

Allow through firewall:
```powershell
# Create firewall rule (requires admin)
New-NetFirewallRule -DisplayName "WebSearch MCP" -Direction Inbound -Program "C:\path\to\websearch-mcp-windows-amd64.exe" -Action Allow
```

**Port Already in Use**

Change the port:
```powershell
$env:PORT = "3000"
.\websearch-mcp-windows-amd64.exe
```

### Linux

#### System Requirements
- **Kernel**: Linux 3.10+ (most modern distributions)
- **Architectures**:
  - Intel/AMD (x86_64): 64-bit processors
  - ARM64: ARM v8 processors (including Raspberry Pi 4+)
- **Dependencies**: None (statically linked)

#### Installation

##### Intel/AMD (x86_64)

```bash
# Download
wget https://github.com/youruser/websearch-mcp/releases/latest/download/websearch-mcp-VERSION-linux-amd64.tar.gz

# Extract
tar -xzf websearch-mcp-VERSION-linux-amd64.tar.gz

# Make executable
chmod +x websearch-mcp-linux-amd64

# Run
./websearch-mcp-linux-amd64
```

##### ARM64

```bash
# Download
wget https://github.com/youruser/websearch-mcp/releases/latest/download/websearch-mcp-VERSION-linux-arm64.tar.gz

# Extract
tar -xzf websearch-mcp-VERSION-linux-arm64.tar.gz

# Make executable
chmod +x websearch-mcp-linux-arm64

# Run
./websearch-mcp-linux-arm64
```

#### Tested Distributions

| Distribution | Version | Status | Notes |
|--------------|---------|--------|-------|
| Ubuntu | 20.04, 22.04, 24.04 | ✅ | Fully supported |
| Debian | 11, 12 | ✅ | Fully supported |
| CentOS/RHEL | 8, 9 | ✅ | Fully supported |
| Fedora | 38, 39, 40 | ✅ | Fully supported |
| Alpine | 3.18+ | ✅ | Fully supported |
| Arch Linux | Rolling | ✅ | Fully supported |
| Raspberry Pi OS | 11+ (64-bit) | ✅ | ARM64 binary required |

#### Common Issues

**Permission Denied**

```bash
# Make executable
chmod +x websearch-mcp-linux-amd64

# If needed, run with sudo
sudo ./websearch-mcp-linux-amd64
```

**Wrong Architecture**

Check your architecture:
```bash
uname -m
# x86_64 → use amd64 binary
# aarch64 or arm64 → use arm64 binary
```

**Firewall Configuration**

```bash
# UFW (Ubuntu/Debian)
sudo ufw allow 8080/tcp

# firewalld (CentOS/RHEL/Fedora)
sudo firewall-cmd --add-port=8080/tcp --permanent
sudo firewall-cmd --reload

# iptables
sudo iptables -A INPUT -p tcp --dport 8080 -j ACCEPT
```

## 🚀 Running as a Service

### systemd (Linux)

Create `/etc/systemd/system/websearch-mcp.service`:

```ini
[Unit]
Description=WebSearch MCP Server
After=network.target

[Service]
Type=simple
User=www-data
WorkingDirectory=/opt/websearch-mcp
ExecStart=/opt/websearch-mcp/websearch-mcp-linux-amd64
Restart=always
RestartSec=10
Environment="PORT=8080"

[Install]
WantedBy=multi-user.target
```

Enable and start:
```bash
sudo systemctl daemon-reload
sudo systemctl enable websearch-mcp
sudo systemctl start websearch-mcp
sudo systemctl status websearch-mcp
```

### launchd (macOS)

Create `~/Library/LaunchAgents/com.websearch-mcp.plist`:

```xml
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>Label</key>
    <string>com.websearch-mcp</string>
    <key>ProgramArguments</key>
    <array>
        <string>/usr/local/bin/websearch-mcp-darwin-arm64</string>
    </array>
    <key>RunAtLoad</key>
    <true/>
    <key>KeepAlive</key>
    <true/>
    <key>StandardOutPath</key>
    <string>/tmp/websearch-mcp.log</string>
    <key>StandardErrorPath</key>
    <string>/tmp/websearch-mcp.error.log</string>
</dict>
</plist>
```

Load and start:
```bash
launchctl load ~/Library/LaunchAgents/com.websearch-mcp.plist
launchctl start com.websearch-mcp
```

### Windows Service

Using NSSM (Non-Sucking Service Manager):

```powershell
# Download NSSM from https://nssm.cc/download

# Install service
nssm install WebSearchMCP "C:\path\to\websearch-mcp-windows-amd64.exe"

# Configure
nssm set WebSearchMCP AppDirectory "C:\path\to"
nssm set WebSearchMCP AppEnvironmentExtra "PORT=8080"

# Start service
nssm start WebSearchMCP
```

## 🔍 Detecting Your Platform

### Automated Detection

```bash
# On Unix-like systems
#!/bin/bash
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case "$ARCH" in
    x86_64|amd64)
        ARCH="amd64"
        ;;
    aarch64|arm64)
        ARCH="arm64"
        ;;
esac

echo "Platform: $OS-$ARCH"
```

### Manual Check

**macOS:**
```bash
# Check architecture
uname -m
# arm64 → Apple Silicon
# x86_64 → Intel

# Check if Rosetta is available
pgrep oahd > /dev/null && echo "Rosetta 2 available"
```

**Windows:**
```powershell
# Check architecture
$env:PROCESSOR_ARCHITECTURE
# AMD64 → Intel/AMD 64-bit
# ARM64 → ARM 64-bit

# Get detailed info
Get-ComputerInfo | Select-Object OsArchitecture
```

**Linux:**
```bash
# Check architecture
uname -m
# x86_64 → Intel/AMD
# aarch64 → ARM64

# Check kernel version
uname -r

# Check distribution
cat /etc/os-release
```

## 🔄 Upgrading

### General Steps

1. **Backup configuration**
   ```bash
   cp .mcp_servers .mcp_servers.backup
   ```

2. **Stop the server**
   ```bash
   # If running as service
   sudo systemctl stop websearch-mcp  # Linux
   launchctl stop com.websearch-mcp   # macOS
   nssm stop WebSearchMCP             # Windows
   
   # If running manually
   pkill websearch-mcp
   ```

3. **Download new version**
   ```bash
   # See installation instructions above for your platform
   ```

4. **Replace binary**
   ```bash
   # Backup old binary
   mv websearch-mcp-darwin-arm64 websearch-mcp-darwin-arm64.old
   
   # Move new binary
   mv websearch-mcp-darwin-arm64.new websearch-mcp-darwin-arm64
   chmod +x websearch-mcp-darwin-arm64
   ```

5. **Restart server**
   ```bash
   # If running as service
   sudo systemctl start websearch-mcp  # Linux
   launchctl start com.websearch-mcp   # macOS
   nssm start WebSearchMCP             # Windows
   ```

## 🧪 Testing Your Installation

### Quick Test

```bash
# Start server in background
./websearch-mcp-* &
SERVER_PID=$!

# Wait for startup
sleep 2

# Test health endpoint
curl http://localhost:8080/health

# Test version endpoint
curl http://localhost:8080/version

# Stop server
kill $SERVER_PID
```

### Full Integration Test

See [USAGE.md](USAGE.md) for comprehensive testing instructions.

## 📞 Support

If you encounter platform-specific issues:

1. Check this document for common issues
2. Review the [Troubleshooting Guide](BUILDING.md#troubleshooting)
3. Open an issue on GitHub with:
   - Platform and architecture
   - OS version
   - Error messages
   - Steps to reproduce

## 🗺️ Future Platform Support

Platforms under consideration:

- **FreeBSD** (amd64, arm64)
- **OpenBSD** (amd64)
- **NetBSD** (amd64)
- **Solaris** (amd64)

Request support for additional platforms by opening a GitHub issue.
