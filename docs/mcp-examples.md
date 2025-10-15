# MCP Examples and Advanced Usage

This guide provides practical examples, troubleshooting steps, and advanced configurations for MCP servers with Tabnine Agents.

## 🎯 Popular MCP Server Examples

### JIRA & Confluence Integration

**Installation:**
```bash
# Using Docker (recommended)
docker pull ghcr.io/sooperset/mcp-atlassian:latest
```

**Configuration:**
```json
{
  "mcpServers": {
    "mcp-atlassian": {
      "command": "docker",
      "args": [
        "run",
        "-i",
        "--rm",
        "-e", "CONFLUENCE_URL",
        "-e", "CONFLUENCE_USERNAME",
        "-e", "CONFLUENCE_API_TOKEN",
        "-e", "JIRA_URL",
        "-e", "JIRA_USERNAME",
        "-e", "JIRA_API_TOKEN",
        "ghcr.io/sooperset/mcp-atlassian:latest"
      ],
      "env": {
        "CONFLUENCE_URL": "https://your-company.atlassian.net/wiki",
        "CONFLUENCE_USERNAME": "your.email@company.com",
        "CONFLUENCE_API_TOKEN": "your_confluence_api_token",
        "JIRA_URL": "https://your-company.atlassian.net",
        "JIRA_USERNAME": "your.email@company.com",
        "JIRA_API_TOKEN": "your_jira_api_token"
      }
    }
  }
}
```

**Usage Examples:**
```
"Create a JIRA ticket for the authentication bug in the login system"
"Search Confluence for documentation about our API rate limiting"
"What are my assigned JIRA tickets for this sprint?"
"Update the project roadmap page in Confluence with the new milestone dates"
```

### GitHub Integration

**Installation:**
```bash
npm install -g @modelcontextprotocol/server-github
```

**Configuration:**
```json
{
  "mcpServers": {
    "github": {
      "command": "mcp-server-github",
      "args": [],
      "env": {
        "GITHUB_PERSONAL_ACCESS_TOKEN": "your_github_token"
      }
    }
  }
}
```

**Usage Examples:**
```
"Show me recent pull requests on the main repository"
"Create a new issue for the performance optimization task"
"What are the open issues assigned to me?"
"Check the commit history for the authentication module"
```

### PostgreSQL Database

**Installation:**
```bash
npm install -g @modelcontextprotocol/server-postgres
```

**Configuration:**
```json
{
  "mcpServers": {
    "postgres": {
      "command": "mcp-server-postgres",
      "args": [],
      "env": {
        "POSTGRES_CONNECTION_STRING": "postgresql://user:pass@localhost:5432/dbname"
      }
    }
  }
}
```

**Usage Examples:**
```
"Show me the schema for the users table"
"Find all customers who haven't logged in for 30 days"
"What's the average order value for this month?"
"Check for any foreign key constraints on the products table"
```

## 🔍 Troubleshooting

### Check Tabnine Logs

**In VS Code:**
1. Press `Ctrl+Shift+P` (or `Cmd+Shift+P` on Mac)
2. Type "Tabnine: Open logs"
3. Look for MCP-related error messages

### Common Issues and Solutions

**❌ "Cannot connect to MCP server"**
```bash
# Check if Docker is running (for Docker-based servers)
docker ps

# Verify network connectivity
ping api.github.com

# Test credentials manually
curl -H "Authorization: token YOUR_TOKEN" https://api.github.com/user
```

**❌ "Authentication failed"**
- Verify API tokens are correct and not expired
- Check that user permissions are sufficient
- Ensure environment variables are set correctly

**❌ "MCP server not found"**
- Verify the server is installed correctly
- Check the command path in configuration
- Ensure all dependencies are installed

### Debugging Steps

1. **Validate Configuration**:
   ```bash
   # Check JSON syntax
   cat .mcp_servers | python -m json.tool
   ```

2. **Test Server Manually**:
   ```bash
   # For Docker-based servers
   docker run -it --rm -e JIRA_URL="$JIRA_URL" ghcr.io/sooperset/mcp-atlassian:latest
   ```

3. **Check Permissions**:
   ```bash
   # Verify file permissions
   ls -la .mcp_servers
   ```

## 🔧 Advanced Configuration

### Multiple MCP Servers

Configure multiple services in one file:

```json
{
  "mcpServers": {
    "jira": {
      "command": "docker",
      "args": ["run", "-i", "--rm", "-e", "JIRA_URL", "-e", "JIRA_API_TOKEN", "mcp-jira:latest"],
      "env": {
        "JIRA_URL": "${JIRA_URL}",
        "JIRA_API_TOKEN": "${JIRA_API_TOKEN}"
      }
    },
    "github": {
      "command": "mcp-server-github",
      "args": [],
      "env": {
        "GITHUB_PERSONAL_ACCESS_TOKEN": "${GITHUB_TOKEN}"
      }
    },
    "postgres": {
      "command": "mcp-server-postgres",
      "args": [],
      "env": {
        "POSTGRES_CONNECTION_STRING": "${DATABASE_URL}"
      }
    }
  }
}
```

### Custom MCP Servers

Create your own MCP server for internal tools:

```json
{
  "mcpServers": {
    "internal-api": {
      "command": "node",
      "args": ["./scripts/internal-mcp-server.js"],
      "env": {
        "INTERNAL_API_KEY": "${INTERNAL_API_KEY}",
        "API_BASE_URL": "https://internal-api.company.com"
      }
    }
  }
}
```

## 🎯 Team Setup

### Shared Configuration

For team environments, create a shared configuration template:

