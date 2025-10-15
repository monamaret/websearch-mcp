# Multi-Platform Support Verification Checklist

Use this checklist to verify that all multi-platform support features are working correctly.

## 📋 Pre-Push Verification

### Local Build Tests

#### Makefile
- [ ] `make help` shows updated documentation references
- [ ] `make build` works on current platform
- [ ] `make build-all-platforms` creates 6 binaries in `dist/`:
  - [ ] `websearch-mcp-darwin-amd64`
  - [ ] `websearch-mcp-darwin-arm64`
  - [ ] `websearch-mcp-windows-amd64.exe`
  - [ ] `websearch-mcp-windows-arm64.exe`
  - [ ] `websearch-mcp-linux-amd64`
  - [ ] `websearch-mcp-linux-arm64`
- [ ] Archives are created:
  - [ ] `.tar.gz` files for Unix platforms
  - [ ] `.zip` files for Windows platforms
- [ ] `dist/checksums.txt` is generated
- [ ] Checksums are valid: `cd dist && shasum -a 256 -c checksums.txt`

#### Manual Builds
- [ ] Can build for macOS Intel: `GOOS=darwin GOARCH=amd64 go build .`
- [ ] Can build for macOS ARM: `GOOS=darwin GOARCH=arm64 go build .`
- [ ] Can build for Windows Intel: `GOOS=windows GOARCH=amd64 go build .`
- [ ] Can build for Windows ARM: `GOOS=windows GOARCH=arm64 go build .`
- [ ] Can build for Linux Intel: `GOOS=linux GOARCH=amd64 go build .`
- [ ] Can build for Linux ARM: `GOOS=linux GOARCH=arm64 go build .`

### Documentation Review

- [ ] All new docs are in `docs/` folder:
  - [ ] `docs/BUILDING.md`
  - [ ] `docs/PLATFORM_SUPPORT.md`
  - [ ] `docs/RELEASE_GUIDE.md`
  - [ ] `docs/QUICK_START.md`
  - [ ] `docs/MULTI_PLATFORM_UPDATE.md`
  - [ ] `docs/README.md`
- [ ] Updated docs are correct:
  - [ ] `README.md`
  - [ ] `docs/WORKFLOWS.md`
- [ ] All links in documentation work
- [ ] No broken references
- [ ] Code examples are correct
- [ ] Platform-specific instructions are clear

### Code Quality

- [ ] `go fmt ./...` shows no changes needed
- [ ] `go vet ./...` shows no issues
- [ ] `go test ./...` passes (if tests exist)
- [ ] No syntax errors in workflows:
  - [ ] `.github/workflows/build.yml` is valid YAML
  - [ ] `.github/workflows/release.yml` is valid YAML

## 🚀 GitHub Actions Verification

### Test Branch Push

1. Create a test branch:
   ```bash
   git checkout -b test-multiplatform
   git push origin test-multiplatform
   ```

2. Verify in GitHub Actions:
   - [ ] Build workflow triggers
   - [ ] Test job completes successfully
   - [ ] All 6 build jobs start:
     - [ ] darwin-amd64
     - [ ] darwin-arm64
     - [ ] windows-amd64
     - [ ] windows-arm64
     - [ ] linux-amd64
     - [ ] linux-arm64
   - [ ] All builds complete successfully
   - [ ] Artifacts are uploaded for each platform
   - [ ] No snapshot release created (not on main branch)

### Main Branch Push

After test branch is verified:

1. Merge to main:
   ```bash
   git checkout main
   git merge test-multiplatform
   git push origin main
   ```

2. Verify in GitHub Actions:
   - [ ] Build workflow triggers
   - [ ] All 6 builds complete
   - [ ] Snapshot release is created
   - [ ] Snapshot release contains:
     - [ ] 6 archive files
     - [ ] 1 checksums.txt
     - [ ] Formatted release notes
     - [ ] Platform-specific instructions

### Release Tag Push

1. Create a test release tag:
   ```bash
   git tag v0.0.1-test
   git push origin v0.0.1-test
   ```

2. Verify in GitHub Actions:
   - [ ] Release workflow triggers
   - [ ] All 6 build jobs complete
   - [ ] Release is created on GitHub
   - [ ] Release contains all 6 archives:
     - [ ] `websearch-mcp-v0.0.1-test-darwin-amd64.tar.gz`
     - [ ] `websearch-mcp-v0.0.1-test-darwin-arm64.tar.gz`
     - [ ] `websearch-mcp-v0.0.1-test-windows-amd64.zip`
     - [ ] `websearch-mcp-v0.0.1-test-windows-arm64.zip`
     - [ ] `websearch-mcp-v0.0.1-test-linux-amd64.tar.gz`
     - [ ] `websearch-mcp-v0.0.1-test-linux-arm64.tar.gz`
   - [ ] Release contains `checksums.txt`
   - [ ] Release notes are formatted correctly
   - [ ] Platform-specific instructions are included

3. Clean up test tag:
   ```bash
   git push --delete origin v0.0.1-test
   git tag -d v0.0.1-test
   # Delete release on GitHub web interface
   ```

## 🧪 Binary Testing

### Download and Test

For each platform you have access to:

#### macOS Intel
- [ ] Download `websearch-mcp-*-darwin-amd64.tar.gz`
- [ ] Extract: `tar -xzf websearch-mcp-*-darwin-amd64.tar.gz`
- [ ] Verify checksum matches
- [ ] Make executable: `chmod +x websearch-mcp-darwin-amd64`
- [ ] Run: `./websearch-mcp-darwin-amd64`
- [ ] Test health: `curl http://localhost:8080/health`
- [ ] Verify version info is embedded

