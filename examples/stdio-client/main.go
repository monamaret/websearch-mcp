//go:build examples
// +build examples

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os/exec"
	"time"
)

type MCPMessage struct {
	JSONRPC string      `json:"jsonrpc"`
	ID      interface{} `json:"id,omitempty"`
	Method  string      `json:"method,omitempty"`
	Params  interface{} `json:"params,omitempty"`
	Result  interface{} `json:"result,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func main() {
	// Start the MCP server in stdio mode
	cmd := exec.Command("../../websearch-mcp", "--stdio")

	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal("Error creating stdin pipe:", err)
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal("Error creating stdout pipe:", err)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal("Error creating stderr pipe:", err)
	}

	// Start the server
	if err := cmd.Start(); err != nil {
		log.Fatal("Error starting server:", err)
	}

	// Read stderr in background
	go func() {
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			log.Printf("[Server] %s", scanner.Text())
		}
	}()

	// Test messages
	testMessages := []MCPMessage{
		{
			JSONRPC: "2.0",
			ID:      1,
			Method:  "initialize",
			Params: map[string]interface{}{
				"protocolVersion": "2024-11-05",
				"capabilities":    map[string]interface{}{},
				"clientInfo": map[string]interface{}{
					"name":    "test-stdio-client",
					"version": "1.0.0",
				},
			},
		},
		{
			JSONRPC: "2.0",
			ID:      2,
			Method:  "tools/list",
		},
		{
			JSONRPC: "2.0",
			ID:      3,
			Method:  "tools/call",
			Params: map[string]interface{}{
				"name": "web_search",
				"arguments": map[string]interface{}{
					"query":       "Go programming tutorial",
					"max_results": 3,
				},
			},
		},
		{
			JSONRPC: "2.0",
			ID:      4,
			Method:  "ping",
		},
	}

	// Send messages and read responses
	encoder := json.NewEncoder(stdin)
	decoder := json.NewDecoder(stdout)

	for i, msg := range testMessages {
		fmt.Printf("\n=== Test %d: %s ===\n", i+1, msg.Method)

		// Send message
		fmt.Println("Sending:")
		jsonBytes, _ := json.MarshalIndent(msg, "", "  ")
		fmt.Println(string(jsonBytes))

		if err := encoder.Encode(msg); err != nil {
			log.Fatal("Error sending message:", err)
		}

		// Read response
		var response MCPMessage
		if err := decoder.Decode(&response); err != nil {
			if err == io.EOF {
				log.Println("Server closed connection")
				break
			}
			log.Fatal("Error reading response:", err)
		}

		// Print response
		fmt.Println("\nReceived:")
		jsonBytes, _ = json.MarshalIndent(response, "", "  ")
		fmt.Println(string(jsonBytes))

		// Wait a bit between messages
		if i < len(testMessages)-1 {
			time.Sleep(500 * time.Millisecond)
		}
	}

	// Close stdin to signal we're done
	stdin.Close()

	// Wait for server to exit
	if err := cmd.Wait(); err != nil {
		// Server may exit with error when stdin closes
		log.Printf("Server exited: %v", err)
	}

	fmt.Println("\nAll tests completed!")
}
