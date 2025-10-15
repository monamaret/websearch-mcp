# Tabnine WebSearch MCP - Quick Reference

## 🚀 One-Minute Setup

```bash
# 1. Build the server
make build

# 2. Setup Tabnine configuration
make setup-tabnine

# 3. Test the integration
make tabnine-demo
```

## 📝 Configuration File

**File**: `.mcp_servers` (in your project root)

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

## 💬 Example Tabnine Prompts

### Research & Documentation
```
"Search for 'React 18 new features' and summarize the key improvements"

"Find the latest security best practices for Node.js applications"

"Look up performance optimization techniques for PostgreSQL databases"
```

### Problem Solving
```
"Search for solutions to Docker memory leak issues"

"Find troubleshooting guides for Kubernetes networking problems"

"Look up common causes of 'CORS error' in web applications"
```

### Learning & Development
```
"Search for Go concurrency patterns and examples"

"Find tutorials on implementing OAuth2 in microservices"

"Look up the latest developments in WebAssembly"
```

### Code Quality & Best Practices
```
"Search for code review checklist templates for Python projects"

"Find examples of clean architecture in Go applications"

"Look up testing strategies for distributed systems"
```

## 🔧 Verification Commands

```bash
# Check if server is running
curl http://localhost:8080/health

# View server statistics
curl http://localhost:8080/stats

# Validate configuration
make validate-tabnine

# Test WebSocket connection
cd examples && go run test-client.go
```

## ⚡ Quick Troubleshooting

| Issue | Solution |
|-------|----------|
| "Cannot see websearch server" | Run `make setup-tabnine` |
| "Connection refused" | Start server with `./start-server.sh` |
| "No search results" | Check internet connection |
| "Port already in use" | Change PORT in `.mcp_servers` |

## 📊 Server Endpoints

- **WebSocket**: `ws://localhost:8080/`
- **Health**: `http://localhost:8080/health`
- **Stats**: `http://localhost:8080/stats`

## 🎯 MCP Tools Available

| Tool | Description | Parameters |
|------|-------------|------------|
| `web_search` | Search the web using DuckDuckGo | `query` (required), `max_results` (1-20, default: 10) |

## 📚 Documentation Links

- **Full Setup Guide**: [TABNINE_SETUP.md](TABNINE_SETUP.md)
- **API Documentation**: [USAGE.md](USAGE.md)
- **Project README**: [README.md](README.md)

## 🚨 Need Help?

1. **Check logs**: Server outputs to console
2. **Run tests**: `go test -v`
3. **Validate setup**: `make validate-tabnine`
4. **Review docs**: Open `TABNINE_SETUP.md`

---

**Pro Tip**: Start your prompts with context about what you're working on for better search results!

Example: *"I'm building a React e-commerce app. Search for best practices for handling payment processing securely."*