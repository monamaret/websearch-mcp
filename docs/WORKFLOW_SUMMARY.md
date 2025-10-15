# Workflow Quick Reference

## What Changed

✅ **Updated `.github/workflows/build.yml`**
- Added automatic binary building with version info
- Creates downloadable artifacts (90-day retention)
- Auto-creates snapshot releases on every `main` push
- Generates checksums for verification

✅ **Updated `Makefile`**
- Added `build-all-platforms` target
- Improved consistency with binary naming
- Better dist directory organization

## Quick Actions

### Download Latest Build

**From Artifacts (internal/team use):**
1. Go to **Actions** → Latest workflow run
2. Scroll to **Artifacts**
3. Download `websearch-mcp-{version}-linux-amd64`

**From Releases (public/external use):**
1. Go to **Releases**
2. Find latest snapshot or official release
3. Download from **Assets**

### Create Official Release

```bash
git tag -a v1.0.0 -m "Release v1.0.0"
git push origin v1.0.0
```

### Build Locally

```bash
# Current platform
make build

# All platforms
make build-all-platforms
```

## GitHub Settings Required

### ✅ Already Set (in workflow files)
- `permissions: contents: write`
- `GITHUB_TOKEN` usage

### 🔧 Verify These Settings

**Actions Permissions:**
1. **Settings** → **Actions** → **General**
2. Under "Workflow permissions":
   - Select "Read and write permissions" ✅
   - Enable "Allow GitHub Actions to create and approve pull requests" (optional)

**That's it!** Everything else is configured in the workflow files.

## Testing the Setup

1. Make a small change to README or code
2. Commit and push to `main`
3. Go to **Actions** tab → Watch the workflow run
4. Check **Releases** → Verify snapshot release was created
5. Download the binary and test it

## File Locations

- Main workflow: `.github/workflows/build.yml`
- Release workflow: `.github/workflows/release.yml`
- Build config: `Makefile`
- Full guide: `GITHUB_SETUP.md`
