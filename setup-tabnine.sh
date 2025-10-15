#!/bin/bash

# Tabnine MCP Setup Script for WebSearch Server

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${GREEN}🔧 Tabnine WebSearch MCP Setup${NC}"
echo -e "${YELLOW}=============================${NC}"
echo ""

# Check if we're in the correct directory
if [ ! -f "main.go" ] || [ ! -f "go.mod" ]; then
    echo -e "${RED}❌ Please run this script from the websearch-mcp project directory${NC}"
    exit 1
fi

# Step 1: Build the server
echo -e "${YELLOW}📦 Step 1: Building WebSearch MCP Server...${NC}"
# Ensure server binary exists
if [ ! -f "./websearch-mcp" ]; then
    echo -e "${BLUE}Building server binary (single binary)...${NC}"
    CGO_ENABLED=0 go build -ldflags="-w -s" -o websearch-mcp .
    echo -e "${GREEN}✅ Server built successfully${NC}"
else
    echo -e "${GREEN}✅ Server binary found${NC}"
fi
echo ""

# Step 2: Create .mcp_servers file
echo -e "${YELLOW}📝 Step 2: Creating Tabnine MCP Configuration...${NC}"

# Determine the absolute path to the websearch-mcp binary
CURRENT_DIR=$(pwd)
BINARY_PATH="$CURRENT_DIR/websearch-mcp"

# Check if .mcp_servers already exists
if [ -f ".mcp_servers" ]; then
    echo -e "${YELLOW}⚠️  .mcp_servers file already exists${NC}"
    read -p "Do you want to overwrite it? (y/N): " -n 1 -r
    echo
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        echo -e "${BLUE}Skipping .mcp_servers creation${NC}"
        echo ""
        echo -e "${YELLOW}Your existing configuration:${NC}"
        cat .mcp_servers
        echo ""
    else
        CREATE_CONFIG=true
    fi
else
    CREATE_CONFIG=true
fi

if [ "$CREATE_CONFIG" = true ]; then
    cat > .mcp_servers << EOF
{
  "mcpServers": {
    "websearch": {
      "command": "$BINARY_PATH",
      "args": [],
      "env": {
        "PORT": "8080"
      }
    }
  }
}
EOF
    echo -e "${GREEN}✅ Created .mcp_servers configuration${NC}"
    echo -e "${BLUE}Configuration saved to: $(pwd)/.mcp_servers${NC}"
fi
echo ""

# Step 3: Test the server
echo -e "${YELLOW}🧪 Step 3: Testing Server Configuration...${NC}"
echo -e "${BLUE}Starting server for quick test...${NC}"

# Start server in background for testing
./websearch-mcp &
SERVER_PID=$!

# Wait for server to start
sleep 2

# Test health endpoint
if curl -s http://localhost:8080/health > /dev/null; then
    echo -e "${GREEN}✅ Server is running and healthy${NC}"
    
    # Test stats endpoint
    echo -e "${BLUE}Server statistics:${NC}"
    curl -s http://localhost:8080/stats | python3 -m json.tool 2>/dev/null || curl -s http://localhost:8080/stats
else
    echo -e "${RED}❌ Server health check failed${NC}"
fi

# Stop test server
kill $SERVER_PID 2>/dev/null || true
echo ""

# Step 4: Display next steps
echo -e "${GREEN}🎉 Setup Complete!${NC}"
echo -e "${YELLOW}=================${NC}"
echo ""
echo -e "${BLUE}Next Steps:${NC}"
echo -e "${YELLOW}1.${NC} Ensure Tabnine Agents is enabled in your IDE"
echo -e "${YELLOW}2.${NC} The .mcp_servers file is ready for Tabnine to discover"
echo -e "${YELLOW}3.${NC} Test the integration by asking your Tabnine Agent:"
echo ""
echo -e "${GREEN}   \"Can you see the websearch MCP server?\"${NC}"
echo -e "${GREEN}   \"Search for 'Go programming tutorials' and give me the top 3 results\"${NC}"
echo ""
echo -e "${BLUE}Configuration Details:${NC}"
echo -e "  📁 Config file: ${YELLOW}$(pwd)/.mcp_servers${NC}"
echo -e "  🔧 Server binary: ${YELLOW}$BINARY_PATH${NC}"
echo -e "  🌐 Server port: ${YELLOW}8080${NC}"
echo -e "  📊 Health check: ${YELLOW}http://localhost:8080/health${NC}"
echo -e "  📈 Statistics: ${YELLOW}http://localhost:8080/stats${NC}"
echo ""
echo -e "${BLUE}Useful Commands:${NC}"
echo -e "  🚀 Start server: ${YELLOW}./start-server.sh${NC}"
echo -e "  🧪 Run tests: ${YELLOW}make test${NC}"
echo -e "  📚 Full guide: ${YELLOW}cat TABNINE_SETUP.md${NC}"
echo ""
echo -e "${GREEN}For detailed configuration options and troubleshooting, see:${NC}"
echo -e "${BLUE}👉 TABNINE_SETUP.md${NC}"
