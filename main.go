package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Build-time variables (set via ldflags)
var (
	version   = "dev"
	buildTime = "unknown"
	gitCommit = "unknown"
)

// Version info structure
type VersionInfo struct {
	Version   string `json:"version"`
	BuildTime string `json:"build_time"`
	GitCommit string `json:"git_commit"`
	GoVersion string `json:"go_version"`
	OS        string `json:"os"`
	Arch      string `json:"arch"`
}

func getVersionInfo() VersionInfo {
	return VersionInfo{
		Version:   version,
		BuildTime: buildTime,
		GitCommit: gitCommit,
		GoVersion: runtime.Version(),
		OS:        runtime.GOOS,
		Arch:      runtime.GOARCH,
	}
}

// Add stats tracking
type ServerStats struct {
	mu              sync.RWMutex
	StartTime       time.Time `json:"start_time"`
	RequestCount    int64     `json:"request_count"`
	SearchCount     int64     `json:"search_count"`
	ConnectionCount int64     `json:"connection_count"`
	ActiveConns     int64     `json:"active_connections"`
	Errors          int64     `json:"errors"`
}

func (s *ServerStats) IncrementRequests() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.RequestCount++
}

func (s *ServerStats) IncrementSearches() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.SearchCount++
}

func (s *ServerStats) IncrementConnections() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.ConnectionCount++
	s.ActiveConns++
}

func (s *ServerStats) DecrementActiveConnections() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.ActiveConns--
}

func (s *ServerStats) IncrementErrors() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Errors++
}

func (s *ServerStats) GetStats() map[string]interface{} {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	uptime := time.Since(s.StartTime)

	return map[string]interface{}{
		"uptime_seconds":     uptime.Seconds(),
		"uptime_human":       uptime.String(),
		"request_count":      s.RequestCount,
		"search_count":       s.SearchCount,
		"connection_count":   s.ConnectionCount,
		"active_connections": s.ActiveConns,
		"errors":             s.Errors,
		"memory": map[string]interface{}{
			"alloc_mb":       float64(m.Alloc) / 1024 / 1024,
			"total_alloc_mb": float64(m.TotalAlloc) / 1024 / 1024,
			"sys_mb":         float64(m.Sys) / 1024 / 1024,
			"gc_cycles":      m.NumGC,
		},
		"goroutines": runtime.NumGoroutine(),
	}
}

// MCP Message types
type MCPMessage struct {
	JSONRPC string      `json:"jsonrpc"`
	ID      interface{} `json:"id,omitempty"`
	Method  string      `json:"method,omitempty"`
	Params  interface{} `json:"params,omitempty"`
	Result  interface{} `json:"result,omitempty"`
	Error   *MCPError   `json:"error,omitempty"`
}

type MCPError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// MCP Server Info
type ServerInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// Tool definitions
type Tool struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	InputSchema interface{} `json:"inputSchema"`
}

type ToolSchema struct {
	Type       string                 `json:"type"`
	Properties map[string]interface{} `json:"properties"`
	Required   []string               `json:"required"`
}

// Search result structures
type SearchResult struct {
	Title       string `json:"title"`
	URL         string `json:"url"`
	Description string `json:"description"`
	Rank        int    `json:"rank"`
}

type SearchResponse struct {
	Query   string         `json:"query"`
	Results []SearchResult `json:"results"`
	Count   int            `json:"count"`
}

// WebSearchServer implements the MCP server
type WebSearchServer struct {
	stats  *ServerStats
	logger *log.Logger
}

func NewWebSearchServer() *WebSearchServer {
	return &WebSearchServer{
		stats: &ServerStats{
			StartTime: time.Now(),
		},
		logger: log.New(os.Stderr, "[MCP] ", log.LstdFlags),
	}
}

func (s *WebSearchServer) handleMessage(msg MCPMessage) *MCPMessage {
	switch msg.Method {
	case "initialize":
		return s.handleInitialize(msg)
	case "tools/list":
		return s.handleToolsList(msg)
	case "tools/call":
		return s.handleToolsCall(msg)
	case "ping":
		return s.handlePing(msg)
	case "stats/get":
		return s.handleGetStats(msg)
	default:
		return &MCPMessage{
			JSONRPC: "2.0",
			ID:      msg.ID,
			Error: &MCPError{
				Code:    -32601,
				Message: "Method not found",
			},
		}
	}
}

func (s *WebSearchServer) handleInitialize(msg MCPMessage) *MCPMessage {
	versionInfo := getVersionInfo()
	return &MCPMessage{
		JSONRPC: "2.0",
		ID:      msg.ID,
		Result: map[string]interface{}{
			"protocolVersion": "2024-11-05",
			"capabilities": map[string]interface{}{
				"tools": map[string]interface{}{},
			},
			"serverInfo": ServerInfo{
				Name:    "websearch-mcp",
				Version: versionInfo.Version,
			},
		},
	}
}

