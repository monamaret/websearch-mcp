.PHONY: build run clean test docker-build docker-run deps

# Version information
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
BUILD_TIME ?= $(shell date -u +%Y-%m-%dT%H:%M:%SZ)
GIT_COMMIT ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")

# Build flags
LDFLAGS = -w -s -X main.version=$(VERSION) -X main.buildTime=$(BUILD_TIME) -X main.gitCommit=$(GIT_COMMIT)

# Default target
all: build

# Build the server
build:
	CGO_ENABLED=0 go build -ldflags="$(LDFLAGS)" -o websearch-mcp .

# Build optimized release binary
build-release:
	CGO_ENABLED=0 go build -ldflags="$(LDFLAGS)" -a -installsuffix cgo -o websearch-mcp .

# Build for multiple platforms
build-all:
	@echo "Building for multiple platforms..."
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="$(LDFLAGS)" -o websearch-mcp-linux-amd64 .
	GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -ldflags="$(LDFLAGS)" -o websearch-mcp-linux-arm64 .
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="$(LDFLAGS)" -o websearch-mcp-darwin-amd64 .
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build -ldflags="$(LDFLAGS)" -o websearch-mcp-darwin-arm64 .
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="$(LDFLAGS)" -o websearch-mcp-windows-amd64.exe .

# Run the server
run:
	go run main.go

# Clean build artifacts
clean:
	rm -f websearch-mcp websearch-mcp-*
	go clean

# Download dependencies
deps:
	go mod download
	go mod tidy

# Test the server (builds and runs test client)
test: build
	@echo "Starting server in background..."
	@./websearch-mcp &
	@SERVER_PID=$$!; \
	sleep 2; \
	echo "Running test client..."; \
	cd examples && go run test-client.go; \
	echo "Stopping server..."; \
	kill $$SERVER_PID

# Build Docker image
docker-build:
	docker build \
		--build-arg VERSION=$(VERSION) \
		--build-arg BUILD_TIME=$(BUILD_TIME) \
		--build-arg GIT_COMMIT=$(GIT_COMMIT) \
		-t websearch-mcp:$(VERSION) \
		-t websearch-mcp:latest \
		.

# Run in Docker
docker-run: docker-build
	docker run -p 8080:8080 websearch-mcp

# Development mode with hot reload (requires air)
dev:
	@if command -v air > /dev/null; then \
		air; \
	else \
		echo "Air not installed. Install with: go install github.com/cosmtrek/air@latest"; \
		echo "Running without hot reload..."; \
		go run main.go; \
	fi

# Format code
fmt:
	go fmt ./...

# Vet code
vet:
	go vet ./...

# Run linter (requires golangci-lint)
lint:
	@if command -v golangci-lint > /dev/null; then \
		golangci-lint run; \
	else \
		echo "golangci-lint not installed. Install with: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
	fi

# Install development tools
install-tools:
	go install github.com/cosmtrek/air@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Tabnine MCP targets
setup-tabnine:
	@echo "Setting up Tabnine MCP configuration..."
	@./setup-tabnine.sh

validate-tabnine:
	@echo "Validating Tabnine MCP setup..."
	@if [ ! -f ".mcp_servers" ]; then \
		echo "❌ .mcp_servers file not found. Run 'make setup-tabnine' first."; \
		exit 1; \
	fi
	@if [ ! -f "./websearch-mcp" ]; then \
		echo "❌ websearch-mcp binary not found. Run 'make build' first."; \
		exit 1; \
	fi
	@echo "✅ Tabnine MCP configuration is valid"

tabnine-demo: build validate-tabnine
	@echo "Starting Tabnine MCP demo..."
	@echo "1. Starting WebSearch MCP server..."
	@./websearch-mcp &
	@SERVER_PID=$$!; \
	sleep 2; \
	echo "2. Testing server health..."; \
	curl -s http://localhost:8080/health | python3 -m json.tool; \
	echo "3. Server is ready for Tabnine integration!"; \
	echo "4. Ask your Tabnine Agent: 'Can you see the websearch MCP server?'"; \
	echo "5. Press Ctrl+C to stop the server"; \
	trap "kill $$SERVER_PID 2>/dev/null || true; echo 'Server stopped.'" INT; \
	wait $$SERVER_PID

# Show version information
version:
	@echo "Version: $(VERSION)"
	@echo "Build Time: $(BUILD_TIME)"
	@echo "Git Commit: $(GIT_COMMIT)"

# Show help
help:
	@echo "Available targets:"
	@echo "  build            - Build the server binary"
	@echo "  build-release    - Build optimized release binary"
	@echo "  build-all        - Build for multiple platforms"
	@echo "  run              - Run the server directly"
	@echo "  clean            - Clean build artifacts"
	@echo "  deps             - Download and tidy dependencies"
	@echo "  test             - Build and test with test client"
	@echo "  docker-build     - Build Docker image with version info"
	@echo "  docker-run       - Run in Docker container"
	@echo "  dev              - Run in development mode with hot reload"
	@echo "  fmt              - Format Go code"
	@echo "  vet              - Run go vet"
	@echo "  lint             - Run golangci-lint"
	@echo "  install-tools    - Install development tools"
	@echo "  version          - Show version information"
	@echo ""
	@echo "Tabnine MCP targets:"
	@echo "  setup-tabnine    - Interactive setup for Tabnine MCP integration"
	@echo "  validate-tabnine - Validate Tabnine MCP configuration"
	@echo "  tabnine-demo     - Start server and show integration demo"
	@echo ""
	@echo "  help             - Show this help message"
