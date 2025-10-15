# GitHub Actions Setup Guide

This document outlines the changes made to enable automatic binary builds and downloadable artifacts for every push to the `main` branch.

## Changes Summary

### 1. Updated `.github/workflows/build.yml`

The build workflow has been enhanced to:

- **Build optimized binaries** with proper version information matching the Makefile standards
- **Create distributable archives** (`.tar.gz`) containing the binary and README
- **Generate checksums** for verification
- **Upload artifacts** that are available for 90 days
- **Create snapshot releases** automatically for every push to `main` branch

#### New Features:

- **Version Information**: Binaries now include version, build time, and git commit hash
- **Two Artifact Types**:
  - `websearch-mcp-binary`: Raw binary file
  - `websearch-mcp-{version}-linux-amd64`: Archive with binary, README, and checksums
- **Snapshot Releases**: Automatic pre-release creation on `main` branch pushes with downloadable assets

### 2. Updated `Makefile`

Added new capabilities:

- **`build-all-platforms`**: Build binaries for multiple platforms locally:
  - Linux (amd64, arm64)
  - macOS (amd64, arm64/Apple Silicon)
  - Windows (amd64)
- **Consistent naming**: Using `BINARY_NAME` variable throughout
- **Better organization**: Cleaner dist directory structure

## How It Works

### On Every Push to Main:

1. **Tests run** to ensure code quality
2. **Binary is built** with version information
3. **Artifacts are created**:
   - Available in the Actions tab under each workflow run
   - Retained for 90 days
4. **Snapshot release is created**:
   - Tagged as `snapshot-{commit-sha}`
   - Marked as pre-release
   - Contains downloadable binary archive
   - Includes checksums for verification

### On Tag Push (v*):

The existing `release.yml` workflow creates official releases when you push a version tag.

## GitHub Configuration Steps

### ✅ Already Configured (No Action Needed)

The following permissions and settings are already configured in the workflow files:

1. **`permissions: contents: write`** - Allows the workflow to create releases
2. **`GITHUB_TOKEN`** - Automatically provided by GitHub Actions

### 🔧 Optional Configuration

#### 1. Adjust Artifact Retention (Optional)

The current retention is set to **90 days**. To change this:

- Edit `.github/workflows/build.yml`
- Find `retention-days: 90`
- Change to your preferred number of days (max 400 for public repos, 90 for private)

#### 2. Protect the Main Branch (Recommended)

To ensure quality before merges:

1. Go to **Settings** → **Branches**
2. Click **Add rule** under "Branch protection rules"
3. Enter `main` as the branch name pattern
4. Enable:
   - ✅ Require a pull request before merging
   - ✅ Require status checks to pass before merging
   - Select the `Test` job from the build workflow
5. Click **Create** or **Save changes**

#### 3. Enable GitHub Releases (Should be enabled by default)

1. Go to your repository **Settings**
2. Scroll down to **Features**
3. Ensure **Releases** is checked

#### 4. Clean Up Old Snapshot Releases (Optional)

Snapshot releases accumulate over time. You can:

**Option A: Manual cleanup**
- Go to **Releases**
- Delete old snapshot releases as needed

**Option B: Automated cleanup** (requires additional workflow)
- Consider implementing a workflow to automatically delete old snapshots
- Keep only the last N snapshot releases

## Accessing Built Binaries

### From GitHub Actions Artifacts:

1. Go to **Actions** tab
2. Click on a completed workflow run
3. Scroll to **Artifacts** section
4. Download the artifact you need

### From Snapshot Releases:

1. Go to **Releases** (or **Code** → **Releases**)
2. Find the snapshot release for your commit
3. Download from the **Assets** section
4. Verify using the provided checksums

### From Official Releases:

1. Create and push a version tag:
   ```bash
   git tag -a v1.0.0 -m "Release version 1.0.0"
   git push origin v1.0.0
   ```
2. The `release.yml` workflow will create an official release
3. Download from the **Releases** page

## Building Locally

### Single Platform:

```bash
# Build for current platform
make build

# Build optimized release binary
make build-release
```

### Multiple Platforms:

```bash
# Build for all supported platforms
make build-all-platforms
```

This creates binaries in the `dist/` directory for:
- Linux (amd64, arm64)
- macOS (Intel, Apple Silicon)
- Windows (amd64)

## Verification

### Verify a downloaded binary:

```bash
# Extract the archive
tar -xzf websearch-mcp-{version}-linux-amd64.tar.gz

# Check the checksum (from checksums.txt)
sha256sum websearch-mcp

# Make executable
chmod +x websearch-mcp

# Run
./websearch-mcp
```

## Workflow Triggers

| Workflow | Trigger | Purpose | Creates Release |
|----------|---------|---------|-----------------|
| `build.yml` | Push to `main` or `develop` | Build and test | Yes (snapshot on main) |
| `build.yml` | Pull request to `main` | Test before merge | No |
| `release.yml` | Push tag `v*` | Official release | Yes (official) |

## Troubleshooting

### Issue: Snapshot releases not being created

**Check:**
1. Go to **Settings** → **Actions** → **General**
2. Under **Workflow permissions**, ensure:
   - ✅ "Read and write permissions" is selected
   - OR "Read repository contents and packages permissions" with manual permission grants

### Issue: Artifacts not visible

**Check:**
1. Workflow completed successfully (green checkmark)
2. Scroll down in the workflow run to the **Artifacts** section
3. Artifacts expire after 90 days

### Issue: Build fails on version information

**Check:**
1. Ensure git history is available (`fetch-depth: 0`)
2. If no tags exist, it will use `dev-{commit}` as version

## Next Steps

1. **Push to main** to trigger your first snapshot build
2. **Check the Actions tab** to monitor the workflow
3. **Verify the snapshot release** is created
4. **Download and test** the binary

For official releases:
```bash
git tag -a v1.0.0 -m "First stable release"
git push origin v1.0.0
```

## Questions?

If you encounter any issues:
1. Check the workflow logs in the **Actions** tab
2. Verify the settings mentioned above
3. Ensure your repository has releases enabled