```json
{
  "mcpServers": {
    "team-jira": {
      "command": "docker",
      "args": ["run", "-i", "--rm", "-e", "JIRA_URL", "-e", "JIRA_API_TOKEN", "mcp-atlassian:latest"],
      "env": {
        "JIRA_URL": "https://company.atlassian.net",
        "JIRA_USERNAME": "${USER_EMAIL}",
        "JIRA_API_TOKEN": "${USER_JIRA_TOKEN}"
      }
    }
  }
}
```

### Environment Setup Script

Create a setup script for new team members:

```bash
#!/bin/bash
# setup-mcp.sh

echo "Setting up MCP for Tabnine Agents..."

# Copy template
cp .mcp_servers.template .mcp_servers

# Set environment variables
echo "Please set the following environment variables:"
echo "export JIRA_API_TOKEN='your_token'"
echo "export GITHUB_TOKEN='your_token'"

# Install dependencies
npm install -g @modelcontextprotocol/server-github
```

## 💡 Real-World Workflows

### Development Workflow

**Sprint Planning:**
```
"Check my assigned JIRA tickets for the current sprint"
"Create a branch for the user authentication feature"
"What are the acceptance criteria for ticket AUTH-123?"
```

**Code Review Process:**
```
"Create a pull request for the authentication feature"
"What feedback did I receive on my last PR?"
"Merge the approved changes to main branch"
```

**Bug Tracking:**
```
"Create a JIRA ticket for the login timeout issue"
"Search for similar authentication bugs in our database"
"Update the bug report with the root cause analysis"
```

### Documentation Workflow

**Knowledge Management:**
```
"Search Confluence for our API documentation standards"
"Create a new page documenting the webhook implementation"
"Update the onboarding guide with the new security requirements"
```

**Project Documentation:**
```
"Document the new microservice architecture in Confluence"
"Update the deployment guide with Docker instructions"
"Create a troubleshooting guide for common database issues"
```

### Data Analysis Workflow

**Database Queries:**
```
"Show me the top 10 customers by order volume this quarter"
"Find all incomplete orders from the last 48 hours"
"What's the average response time for our API endpoints?"
```

**Performance Monitoring:**
```
"Check for any slow queries in the performance schema"
"Analyze user engagement metrics for the new feature"
"Generate a report of system health metrics"
```

## 🚀 Advanced Use Cases

### CI/CD Integration

```json
{
  "mcpServers": {
    "jenkins": {
      "command": "mcp-server-jenkins",
      "args": [],
      "env": {
        "JENKINS_URL": "https://ci.company.com",
        "JENKINS_TOKEN": "${JENKINS_API_TOKEN}"
      }
    }
  }
}
```

**Usage:**
```
"Trigger a build for the staging environment"
"Check the status of the latest deployment pipeline"
"Deploy the feature branch to the testing environment"
```

### Monitoring and Alerting

```json
{
  "mcpServers": {
    "grafana": {
      "command": "mcp-server-grafana",
      "args": [],
      "env": {
        "GRAFANA_URL": "https://monitoring.company.com",
        "GRAFANA_API_KEY": "${GRAFANA_API_KEY}"
      }
    }
  }
}
```

**Usage:**
```
"Show me the current system performance metrics"
"Check for any active alerts in the monitoring system"
"Generate a performance report for the last 24 hours"
```

### Customer Support Integration

```json
{
  "mcpServers": {
    "zendesk": {
      "command": "mcp-server-zendesk",
      "args": [],
      "env": {
        "ZENDESK_SUBDOMAIN": "company",
        "ZENDESK_EMAIL": "${SUPPORT_EMAIL}",
        "ZENDESK_TOKEN": "${ZENDESK_API_TOKEN}"
      }
    }
  }
}
```

**Usage:**
```
"Check for high-priority support tickets"
"Create a ticket for the reported API issue"
"Update the customer about the bug fix deployment"
```

## 📊 Performance Optimization

### Connection Pooling

For database connections, configure connection pooling:

```json
{
  "mcpServers": {
    "postgres": {
      "command": "mcp-server-postgres",
      "args": ["--pool-size=10", "--max-connections=20"],
      "env": {
        "POSTGRES_CONNECTION_STRING": "${DATABASE_URL}",
        "PGPOOL_ENABLED": "true"
      }
    }
  }
}
```

### Caching Strategies

Implement caching for frequently accessed data:

```json
{
  "mcpServers": {
    "github": {
      "command": "mcp-server-github",
      "args": ["--cache-ttl=300"],
      "env": {
        "GITHUB_PERSONAL_ACCESS_TOKEN": "${GITHUB_TOKEN}",
        "ENABLE_CACHE": "true"
      }
    }
  }
}
```

### Rate Limiting

Configure rate limiting to avoid API throttling:

```json
{
  "mcpServers": {
    "jira": {
      "command": "docker",
      "args": ["run", "-i", "--rm", "-e", "RATE_LIMIT=100", "mcp-jira:latest"],
      "env": {
        "JIRA_URL": "${JIRA_URL}",
        "JIRA_API_TOKEN": "${JIRA_API_TOKEN}",
        "RATE_LIMIT_PER_MINUTE": "100"
      }
    }
  }
}
```

## 🔮 Next Steps

- Explore [MCP Server Development](https://github.com/modelcontextprotocol) for creating custom servers
- Check [Guidelines Setup](./guidelines-setup.md) to customize agent behavior
- Join the [MCP Community](https://github.com/modelcontextprotocol/servers) for more examples

## 📚 Additional Resources

- [Official MCP Documentation](https://modelcontextprotocol.io)
- [MCP Server Registry](https://github.com/modelcontextprotocol/servers)
- [Tabnine Agents Documentation](https://docs.tabnine.com/agents)