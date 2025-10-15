package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

func TestWebSearchServer_Initialize(t *testing.T) {
	server := NewWebSearchServer()

	msg := MCPMessage{
		JSONRPC: "2.0",
		ID:      1,
		Method:  "initialize",
		Params: map[string]interface{}{
			"protocolVersion": "2024-11-05",
			"capabilities":    map[string]interface{}{},
			"clientInfo": map[string]interface{}{
				"name":    "test-client",
				"version": "1.0.0",
			},
		},
	}

	response := server.handleMessage(msg)

	if response == nil {
		t.Fatal("Expected response, got nil")
	}

	if response.Error != nil {
		t.Fatalf("Expected no error, got: %v", response.Error)
	}

	result, ok := response.Result.(map[string]interface{})
	if !ok {
		t.Fatal("Expected result to be a map")
	}

	if result["protocolVersion"] != "2024-11-05" {
		t.Errorf("Expected protocolVersion '2024-11-05', got %v", result["protocolVersion"])
	}
}

func TestWebSearchServer_ToolsList(t *testing.T) {
	server := NewWebSearchServer()

	msg := MCPMessage{
		JSONRPC: "2.0",
		ID:      2,
		Method:  "tools/list",
	}

	response := server.handleMessage(msg)

	if response == nil {
		t.Fatal("Expected response, got nil")
	}

	if response.Error != nil {
		t.Fatalf("Expected no error, got: %v", response.Error)
	}

	result, ok := response.Result.(map[string]interface{})
	if !ok {
		t.Fatal("Expected result to be a map")
	}

	tools, ok := result["tools"].([]Tool)
	if !ok {
		t.Fatal("Expected tools to be a slice of Tool")
	}

	if len(tools) != 1 {
		t.Fatalf("Expected 1 tool, got %d", len(tools))
	}

	if tools[0].Name != "web_search" {
		t.Errorf("Expected tool name 'web_search', got %s", tools[0].Name)
	}
}

func TestWebSearchServer_Ping(t *testing.T) {
	server := NewWebSearchServer()

	msg := MCPMessage{
		JSONRPC: "2.0",
		ID:      4,
		Method:  "ping",
	}

	response := server.handleMessage(msg)

	if response == nil {
		t.Fatal("Expected response, got nil")
	}

	if response.Error != nil {
		t.Fatalf("Expected no error, got: %v", response.Error)
	}

	if response.Result != "pong" {
		t.Errorf("Expected result 'pong', got %v", response.Result)
	}
}

func TestWebSearchServer_InvalidMethod(t *testing.T) {
	server := NewWebSearchServer()

	msg := MCPMessage{
		JSONRPC: "2.0",
		ID:      5,
		Method:  "invalid_method",
	}

	response := server.handleMessage(msg)

	if response == nil {
		t.Fatal("Expected response, got nil")
	}

	if response.Error == nil {
		t.Fatal("Expected error, got nil")
	}

	if response.Error.Code != -32601 {
		t.Errorf("Expected error code -32601, got %d", response.Error.Code)
	}
}

func TestWebSearchServer_FormatSearchResults(t *testing.T) {
	server := NewWebSearchServer()

	results := &SearchResponse{
		Query: "test query",
		Results: []SearchResult{
			{
				Title:       "Test Result 1",
				URL:         "https://example.com/1",
				Description: "This is a test result",
				Rank:        1,
			},
			{
				Title:       "Test Result 2",
				URL:         "https://example.com/2",
				Description: "This is another test result",
				Rank:        2,
			},
		},
		Count: 2,
	}

	formatted := server.formatSearchResults(results)

	if !strings.Contains(formatted, "test query") {
		t.Error("Formatted results should contain the query")
	}

	if !strings.Contains(formatted, "Test Result 1") {
		t.Error("Formatted results should contain the first result title")
	}

	if !strings.Contains(formatted, "https://example.com/1") {
		t.Error("Formatted results should contain the first result URL")
	}

	if !strings.Contains(formatted, "Found 2 results") {
		t.Error("Formatted results should contain the result count")
	}
}

func TestWebSearchServer_FormatSearchResultsEmpty(t *testing.T) {
	server := NewWebSearchServer()

	results := &SearchResponse{
		Query:   "empty query",
		Results: []SearchResult{},
		Count:   0,
	}

	formatted := server.formatSearchResults(results)

	expected := "No results found for query: empty query"
	if formatted != expected {
		t.Errorf("Expected '%s', got '%s'", expected, formatted)
	}
}

func TestHealthEndpoint(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":    "healthy",
			"service":   "websearch-mcp",
			"version":   "1.0.0",
			"timestamp": time.Now().UTC().Format(time.RFC3339),
		})
	})

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response map[string]interface{}
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if response["status"] != "healthy" {
		t.Errorf("Expected status 'healthy', got %v", response["status"])
	}

	if response["service"] != "websearch-mcp" {
		t.Errorf("Expected service 'websearch-mcp', got %v", response["service"])
	}
}

// Integration test for WebSocket connection
func TestWebSocketConnection(t *testing.T) {
	server := NewWebSearchServer()

	// Create test server
	s := httptest.NewServer(http.HandlerFunc(server.handleConnection))
	defer s.Close()

	// Convert http://127.0.0.1 to ws://127.0.0.1
	u := "ws" + strings.TrimPrefix(s.URL, "http")

	// Connect to the server
	ws, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer ws.Close()

	// Send initialize message
	initMsg := MCPMessage{
		JSONRPC: "2.0",
		ID:      1,
		Method:  "initialize",
		Params: map[string]interface{}{
			"protocolVersion": "2024-11-05",
			"capabilities":    map[string]interface{}{},
			"clientInfo": map[string]interface{}{
				"name":    "test-client",
				"version": "1.0.0",
			},
		},
	}

	if err := ws.WriteJSON(initMsg); err != nil {
		t.Fatalf("Failed to send message: %v", err)
	}

	// Read response
	var response MCPMessage
	if err := ws.ReadJSON(&response); err != nil {
		t.Fatalf("Failed to read response: %v", err)
	}

	if response.Error != nil {
		t.Fatalf("Received error: %v", response.Error)
	}

	if response.Result == nil {
		t.Fatal("Expected result, got nil")
	}
}