#### macOS Apple Silicon
- [ ] Download `websearch-mcp-*-darwin-arm64.tar.gz`
- [ ] Extract: `tar -xzf websearch-mcp-*-darwin-arm64.tar.gz`
- [ ] Verify checksum matches
- [ ] Make executable: `chmod +x websearch-mcp-darwin-arm64`
- [ ] Run: `./websearch-mcp-darwin-arm64`
- [ ] Test health: `curl http://localhost:8080/health`
- [ ] Verify version info is embedded

#### Windows Intel
- [ ] Download `websearch-mcp-*-windows-amd64.zip`
- [ ] Extract the zip file
- [ ] Verify checksum matches
- [ ] Run: `.\websearch-mcp-windows-amd64.exe`
- [ ] Test health: `Invoke-WebRequest http://localhost:8080/health`
- [ ] Verify version info is embedded

#### Windows ARM
- [ ] Download `websearch-mcp-*-windows-arm64.zip`
- [ ] Extract the zip file
- [ ] Verify checksum matches
- [ ] Run: `.\websearch-mcp-windows-arm64.exe`
- [ ] Test health: `Invoke-WebRequest http://localhost:8080/health`
- [ ] Verify version info is embedded

#### Linux Intel
- [ ] Download `websearch-mcp-*-linux-amd64.tar.gz`
- [ ] Extract: `tar -xzf websearch-mcp-*-linux-amd64.tar.gz`
- [ ] Verify checksum matches
- [ ] Make executable: `chmod +x websearch-mcp-linux-amd64`
- [ ] Run: `./websearch-mcp-linux-amd64`
- [ ] Test health: `curl http://localhost:8080/health`
- [ ] Verify version info is embedded

#### Linux ARM
- [ ] Download `websearch-mcp-*-linux-arm64.tar.gz`
- [ ] Extract: `tar -xzf websearch-mcp-*-linux-arm64.tar.gz`
- [ ] Verify checksum matches
- [ ] Make executable: `chmod +x websearch-mcp-linux-arm64`
- [ ] Run: `./websearch-mcp-linux-arm64`
- [ ] Test health: `curl http://localhost:8080/health`
- [ ] Verify version info is embedded

### Checksum Verification

- [ ] Download `checksums.txt`
- [ ] Verify each archive:
  ```bash
  # macOS/Linux
  sha256sum -c checksums.txt
  
  # Windows (PowerShell)
  Get-FileHash websearch-mcp-*.zip -Algorithm SHA256
  ```
- [ ] All checksums match

## 📚 Documentation Verification

### Quick Start Guide
- [ ] Instructions are clear for each platform
- [ ] Download links are correct format
- [ ] Commands work as documented
- [ ] Troubleshooting tips are helpful

### Platform Support
- [ ] All platforms listed in table
- [ ] System requirements are accurate
- [ ] Installation steps work
- [ ] Service setup instructions work (test on one platform)

### Building Guide
- [ ] Prerequisites are complete
- [ ] Build commands work
- [ ] Cross-compilation examples work
- [ ] Troubleshooting covers common issues

### Release Guide
- [ ] Pre-release checklist is complete
- [ ] Release process is clear
- [ ] Commands are correct
- [ ] Rollback procedures make sense

### Workflows Documentation
- [ ] Matrix build explanation is clear
- [ ] Platform details are accurate
- [ ] Troubleshooting is helpful
- [ ] Links to resources work

## 🔗 Integration Testing

### Tabnine Integration
- [ ] Create `.mcp_servers` file with platform-specific binary
- [ ] Start server
- [ ] Test with Tabnine Agent
- [ ] Verify search works
- [ ] Check logs for errors

### MCP Protocol
- [ ] Initialize request works
- [ ] Tools list request works
- [ ] Web search tool works
- [ ] Error handling works

## 📊 Metrics to Monitor

After release, monitor:

### GitHub Actions
- [ ] Build success rate > 95%
- [ ] Average build time < 10 minutes
- [ ] No persistent failures

### Downloads
- [ ] Which platforms are most popular
- [ ] Download completion rate
- [ ] Checksum verification usage

### Issues
- [ ] Platform-specific bug reports
- [ ] Installation issues
- [ ] Documentation gaps

## ✅ Final Checks

Before announcing the release:

- [ ] All builds passing on main branch
- [ ] Latest release has all 6 archives
- [ ] Documentation is up to date
- [ ] CHANGES.md includes latest updates
- [ ] README.md has correct version
- [ ] All links tested
- [ ] Known issues documented

## 🚨 Rollback Plan

If critical issues are found:

1. Immediate:
   - [ ] Mark release as pre-release on GitHub
   - [ ] Add warning to release notes
   - [ ] Document the issue

2. Short-term:
   - [ ] Create hotfix branch
   - [ ] Fix the issue
   - [ ] Create new patch release

3. Long-term:
   - [ ] Review what went wrong
   - [ ] Update testing procedures
   - [ ] Improve documentation

## 📝 Notes

Use this section to track any issues or observations during verification:

```
Date: ___________
Tester: ___________

Issues Found:
- 

Observations:
- 

Recommended Actions:
- 

```

## ✨ Success Criteria

All items checked = Ready for production! 🎉

- Build system: All Makefile targets work
- CI/CD: All workflows pass
- Binaries: All 6 platforms build and run
- Documentation: All docs are complete and accurate
- Integration: Tabnine integration works
- Release: Test release completes successfully

---

**Last Updated**: 2024-01-15  
**Version**: 1.0.0  
**Status**: Ready for verification
