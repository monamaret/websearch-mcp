#!/bin/bash

# WebSearch MCP Server Start Script

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Default mode and port
MODE=${MCP_MODE:-http}
PORT=${PORT:-8080}

echo -e "${GREEN}Starting WebSearch MCP Server...${NC}"

# Check if binary exists
if [ ! -f "./websearch-mcp" ]; then
    echo -e "${YELLOW}Binary not found. Building...${NC}"
    CGO_ENABLED=0 go build -ldflags="-w -s" -o websearch-mcp .
    echo -e "${GREEN}Build complete!${NC}"
    echo ""
fi

# Start server based on mode
if [ "$MODE" = "stdio" ]; then
    echo -e "${YELLOW}Mode: Stdio (for MCP clients)${NC}"
    echo -e "${YELLOW}Note: Server will read from stdin and write to stdout${NC}"
    echo -e "${YELLOW}Logs will be written to stderr${NC}"
    echo ""
    echo -e "${GREEN}Server starting in stdio mode...${NC}"
    echo -e "${YELLOW}Press Ctrl+C to stop${NC}"
    echo ""
    exec ./websearch-mcp --stdio
else
    echo -e "${YELLOW}Mode: HTTP (for testing/debugging)${NC}"
    echo -e "${YELLOW}Port: ${PORT}${NC}"
    echo -e "${YELLOW}Health check: http://localhost:${PORT}/health${NC}"
    echo -e "${YELLOW}Stats: http://localhost:${PORT}/stats${NC}"
    echo -e "${YELLOW}Version: http://localhost:${PORT}/version${NC}"
    echo ""
    echo -e "${GREEN}Server starting in HTTP mode...${NC}"
    echo -e "${YELLOW}Press Ctrl+C to stop${NC}"
    echo ""
    export PORT=$PORT
    exec ./websearch-mcp --http $PORT
fi