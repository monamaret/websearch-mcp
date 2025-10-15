# GitHub Workflows Documentation

This repository includes GitHub Actions workflows for automated building, testing, and releasing of the WebSearch MCP Server with support for multiple platforms and architectures.

## 📋 Available Workflows

### 1. Build and Test (`.github/workflows/build.yml`)

**Triggers:**
- Push to `main` or `develop` branches
- Pull requests to `main` branch

**Jobs:**
- **Test**: Runs unit tests on Ubuntu
- **Build**: Creates optimized binaries for all supported platforms using a build matrix

**Build Matrix:**
The workflow builds binaries for the following platforms:
- **macOS**
  - Intel (x86_64/amd64)
  - Apple Silicon (ARM64)
- **Windows**
  - Intel (x86_64/amd64)
  - ARM64
- **Linux**
  - Intel (x86_64/amd64)
  - ARM64

**Artifacts:**
Each platform produces:
- Binary: `websearch-mcp-{platform}-{arch}{ext}`
- Archive: `websearch-mcp-{version}-{platform}-{arch}.tar.gz` (Unix) or `.zip` (Windows)
- Checksums: `checksums.txt`

**Snapshot Releases:**
When pushing to `main`, the workflow automatically creates a snapshot release with:
- All platform binaries
- Combined checksums file
- Detailed installation instructions for each platform

### 2. Release (`.github/workflows/release.yml`)

**Triggers:**
- Git tags matching `v*` (e.g., `v1.0.0`, `v1.2.3`)

**Features:**
- Builds optimized binaries for all platforms with embedded version info
- Creates GitHub releases with comprehensive documentation
- Includes installation instructions for each platform
- Provides verification checksums

**Release Assets:**
Each release includes 6 binary archives:
- `websearch-mcp-{version}-darwin-amd64.tar.gz` (macOS Intel)
- `websearch-mcp-{version}-darwin-arm64.tar.gz` (macOS Apple Silicon)
- `websearch-mcp-{version}-windows-amd64.zip` (Windows Intel)
- `websearch-mcp-{version}-windows-arm64.zip` (Windows ARM)
- `websearch-mcp-{version}-linux-amd64.tar.gz` (Linux Intel)
- `websearch-mcp-{version}-linux-arm64.tar.gz` (Linux ARM)
- `checksums.txt` (SHA256 checksums for verification)

## 🔧 Setup Requirements

**Repository Settings:**
1. Enable GitHub Actions in repository settings
2. Configure branch protection for `main` branch:
   - Require status checks before merging
   - Include workflow checks: `test`, `build`

**Required Permissions:**
- `contents: write` (for creating releases and uploading artifacts)

## 🚀 Release Process

### Creating a Release

1. **Prepare for release:**
   ```bash
   git checkout main
   git pull origin main
   ```

2. **Create and push a tag:**
   ```bash
   git tag v1.0.0
   git push origin v1.0.0
   ```

3. **Automated process:**
   - Release workflow triggers automatically
   - Binaries are built in parallel for all platforms
   - Archives are created (tar.gz for Unix, zip for Windows)
   - Checksums are generated
   - GitHub release is created with all assets attached

### Version Information

All builds include version information injected at build time:

```go
// These variables are set via ldflags during build
var (
    version   = "dev"        // Git tag or "dev"
    buildTime = "unknown"    // RFC3339 timestamp
    gitCommit = "unknown"    // Short git commit hash
)
```

Access version info via:
- Health endpoint: `GET http://localhost:8080/health`
- Version endpoint: `GET http://localhost:8080/version`
- MCP initialize: Included in server info

## 🏗️ Platform-Specific Build Details

### macOS Builds
- Run on: `macos-latest` GitHub runners
- Cross-compilation: Both Intel and ARM builds run on the same runner
- Archive format: `.tar.gz`
- Binary naming: `websearch-mcp-darwin-{amd64|arm64}`

### Windows Builds
- Run on: `windows-latest` GitHub runners
- Cross-compilation: Both Intel and ARM builds run on the same runner
- Archive format: `.zip`
- Binary naming: `websearch-mcp-windows-{amd64|arm64}.exe`

### Linux Builds
- Run on: `ubuntu-latest` GitHub runners
- Cross-compilation: Both Intel and ARM builds run on the same runner
- Archive format: `.tar.gz`
- Binary naming: `websearch-mcp-linux-{amd64|arm64}`

## 🧪 Testing Workflows

### Local Testing

You can build for all platforms locally using the Makefile:

```bash
# Build for all platforms
make build-all-platforms

# Output will be in the dist/ directory
ls -lh dist/
```

