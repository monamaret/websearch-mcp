# Release Guide

This guide explains how to create and manage releases for the WebSearch MCP Server.

## 📋 Table of Contents

- [Quick Release](#quick-release)
- [Pre-Release Checklist](#pre-release-checklist)
- [Creating a Release](#creating-a-release)
- [Release Types](#release-types)
- [Post-Release Tasks](#post-release-tasks)
- [Rollback Procedure](#rollback-procedure)
- [Release Assets](#release-assets)

## Quick Release

For those familiar with the process:

```bash
# 1. Update version references (if needed)
# 2. Commit all changes
git add .
git commit -m "Prepare for release v1.2.3"
git push

# 3. Create and push tag
git tag v1.2.3
git push origin v1.2.3

# 4. GitHub Actions will automatically:
#    - Build binaries for all platforms
#    - Create release with all assets
#    - Generate release notes
```

## Pre-Release Checklist

Before creating a release, ensure:

### Code Quality
- [ ] All tests pass locally: `go test ./...`
- [ ] Code is formatted: `make fmt`
- [ ] Linter passes: `make lint` (if golangci-lint installed)
- [ ] No known bugs or critical issues

### Documentation
- [ ] Update `CHANGES.md` with release notes
- [ ] Update version numbers in documentation if needed
- [ ] README.md is up to date
- [ ] All new features are documented

### Version Numbering
- [ ] Follow [Semantic Versioning](https://semver.org/)
  - MAJOR: Breaking changes
  - MINOR: New features (backward compatible)
  - PATCH: Bug fixes (backward compatible)

### Testing
- [ ] Test build locally: `make build-all-platforms`
- [ ] Verify binaries work on target platforms
- [ ] Test MCP integration with Tabnine

### Repository State
- [ ] All changes committed to main branch
- [ ] Branch is up to date with remote
- [ ] No uncommitted changes

## Creating a Release

### Step 1: Prepare the Release

1. **Update CHANGES.md**
   ```bash
   # Add entry to docs/CHANGES.md
   vim docs/CHANGES.md
   ```

   Example entry:
   ```markdown
   ## [1.2.3] - 2024-01-15
   
   ### Added
   - New feature X
   - Enhanced Y functionality
   
   ### Fixed
   - Bug in Z component
   - Performance issue with search
   
   ### Changed
   - Updated dependency versions
   ```

2. **Commit changes**
   ```bash
   git add docs/CHANGES.md
   git commit -m "Update CHANGES.md for v1.2.3"
   git push origin main
   ```

### Step 2: Create and Push Tag

1. **Create annotated tag**
   ```bash
   # For releases
   git tag -a v1.2.3 -m "Release version 1.2.3"
   
   # For pre-releases
   git tag -a v1.2.3-beta.1 -m "Release version 1.2.3-beta.1"
   ```

2. **Push tag to GitHub**
   ```bash
   git push origin v1.2.3
   ```

### Step 3: Monitor Release Build

1. **Watch GitHub Actions**
   - Go to: `https://github.com/yourusername/websearch-mcp/actions`
   - Find the "Release" workflow
   - Monitor build progress

2. **Verify builds complete**
   - All 6 platform builds succeed (darwin-amd64, darwin-arm64, windows-amd64, windows-arm64, linux-amd64, linux-arm64)
   - Archives are created
   - Release is published

### Step 4: Verify Release

1. **Check release page**
   - Navigate to: `https://github.com/yourusername/websearch-mcp/releases`
   - Verify latest release appears
   - Check all assets are present

2. **Verify assets**
   Expected files:
   - `websearch-mcp-v1.2.3-darwin-amd64.tar.gz`
   - `websearch-mcp-v1.2.3-darwin-arm64.tar.gz`
   - `websearch-mcp-v1.2.3-windows-amd64.zip`
   - `websearch-mcp-v1.2.3-windows-arm64.zip`
   - `websearch-mcp-v1.2.3-linux-amd64.tar.gz`
   - `websearch-mcp-v1.2.3-linux-arm64.tar.gz`
   - `checksums.txt`

3. **Test download and run**
   ```bash
   # Download for your platform
   curl -LO https://github.com/youruser/websearch-mcp/releases/download/v1.2.3/websearch-mcp-v1.2.3-darwin-arm64.tar.gz
   
   # Extract
   tar -xzf websearch-mcp-v1.2.3-darwin-arm64.tar.gz
   
   # Run
   ./websearch-mcp-darwin-arm64
   ```

## Release Types

### Stable Release

For production-ready releases:

```bash
# Version format: vMAJOR.MINOR.PATCH
git tag -a v1.2.3 -m "Release version 1.2.3"
git push origin v1.2.3
```

**Release flags:**
- `draft: false`
- `prerelease: false`

### Pre-Release (Beta/RC)

For testing before stable release:

```bash
# Beta release
git tag -a v1.2.3-beta.1 -m "Beta release 1.2.3-beta.1"
git push origin v1.2.3-beta.1

# Release candidate
git tag -a v1.2.3-rc.1 -m "Release candidate 1.2.3-rc.1"
git push origin v1.2.3-rc.1
```

**Release flags:**
- `draft: false`
- `prerelease: true`

### Snapshot Release

Automatic builds from main branch:

- Created automatically on push to main
- Tagged as: `snapshot-{commit-sha}`
- Marked as pre-release
- Not meant for production use

## Post-Release Tasks

### 1. Update Documentation

If needed, update documentation links:
```bash
# Update README.md with new version
sed -i 's/VERSION/v1.2.3/g' README.md
git commit -am "Update version references to v1.2.3"
git push
```

### 2. Announce Release

- Update project website (if applicable)
- Announce on social media/blog
- Notify users via mailing list
- Update package managers (if applicable)

### 3. Monitor Issues

- Watch for bug reports
- Check download metrics
- Monitor user feedback

### 4. Plan Next Release

- Create milestone for next version
- Assign issues to milestone
- Update project roadmap

## Rollback Procedure

If a release has critical issues:

### Option 1: Delete Release (Recommended for critical bugs)

```bash
# Delete remote tag
git push --delete origin v1.2.3

# Delete local tag
git tag -d v1.2.3

# Delete release on GitHub
# Go to Releases page → Click release → Delete release
```

### Option 2: Mark as Pre-release

If release is already downloaded:
1. Edit release on GitHub
2. Check "This is a pre-release"
3. Add warning to release notes
4. Create hotfix release

### Option 3: Hotfix Release

For urgent fixes:

```bash
# Create hotfix branch from tag
git checkout -b hotfix-1.2.4 v1.2.3

# Fix issue
git commit -am "Fix critical bug"

# Create new patch release
git tag -a v1.2.4 -m "Hotfix release 1.2.4"
git push origin v1.2.4

# Merge back to main
git checkout main
git merge hotfix-1.2.4
git push
```

## Release Assets

### Asset Naming Convention

```
websearch-mcp-{version}-{platform}-{arch}.{ext}
```

Examples:
- `websearch-mcp-v1.2.3-darwin-arm64.tar.gz`
- `websearch-mcp-v1.2.3-windows-amd64.zip`
- `websearch-mcp-v1.2.3-linux-amd64.tar.gz`

### Asset Contents

Each archive contains:
- Binary executable (platform-specific name)
- `README.md` documentation

### Checksums

`checksums.txt` contains SHA256 hashes for all archives:
```
a1b2c3d4... websearch-mcp-v1.2.3-darwin-amd64.tar.gz
e5f6g7h8... websearch-mcp-v1.2.3-darwin-arm64.tar.gz
...
```

Users can verify downloads:
```bash
# Linux/macOS
sha256sum -c checksums.txt

# Windows (PowerShell)
Get-FileHash websearch-mcp-*.zip -Algorithm SHA256
```

## Troubleshooting

### Build Fails for Specific Platform

**Problem:** One or more platform builds fail

**Solution:**
1. Check build logs in GitHub Actions
2. Test build locally:
   ```bash
   GOOS=darwin GOARCH=arm64 go build .
   ```
3. Fix issues and create new tag

### Release Not Created

**Problem:** Tag pushed but no release created

**Solution:**
1. Check tag format matches `v*`
2. Verify workflow permissions (needs `contents: write`)
3. Check workflow runs in Actions tab
4. Manually trigger workflow if needed

### Missing Assets

**Problem:** Release created but some assets missing

**Solution:**
1. Check build matrix in workflow
2. Verify all builds completed successfully
3. Re-run failed jobs in Actions tab
4. If issue persists, delete and recreate release

### Incorrect Version in Binary

**Problem:** Binary reports wrong version

**Solution:**
1. Ensure tag is annotated: `git tag -a v1.2.3`
2. Verify ldflags in build command
3. Check version detection in workflow

## Best Practices

1. **Always use annotated tags**
   ```bash
   # Good
   git tag -a v1.2.3 -m "Release 1.2.3"
   
   # Bad (lightweight tag)
   git tag v1.2.3
   ```

2. **Follow semantic versioning strictly**
   - Breaking changes → Major version
   - New features → Minor version
   - Bug fixes → Patch version

3. **Test before releasing**
   - Build and test all platforms locally
   - Run integration tests
   - Verify documentation

4. **Write clear release notes**
   - Highlight breaking changes
   - List new features
   - Document bug fixes
   - Include upgrade instructions if needed

5. **Maintain CHANGES.md**
   - Update before each release
   - Follow Keep a Changelog format
   - Include dates

6. **Monitor after release**
   - Watch for issues
   - Respond to user feedback
   - Be ready to hotfix if needed

## Release Schedule

Recommended release cadence:

- **Major releases**: When breaking changes accumulate
- **Minor releases**: Monthly (if new features)
- **Patch releases**: As needed for bugs
- **Snapshot builds**: Automatic on every main branch push

## Contact

For questions about the release process:
- Open an issue on GitHub
- Contact maintainers
- Check CI/CD documentation

## See Also

- [Semantic Versioning](https://semver.org/)
- [Keep a Changelog](https://keepachangelog.com/)
- [GitHub Releases Documentation](https://docs.github.com/en/repositories/releasing-projects-on-github)
- [WORKFLOWS.md](WORKFLOWS.md) - CI/CD workflows
- [BUILDING.md](BUILDING.md) - Build instructions
