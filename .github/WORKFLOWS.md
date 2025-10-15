# GitHub Workflows Documentation

This repository includes simplified GitHub Actions workflows for automated building, testing, and releasing of the WebSearch MCP Server.

## 📋 Available Workflows

### 1. Build and Test (`.github/workflows/build.yml`)

**Triggers:**
- Push to `main` or `develop` branches
- Pull requests to `main` branch

**Jobs:**
- **Test**: Runs unit tests
- **Build**: Creates binaries for Linux and macOS (AMD64/ARM64 architectures)
- **Docker**: Builds multi-platform Docker images (build only, no push)

**Artifacts:**
- Cross-platform binaries
- Docker images (for testing purposes)

### 2. Release (`.github/workflows/release.yml`)

**Triggers:**
- Git tags matching `v*` (e.g., `v1.0.0`)

**Features:**
- Creates release binaries for Linux and macOS (AMD64/ARM64)
- Creates GitHub releases with auto-generated descriptions
- Includes documentation in release archives

**Assets:**
- Compressed archives (`.tar.gz` format)
- GitHub releases with basic documentation

## 🔧 Setup Requirements

### Repository Settings

1. **Enable GitHub Actions** in repository settings
2. **Configure branch protection** for `main` branch:
   - Require status checks before merging
   - Include workflow checks: `test`, `build`

## 🚀 Release Process

### Creating a Release

1. **Prepare for release:**
   ```bash
   # Ensure all changes are committed and pushed
   git checkout main
   git pull origin main
   ```

2. **Create and push a tag:**
   ```bash
   # Create a new tag (use semantic versioning)
   git tag v1.0.0
   git push origin v1.0.0
   ```

3. **Automated process:**
   - Release workflow triggers automatically
   - Binaries are built for supported platforms
   - GitHub release is created with basic information

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
- **Health endpoint**: `GET http://localhost:8080/health`
- **Version endpoint**: `GET http://localhost:8080/version`
- **MCP initialize**: Included in server info

## 🧪 Testing Workflows

### Local Testing

Test workflow components locally:

```bash
# Test build process
go build -o websearch-mcp .

# Test Docker build
docker build -t websearch-mcp:test .

# Run tests
go test -v ./...
```

### Workflow Debugging

**View workflow runs:**
- Go to Actions tab in GitHub repository
- Click on specific workflow run
- Expand job steps to see detailed logs

**Common issues:**
- **Build failures**: Check Go version compatibility
- **Test failures**: Review test output in workflow logs

## 📊 Workflow Outputs

### Build Artifacts

**Binary naming convention:**
```
websearch-mcp-{os}-{arch}
```

**Examples:**
- `websearch-mcp-linux-amd64`
- `websearch-mcp-darwin-arm64`

### Release Archives

**Contents:**
- Binary executable
- README.md documentation

**Archive format:**
- `.tar.gz` for all platforms

## 🔄 Maintenance

### Regular Updates

**Manual maintenance:**
- Review and update Go version in workflows when new releases are available
- Monitor workflow performance and optimize as needed

### Workflow Modifications

When modifying workflows:

1. **Test changes** in a fork or feature branch
2. **Use workflow_dispatch** for manual testing
3. **Monitor resource usage** to stay within GitHub Actions limits
4. **Document changes** in this file

## 📚 Additional Resources

- [GitHub Actions Documentation](https://docs.github.com/en/actions)
- [Semantic Versioning](https://semver.org/)

## 🎯 Best Practices

1. **Use semantic versioning** for releases
2. **Keep workflows simple** and focused
3. **Test locally** before pushing changes
4. **Document breaking changes** in release notes
5. **Monitor workflow execution** for performance issues