Individual platform builds:
```bash
# macOS Intel
GOOS=darwin GOARCH=amd64 go build -o websearch-mcp-darwin-amd64 .

# macOS Apple Silicon
GOOS=darwin GOARCH=arm64 go build -o websearch-mcp-darwin-arm64 .

# Windows Intel
GOOS=windows GOARCH=amd64 go build -o websearch-mcp-windows-amd64.exe .

# Windows ARM
GOOS=windows GOARCH=arm64 go build -o websearch-mcp-windows-arm64.exe .

# Linux Intel
GOOS=linux GOARCH=amd64 go build -o websearch-mcp-linux-amd64 .

# Linux ARM
GOOS=linux GOARCH=arm64 go build -o websearch-mcp-linux-arm64 .
```

Run tests:
```bash
go test -v ./...
```

### Workflow Debugging

- View runs in the Actions tab
- Expand job steps to see detailed logs
- Matrix builds run in parallel for faster execution

**Common issues:**
- Build failures: Check Go version compatibility
- Test failures: Review test output in workflow logs
- Archive creation issues: Verify tar/zip tools availability
- Cross-compilation errors: Usually indicate platform-specific code issues

## 📊 Workflow Outputs

### Build Artifacts

Each platform build produces:
- **Binary**: Platform-specific executable
- **Archive**: Compressed archive with binary and README
- **Checksums**: SHA256 checksums for verification

### Release Archives

Contents of each archive:
- Binary executable (platform-specific)
- `README.md` documentation

Archive naming convention:
- Unix: `websearch-mcp-{version}-{platform}-{arch}.tar.gz`
- Windows: `websearch-mcp-{version}-windows-{arch}.zip`

## 🔄 Maintenance

### Regular Updates

1. **Go version updates:**
   ```yaml
   # Update in both build.yml and release.yml
   - name: Set up Go
     uses: actions/setup-go@v4
     with:
       go-version: '1.21'  # Update this version
   ```

2. **Runner updates:**
   - Monitor for new runner images
   - Test with updated runners before merging

3. **Dependency updates:**
   ```bash
   go get -u ./...
   go mod tidy
   ```

### Adding New Platforms

To add support for a new platform:

1. **Update build matrix in workflows:**
   ```yaml
   - os: ubuntu-latest
     platform: linux-riscv64
     goos: linux
     goarch: riscv64
     ext: ''
   ```

2. **Update Makefile:**
   ```makefile
   @echo "Building for Linux (RISC-V)..."
   GOOS=linux GOARCH=riscv64 CGO_ENABLED=0 go build -ldflags="$(LDFLAGS)" -a -installsuffix cgo -o dist/$(BINARY_NAME)-linux-riscv64 .
   ```

3. **Update documentation** (this file and README.md)

### Workflow Modifications

1. Test changes in a fork or feature branch
2. Use `workflow_dispatch` for manual testing
3. Monitor resource usage to stay within GitHub Actions limits
4. Document changes in this file

## 📚 Additional Resources

- [GitHub Actions Documentation](https://docs.github.com/en/actions)
- [Go Cross Compilation](https://go.dev/doc/install/source#environment)
- [Semantic Versioning](https://semver.org/)
- [GitHub Actions: Build Matrix](https://docs.github.com/en/actions/using-jobs/using-a-matrix-for-your-jobs)

## 🎯 Best Practices

1. **Use semantic versioning** for releases (MAJOR.MINOR.PATCH)
2. **Test locally** before pushing changes
3. **Keep workflows DRY** by using reusable components
4. **Monitor build times** and optimize as needed
5. **Document breaking changes** in release notes
6. **Verify checksums** after downloading releases
7. **Test each platform binary** in the target environment
8. **Keep dependencies updated** to avoid security issues
9. **Use platform-specific features** only when necessary
10. **Maintain consistent naming** across platforms

## 🔍 Troubleshooting

### Build Matrix Issues

**Problem**: Build fails for specific platform
- Check Go's support for that platform/architecture combination
- Verify CGO is disabled for cross-compilation
- Review platform-specific code

**Problem**: Archive creation fails
- Windows: Ensure 7z is available
- Unix: Ensure tar is available
- Check file paths are correct

### Release Issues

**Problem**: Release creation fails
- Verify `contents: write` permission is set
- Check artifact download completed successfully
- Ensure tag format matches trigger pattern

**Problem**: Checksums don't match
- Verify build was clean (no local modifications)
- Check file was not corrupted during download
- Ensure using correct hash algorithm (SHA256)

### Performance Issues

**Problem**: Builds taking too long
- Matrix builds should run in parallel
- Check for unnecessary steps
- Consider caching Go modules more effectively

## 📈 Metrics

Monitor these metrics for workflow health:
- Build success rate
- Average build time per platform
- Artifact size trends
- Download counts per platform
- Test coverage and pass rate
