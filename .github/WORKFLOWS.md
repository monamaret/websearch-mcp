# GitHub Workflows Documentation

This repository includes several GitHub Actions workflows for automated building, testing, and releasing of the WebSearch MCP Server.

## 📋 Available Workflows

### 1. Build and Test (`.github/workflows/build.yml`)

**Triggers:**
- Push to `main` or `develop` branches
- Pull requests to `main` branch

**Jobs:**
- **Test**: Runs unit tests with race detection and coverage reporting
- **Build**: Creates binaries for multiple platforms (Linux, macOS, Windows on AMD64/ARM64)
- **Docker**: Builds multi-platform Docker images
- **Lint**: Runs golangci-lint for code quality
- **Security**: Performs security scanning with Gosec

**Artifacts:**
- Cross-platform binaries with SHA256 checksums
- Coverage reports uploaded to Codecov
- Docker images (when not on PR)

### 2. Release (`.github/workflows/release.yml`)

**Triggers:**
- Git tags matching `v*` (e.g., `v1.0.0`)

**Features:**
- Creates release binaries for all supported platforms
- Generates SHA256 checksums for verification
- Builds and pushes Docker images to Docker Hub and GitHub Container Registry
- Creates GitHub releases with auto-generated changelogs
- Includes documentation and examples in release archives

**Assets:**
- Compressed archives (`.tar.gz` for Unix, `.zip` for Windows)
- SHA256 checksum files
- Docker images tagged with version and `latest`

### 3. GoReleaser (`.github/workflows/goreleaser.yml`)

**Triggers:**
- Git tags matching `v*`

**Features:**
- Advanced release automation using GoReleaser
- Homebrew formula generation (if configured)
- Linux package generation (DEB, RPM, APK)
- Docker multi-platform builds
- Comprehensive changelog generation

### 4. Security and Dependencies (`.github/workflows/security.yml`)

**Triggers:**
- Weekly schedule (Mondays at 6 AM UTC)
- Manual workflow dispatch

**Features:**
- Go vulnerability scanning with `govulncheck`
- Dependency review for pull requests
- Nancy security scanning
- CodeQL analysis for code security
- Automated dependency updates via pull requests

## 🔧 Setup Requirements

### Required Secrets

For full functionality, configure these secrets in your repository:

```bash
# Docker Hub (optional, for Docker image publishing)
DOCKER_USERNAME=your-dockerhub-username
DOCKER_PASSWORD=your-dockerhub-password

# Homebrew tap (optional, for Homebrew formula)
HOMEBREW_TAP_GITHUB_TOKEN=github-token-with-repo-access
```

### Repository Settings

1. **Enable GitHub Actions** in repository settings
2. **Configure branch protection** for `main` branch:
   - Require status checks before merging
   - Include workflow checks: `test`, `build`, `lint`, `security`
3. **Enable vulnerability alerts** in security settings

## 🚀 Release Process

### Creating a Release

1. **Prepare for release:**
   ```bash
   # Ensure all changes are committed and pushed
   git checkout main
   git pull origin main
   
   # Update version in documentation if needed
   ```

2. **Create and push a tag:**
   ```bash
   # Create a new tag (use semantic versioning)
   git tag v1.0.0
   git push origin v1.0.0
   ```

3. **Automated process:**
   - Release workflow triggers automatically
   - Binaries are built for all platforms
   - Docker images are created and published
   - GitHub release is created with changelog
   - GoReleaser handles additional packaging

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
make build-all

# Test Docker build
make docker-build

# Run security checks
go install github.com/securecodewarrior/gosec/v2/cmd/gosec@latest
gosec ./...

# Run vulnerability check
go install golang.org/x/vuln/cmd/govulncheck@latest
govulncheck ./...
```

### Workflow Debugging

**View workflow runs:**
- Go to Actions tab in GitHub repository
- Click on specific workflow run
- Expand job steps to see detailed logs

**Common issues:**
- **Build failures**: Check Go version compatibility
- **Docker push failures**: Verify Docker Hub credentials
- **Release failures**: Ensure tag follows semantic versioning

## 📊 Workflow Outputs

### Build Artifacts

**Binary naming convention:**
```
websearch-mcp-{version}-{os}-{arch}[.exe]
```

**Examples:**
- `websearch-mcp-v1.0.0-linux-amd64`
- `websearch-mcp-v1.0.0-darwin-arm64`
- `websearch-mcp-v1.0.0-windows-amd64.exe`

### Docker Images

**Image tags:**
- `websearch-mcp:latest` (latest release)
- `websearch-mcp:v1.0.0` (specific version)
- `websearch-mcp:v1` (major version)
- `websearch-mcp:v1.0` (minor version)
- `ghcr.io/owner/websearch-mcp:*` (GitHub Container Registry)

### Release Archives

**Contents:**
- Binary executable
- Documentation files (README.md, USAGE.md, etc.)
- Configuration examples
- SHA256 checksums

## 🔄 Maintenance

### Regular Updates

**Weekly automated tasks:**
- Dependency vulnerability scanning
- Dependency update pull requests
- Security analysis

**Manual maintenance:**
- Review and merge dependency updates
- Update Go version in workflows when new releases are available
- Monitor workflow performance and optimize as needed

### Workflow Modifications

When modifying workflows:

1. **Test changes** in a fork or feature branch
2. **Use workflow_dispatch** for manual testing
3. **Monitor resource usage** to stay within GitHub Actions limits
4. **Document changes** in this file

## 📚 Additional Resources

- [GitHub Actions Documentation](https://docs.github.com/en/actions)
- [GoReleaser Documentation](https://goreleaser.com/)
- [Docker GitHub Actions](https://github.com/docker/build-push-action)
- [Semantic Versioning](https://semver.org/)

## 🎯 Best Practices

1. **Use semantic versioning** for releases
2. **Include comprehensive tests** before releases
3. **Monitor security vulnerabilities** regularly
4. **Keep dependencies updated** via automated PRs
5. **Document breaking changes** in release notes
6. **Test workflows** in feature branches when possible