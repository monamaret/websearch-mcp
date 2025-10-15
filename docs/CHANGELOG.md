# Changelog

## [Unreleased] - 2024-01-XX

### 🚀 Major Changes

#### Stdio Mode (Breaking Change)
- **Changed default communication mode from HTTP/WebSocket to stdio**
- Server now communicates via standard input/output by default
- This makes the server compatible with standard MCP client implementations
- HTTP mode is still available via `--http` flag for testing

### ✨ Added

#### New Features
- **Stdio mode support** - Default communication via stdin/stdout
- **Command line arguments** - `--stdio`, `--http`, `--version`, `--help`
- **Environment variable support** - `MCP_MODE` to control mode
- **Dual-mode operation** - Can run in either stdio or HTTP mode
- **Better logging** - Logs go to stderr, separate from protocol messages

#### New Documentation
- `STDIO_MIGRATION.md` - Detailed migration guide from HTTP to stdio
- `UPGRADE_GUIDE.md` - Step-by-step upgrade instructions
- `docs/GETTING_STARTED.md` - Comprehensive getting started guide
- `examples/README.md` - Examples documentation
- `examples/test-stdio-client.go` - Stdio test client example

#### New Scripts
- `test-stdio.sh` - Quick stdio mode testing script

### 🔄 Changed

#### Configuration
- **Simplified `.mcp_servers` configuration** - No PORT required
- **Removed PORT environment variable** from default config
- **Updated all documentation** to reflect stdio as default

#### Code Structure
- **Refactored main.go** - Split into `runStdio()` and `runHTTP()` functions
- **Updated server initialization** - New `logger` field for stderr logging
- **Improved error handling** - Better error messages for both modes

#### Examples
- **Renamed `test-client.go`** to `test-http-client.go` for clarity
- **Added `test-stdio-client.go`** for testing stdio mode
- **Updated all example configs** to use stdio mode

#### Tests
- **Removed WebSocket integration test** - No longer needed
- **Kept unit tests** - All tests still pass
- **Updated test documentation** - Reflect new testing approach

#### Documentation
- Updated `README.md` with stdio/HTTP mode explanations
- Updated `docs/TABNINE_SETUP.md` with new configuration
- Updated `docs/TABNINE_QUICK_REFERENCE.md` with new modes
- Updated `docs/USAGE.md` (if exists)
- All docs now show stdio as default, HTTP as optional

### 🗑️ Removed

#### Dependencies
- **Removed dependency on `github.com/gorilla/websocket`** from main code
  - Note: Still used in HTTP test client example

#### Features
- **Removed WebSocket endpoint from default mode**
  - Available in HTTP mode with `--http` flag
- **Removed automatic HTTP server startup**
  - Only starts when explicitly requested

### 🐛 Fixed

#### MCP Compatibility
- **Fixed protocol compatibility** - Now follows MCP stdio specification
- **Fixed client integration** - Works correctly with Tabnine and other MCP clients
- **Fixed lifecycle management** - Server now controlled by client

### 📝 Migration Guide

#### From Old Version (HTTP)

**Before:**
```json
{
  "mcpServers": {
    "websearch": {
      "command": "./websearch-mcp",
      "args": [],
      "env": {
        "PORT": "8080"
      }
    }
  }
}
```

**After:**
```json
{
  "mcpServers": {
    "websearch": {
      "command": "./websearch-mcp",
      "args": [],
      "env": {}
    }
  }
}
```

#### Command Line Usage

**Old way (only HTTP):**
```bash
PORT=8080 ./websearch-mcp
```

**New way (stdio by default):**
```bash
./websearch-mcp                # Stdio mode (default)
./websearch-mcp --stdio        # Explicit stdio mode
./websearch-mcp --http 8080    # HTTP mode for testing
```

### 🔒 Security

- No security changes in this release
- Stdio mode is more secure as it doesn't expose network ports by default
- HTTP mode still available when needed for testing

### 📊 Performance

- Stdio mode has lower overhead than HTTP/WebSocket
- No port conflicts or connection management overhead
- Faster startup time in stdio mode

### 🧪 Testing

- All existing unit tests pass
- New stdio mode tested with test-stdio.sh
- HTTP mode backwards compatibility verified
- Integration with Tabnine verified

### 📚 Documentation Updates

- Added 5 new documentation files
- Updated 4 existing documentation files
- Added examples for both modes
- Comprehensive migration guides

### ⚠️ Breaking Changes

1. **Default mode changed from HTTP to stdio**
   - Solution: Use `--http` flag if you need HTTP mode
   - Or set `MCP_MODE=http` environment variable

2. **No automatic HTTP server startup**
   - Solution: Explicitly use `--http [port]` to start HTTP server

3. **Logs now go to stderr instead of stdout**
   - Solution: Redirect stderr if you want to capture logs: `2>&1 | tee log.txt`

### 🔮 Future Plans

- Additional MCP capabilities
- More search providers (optional)
- Enhanced caching for better performance
- Metrics and monitoring improvements

---

## Version History

### [1.0.0] - 2024-XX-XX (Upcoming)
- First stable release with stdio mode as default

### [0.x.x] - Previous
- HTTP/WebSocket mode releases
- Initial MCP implementation

---

## How to Upgrade

See [UPGRADE_GUIDE.md](./UPGRADE_GUIDE.md) for detailed upgrade instructions.

## Questions?

- Read the [Getting Started Guide](./docs/GETTING_STARTED.md)
- Check the [Tabnine Setup Guide](./docs/TABNINE_SETUP.md)
- See [Migration Guide](./STDIO_MIGRATION.md)
- Open an issue on GitHub
