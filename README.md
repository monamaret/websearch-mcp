# WebSearch MCP Server

A lightweight Model Context Protocol (MCP) server written in Go that adds web search (DuckDuckGo) to MCP-compatible clients like Tabnine Agents.

## What it does (quick overview)
- Web search via DuckDuckGo with clean, structured results
- Works out of the box with Tabnine Agents (stdio mode by default)
- Optional HTTP/WebSocket mode for testing and debugging
- Single portable binary with multi-platform support
- No API keys required

## Install and set up with Tabnine (all OS)

Below are end-to-end instructions to download the binary, place it in your project, and configure Tabnine using the same structure as docs/mcp-server-config.json.

### 1) Download the binary for your OS

Go to the Releases page and download the binary for your platform:
- https://github.com/yourusername/websearch-mcp/releases/latest

Binary naming (adjust VERSION to the release version):
- macOS Apple Silicon: websearch-mcp-VERSION-darwin-arm64
- macOS Intel: websearch-mcp-VERSION-darwin-amd64
- Linux x86_64: websearch-mcp-VERSION-linux-amd64
- Linux ARM64: websearch-mcp-VERSION-linux-arm64
- Windows x86_64: websearch-mcp-VERSION-windows-amd64.exe
- Windows ARM64: websearch-mcp-VERSION-windows-arm64.exe

Examples:

macOS (Apple Silicon)
```bash
curl -LO https://github.com/yourusername/websearch-mcp/releases/latest/download/websearch-mcp-VERSION-darwin-arm64.tar.gz
tar -xzf websearch-mcp-VERSION-darwin-arm64.tar.gz
mv websearch-mcp-darwin-arm64 websearch-mcp
chmod +x websearch-mcp
```

macOS (Intel)
```bash
curl -LO https://github.com/yourusername/websearch-mcp/releases/latest/download/websearch-mcp-VERSION-darwin-amd64.tar.gz
tar -xzf websearch-mcp-VERSION-darwin-amd64.tar.gz
mv websearch-mcp-darwin-amd64 websearch-mcp
chmod +x websearch-mcp
```

Linux (x86_64)
```bash
wget https://github.com/yourusername/websearch-mcp/releases/latest/download/websearch-mcp-VERSION-linux-amd64.tar.gz
tar -xzf websearch-mcp-VERSION-linux-amd64.tar.gz
mv websearch-mcp-linux-amd64 websearch-mcp
chmod +x websearch-mcp
```

Linux (ARM64)
```bash
wget https://github.com/yourusername/websearch-mcp/releases/latest/download/websearch-mcp-VERSION-linux-arm64.tar.gz
tar -xzf websearch-mcp-VERSION-linux-arm64.tar.gz
mv websearch-mcp-linux-arm64 websearch-mcp
chmod +x websearch-mcp
```

Windows (PowerShell)
```powershell
# Download the appropriate .zip from the Releases page, extract it, then rename:
#   websearch-mcp-windows-amd64.exe   -> websearch-mcp.exe
#   or
#   websearch-mcp-windows-arm64.exe   -> websearch-mcp.exe
```

Place the resulting binary in your project root alongside your code (or any folder you reference in the next step).

### 2) Add Tabnine MCP config

Copy docs/mcp-server-config.json to your project root and adjust the command path to the binary you downloaded. You can rename it to .mcp.json (recommended) or keep any file name your Tabnine setup expects, as long as it contains the mcpServers block.

Base template (from docs/mcp-server-config.json):
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

Notes:
- command should point to the binary in your environment.
- The env.PORT is only used for HTTP mode. Tabnine uses stdio mode by default; you can keep or remove it—it will be ignored unless you run in HTTP mode.

OS-specific command examples:

macOS (Apple Silicon)
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

macOS (Intel)
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

Linux (x86_64 or ARM64)
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

Windows (x86_64 or ARM64)
```json
{
  "mcpServers": {
    "websearch": {
      "command": ".\\websearch-mcp.exe",
      "args": [],
      "env": {}
    }
  }
}
```

Place this config in your project root (for example, as .mcp.json) so Tabnine Agents can detect and launch the MCP server.

### 3) Test in Tabnine

Open Tabnine Chat and ask something that triggers the web search tool, for example:
- "Search for 'Go concurrency patterns' and summarize the top 5 results."

If everything is configured correctly, Tabnine will start the server in stdio mode and return structured search results.

## Optional: Build from source

