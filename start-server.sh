#!/bin/bash

# WebSearch MCP Server Start Script

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Default port
PORT=${PORT:-8080}

echo -e "${GREEN}Starting WebSearch MCP Server...${NC}"
echo -e "${YELLOW}Port: ${PORT}${NC}"
echo -e "${YELLOW}WebSocket endpoint: ws://localhost:${PORT}/${NC}"
echo -e "${YELLOW}Health check: http://localhost:${PORT}/health${NC}"
echo ""

# Check if binary exists
if [ ! -f "./websearch-mcp" ]; then
    echo -e "${YELLOW}Binary not found. Building...${NC}"
    go build -o websearch-mcp .
    echo -e "${GREEN}Build complete!${NC}"
    echo ""
fi

# Start the server
echo -e "${GREEN}Server starting...${NC}"
echo -e "${YELLOW}Press Ctrl+C to stop${NC}"
echo ""

export PORT=$PORT
exec ./websearch-mcp