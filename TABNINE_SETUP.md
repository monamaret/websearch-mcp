# Tabnine MCP Setup Guide - WebSearch Server

This guide will help you configure the WebSearch MCP server to work with Tabnine Agents, enabling your AI assistant to search the web directly within your development workflow.

## 🚀 Quick Setup

### Prerequisites

- **Tabnine with Agent Mode**: Ensure you have Tabnine Agents enabled
- **Go 1.21+**: Required to build the MCP server
- **Network Access**: For web search functionality

### Step 1: Build the WebSearch MCP Server

```bash
# Clone or navigate to the websearch-mcp directory
cd /path/to/websearch-mcp

# Build the server
go build -o websearch-mcp .

# Or use the Makefile
make build
```

### Step 2: Create Tabnine MCP Configuration

Create a `.mcp_servers` file in your project root directory:

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

### Step 3: Test the Integration

1. **Start your Tabnine Agent** with MCP support enabled
2. **Verify connection** by asking:
   ```
   "Can you see the websearch MCP server?"
   ```
3. **Test search functionality**:
   ```
   "Search for 'Go programming best practices' and give me the top 5 results"
   ```

## 📋 Configuration Options

### Basic Configuration

The minimal configuration for the websearch MCP server:

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

### Advanced Configuration

For production environments with custom settings:

```json
{
  "mcpServers": {
    "websearch": {
      "command": "/usr/local/bin/websearch-mcp",
      "args": [],
      "env": {
        "PORT": "${WEBSEARCH_PORT:-8080}",
        "LOG_LEVEL": "${LOG_LEVEL:-info}",
        "TIMEOUT": "${REQUEST_TIMEOUT:-30}"
      }
    }
  }
}
```

### Docker Configuration

If you prefer to run the server in Docker:

```json
{
  "mcpServers": {
    "websearch-docker": {
      "command": "docker",
      "args": [
        "run",
        "-i",
        "--rm",
        "-p", "8080:8080",
        "-e", "PORT=8080",
        "websearch-mcp:latest"
      ],
      "env": {}
    }
  }
}
```

## 🛠️ Deployment Scenarios

### Development Setup

For local development with hot reload:

```bash
# Terminal 1: Start the MCP server
make dev

# Terminal 2: Test with Tabnine Agent
# Your agent can now perform web searches
```

### Production Deployment

1. **Build optimized binary:**
   ```bash
   CGO_ENABLED=0 go build -ldflags="-w -s" -o websearch-mcp-prod .
   ```

2. **Install as system service:**
   ```bash
   sudo cp websearch-mcp-prod /usr/local/bin/websearch-mcp
   sudo chmod +x /usr/local/bin/websearch-mcp
   ```

3. **Update `.mcp_servers` configuration:**
   ```json
   {
     "mcpServers": {
       "websearch": {
         "command": "/usr/local/bin/websearch-mcp",
         "args": [],
         "env": {
           "PORT": "8080"
         }
       }
     }
   }
   ```

### Multi-Tool Setup

Combine websearch with other MCP tools:

```json
{
  "mcpServers": {
    "websearch": {
      "command": "./websearch-mcp",
      "args": [],
      "env": {
        "PORT": "8080"
      }
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

## 🎯 Usage Examples

Once configured, you can use the websearch functionality in natural language:

### Research and Documentation

```
You: "Search for the latest React performance optimization techniques"
Agent: Uses websearch MCP → Returns current best practices and articles

You: "Find documentation about Go's new generics features"
Agent: Searches web → Provides links to official docs and tutorials
```

### Problem Solving

```
You: "Search for solutions to 'connection timeout' errors in Docker"
Agent: Searches → Returns troubleshooting guides and Stack Overflow solutions

You: "Look up common patterns for implementing authentication in microservices"
Agent: Finds → Security best practices and implementation examples
```

### Staying Current

```
You: "What are the latest developments in WebAssembly?"
Agent: Searches → Recent news, updates, and technical articles

You: "Find recent vulnerability reports for Node.js"
Agent: Searches → Security advisories and patch information
```

## 🔧 Troubleshooting

### Common Issues

**Agent Cannot See WebSearch Server**
- Verify the `.mcp_servers` file is in the correct location
- Check that the websearch-mcp binary is built and executable
- Ensure the specified port (8080) is available

**No Search Results**
- Check internet connectivity
- Verify DuckDuckGo is accessible from your network
- Review server logs for error messages

**Connection Errors**
- Confirm the server is running on the specified port
- Check firewall settings
- Validate WebSocket connections are allowed

### Debug Commands

```bash
# Check if server starts correctly
./websearch-mcp

# Test health endpoint
curl http://localhost:8080/health

# View server statistics
curl http://localhost:8080/stats

# Run comprehensive tests
go test -v
```

### Logging

Enable detailed logging by checking the server output:

```bash
# Start with verbose output
./websearch-mcp 2>&1 | tee websearch.log
```

## 🔐 Security Considerations

### Network Security
- The websearch server makes outbound HTTPS requests to DuckDuckGo
- No incoming external connections are required (only WebSocket from Tabnine)
- Consider firewall rules for production deployments

### Data Privacy
- Search queries are sent to DuckDuckGo (privacy-focused search engine)
- No search history is stored by the MCP server
- All communication between Tabnine and MCP server is local

### Access Control
- The MCP server only provides search functionality
- No file system access or system command execution
- Suitable for restricted environments

## 📊 Monitoring

### Health Checks

Monitor server health using the built-in endpoints:

```bash
# Basic health check
curl http://localhost:8080/health

# Detailed statistics
curl http://localhost:8080/stats
```

### Performance Monitoring

The server provides metrics including:
- Request count and response times
- Search operation statistics
- Memory usage and garbage collection
- Active connection count

## 🔮 Advanced Usage

### Environment Variables

Configure the server behavior using environment variables:

```bash
export PORT=8080              # Server port
export LOG_LEVEL=info         # Logging level
export REQUEST_TIMEOUT=30     # HTTP timeout in seconds
```

### Multiple Instances

Run multiple websearch servers for load balancing:

```json
{
  "mcpServers": {
    "websearch-primary": {
      "command": "./websearch-mcp",
      "env": { "PORT": "8080" }
    },
    "websearch-backup": {
      "command": "./websearch-mcp",
      "env": { "PORT": "8081" }
    }
  }
}
```

### Integration with CI/CD

Include websearch functionality in automated workflows:

```yaml
# Example GitHub Action step
- name: Start WebSearch MCP
  run: |
    ./websearch-mcp &
    sleep 2
    # Run your Tabnine Agent tasks that need web search
```

## ✅ Verification Checklist

- [ ] WebSearch MCP server builds successfully
- [ ] `.mcp_servers` file is correctly configured
- [ ] Tabnine Agent can connect to the MCP server
- [ ] Web search functionality works as expected
- [ ] Health endpoints respond correctly
- [ ] Server logs show no errors

## 🚀 Next Steps

With websearch MCP configured, explore:

1. **Combine with other MCP tools** for comprehensive workflows
2. **Create custom search workflows** using Tabnine Agent capabilities
3. **Monitor usage patterns** to optimize your development process
4. **Share configurations** with your team for consistent setups

Start using your enhanced Tabnine Agent:

```
"Search for the latest security best practices for Go web applications and help me implement them in this project"
```

Your agent now has the power of web search integrated directly into your development workflow! 🎉