```bash
git clone <repository-url>
cd websearch-mcp
make build           # build for your current platform
# or
make build-all-platforms
```
See docs/BUILDING.md for details.

## HTTP mode (optional, for testing)
- Run: ./websearch-mcp --http 8080 (or set MCP_MODE=http and PORT=8080)
- Endpoints: /health, /stats, /version
- WebSocket: ws://localhost:8080/

Tabnine does not need HTTP mode; it uses stdio mode by default.

## Docs
- docs/BUILDING.md
- docs/PLATFORM_SUPPORT.md
- docs/USAGE.md
- docs/TABNINE_SETUP.md
- docs/TABNINE_QUICK_REFERENCE.md
- docs/WORKFLOWS.md
- docs/mcp-introduction.md

## License
MIT License

## Features

- **Web Search**: Search the web using DuckDuckGo
- **MCP Compliant**: Implements the Model Context Protocol specification
- **Tabnine Ready**: Pre-configured for Tabnine Agents integration
- **Stdio Communication**: Direct communication via standard input/output (default)
- **HTTP Mode**: Optional HTTP/WebSocket mode for testing and debugging
- **Configurable Results**: Specify maximum number of search results
- **Clean Response Format**: Structured search results with titles, URLs, and descriptions
- **Performance Monitoring**: Built-in health checks and statistics
- **Lightweight Binary**: Single portable Go binary
- **Multi-Platform Support**: Pre-built binaries for macOS (Intel & Apple Silicon), Windows (Intel & ARM), and Linux (Intel & ARM)

## 🚀 Quick Start

### Download Pre-built Binary

