# Stdio Migration - MCP Server Update

## Summary

The WebSearch MCP server has been updated to use **stdio (standard input/output)** communication by default, instead of HTTP/WebSocket. This change makes the server compatible with standard MCP client implementations, including Tabnine Agents.

## What Changed

### Before (HTTP/WebSocket Mode)
- Server started as a long-running HTTP server on port 8080
- Used WebSocket for bidirectional communication
- Required explicit port configuration
- Had to be started manually before use

### After (Stdio Mode)
- Server communicates via stdin/stdout by default
- Launched automatically by MCP clients (like Tabnine)
- No port configuration needed
- HTTP mode still available for testing via `--http` flag

## Key Changes

### 1. Communication Protocol
- **Default**: Stdio mode (line-delimited JSON-RPC over stdin/stdout)
- **Optional**: HTTP mode with `--http` flag for testing and debugging

### 2. Configuration Updates
The `.mcp_servers` configuration is now simpler:

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

### 3. Command Line Options

New command line flags:
- `--stdio` - Run in stdio mode (default)
- `--http [port]` - Run in HTTP mode on specified port
- `--version`, `-v` - Show version information
- `--help`, `-h` - Show help message

Environment variables:
- `MCP_MODE` - Set to 'http' or 'stdio' (default: stdio)
- `PORT` - Port for HTTP mode (default: 8080)

### 4. Code Changes

#### Removed Dependencies
- `github.com/gorilla/websocket` - No longer needed for stdio mode

#### New Functions
- `runStdio()` - Handles stdio communication
- `runHTTP()` - Handles HTTP mode (refactored from main)

#### Updated Behavior
- Logs now go to stderr (not mixed with stdout)
- Server starts in stdio mode by default
- HTTP mode preserved for backward compatibility and testing

## Migration Guide

### For Tabnine Users

1. **Update your `.mcp_servers` file** to remove the PORT environment variable:
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

2. **Rebuild the server** (if building from source):
   ```bash
   make build
   ```

3. **No other changes needed** - Tabnine will automatically launch the server when needed

### For HTTP Mode Users (Testing/Development)

If you need HTTP mode for testing:

```bash
# Run in HTTP mode
./websearch-mcp --http 8080

# Or use environment variables
MCP_MODE=http PORT=8080 ./websearch-mcp
```

## Benefits of Stdio Mode

1. **Simpler Configuration** - No port management needed
2. **Better Integration** - Standard MCP protocol implementation
3. **Automatic Lifecycle** - Client controls server startup/shutdown
4. **No Port Conflicts** - Each instance runs independently
5. **Cleaner Logs** - Logs separated from protocol messages

## Testing

### Test Stdio Mode

```bash
# Run the test script
./test-stdio.sh

# Or test manually
echo '{"jsonrpc":"2.0","id":1,"method":"initialize","params":{}}' | ./websearch-mcp
```

### Test HTTP Mode

```bash
# Start server in HTTP mode
./websearch-mcp --http 8080

# In another terminal, test endpoints
curl http://localhost:8080/health
curl http://localhost:8080/stats
```

## Backward Compatibility

The HTTP mode is preserved and can be used by:
1. Using the `--http` flag
2. Setting `MCP_MODE=http` environment variable
3. Configuring args in `.mcp_servers`: `"args": ["--http", "8080"]`

## Troubleshooting

### Server Not Responding
- Check that the binary path in `.mcp_servers` is correct
- Verify the binary has execute permissions
- Check Tabnine logs for error messages

### Testing Connection
```bash
# Test stdio mode manually
echo '{"jsonrpc":"2.0","id":1,"method":"ping"}' | ./websearch-mcp

# Should output: {"jsonrpc":"2.0","id":1,"result":"pong"}
```

### Logs Not Visible
In stdio mode, logs go to stderr. To see them:
```bash
./websearch-mcp 2>&1 | tee server.log
```

## References

- [MCP Specification](https://modelcontextprotocol.io/)
- [Tabnine MCP Setup](docs/TABNINE_SETUP.md)
- [Quick Reference](docs/TABNINE_QUICK_REFERENCE.md)
