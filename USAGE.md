# WebSearch MCP Server Usage Guide

## Quick Start

1. **Build and run the server:**
   ```bash
   ./start-server.sh
   ```

2. **Test the server:**
   ```bash
   ./demo.sh
   ```

3. **Run unit tests:**
   ```bash
   make test
   ```

## MCP Client Integration

### Connection

Connect to the WebSocket endpoint:
```
ws://localhost:8080/
```

### Message Flow

1. **Initialize the connection:**
   ```json
   {
     "jsonrpc": "2.0",
     "id": 1,
     "method": "initialize",
     "params": {
       "protocolVersion": "2024-11-05",
       "capabilities": {},
       "clientInfo": {
         "name": "your-client",
         "version": "1.0.0"
       }
     }
   }
   ```

2. **List available tools:**
   ```json
   {
     "jsonrpc": "2.0",
     "id": 2,
     "method": "tools/list"
   }
   ```

3. **Perform a web search:**
   ```json
   {
     "jsonrpc": "2.0",
     "id": 3,
     "method": "tools/call",
     "params": {
       "name": "web_search",
       "arguments": {
         "query": "latest AI developments",
         "max_results": 5
       }
     }
   }
   ```

## Web Search Tool

### Parameters

- **query** (required): Search terms
- **max_results** (optional): Number of results (1-20, default: 10)

### Response Format

The search results are returned as formatted text containing:
- Search query
- Number of results found
- For each result:
  - Rank number
  - Title
  - URL
  - Description (if available)

### Example Response

```
Search results for: Go programming tutorial
Found 3 results:

1. Go Programming Tutorial - Complete Guide
   URL: https://example.com/go-tutorial
   Description: Learn Go programming from basics to advanced concepts...

2. Official Go Documentation
   URL: https://golang.org/doc/
   Description: The official Go programming language documentation...

3. Go by Example
   URL: https://gobyexample.com/
   Description: Go by Example is a hands-on introduction to Go...
```

## Development

### Running in Development Mode

```bash
make dev
```

This requires [air](https://github.com/cosmtrek/air) for hot reloading.

### Building from source

```bash
CGO_ENABLED=0 go build -ldflags="-w -s" -o websearch-mcp .
```

### Environment Variables

- `PORT`: Server port (default: 8080)

### Health Check

Check server health at:
```
GET http://localhost:8080/health
```

Response:
```json
{
  "status": "healthy",
  "service": "websearch-mcp",
  "version": "1.0.0",
  "timestamp": "2024-01-15T10:30:00Z"
}
```

## Troubleshooting

### Common Issues

1. **Connection refused:**
   - Ensure the server is running
   - Check the port is not in use
   - Verify firewall settings

2. **No search results:**
   - Check internet connectivity
   - Verify DuckDuckGo is accessible
   - Try different search terms

3. **WebSocket errors:**
   - Ensure client supports WebSocket
   - Check for proxy/firewall blocking WebSocket connections
   - Verify correct endpoint URL

### Debug Mode

Run with verbose logging:
```bash
go run main.go
```

### Testing with curl

Test WebSocket connection (requires `websocat` or similar):
```bash
# Install websocat
cargo install websocat

# Connect and send messages
echo '{"jsonrpc":"2.0","id":1,"method":"ping"}' | websocat ws://localhost:8080/
```

## Integration Examples

### Python Client

```python
import asyncio
import websockets
import json

async def test_mcp_server():
    uri = "ws://localhost:8080/"
    async with websockets.connect(uri) as websocket:
        # Initialize
        init_msg = {
            "jsonrpc": "2.0",
            "id": 1,
            "method": "initialize",
            "params": {
                "protocolVersion": "2024-11-05",
                "capabilities": {},
                "clientInfo": {"name": "python-client", "version": "1.0.0"}
            }
        }
        await websocket.send(json.dumps(init_msg))
        response = await websocket.recv()
        print("Initialize:", response)
        
        # Search
        search_msg = {
            "jsonrpc": "2.0",
            "id": 2,
            "method": "tools/call",
            "params": {
                "name": "web_search",
                "arguments": {"query": "Python programming", "max_results": 3}
            }
        }
        await websocket.send(json.dumps(search_msg))
        response = await websocket.recv()
        print("Search:", response)

asyncio.run(test_mcp_server())
```

### JavaScript/Node.js Client

```javascript
const WebSocket = require('ws');

const ws = new WebSocket('ws://localhost:8080/');

ws.on('open', function open() {
    // Initialize
    ws.send(JSON.stringify({
        jsonrpc: "2.0",
        id: 1,
        method: "initialize",
        params: {
            protocolVersion: "2024-11-05",
            capabilities: {},
            clientInfo: { name: "js-client", version: "1.0.0" }
        }
    }));
});

ws.on('message', function message(data) {
    const response = JSON.parse(data);
    console.log('Received:', response);
    
    if (response.id === 1) {
        // Send search request
        ws.send(JSON.stringify({
            jsonrpc: "2.0",
            id: 2,
            method: "tools/call",
            params: {
                name: "web_search",
                arguments: { query: "JavaScript tutorial", max_results: 3 }
            }
        }));
    }
});
```