Download the appropriate binary for your platform from the [latest release](https://github.com/yourusername/websearch-mcp/releases/latest):

#### macOS
```bash
# Apple Silicon (M1/M2/M3)
curl -LO https://github.com/youruser/websearch-mcp/releases/latest/download/websearch-mcp-VERSION-darwin-arm64.tar.gz
tar -xzf websearch-mcp-VERSION-darwin-arm64.tar.gz
chmod +x websearch-mcp-darwin-arm64

# Intel
curl -LO https://github.com/youruser/websearch-mcp/releases/latest/download/websearch-mcp-VERSION-darwin-amd64.tar.gz
tar -xzf websearch-mcp-VERSION-darwin-amd64.tar.gz
chmod +x websearch-mcp-darwin-amd64
```

#### Windows
```powershell
# Download from releases page and extract

# Intel/AMD64
# websearch-mcp-windows-amd64.exe

# ARM64
# websearch-mcp-windows-arm64.exe
```

#### Linux
```bash
# Intel/AMD64
wget https://github.com/youruser/websearch-mcp/releases/latest/download/websearch-mcp-VERSION-linux-amd64.tar.gz
tar -xzf websearch-mcp-VERSION-linux-amd64.tar.gz
chmod +x websearch-mcp-linux-amd64

# ARM64
wget https://github.com/youruser/websearch-mcp/releases/latest/download/websearch-mcp-VERSION-linux-arm64.tar.gz
tar -xzf websearch-mcp-VERSION-linux-arm64.tar.gz
chmod +x websearch-mcp-linux-arm64
```

### Build from Source

If you prefer to build from source:

```bash
# Clone the repository
git clone <repository-url>
cd websearch-mcp

# Build for your platform
make build

# Or build for all platforms
make build-all-platforms
```

See [docs/BUILDING.md](docs/BUILDING.md) for detailed build instructions.

## 📦 Supported Platforms

| Platform | Architecture | Status |
|----------|--------------|--------|
| macOS 11+ | Apple Silicon (ARM64) | ✅ Fully Supported |
| macOS 11+ | Intel (x86_64) | ✅ Fully Supported |
| Windows 10/11 | Intel/AMD (x86_64) | ✅ Fully Supported |
| Windows 10/11 | ARM64 | ✅ Supported |
| Linux | Intel/AMD (x86_64) | ✅ Fully Supported |
| Linux | ARM64 | ✅ Fully Supported |

See [docs/PLATFORM_SUPPORT.md](docs/PLATFORM_SUPPORT.md) for detailed platform information.

## Quick Start with Tabnine

### 1. Download or Build the Server

Download the binary for your platform (see above) or build from source.

### 2. Configure Tabnine MCP

Create a `.mcp_servers` file in your project root:

```json
{
  "mcpServers": {
    "websearch": {
      "command": "./websearch-mcp-darwin-arm64",
      "args": [],
      "env": {}
    }
  }
}
```

**Note:** Update the `command` path to match your platform's binary name:
- macOS Apple Silicon: `./websearch-mcp-darwin-arm64`
- macOS Intel: `./websearch-mcp-darwin-amd64`
- Windows Intel: `websearch-mcp-windows-amd64.exe`
- Windows ARM: `websearch-mcp-windows-arm64.exe`
- Linux Intel: `./websearch-mcp-linux-amd64`
- Linux ARM: `./websearch-mcp-linux-arm64`

The server runs in **stdio mode** by default, which means it communicates directly with Tabnine via standard input/output. No HTTP server is started, and the server is launched automatically when Tabnine needs it.

### 3. Test Integration

Ask your Tabnine Agent:
```
"Search for 'Go programming best practices' and give me the top 5 results"
```

For detailed setup instructions, see [docs/TABNINE_SETUP.md](docs/TABNINE_SETUP.md).

## Usage

### Running the Server

#### Stdio Mode (Default - for MCP clients)

The server runs in stdio mode by default when called by MCP clients like Tabnine:

```bash
# The server is launched automatically by Tabnine
# You can test it manually by running:
./websearch-mcp-darwin-arm64

# Type JSON-RPC messages (see examples below)
```

#### HTTP Mode (For testing and debugging)

```bash
# macOS/Linux
./websearch-mcp-darwin-arm64 --http 8080

# Windows
websearch-mcp-windows-amd64.exe --http 8080

# Or use environment variable
MCP_MODE=http PORT=8080 ./websearch-mcp-darwin-arm64
```

#### Command Line Options

```bash
# Show help
./websearch-mcp --help

# Show version
./websearch-mcp --version

# Run in stdio mode (default)
./websearch-mcp --stdio

# Run in HTTP mode
./websearch-mcp --http [port]
```

### Endpoints (HTTP Mode Only)

When running in HTTP mode, the server exposes:
- Health: `http://localhost:8080/health`
- Stats: `http://localhost:8080/stats`
- Version: `http://localhost:8080/version`

## 📚 Documentation

- **[Building](docs/BUILDING.md)**: Detailed build instructions for all platforms
- **[Platform Support](docs/PLATFORM_SUPPORT.md)**: Platform-specific installation and compatibility
- **[Usage Guide](docs/USAGE.md)**: How to use the server
- **[Tabnine Setup](docs/TABNINE_SETUP.md)**: Integration with Tabnine Agents
- **[Tabnine Quick Reference](docs/TABNINE_QUICK_REFERENCE.md)**: Quick reference for Tabnine integration
- **[Workflows](docs/WORKFLOWS.md)**: CI/CD and release processes
- **[MCP Introduction](docs/mcp-introduction.md)**: Understanding the Model Context Protocol

## MCP Protocol Implementation

### Communication Modes

#### Stdio Mode (Default)
- Messages are sent line-by-line via stdin
- Responses are written line-by-line to stdout
- Logs go to stderr
- Best for integration with MCP clients

#### HTTP Mode
- WebSocket communication at `ws://localhost:8080/`
- HTTP endpoints for health checks and stats
- Best for testing and debugging

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
        "text": "Search results for: Go programming best practices\\nFound 8 results:\\n\\n1. Go Best Practices\\n   URL: https://example.com/go-best-practices\\n   Description: A comprehensive guide to Go programming best practices...\\n\\n..."
      }
    ]
  }
}
```

## Configuration

### Environment Variables

- PORT: Server port (default: 8080)
- MCP_MODE: Communication mode ('stdio' or 'http', default: 'stdio')

## Development

### Running in Development Mode

```bash
go run main.go
```

### Testing the Server

You can test the server using a WebSocket client or the provided test scripts.

## Dependencies

- github.com/PuerkitoBio/goquery: HTML parsing for web scraping
- github.com/gorilla/websocket: WebSocket implementation

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
4. **Wrong Architecture**: Download the correct binary for your platform (see [Platform Support](docs/PLATFORM_SUPPORT.md))

### Platform-Specific Issues

See [docs/PLATFORM_SUPPORT.md](docs/PLATFORM_SUPPORT.md) for platform-specific troubleshooting.

### Logging

The server logs all connections, disconnections, and errors to help with debugging.

## 🔗 Related Resources

- [Model Context Protocol Specification](https://modelcontextprotocol.io/)
- [Tabnine Documentation](https://www.tabnine.com/docs)
- [DuckDuckGo Search](https://duckduckgo.com/)
