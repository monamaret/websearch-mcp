# Quick Start Guide

Get up and running with WebSearch MCP Server in 5 minutes.

## 🎯 Choose Your Platform

Select your platform below and follow the instructions:

- [macOS](#macos)
- [Windows](#windows)
- [Linux](#linux)

---

## macOS

### Step 1: Determine Your Mac Type

**Apple Silicon (M1/M2/M3)** - Newer Macs (2020+)  
**Intel** - Older Macs (before 2020)

**How to check:**
```bash
# Run this command
uname -m

# Output:
# arm64     → You have Apple Silicon
# x86_64    → You have Intel
```

### Step 2: Download Binary

#### Apple Silicon (M1/M2/M3)
```bash
# Download
curl -LO https://github.com/youruser/websearch-mcp/releases/latest/download/websearch-mcp-VERSION-darwin-arm64.tar.gz

# Extract
tar -xzf websearch-mcp-VERSION-darwin-arm64.tar.gz

# Make executable
chmod +x websearch-mcp-darwin-arm64
```

#### Intel
```bash
# Download
curl -LO https://github.com/youruser/websearch-mcp/releases/latest/download/websearch-mcp-VERSION-darwin-amd64.tar.gz

# Extract
tar -xzf websearch-mcp-VERSION-darwin-amd64.tar.gz

# Make executable
chmod +x websearch-mcp-darwin-amd64
```

### Step 3: Run the Server

```bash
# Apple Silicon
./websearch-mcp-darwin-arm64

# Intel
./websearch-mcp-darwin-amd64
```

### Step 4: Verify It's Running

Open another terminal and run:
```bash
curl http://localhost:8080/health
```

You should see a JSON response with server status.

### Troubleshooting

**"Cannot open because developer cannot be verified"**

Run this command:
```bash
# Apple Silicon
xattr -d com.apple.quarantine websearch-mcp-darwin-arm64

# Intel
xattr -d com.apple.quarantine websearch-mcp-darwin-amd64
```

Then try running again.

---

## Windows

### Step 1: Determine Your Windows Type

**Most Windows PCs** use Intel/AMD (x86_64)  
**Surface Pro X and some tablets** use ARM64

**How to check:**
```powershell
# Run this in PowerShell
$env:PROCESSOR_ARCHITECTURE

# Output:
# AMD64  → You have Intel/AMD (x86_64)
# ARM64  → You have ARM
```

### Step 2: Download Binary

#### Intel/AMD (Most Common)

1. Go to: https://github.com/youruser/websearch-mcp/releases/latest
2. Download: `websearch-mcp-VERSION-windows-amd64.zip`
3. Extract the ZIP file (Right-click → Extract All)

**Or using PowerShell:**
```powershell
# Download
Invoke-WebRequest -Uri "https://github.com/youruser/websearch-mcp/releases/latest/download/websearch-mcp-VERSION-windows-amd64.zip" -OutFile "websearch-mcp.zip"

# Extract
Expand-Archive -Path websearch-mcp.zip -DestinationPath .
```

#### ARM64

1. Go to: https://github.com/youruser/websearch-mcp/releases/latest
2. Download: `websearch-mcp-VERSION-windows-arm64.zip`
3. Extract the ZIP file (Right-click → Extract All)

### Step 3: Run the Server

**Option A: Double-click**
- Double-click `websearch-mcp-windows-amd64.exe` in File Explorer

**Option B: Command Line**

PowerShell:
```powershell
.\websearch-mcp-windows-amd64.exe
```

Command Prompt:
```cmd
websearch-mcp-windows-amd64.exe
```

### Step 4: Verify It's Running

Open another PowerShell/Command Prompt:
```powershell
# PowerShell
Invoke-WebRequest http://localhost:8080/health

# Or using curl (Windows 10+)
curl http://localhost:8080/health
```

### Troubleshooting

**Windows Defender SmartScreen Warning**

1. Click "More info"
2. Click "Run anyway"

**Firewall Blocking**

If prompted, click "Allow access" for the firewall.

---

## Linux

### Step 1: Determine Your Architecture

```bash
# Run this command
uname -m

# Output:
# x86_64    → Intel/AMD (most common)
# aarch64   → ARM64
# arm64     → ARM64
```

### Step 2: Download Binary

#### Intel/AMD (x86_64)
```bash
# Download
wget https://github.com/youruser/websearch-mcp/releases/latest/download/websearch-mcp-VERSION-linux-amd64.tar.gz

# Or with curl
curl -LO https://github.com/youruser/websearch-mcp/releases/latest/download/websearch-mcp-VERSION-linux-amd64.tar.gz

# Extract
tar -xzf websearch-mcp-VERSION-linux-amd64.tar.gz

# Make executable
chmod +x websearch-mcp-linux-amd64
```

#### ARM64
```bash
# Download
wget https://github.com/youruser/websearch-mcp/releases/latest/download/websearch-mcp-VERSION-linux-arm64.tar.gz

# Extract
tar -xzf websearch-mcp-VERSION-linux-arm64.tar.gz

# Make executable
chmod +x websearch-mcp-linux-arm64
```

### Step 3: Run the Server

```bash
# Intel/AMD
./websearch-mcp-linux-amd64

# ARM64
./websearch-mcp-linux-arm64
```

### Step 4: Verify It's Running

Open another terminal:
```bash
curl http://localhost:8080/health
```

### Troubleshooting

**Permission Denied**
```bash
chmod +x websearch-mcp-linux-amd64
```

**Port Already in Use**
```bash
# Use a different port
PORT=3000 ./websearch-mcp-linux-amd64
```

---

## 🔧 Configuration for Tabnine

After the server is running, configure Tabnine to use it.

### Create `.mcp_servers` File

Create a file named `.mcp_servers` in your project root:

#### macOS Apple Silicon
```json
{
  "mcpServers": {
    "websearch": {
      "command": "/full/path/to/websearch-mcp-darwin-arm64",
      "args": [],
      "env": {
        "PORT": "8080"
      }
    }
  }
}
```

#### macOS Intel
```json
{
  "mcpServers": {
    "websearch": {
      "command": "/full/path/to/websearch-mcp-darwin-amd64",
      "args": [],
      "env": {
        "PORT": "8080"
      }
    }
  }
}
```

#### Windows Intel
```json
{
  "mcpServers": {
    "websearch": {
      "command": "C:\\full\\path\\to\\websearch-mcp-windows-amd64.exe",
      "args": [],
      "env": {
        "PORT": "8080"
      }
    }
  }
}
```

#### Windows ARM
```json
{
  "mcpServers": {
    "websearch": {
      "command": "C:\\full\\path\\to\\websearch-mcp-windows-arm64.exe",
      "args": [],
      "env": {
        "PORT": "8080"
      }
    }
  }
}
```

#### Linux Intel/AMD
```json
{
  "mcpServers": {
    "websearch": {
      "command": "/full/path/to/websearch-mcp-linux-amd64",
      "args": [],
      "env": {
        "PORT": "8080"
      }
    }
  }
}
```

#### Linux ARM
```json
{
  "mcpServers": {
    "websearch": {
      "command": "/full/path/to/websearch-mcp-linux-arm64",
      "args": [],
      "env": {
        "PORT": "8080"
      }
    }
  }
}
```

### Test Tabnine Integration

Ask your Tabnine Agent:
```
"Search for 'Go programming best practices' and give me the top 5 results"
```

---

## 🧪 Basic Testing

### Test Health Endpoint

**macOS/Linux:**
```bash
curl http://localhost:8080/health
```

**Windows (PowerShell):**
```powershell
Invoke-WebRequest http://localhost:8080/health | Select-Object -Expand Content
```

**Expected response:**
```json
{
  "status": "healthy",
  "version": "v1.2.3",
  "uptime": "5m30s"
}
```

### Test Version Endpoint

**macOS/Linux:**
```bash
curl http://localhost:8080/version
```

**Windows (PowerShell):**
```powershell
Invoke-WebRequest http://localhost:8080/version | Select-Object -Expand Content
```

---

## 🚀 Next Steps

### Run as a Service

For production use, run the server as a system service:

- **Linux**: See [docs/PLATFORM_SUPPORT.md#systemd-linux](PLATFORM_SUPPORT.md#systemd-linux)
- **macOS**: See [docs/PLATFORM_SUPPORT.md#launchd-macos](PLATFORM_SUPPORT.md#launchd-macos)
- **Windows**: See [docs/PLATFORM_SUPPORT.md#windows-service](PLATFORM_SUPPORT.md#windows-service)

### Configure Port

Change the default port (8080):

**macOS/Linux:**
```bash
PORT=3000 ./websearch-mcp-darwin-arm64
```

**Windows (PowerShell):**
```powershell
$env:PORT = "3000"
.\websearch-mcp-windows-amd64.exe
```

### View Logs

The server logs to stdout. To save logs to a file:

**macOS/Linux:**
```bash
./websearch-mcp-darwin-arm64 > server.log 2>&1 &
```

**Windows (PowerShell):**
```powershell
.\websearch-mcp-windows-amd64.exe > server.log 2>&1
```

---

## 📚 Further Reading

- **[Full Documentation](../README.md)** - Complete feature list and API docs
- **[Platform Support](PLATFORM_SUPPORT.md)** - Detailed platform information
- **[Tabnine Setup](TABNINE_SETUP.md)** - Comprehensive Tabnine integration guide
- **[Building from Source](BUILDING.md)** - Build your own binaries
- **[Usage Guide](USAGE.md)** - Advanced usage and configuration

---

## ❓ Common Questions

**Q: Which binary should I download?**

A: Check your platform:
- Mac with M1/M2/M3 → darwin-arm64
- Mac with Intel → darwin-amd64
- Windows PC → windows-amd64 (most common)
- Windows tablet/Surface Pro X → windows-arm64
- Linux PC/Server → linux-amd64 (most common)
- Raspberry Pi 4+ → linux-arm64

**Q: How do I stop the server?**

A: Press `Ctrl+C` in the terminal where it's running.

**Q: Can I run multiple instances?**

A: Yes, but use different ports:
```bash
PORT=8080 ./websearch-mcp-darwin-arm64 &
PORT=8081 ./websearch-mcp-darwin-arm64 &
```

**Q: How do I update to a new version?**

A:
1. Stop the server
2. Download the new version
3. Replace the old binary
4. Start the server again

**Q: Is it safe to download these binaries?**

A: Yes! Verify using the SHA256 checksums in `checksums.txt` from the release page.

---

## 🆘 Getting Help

If you run into issues:

1. Check the [Troubleshooting](#troubleshooting) section for your platform above
2. Review [docs/PLATFORM_SUPPORT.md](PLATFORM_SUPPORT.md) for detailed help
3. Search [GitHub Issues](https://github.com/youruser/websearch-mcp/issues)
4. Open a new issue with:
   - Your platform and architecture
   - Steps you followed
   - Error messages
   - What you expected to happen

---

**Ready to start?** Scroll to your platform above and get started! ⬆️