func (s *WebSearchServer) handleToolsList(msg MCPMessage) *MCPMessage {
	tools := []Tool{
		{
			Name:        "web_search",
			Description: "Search the web for information using DuckDuckGo",
			InputSchema: ToolSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"query": map[string]interface{}{
						"type":        "string",
						"description": "The search query to execute",
					},
					"max_results": map[string]interface{}{
						"type":        "integer",
						"description": "Maximum number of results to return (default: 10)",
						"default":     10,
						"minimum":     1,
						"maximum":     20,
					},
				},
				Required: []string{"query"},
			},
		},
	}

	return &MCPMessage{
		JSONRPC: "2.0",
		ID:      msg.ID,
		Result: map[string]interface{}{
			"tools": tools,
		},
	}
}

func (s *WebSearchServer) handleToolsCall(msg MCPMessage) *MCPMessage {
	params, ok := msg.Params.(map[string]interface{})
	if !ok {
		return &MCPMessage{
			JSONRPC: "2.0",
			ID:      msg.ID,
			Error: &MCPError{
				Code:    -32602,
				Message: "Invalid params",
			},
		}
	}

	name, ok := params["name"].(string)
	if !ok {
		return &MCPMessage{
			JSONRPC: "2.0",
			ID:      msg.ID,
			Error: &MCPError{
				Code:    -32602,
				Message: "Missing tool name",
			},
		}
	}

	arguments, ok := params["arguments"].(map[string]interface{})
	if !ok {
		return &MCPMessage{
			JSONRPC: "2.0",
			ID:      msg.ID,
			Error: &MCPError{
				Code:    -32602,
				Message: "Missing tool arguments",
			},
		}
	}

	switch name {
	case "web_search":
		return s.handleWebSearch(msg, arguments)
	default:
		return &MCPMessage{
			JSONRPC: "2.0",
			ID:      msg.ID,
			Error: &MCPError{
				Code:    -32601,
				Message: "Tool not found",
			},
		}
	}
}

func (s *WebSearchServer) handleWebSearch(msg MCPMessage, args map[string]interface{}) *MCPMessage {
	query, ok := args["query"].(string)
	if !ok || query == "" {
		return &MCPMessage{
			JSONRPC: "2.0",
			ID:      msg.ID,
			Error: &MCPError{
				Code:    -32602,
				Message: "Query parameter is required",
			},
		}
	}

	maxResults := 10
	if mr, ok := args["max_results"].(float64); ok {
		maxResults = int(mr)
	}

	s.stats.IncrementSearches()
	results, err := s.performWebSearch(query, maxResults)
	if err != nil {
		s.stats.IncrementErrors()
		return &MCPMessage{
			JSONRPC: "2.0",
			ID:      msg.ID,
			Error: &MCPError{
				Code:    -32603,
				Message: fmt.Sprintf("Search failed: %v", err),
			},
		}
	}

	return &MCPMessage{
		JSONRPC: "2.0",
		ID:      msg.ID,
		Result: map[string]interface{}{
			"content": []map[string]interface{}{
				{
					"type": "text",
					"text": s.formatSearchResults(results),
				},
			},
		},
	}
}

func (s *WebSearchServer) performWebSearch(query string, maxResults int) (*SearchResponse, error) {
	// Use DuckDuckGo for web search
	searchURL := fmt.Sprintf("https://html.duckduckgo.com/html/?q=%s", url.QueryEscape(query))

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	req, err := http.NewRequest("GET", searchURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers to mimic a browser
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Accept-Encoding", "gzip, deflate")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to perform search: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search request failed with status: %d", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML: %w", err)
	}

	var results []SearchResult
	rank := 1

	doc.Find(".result").Each(func(i int, sel *goquery.Selection) {
		if rank > maxResults {
			return
		}

		titleLink := sel.Find(".result__title a")
		title := strings.TrimSpace(titleLink.Text())
		href, exists := titleLink.Attr("href")

		if !exists || title == "" {
			return
		}

		// Clean up the URL (DuckDuckGo uses redirect URLs)
		if strings.HasPrefix(href, "//duckduckgo.com/l/?uddg=") {
			return // Skip these redirect URLs
		}

		description := strings.TrimSpace(sel.Find(".result__snippet").Text())

		results = append(results, SearchResult{
			Title:       title,
			URL:         href,
			Description: description,
			Rank:        rank,
		})
		rank++
	})

	return &SearchResponse{
		Query:   query,
		Results: results,
		Count:   len(results),
	}, nil
}

func (s *WebSearchServer) formatSearchResults(response *SearchResponse) string {
	if response.Count == 0 {
		return fmt.Sprintf("No results found for query: %s", response.Query)
	}

	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("Search results for: %s\n", response.Query))
	builder.WriteString(fmt.Sprintf("Found %d results:\n\n", response.Count))

	for _, result := range response.Results {
		builder.WriteString(fmt.Sprintf("%d. %s\n", result.Rank, result.Title))
		builder.WriteString(fmt.Sprintf("   URL: %s\n", result.URL))
		if result.Description != "" {
			builder.WriteString(fmt.Sprintf("   Description: %s\n", result.Description))
		}
		builder.WriteString("\n")
	}

	return builder.String()
}

