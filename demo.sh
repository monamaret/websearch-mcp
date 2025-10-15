#!/bin/bash

# WebSearch MCP Server Demo Script

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${GREEN}WebSearch MCP Server Demo${NC}"
echo -e "${YELLOW}=========================${NC}"
echo ""

# Build the server if needed
if [ ! -f "./websearch-mcp" ]; then
    echo -e "${YELLOW}Building server...${NC}"
    go build -o websearch-mcp .
    echo -e "${GREEN}Build complete!${NC}"
    echo ""
fi

# Start server in background
echo -e "${YELLOW}Starting MCP server...${NC}"
./websearch-mcp &
SERVER_PID=$!

# Wait for server to start
sleep 2

# Check if server is running
if ! kill -0 $SERVER_PID 2>/dev/null; then
    echo -e "${RED}Failed to start server${NC}"
    exit 1
fi

echo -e "${GREEN}Server started successfully (PID: $SERVER_PID)${NC}"
echo -e "${BLUE}WebSocket endpoint: ws://localhost:8080/${NC}"
echo -e "${BLUE}Health check: http://localhost:8080/health${NC}"
echo ""

# Test health endpoint
echo -e "${YELLOW}Testing health endpoint...${NC}"
HEALTH_RESPONSE=$(curl -s http://localhost:8080/health)
echo -e "${GREEN}Health check response:${NC}"
echo "$HEALTH_RESPONSE" | python3 -m json.tool 2>/dev/null || echo "$HEALTH_RESPONSE"
echo ""

# Build and run test client
echo -e "${YELLOW}Building test client...${NC}"
cd examples
go build -o test-client test-client.go
echo -e "${GREEN}Test client built!${NC}"
echo ""

echo -e "${YELLOW}Running test client to demonstrate MCP functionality...${NC}"
echo -e "${BLUE}This will test:${NC}"
echo -e "${BLUE}  1. Server initialization${NC}"
echo -e "${BLUE}  2. Tools listing${NC}"
echo -e "${BLUE}  3. Web search functionality${NC}"
echo -e "${BLUE}  4. Ping/pong${NC}"
echo ""

# Run test client
./test-client

echo ""
echo -e "${YELLOW}Demo complete! Stopping server...${NC}"

# Cleanup
kill $SERVER_PID 2>/dev/null || true
rm -f test-client
cd ..

echo -e "${GREEN}Server stopped. Demo finished!${NC}"
echo ""
echo -e "${YELLOW}To run the server manually:${NC}"
echo -e "${BLUE}  ./start-server.sh${NC}"
echo ""
echo -e "${YELLOW}To run tests:${NC}"
echo -e "${BLUE}  go test -v${NC}"
echo ""
echo -e "${YELLOW}To build Docker image:${NC}"
echo -e "${BLUE}  make docker-build${NC}"