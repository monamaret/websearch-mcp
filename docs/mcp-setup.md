# Setting up MCP with Tabnine Agents

This guide provides step-by-step instructions for configuring Model Context Protocol (MCP) servers to work with Tabnine Agents, enabling seamless integration with external tools and services.

## 🚀 Quick Setup Overview

1. **Create MCP Configuration**: Set up `.mcp_servers` file
2. **Install MCP Server**: Choose and install your desired MCP server
3. **Configure Authentication**: Set up API keys and credentials
4. **Test Connection**: Verify integration works with Tabnine
5. **Start Using**: Begin leveraging external tools in your workflows

## 📁 Step 1: Create MCP Configuration File

Create a `.mcp_servers` file in your project root directory. This file will contain the configuration for all your MCP servers.

## 🔧 Step 2: Basic Configuration Structure

The `.mcp_servers` file uses JSON format:

```json
{
  "mcpServers": {
    "server-name": {
      "command": "command-to-run-server",
      "args": ["argument1", "argument2"],
      "env": {
        "ENV_VARIABLE": "value"
      }
    }
  }
}
```

### Configuration Parameters

- **`command`**: The executable that runs the MCP server
- **`args`**: Array of command-line arguments
- **`env`**: Environment variables for the server (API keys, URLs, etc.)

## 🔐 Step 3: Authentication Setup

### API Token Creation

**For JIRA/Confluence:**
1. Go to Atlassian Account Settings
2. Navigate to Security → API Tokens
3. Create a new token
4. Copy and store securely

**For GitHub:**
1. Go to GitHub Settings → Developer Settings
2. Personal Access Tokens → Tokens (classic)
3. Generate new token with appropriate scopes
4. Copy and store securely

**For Databases:**
1. Create a dedicated service account
2. Grant minimum required permissions
3. Use connection strings with authentication

### Environment Variables (Recommended)

Instead of hardcoding credentials, use environment variables:

```json
{
  "mcpServers": {
    "mcp-atlassian": {
      "command": "docker",
      "args": ["run", "-i", "--rm", "-e", "JIRA_URL", "-e", "JIRA_API_TOKEN", "mcp-atlassian:latest"],
      "env": {
        "JIRA_URL": "${JIRA_URL}",
        "JIRA_API_TOKEN": "${JIRA_API_TOKEN}"
      }
    }
  }
}
```

Then set environment variables:
```bash
export JIRA_URL="https://your-company.atlassian.net"
export JIRA_API_TOKEN="your_token_here"
```

## ✅ Step 4: Verify MCP Server Connection

### Test the Connection

Ask your Tabnine Agent:

```
"Are you able to see the JIRA MCP server?"
```

**Expected Response:**
- ✅ "Yes, I can see the JIRA MCP server and access its capabilities"
- ❌ "No, I cannot connect to the JIRA MCP server"

### Basic Functionality Test

Try a simple operation:

```
"List my recent JIRA issues"
"Search for 'authentication' in Confluence"
"Show me the database schema"
```

## 📋 Best Practices

### Security
- **Never commit API tokens** to version control
- **Use environment variables** for all sensitive data
- **Implement least privilege** access for service accounts
- **Regularly rotate** API tokens

### Performance
- **Start with essential tools** to avoid overwhelming the agent
- **Monitor response times** and optimize configurations
- **Use caching** where appropriate
- **Consider rate limiting** for API-heavy operations

### Maintenance
- **Keep MCP servers updated** to the latest versions
- **Document your configuration** for team members
- **Test integrations regularly** to catch issues early
- **Monitor usage patterns** to optimize workflows

## 🔮 Next Steps

Now that you have MCP configured, explore:

- [MCP Examples](./mcp-examples.md) - Real-world use cases and workflows
- [Guidelines Setup](./guidelines-setup.md) - Customize agent behavior

Or start using your MCP integration:

```
"Create a JIRA ticket for the authentication feature we just implemented"
"Update the project documentation in Confluence with the new API endpoints"
"Deploy this feature to staging using our CI/CD pipeline"
```