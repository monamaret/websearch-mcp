# Migration Checklist - HTTP to Stdio Mode

Use this checklist to ensure a smooth migration from HTTP mode to stdio mode.

## Pre-Migration

- [ ] **Backup your current configuration**
  ```bash
  cp .mcp_servers .mcp_servers.backup
  ```

- [ ] **Note your current setup**
  - Server binary location: `__________`
  - Port number (if applicable): `__________`
  - Any custom environment variables: `__________`

- [ ] **Read the documentation**
  - [ ] [STDIO_MIGRATION.md](./STDIO_MIGRATION.md)
  - [ ] [UPGRADE_GUIDE.md](./UPGRADE_GUIDE.md)
  - [ ] [docs/GETTING_STARTED.md](./docs/GETTING_STARTED.md)

## Migration Steps

### 1. Update Code

- [ ] **Pull latest changes** (if using git)
  ```bash
  git pull origin main
  ```

- [ ] **Rebuild the binary**
  ```bash
  make build
  # or
  go build -o websearch-mcp .
  ```

- [ ] **Verify build**
  ```bash
  ./websearch-mcp --version
  ```

### 2. Update Configuration

- [ ] **Edit `.mcp_servers` file**
  
  **Remove this:**
  ```json
  "env": {
    "PORT": "8080"
  }
  ```
  
  **Replace with:**
  ```json
  "env": {}
  ```

- [ ] **Verify configuration syntax**
  ```bash
  cat .mcp_servers | python -m json.tool
  ```

### 3. Test the Server

- [ ] **Test stdio mode manually**
  ```bash
  ./test-stdio.sh
  ```
  
  **Expected:** All 3 tests should pass ✅

- [ ] **Test with echo command**
  ```bash
  echo '{"jsonrpc":"2.0","id":1,"method":"ping"}' | ./websearch-mcp
  ```
  
  **Expected:** `{"jsonrpc":"2.0","id":1,"result":"pong"}`

- [ ] **Test HTTP mode (optional)**
  ```bash
  ./websearch-mcp --http 8080 &
  sleep 2
  curl http://localhost:8080/health
  pkill websearch-mcp
  ```
  
  **Expected:** Health check returns "healthy"

### 4. Integration Testing

- [ ] **Restart Tabnine**
  - Close and reopen your IDE
  - Or restart Tabnine extension

- [ ] **Test with Tabnine**
  
  Ask Tabnine:
  ```
  "Can you see the websearch MCP server?"
  ```
  
  **Expected:** Tabnine confirms it can see the server ✅

- [ ] **Perform a search**
  
  Ask Tabnine:
  ```
  "Search for 'Go programming best practices' and give me the top 3 results"
  ```
  
  **Expected:** Tabnine returns search results ✅

### 5. Verify Everything Works

- [ ] **Check server logs** (in Tabnine logs)
  - No error messages ✅
  - Server starts and stops cleanly ✅

- [ ] **Test multiple searches**
  - [ ] First search works
  - [ ] Second search works
  - [ ] Third search works

- [ ] **Test edge cases**
  - [ ] Empty query (should return error)
  - [ ] Very long query
  - [ ] Special characters in query

## Post-Migration

### Cleanup

- [ ] **Remove old startup scripts** (if any)
  ```bash
  # Only if you had custom startup scripts for HTTP mode
  rm old-start-script.sh  # Example
  ```

- [ ] **Update team documentation** (if applicable)
  - Share new `.mcp_servers` configuration
  - Update team wiki/docs
  - Notify team members

- [ ] **Remove backup** (after confirming everything works)
  ```bash
  # Wait a few days before doing this
  rm .mcp_servers.backup
  ```

### Validation

- [ ] **Verify performance**
  - Searches return results quickly ✅
  - No timeouts or errors ✅

- [ ] **Check resource usage**
  - Server doesn't consume excessive memory ✅
  - No zombie processes ✅

- [ ] **Test over time**
  - Works after IDE restart ✅
  - Works after system restart ✅
  - Works after a few days of use ✅

## Troubleshooting Checklist

If something doesn't work:

- [ ] **Check binary permissions**
  ```bash
  ls -l websearch-mcp
  # Should show: -rwxr-xr-x
  chmod +x websearch-mcp  # If needed
  ```

- [ ] **Check binary path in config**
  ```bash
  cat .mcp_servers | grep command
  # Verify path is correct
  ```

- [ ] **Test binary directly**
  ```bash
  ./websearch-mcp --version
  ./websearch-mcp --help
  echo '{"jsonrpc":"2.0","id":1,"method":"ping"}' | ./websearch-mcp
  ```

- [ ] **Check Tabnine logs**
  - Look for MCP-related errors
  - Note any error messages

- [ ] **Try HTTP mode** (temporary troubleshooting)
  ```json
  {
    "mcpServers": {
      "websearch": {
        "command": "./websearch-mcp",
        "args": ["--http", "8080"],
        "env": {}
      }
    }
  }
  ```

- [ ] **Restart everything**
  - Restart IDE
  - Restart Tabnine
  - Restart computer (if needed)

## Rollback Plan

If you need to rollback:

### Option 1: Use HTTP Mode

- [ ] **Update `.mcp_servers` to use HTTP mode**
  ```json
  {
    "mcpServers": {
      "websearch": {
        "command": "./websearch-mcp",
        "args": ["--http", "8080"],
        "env": {}
      }
    }
  }
  ```

### Option 2: Restore Backup

- [ ] **Restore old configuration**
  ```bash
  cp .mcp_servers.backup .mcp_servers
  ```

- [ ] **Use old binary** (if you kept it)
  ```bash
  cp websearch-mcp.old websearch-mcp
  ```

## Success Criteria

Migration is successful when:

- ✅ Server starts automatically when Tabnine needs it
- ✅ Searches return results correctly
- ✅ No manual server management needed
- ✅ No errors in Tabnine logs
- ✅ Multiple searches work without issues
- ✅ Server stops cleanly when Tabnine closes

## Additional Resources

- **Documentation**
  - [STDIO_MIGRATION.md](./STDIO_MIGRATION.md) - Detailed migration guide
  - [UPGRADE_GUIDE.md](./UPGRADE_GUIDE.md) - Step-by-step upgrade
  - [docs/GETTING_STARTED.md](./docs/GETTING_STARTED.md) - Getting started
  - [SUMMARY.md](./SUMMARY.md) - Summary of changes

- **Examples**
  - [examples/README.md](./examples/README.md) - Example configurations
  - [examples/test-stdio-client.go](./examples/test-stdio-client.go) - Test client

- **Testing**
  - [test-stdio.sh](./test-stdio.sh) - Quick test script

## Notes

- **Time Required:** 10-15 minutes
- **Difficulty:** Easy
- **Risk Level:** Low (can rollback easily)
- **Team Impact:** Minimal (update config file)

## Questions?

- [ ] Checked documentation above
- [ ] Tried troubleshooting steps
- [ ] Searched existing issues
- [ ] Ready to open new issue if needed

---

## Migration Date

**Started:** `__________`  
**Completed:** `__________`  
**Tested By:** `__________`  
**Status:** ⬜ Not Started | ⬜ In Progress | ⬜ Complete | ⬜ Rolled Back

## Sign-Off

- [ ] Configuration updated
- [ ] Tests passed
- [ ] Integration verified
- [ ] Team notified (if applicable)
- [ ] Documentation updated (if applicable)

**Migrated by:** `__________`  
**Date:** `__________`

---

**Congratulations on completing the migration! 🎉**
