# WebSearch MCP Server

A Model Context Protocol (MCP) server implementation in Go that provides web search capabilities using DuckDuckGo. Designed for seamless integration with Tabnine Agents and other MCP-compatible AI systems.

## Features

- **Web Search**: Search the web using DuckDuckGo
- **MCP Compliant**: Implements the Model Context Protocol specification
- **Tabnine Ready**: Pre-configured for Tabnine Agents integration
- **WebSocket Communication**: Real-time communication via WebSocket
- **Configurable Results**: Specify maximum number of search results
- **Clean Response Format**: Structured search results with titles, URLs, and descriptions
- **Performance Monitoring**: Built-in health checks and statistics
- **Docker Support**: Containerized deployment ready

## Quick Start with Tabnine

### 1. Build the Server
```bash
go build -o websearch-mcp .
# or
make build
```

### 2. Configure Tabnine MCP
Create a `.mcp_servers` file in your project root:

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

### 3. Test Integration
Ask your Tabnine Agent:
```
"Search for 'Go programming best practices' and give me the top 5 results"
```

For detailed setup instructions, see [TABNINE_SETUP.md](TABNINE_SETUP.md).

## Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd websearch-mcp
```

2. Install dependencies:
```bash
go mod tidy
```

3. Build the server:
```bash
go build -o websearch-mcp
```

## Usage

### Running the Server

```bash
./websearch-mcp
```

The server will start on port 8080 by default. You can specify a different port using the `PORT` environment variable:

```bash
PORT=3000 ./websearch-mcp
```

### WebSocket Endpoint

The server exposes a WebSocket endpoint at:
```
ws://localhost:8080/
```

## MCP Protocol Implementation

### Supported Methods

1. **initialize**: Initialize the MCP connection
2. **tools/list**: List available tools
3. **tools/call**: Execute a tool
4. **ping**: Health check

### Available Tools

#### web_search

Search the web for information using DuckDuckGo.

**Parameters:**
- `query` (string, required): The search query to execute
- `max_results` (integer, optional): Maximum number of results to return (default: 10, max: 20)

**Example:**
```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "method": "tools/call",
  "params": {
    "name": "web_search",
    "arguments": {
      "query": "latest developments in AI",
      "max_results": 5
    }
  }
}
```

## API Examples

### Initialize Connection

```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "method": "initialize",
  "params": {
    "protocolVersion": "2024-11-05",
    "capabilities": {},
    "clientInfo": {
      "name": "example-client",
      "version": "1.0.0"
    }
  }
}
```

### List Available Tools

```json
{
  "jsonrpc": "2.0",
  "id": 2,
  "method": "tools/list"
}
```

### Perform Web Search

```json
{
  "jsonrpc": "2.0",
  "id": 3,
  "method": "tools/call",
  "params": {
    "name": "web_search",
    "arguments": {
      "query": "Go programming best practices",
      "max_results": 8
    }
  }
}
```

## Response Format

Search results are returned in the following format:

```json
{
  "jsonrpc": "2.0",
  "id": 3,
  "result": {
    "content": [
      {
        "type": "text",
        "text": "Search results for: Go programming best practices\nFound 8 results:\n\n1. Go Best Practices\n   URL: https://example.com/go-best-practices\n   Description: A comprehensive guide to Go programming best practices...\n\n..."
      }
    ]
  }
}
```

## Configuration

### Environment Variables

- `PORT`: Server port (default: 8080)

## Development

### Running in Development Mode

```bash
go run main.go
```

### Testing the Server

You can test the server using a WebSocket client or the provided test scripts.

## Dependencies

- `github.com/PuerkitoBio/goquery`: HTML parsing for web scraping
- `github.com/gorilla/websocket`: WebSocket implementation

## Security Considerations

- The server accepts connections from all origins for MCP compatibility
- DuckDuckGo is used as the search provider to avoid API key requirements
- Request timeouts are configured to prevent hanging connections
- The server includes graceful shutdown handling

## License

MIT License

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## Troubleshooting

### Common Issues

1. **Connection Refused**: Ensure the server is running and the port is correct
2. **No Search Results**: Check your internet connection and verify DuckDuckGo is accessible
3. **WebSocket Errors**: Ensure your client supports WebSocket connections

### Logging

The server logs all connections, disconnections, and errors to help with debugging.