func (s *WebSearchServer) handlePing(msg MCPMessage) *MCPMessage {
	return &MCPMessage{
		JSONRPC: "2.0",
		ID:      msg.ID,
		Result:  "pong",
	}
}

func (s *WebSearchServer) handleGetStats(msg MCPMessage) *MCPMessage {
	return &MCPMessage{
		JSONRPC: "2.0",
		ID:      msg.ID,
		Result:  s.stats.GetStats(),
	}
}

// Run stdio mode - communicate via standard input/output
func (s *WebSearchServer) runStdio() error {
	s.logger.Println("Starting MCP server in stdio mode")

	reader := bufio.NewReader(os.Stdin)
	encoder := json.NewEncoder(os.Stdout)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				s.logger.Println("Received EOF, shutting down")
				return nil
			}
			return fmt.Errorf("error reading from stdin: %w", err)
		}

		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		var msg MCPMessage
		if err := json.Unmarshal([]byte(line), &msg); err != nil {
			s.logger.Printf("Error unmarshaling message: %v", err)
			s.stats.IncrementErrors()
			continue
		}

		s.stats.IncrementRequests()
		response := s.handleMessage(msg)

		if response != nil {
			if err := encoder.Encode(response); err != nil {
				s.logger.Printf("Error encoding response: %v", err)
				s.stats.IncrementErrors()
				return fmt.Errorf("error writing to stdout: %w", err)
			}
		}
	}
}

// Run HTTP mode (for backward compatibility and testing)
func (s *WebSearchServer) runHTTP(port string) error {
	s.logger.Printf("Starting MCP server in HTTP mode on port %s", port)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		versionInfo := getVersionInfo()
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":    "healthy",
			"service":   "websearch-mcp",
			"version":   versionInfo.Version,
			"timestamp": time.Now().UTC().Format(time.RFC3339),
			"build_info": map[string]interface{}{
				"version":    versionInfo.Version,
				"build_time": versionInfo.BuildTime,
				"git_commit": versionInfo.GitCommit,
				"go_version": versionInfo.GoVersion,
			},
		})
	})

	http.HandleFunc("/stats", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(s.stats.GetStats())
	})

	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(getVersionInfo())
	})

	srv := &http.Server{
		Addr:         ":" + port,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Graceful shutdown
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c

		s.logger.Println("Shutting down server...")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			s.logger.Printf("Server shutdown error: %v", err)
		}
	}()

	s.logger.Printf("Health endpoint: http://localhost:%s/health", port)
	s.logger.Printf("Stats endpoint: http://localhost:%s/stats", port)
	s.logger.Printf("Version endpoint: http://localhost:%s/version", port)

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("server failed to start: %w", err)
	}

	s.logger.Println("Server stopped")
	return nil
}

func main() {
	server := NewWebSearchServer()

	// Check if we should run in HTTP mode (for backward compatibility)
	// If PORT environment variable is set or --http flag is provided, run in HTTP mode
	// Otherwise, run in stdio mode (default for MCP)
	mode := os.Getenv("MCP_MODE")
	port := os.Getenv("PORT")

	// Check for command line arguments
	for i, arg := range os.Args {
		if arg == "--http" {
			mode = "http"
			if i+1 < len(os.Args) && !strings.HasPrefix(os.Args[i+1], "--") {
				port = os.Args[i+1]
			}
		} else if arg == "--stdio" {
			mode = "stdio"
		} else if arg == "--version" || arg == "-v" {
			versionInfo := getVersionInfo()
			fmt.Printf("websearch-mcp version %s\n", versionInfo.Version)
			fmt.Printf("  Build Time: %s\n", versionInfo.BuildTime)
			fmt.Printf("  Git Commit: %s\n", versionInfo.GitCommit)
			fmt.Printf("  Go Version: %s\n", versionInfo.GoVersion)
			fmt.Printf("  OS/Arch: %s/%s\n", versionInfo.OS, versionInfo.Arch)
			return
		} else if arg == "--help" || arg == "-h" {
			fmt.Println("WebSearch MCP Server")
			fmt.Println()
			fmt.Println("Usage:")
			fmt.Println("  websearch-mcp [options]")
			fmt.Println()
			fmt.Println("Options:")
			fmt.Println("  --stdio           Run in stdio mode (default)")
			fmt.Println("  --http [port]     Run in HTTP mode on specified port (default: 8080)")
			fmt.Println("  --version, -v     Show version information")
			fmt.Println("  --help, -h        Show this help message")
			fmt.Println()
			fmt.Println("Environment Variables:")
			fmt.Println("  MCP_MODE          Set to 'http' or 'stdio' (default: stdio)")
			fmt.Println("  PORT              Port for HTTP mode (default: 8080)")
			return
		}
	}

	// Default to stdio mode if not specified
	if mode == "" {
		mode = "stdio"
	}

	var err error
	if mode == "http" {
		if port == "" {
			port = "8080"
		}
		err = server.runHTTP(port)
	} else {
		err = server.runStdio()
	}

	if err != nil {
		server.logger.Fatalf("Server error: %v", err)
	}
}
