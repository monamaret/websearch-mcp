# WebSearch MCP Examples

This directory contains example configurations and test clients for the WebSearch MCP server.

Note: Example clients are excluded from normal builds via Go build tags. To build or run them, use the `examples` build tag.

## Configuration Examples

### `.mcp_servers` - Basic Configuration

Basic stdio mode configuration for Tabnine:

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

**Usage:** Copy this file to your project root to enable WebSearch MCP with Tabnine.

### `.mcp_servers.production` - Production Configuration

Production configuration with absolute paths:

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

**Usage:** Use this for production deployments where the binary is installed system-wide.

### `.mcp_servers.docker` - Docker Configuration

Run the MCP server in a Docker container:

```json
{
  "mcpServers": {
    "websearch-docker": {
      "command": "docker",
      "args": [
        "run",
        "-i",
        "--rm",
        "websearch-mcp:latest"
      ],
      "env": {}
    }
  }
}
```

**Usage:** Use this if you prefer running the server in Docker.

### `.mcp_servers.multi` - Multi-Tool Configuration

Example of combining websearch with other MCP tools:

```json
{
  "mcpServers": {
    "websearch": {
      "command": "./websearch-mcp",
      "args": [],
      "env": {}
    },
    "filesystem": {
      "command": "npx",
      "args": [
        "@modelcontextprotocol/server-filesystem",
        "/path/to/your/project"
      ],
      "env": {}
    },
    "git": {
      "command": "npx",
      "args": [
        "@modelcontextprotocol/server-git",
        "/path/to/your/git/repo"
      ],
      "env": {}
    }
  }
}
```

**Usage:** Use this to combine multiple MCP tools in your workflow.

## Test Clients

### `stdio-client/` - Stdio Test Client (Recommended)

Test client for stdio mode (the default):

```bash
# Build the main server first
cd ..
go build -o websearch-mcp .

# Run the stdio test client (requires build tag)
cd examples/stdio-client
go run -tags=examples .
```

This client:
- Launches the MCP server as a subprocess
- Sends test messages via stdin
- Reads responses from stdout
- Displays server logs from stderr

**Output:**
```
=== Test 1: initialize ===
Sending:
{
  "jsonrpc": "2.0",
  "id": 1,
  "method": "initialize",
  ...
}

Received:
{
  "jsonrpc": "2.0",
  "id": 1,
  "result": {
    "protocolVersion": "2024-11-05",
    ...
  }
}
```

### `http-client/` - HTTP Test Client

Test client for HTTP/WebSocket mode (for testing/debugging):

```bash
# Start the server in HTTP mode
cd ..
./websearch-mcp --http 8080

# In another terminal, run the HTTP test client (requires build tag)
cd examples/http-client
go run -tags=examples .
```

This client:
- Connects via WebSocket to localhost:8080
- Sends test messages
- Receives and displays responses
- Useful for debugging HTTP mode

### Building example binaries

```bash
# Build stdio example
cd examples/stdio-client
go build -tags=examples -o test-stdio .

# Build HTTP example
cd examples/http-client
go build -tags=examples -o test-http .
```

## Quick Test

### Test Stdio Mode (Recommended)

```bash
# From examples directory
cd ..
./websearch-mcp --stdio <<'EOF'
{"jsonrpc":"2.0","id":1,"method":"ping"}
EOF
```

Expected output:
```json
{"jsonrpc":"2.0","id":1,"result":"pong"}
```

### Test HTTP Mode

```bash
# Terminal 1: Start server
cd ..
./websearch-mcp --http 8080

# Terminal 2: Test with curl
curl http://localhost:8080/health
curl http://localhost:8080/stats
```

## Common Use Cases

### Use Case 1: Tabnine Integration

Copy the basic configuration:
```bash
cp .mcp_servers ../.mcp_servers
```

Then ask Tabnine:
```
"Search for Go programming best practices"
```

### Use Case 2: Manual Testing

Run the stdio test client:
```bash
cd ..
go build -o websearch-mcp .
cd examples/stdio-client
go run -tags=examples .
```

### Use Case 3: HTTP Mode Testing

Start server and use test client:
```bash
# Terminal 1
cd ..
./websearch-mcp --http 8080

# Terminal 2
cd examples/http-client
go run -tags=examples .
```

### Use Case 4: Production Deployment

1. Build and install:
   ```bash
   cd ..
   CGO_ENABLED=0 go build -ldflags="-w -s" -o websearch-mcp .
   sudo cp websearch-mcp /usr/local/bin/
   sudo chmod +x /usr/local/bin/websearch-mcp
   ```

2. Use production config:
   ```bash
   cp examples/.mcp_servers.production .mcp_servers
   ```

## Troubleshooting

### Issue: "command not found"

**Solution:** Check the `command` path in your `.mcp_servers` file matches your binary location.

### Issue: "permission denied"

**Solution:** Make the binary executable:
```bash
chmod +x websearch-mcp
```

### Issue: "server not responding"

**Solution for stdio mode:**
```bash
# Test manually
echo '{"jsonrpc":"2.0","id":1,"method":"ping"}' | ./websearch-mcp
```

**Solution for HTTP mode:**
```bash
# Check if server is running
curl http://localhost:8080/health
```

## Building Test Clients

### Build stdio client:
```bash
cd examples/stdio-client
go build -tags=examples -o test-stdio .
./test-stdio
```

### Build HTTP client:
```bash
cd examples/http-client
go build -tags=examples -o test-http .
./test-http
```

## Next Steps

- **Read the full setup guide:** [../docs/TABNINE_SETUP.md](../docs/TABNINE_SETUP.md)
- **Quick reference:** [../docs/TABNINE_QUICK_REFERENCE.md](../docs/TABNINE_QUICK_REFERENCE.md)
- **Getting started:** [../docs/GETTING_STARTED.md](../docs/GETTING_STARTED.md)

## Notes

- **Stdio mode** is the default and recommended mode for MCP clients
- **HTTP mode** is available for testing and debugging with `--http` flag
- All configuration files use stdio mode unless explicitly configured for HTTP
- The server automatically detects which mode to use based on how it's called
