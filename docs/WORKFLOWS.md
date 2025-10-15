# GitHub Workflows Documentation

This repository includes simplified GitHub Actions workflows for automated building, testing, and releasing of the WebSearch MCP Server.

## 📋 Available Workflows

### 1. Build and Test (`.github/workflows/build.yml`)

Triggers:
- Push to `main` or `develop` branches
- Pull requests to `main` branch

Jobs:
- Test: Runs unit tests
- Build: Creates a single optimized Go binary artifact (no OS/arch matrix)

Artifacts:
- Single binary: `websearch-mcp`

### 2. Release (`.github/workflows/release.yml`)

Triggers:
- Git tags matching `v*` (e.g., `v1.0.0`)

Features:
- Builds a single optimized binary with embedded version info
- Creates GitHub releases with auto-generated descriptions
- Includes documentation in a single release archive

Assets:
- Compressed archive: `websearch-mcp-<version>.tar.gz`

Note: Docker images are no longer built as part of CI.

## 🔧 Setup Requirements

Repository Settings:
1. Enable GitHub Actions in repository settings
2. Configure branch protection for `main` branch:
   - Require status checks before merging
   - Include workflow checks: `test`, `build`

## 🚀 Release Process

Creating a Release:
1. Prepare for release:
   ```bash
   git checkout main
   git pull origin main
   ```
2. Create and push a tag:
   ```bash
   git tag v1.0.0
   git push origin v1.0.0
   ```
3. Automated process:
   - Release workflow triggers automatically
   - A single binary is built and archived
   - GitHub release is created with the archive attached

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
- Health endpoint: GET http://localhost:8080/health
- Version endpoint: GET http://localhost:8080/version
- MCP initialize: Included in server info

## 🧪 Testing Workflows

Local Testing:

```bash
# Test build process
go build -o websearch-mcp .

# Run tests
go test -v ./...
```

Workflow Debugging:
- View runs in the Actions tab
- Expand job steps to see detailed logs

Common issues:
- Build failures: Check Go version compatibility
- Test failures: Review test output in workflow logs

## 📊 Workflow Outputs

Build Artifacts:
- Binary name: `websearch-mcp`

Release Archives:
- Contents:
  - Binary executable
  - README.md documentation
- Archive format: `.tar.gz`

## 🔄 Maintenance

Regular Updates:
- Review and update Go version in workflows when new releases are available
- Monitor workflow performance and optimize as needed

Workflow Modifications:
1. Test changes in a fork or feature branch
2. Use workflow_dispatch for manual testing
3. Monitor resource usage to stay within GitHub Actions limits
4. Document changes in this file

## 📚 Additional Resources

- GitHub Actions Documentation: https://docs.github.com/en/actions
- Semantic Versioning: https://semver.org/

## 🎯 Best Practices

1. Use semantic versioning for releases
2. Keep workflows simple and focused
3. Test locally before pushing changes
4. Document breaking changes in release notes
5. Monitor workflow execution for performance issues
