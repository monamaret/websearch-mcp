# Upgrade Guide: HTTP to Stdio Mode

## Overview

This guide helps you upgrade from the old HTTP/WebSocket-based MCP server to the new stdio-based implementation.

## What's New

### ✅ Stdio Mode (Default)
- **Direct communication** via standard input/output
- **Automatic lifecycle** managed by MCP clients
- **No port configuration** needed
- **Better compatibility** with MCP protocol

### ✅ HTTP Mode (Optional)
- **Testing and debugging** support
- **Health and stats endpoints** for monitoring
- **Backward compatible** with existing setups

## Quick Migration (3 Steps)

### Step 1: Update Configuration

Edit your `.mcp_servers` file:

**Before:**
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

**After:**
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

### Step 2: Rebuild (if building from source)

```bash
make build
# or
go build -o websearch-mcp .
```

### Step 3: Restart Tabnine

Restart your Tabnine Agent to pick up the new configuration.

## That's It!

The server will now run in stdio mode automatically when called by Tabnine.

---

## Detailed Changes

### Configuration Changes

#### Stdio Mode (Recommended)
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

#### HTTP Mode (For Testing)
If you need HTTP mode for testing or debugging:

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

Or use environment variables:

```json
{
  "mcpServers": {
    "websearch": {
      "command": "./websearch-mcp",
      "args": [],
      "env": {
        "MCP_MODE": "http",
        "PORT": "8080"
      }
    }
  }
}
```

### Command Line Changes

#### New Flags
- `--stdio` - Run in stdio mode (default)
- `--http [port]` - Run in HTTP mode
- `--version`, `-v` - Show version
- `--help`, `-h` - Show help

#### Examples
```bash
# Stdio mode (default)
./websearch-mcp
./websearch-mcp --stdio

# HTTP mode
./websearch-mcp --http 8080
./websearch-mcp --http

# Using environment variables
MCP_MODE=http PORT=8080 ./websearch-mcp
```

### Behavioral Changes

| Aspect | Before (HTTP) | After (Stdio) |
|--------|---------------|---------------|
| **Startup** | Manual start required | Auto-started by client |
| **Port** | Required configuration | Not needed |
| **Logs** | Mixed with responses | Separate (stderr) |
| **Communication** | WebSocket | stdin/stdout |
| **Testing** | HTTP endpoints | JSON-RPC via stdin |

## Testing Your Setup

### Test Stdio Mode

```bash
# Test with echo
echo '{"jsonrpc":"2.0","id":1,"method":"ping"}' | ./websearch-mcp

# Expected output:
# {"jsonrpc":"2.0","id":1,"result":"pong"}

# Run test script
./test-stdio.sh
```

### Test HTTP Mode

```bash
# Start in HTTP mode
./websearch-mcp --http 8080

# In another terminal
curl http://localhost:8080/health
curl http://localhost:8080/stats
```

### Test with Tabnine

Ask your Tabnine Agent:
```
"Search for 'Go programming best practices' and summarize the top 3 results"
```

## Troubleshooting

### Issue: "Cannot see websearch server"

**Solution:**
1. Check `.mcp_servers` file exists in your project root
2. Verify the `command` path is correct
3. Ensure the binary has execute permissions: `chmod +x websearch-mcp`

### Issue: "Permission denied"

**Solution:**
```bash
chmod +x websearch-mcp
```

### Issue: "Server not responding"

**Solutions:**
1. Check Tabnine logs for error messages
2. Test manually: `echo '{"jsonrpc":"2.0","id":1,"method":"ping"}' | ./websearch-mcp`
3. Verify you're using the correct binary for your platform

### Issue: "Want to use HTTP mode"

**Solution:**
Add `--http` flag to args in `.mcp_servers`:
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

### Issue: "Want to see logs"

**Solution:**
In stdio mode, logs go to stderr. To capture them:
```bash
./websearch-mcp 2>&1 | tee server.log
```

## Benefits Summary

### ✅ Simpler Configuration
No port management, fewer environment variables

### ✅ Better MCP Compliance
Standard stdio communication matches MCP specification

### ✅ Automatic Lifecycle
Client handles startup/shutdown automatically

### ✅ No Port Conflicts
Each MCP client instance runs independently

### ✅ Cleaner Logs
Protocol messages and logs are separated

### ✅ Backward Compatible
HTTP mode still available when needed

## Platform-Specific Notes

### macOS
```bash
# Build
make build

# Test
./websearch-mcp --version
```

### Windows
```powershell
# Build
go build -o websearch-mcp.exe .

# Test
.\websearch-mcp.exe --version
```

### Linux
```bash
# Build
make build

# Test
./websearch-mcp --version
```

## Need Help?

1. **Read the docs:**
   - [STDIO_MIGRATION.md](./STDIO_MIGRATION.md) - Detailed migration guide
   - [docs/TABNINE_SETUP.md](./docs/TABNINE_SETUP.md) - Full setup guide
   - [docs/TABNINE_QUICK_REFERENCE.md](./docs/TABNINE_QUICK_REFERENCE.md) - Quick reference

2. **Test the server:**
   ```bash
   ./test-stdio.sh
   ```

3. **Check logs:**
   - Tabnine logs (in your IDE)
   - Server logs (stderr)

4. **Open an issue:**
   - Provide your configuration
   - Include error messages
   - Mention your platform

## Rollback (If Needed)

If you need to rollback to HTTP mode:

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

This maintains the old behavior while using the new binary.

---

**Ready to upgrade?** Follow the [Quick Migration](#quick-migration-3-steps) steps above!
