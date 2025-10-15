#!/bin/bash
# Test script for MCP server in stdio mode

echo "Testing WebSearch MCP Server in stdio mode..."
echo ""

# Test 1: Initialize
echo "Test 1: Initialize"
echo '{"jsonrpc":"2.0","id":1,"method":"initialize","params":{"protocolVersion":"2024-11-05","capabilities":{},"clientInfo":{"name":"test-client","version":"1.0.0"}}}' | ./websearch-mcp --stdio 2>/dev/null
echo ""

# Test 2: List tools
echo "Test 2: List tools"
echo '{"jsonrpc":"2.0","id":2,"method":"tools/list"}' | ./websearch-mcp --stdio 2>/dev/null
echo ""

# Test 3: Ping
echo "Test 3: Ping"
echo '{"jsonrpc":"2.0","id":3,"method":"ping"}' | ./websearch-mcp --stdio 2>/dev/null
echo ""

echo "Tests completed!"
