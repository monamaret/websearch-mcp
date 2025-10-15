# Getting Started with WebSearch MCP

## What is WebSearch MCP?

WebSearch MCP is a tool that enables AI coding assistants (like Tabnine) to search the web using DuckDuckGo. It implements the Model Context Protocol (MCP), allowing seamless integration with AI agents.

## Quick Start (2 Minutes)

### 1. Get the Binary

Choose one option:

**Option A: Download Pre-built Binary**
- Go to [Releases](https://github.com/youruser/websearch-mcp/releases)
- Download for your platform (macOS, Windows, or Linux)
- Extract and make executable

**Option B: Build from Source**
```bash
git clone <repository-url>
cd websearch-mcp
make build
```

### 2. Configure Tabnine

Create `.mcp_servers` in your project root:

```json
{
  "mcpServers": {
    "websearch": {
      "command": "./websearch-mcp",
      "args": [],
      "env": {}
    }
  }
}
```

**Important:** Update `command` to match your binary path:
- macOS (M1/M2/M3): `./websearch-mcp-darwin-arm64`
- macOS (Intel): `./websearch-mcp-darwin-amd64`
- Windows: `websearch-mcp-windows-amd64.exe`
- Linux: `./websearch-mcp-linux-amd64`

### 3. Test It

Ask your Tabnine Agent:
```
"Search for 'Go programming best practices' and give me the top 5 results"
```

Done! 🎉

## How It Works

```
┌─────────────┐         ┌──────────────┐         ┌─────────────┐
│   Tabnine   │ ◄────► │ WebSearch MCP │ ◄────► │ DuckDuckGo  │
│   Agent     │  stdio  │    Server     │  HTTPS  │   Search    │
└─────────────┘         └──────────────┘         └─────────────┘
```

1. **You ask** Tabnine a question
2. **Tabnine launches** the WebSearch MCP server
3. **Server searches** DuckDuckGo
4. **Results return** to Tabnine
5. **Tabnine responds** with search results

## Example Prompts

### Research
```
"Search for the latest React 18 features and summarize them"
"Find documentation about Go's new generics"
"Look up common Git workflow patterns"
```

### Troubleshooting
```
"Search for solutions to Docker networking issues"
"Find fixes for 'CORS error' in web applications"
"Look up PostgreSQL performance optimization tips"
```

### Learning
```
"Search for Kubernetes best practices"
"Find tutorials on implementing OAuth2"
"Look up examples of clean architecture in Go"
```

## Modes of Operation

### Stdio Mode (Default)

**Best for:** Production use with Tabnine

- Automatic startup by Tabnine
- No configuration needed
- Direct stdin/stdout communication

```bash
# Tabnine calls automatically
./websearch-mcp
```

### HTTP Mode (Testing)

**Best for:** Manual testing and debugging

- Health check endpoints
- Stats and monitoring
- Manual API testing

```bash
# Start manually
./websearch-mcp --http 8080

# Test endpoints
curl http://localhost:8080/health
curl http://localhost:8080/stats
```

## Configuration Options

### Minimal (Recommended)
```json
{
  "mcpServers": {
    "websearch": {
      "command": "./websearch-mcp",
      "args": [],
      "env": {}
    }
  }
}
```

### With Absolute Path
```json
{
  "mcpServers": {
    "websearch": {
      "command": "/usr/local/bin/websearch-mcp",
      "args": [],
      "env": {}
    }
  }
}
```

### HTTP Mode
```json
{
  "mcpServers": {
    "websearch": {
      "command": "./websearch-mcp",
      "args": ["--http", "8080"],
      "env": {}
    }
  }
}
```

## Verification

### Check Installation

```bash
# Show version
./websearch-mcp --version

# Show help
./websearch-mcp --help
```

### Test Manually

```bash
# Test stdio mode
echo '{"jsonrpc":"2.0","id":1,"method":"ping"}' | ./websearch-mcp

# Expected: {"jsonrpc":"2.0","id":1,"result":"pong"}
```

### Test with Tabnine

Ask Tabnine:
```
"Can you see the websearch MCP server?"
```

If yes, you're all set! ✅

## Troubleshooting

### Can't Find Server

**Problem:** Tabnine can't see the websearch server

**Solutions:**
1. Check `.mcp_servers` exists in project root
2. Verify `command` path is correct
3. Make binary executable: `chmod +x websearch-mcp`

### Permission Denied

**Problem:** Permission error when running

**Solution:**
```bash
chmod +x websearch-mcp
```

### No Search Results

**Problem:** Searches return no results

**Solutions:**
1. Check internet connection
2. Verify DuckDuckGo is accessible
3. Try a different search query

### Binary Not Found

**Problem:** Command not found error

**Solutions:**
1. Use absolute path: `/full/path/to/websearch-mcp`
2. Check binary is in the right location
3. Verify binary name matches your platform

## Platform-Specific Setup

### macOS

```bash
# Download and extract
curl -LO https://github.com/.../websearch-mcp-darwin-arm64.tar.gz
tar -xzf websearch-mcp-darwin-arm64.tar.gz
chmod +x websearch-mcp-darwin-arm64

# Configure
echo '{
  "mcpServers": {
    "websearch": {
      "command": "./websearch-mcp-darwin-arm64",
      "args": [],
      "env": {}
    }
  }
}' > .mcp_servers
```

### Windows

```powershell
# Extract the downloaded zip file
# Then configure

@"
{
  "mcpServers": {
    "websearch": {
      "command": "websearch-mcp-windows-amd64.exe",
      "args": [],
      "env": {}
    }
  }
}
"@ | Out-File -FilePath .mcp_servers -Encoding utf8
```

### Linux

```bash
# Download and extract
wget https://github.com/.../websearch-mcp-linux-amd64.tar.gz
tar -xzf websearch-mcp-linux-amd64.tar.gz
chmod +x websearch-mcp-linux-amd64

# Configure
cat > .mcp_servers << 'EOF'
{
  "mcpServers": {
    "websearch": {
      "command": "./websearch-mcp-linux-amd64",
      "args": [],
      "env": {}
    }
  }
}
EOF
```

## Next Steps

### Learn More
- [Full Tabnine Setup Guide](./TABNINE_SETUP.md)
- [Quick Reference](./TABNINE_QUICK_REFERENCE.md)
- [MCP Introduction](./mcp-introduction.md)

### Build from Source
- [Building Guide](./BUILDING.md)
- [Platform Support](./PLATFORM_SUPPORT.md)

### Advanced Usage
- [Usage Guide](./USAGE.md)
- [MCP Examples](./mcp-examples.md)

## FAQ

**Q: Do I need to start the server manually?**
A: No! Tabnine starts it automatically in stdio mode.

**Q: What port does it use?**
A: None in stdio mode. Only HTTP mode uses a port.

**Q: Is it secure?**
A: Yes. It only searches DuckDuckGo (HTTPS) and runs locally.

**Q: Does it store search history?**
A: No. No data is stored.

**Q: Can I use multiple MCP servers?**
A: Yes! Add more entries to `.mcp_servers`.

**Q: How do I update?**
A: Download new binary and replace the old one.

## Support

- **Documentation**: See `docs/` folder
- **Issues**: Open a GitHub issue
- **Examples**: See `examples/` folder

---

**Ready to search the web with your AI assistant?** Follow the [Quick Start](#quick-start-2-minutes) guide above!
