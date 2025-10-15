# GitHub Actions Setup Checklist

Use this checklist to verify your workflow is properly configured and working.

## Pre-Flight Checks

### ✅ Repository Settings

- [ ] **GitHub Actions enabled**
  - Go to: Settings → Actions → General
  - Verify: "Allow all actions and reusable workflows" is selected
  
- [ ] **Workflow permissions configured**
  - Go to: Settings → Actions → General → Workflow permissions
  - Select: **"Read and write permissions"** ✅
  - Optional: Enable "Allow GitHub Actions to create and approve pull requests"
  - Click: **Save**

- [ ] **Releases feature enabled**
  - Go to: Settings → Features
  - Verify: "Releases" checkbox is checked ✅

### ✅ Workflow Files Present

- [ ] `.github/workflows/build.yml` exists and updated
- [ ] `.github/workflows/release.yml` exists
- [ ] Files have correct YAML syntax (no errors)

### ✅ Local Environment

- [ ] Go 1.21+ installed (`go version`)
- [ ] Git repository initialized
- [ ] Committed all workflow changes
- [ ] On `main` branch (or ready to merge)

## First Run Verification

### Step 1: Push to Main

```bash
# Ensure you're on main branch
git checkout main

# Stage the workflow changes
git add .github/workflows/build.yml
git add Makefile
git add GITHUB_SETUP.md .github/

# Commit
git commit -m "feat: Add automated binary builds and snapshot releases"

# Push to trigger workflow
git push origin main
```

**Check this:**
- [ ] Push completed without errors
- [ ] Received confirmation on terminal

### Step 2: Monitor Workflow Run

1. **Go to Actions tab**
   - URL: `https://github.com/YOUR_USERNAME/websearch-mcp/actions`
   
2. **Find the workflow run**
   - [ ] "Build and Test" workflow is running or completed
   - [ ] Commit message matches your push
   - [ ] Status shows progress or ✅ green checkmark

3. **Check each job**
   - [ ] **Test job**: ✅ Passed
   - [ ] **Build job**: ✅ Passed
   - [ ] **Create Snapshot Release job**: ✅ Passed (only on main)

4. **Check build logs**
   - [ ] "Build binary" step shows successful compilation
   - [ ] Version info displayed correctly
   - [ ] Archive creation successful
   - [ ] Checksums generated

### Step 3: Verify Artifacts

**In the workflow run page:**

1. **Scroll to Artifacts section**
   - [ ] See: `websearch-mcp-binary` artifact
   - [ ] See: `websearch-mcp-{version}-linux-amd64` artifact
   - [ ] Both show file sizes > 0
   
2. **Download and test** (optional)
   ```bash
   # Download the archive artifact
   # Extract locally
   tar -xzf websearch-mcp-*.tar.gz
   
   # Verify contents
   ls -la
   # Should see:
   # - websearch-mcp (binary)
   # - README.md
   ```

### Step 4: Verify Snapshot Release

1. **Go to Releases page**
   - URL: `https://github.com/YOUR_USERNAME/websearch-mcp/releases`
   
2. **Check for snapshot release**
   - [ ] Release titled "Snapshot Build (main @ {sha})" exists
   - [ ] Has orange "Pre-release" badge
   - [ ] Tag format: `snapshot-{commit-sha}`
   - [ ] Shows commit hash in description
   - [ ] Has Assets section

3. **Verify assets**
   - [ ] See: `websearch-mcp-{version}-linux-amd64.tar.gz`
   - [ ] See: `checksums.txt`
   - [ ] File sizes look correct
   - [ ] Download links work

### Step 5: Download and Test Binary

```bash
# Download the release asset
wget https://github.com/YOUR_USERNAME/websearch-mcp/releases/download/snapshot-XXX/websearch-mcp-VERSION-linux-amd64.tar.gz

# Verify checksum
wget https://github.com/YOUR_USERNAME/websearch-mcp/releases/download/snapshot-XXX/checksums.txt
sha256sum -c checksums.txt

# Extract
tar -xzf websearch-mcp-*-linux-amd64.tar.gz

# Make executable
chmod +x websearch-mcp

# Test run
./websearch-mcp
```

**Verify:**
- [ ] Binary downloads successfully
- [ ] Checksum verification passes
- [ ] Binary is executable
- [ ] Binary runs without errors
- [ ] Version info embedded (if you added --version flag)

## Additional Verifications

### ✅ Pull Request Workflow

1. **Create a test branch**
   ```bash
   git checkout -b test-pr
   echo "# Test" >> README.md
   git add README.md
   git commit -m "test: Verify PR workflow"
   git push origin test-pr
   ```

2. **Create Pull Request**
   - [ ] Go to GitHub → Pull Requests → New
   - [ ] Base: `main` ← Compare: `test-pr`
   - [ ] Create pull request

3. **Check workflow**
   - [ ] "Build and Test" workflow triggers
   - [ ] Tests run automatically
   - [ ] Build completes
   - [ ] **No** snapshot release created (correct!)
   - [ ] Artifacts are uploaded

4. **Cleanup**
   ```bash
   # Close PR, delete branch
   git checkout main
   git branch -D test-pr
   ```

### ✅ Official Release Workflow

**Test the release workflow:**

1. **Create version tag**
   ```bash
   git checkout main
   git pull origin main
   
   # Create tag
   git tag -a v0.1.0 -m "Test release v0.1.0"
   
   # Push tag
   git push origin v0.1.0
   ```

2. **Monitor workflow**
   - [ ] "Release" workflow triggers
   - [ ] Build job completes
   - [ ] Create Release job completes

3. **Verify official release**
   - [ ] Release appears on Releases page
   - [ ] Title: "WebSearch MCP Server v0.1.0"
   - [ ] **NOT** marked as pre-release (it's official)
   - [ ] Tag: `v0.1.0`
   - [ ] Has binary assets
   - [ ] Has release notes

4. **Cleanup test release** (optional)
   ```bash
   # Delete the release via GitHub UI or:
   gh release delete v0.1.0 --yes
   git push origin :refs/tags/v0.1.0
   git tag -d v0.1.0
   ```

## Troubleshooting Checklist

### ❌ Workflow doesn't trigger

**Check:**
- [ ] Workflow file in `.github/workflows/` directory
- [ ] File has `.yml` or `.yaml` extension
- [ ] YAML syntax is valid (use YAML validator)
- [ ] Pushed to correct branch (`main`)
- [ ] Actions are enabled in repo settings

### ❌ Snapshot release not created

**Check:**
- [ ] Workflow permissions set to "Read and write"
- [ ] Pushed to `main` branch (not `develop` or other)
- [ ] `create-snapshot-release` job shows in workflow run
- [ ] No errors in release creation step logs
- [ ] `GITHUB_TOKEN` has required permissions

**Fix:**
```yaml
# In build.yml, verify this is present:
permissions:
  contents: write
```

### ❌ Build fails

**Common issues:**

1. **Go version mismatch**
   - [ ] Workflow uses Go 1.21
   - [ ] Your local Go version compatible
   - [ ] `go.mod` specifies correct version

2. **Missing dependencies**
   - [ ] `go.mod` and `go.sum` committed
   - [ ] Dependencies downloadable
   - [ ] No private dependencies without access

3. **Build flags incorrect**
   - [ ] Check build step logs
   - [ ] Verify ldflags syntax
   - [ ] Version variables exist in main.go

### ❌ Artifacts empty or missing

**Check:**
- [ ] Build step completed successfully
- [ ] Binary file created (check logs)
- [ ] Archive creation step passed
- [ ] Upload artifact step completed
- [ ] Correct file paths in workflow

### ❌ Checksums don't match

**Possible causes:**
- [ ] Downloaded wrong file version
- [ ] File corrupted during download
- [ ] Checksum file from different build

**Fix:**
Re-download both the binary archive and checksums.txt from the same release.

## Success Criteria

All of these should be true:

- [x] ✅ Workflow runs on push to main
- [x] ✅ All jobs pass (green checkmarks)
- [x] ✅ Artifacts appear in workflow run
- [x] ✅ Snapshot release created automatically
- [x] ✅ Binary is downloadable from Releases
- [x] ✅ Binary runs successfully
- [x] ✅ Checksums verify correctly
- [x] ✅ Version info embedded in binary

## Next Steps After Verification

Once everything is working:

1. **Document for team**
   - Share `GITHUB_SETUP.md` with team
   - Add download instructions to README
   - Update contribution guidelines

2. **Set up branch protection** (optional)
   - Require PR reviews
   - Require status checks
   - Prevent direct pushes to main

3. **Plan release strategy**
   - Decide on versioning scheme (semver)
   - Create release schedule
   - Document release process

4. **Consider enhancements**
   - Add multi-platform builds to CI
   - Set up automated testing
   - Add code coverage reports
   - Implement changelog generation

## Support Resources

- **Workflow logs**: Actions tab → Select run → View logs
- **GitHub docs**: https://docs.github.com/en/actions
- **Go build docs**: https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies

## Questions?

If something doesn't work:

1. Check workflow logs for error messages
2. Verify all items in this checklist
3. Review the detailed guides:
   - `GITHUB_SETUP.md` - Full setup guide
   - `WORKFLOW_SUMMARY.md` - Quick reference
   - `CHANGES.md` - What changed
   - `WORKFLOW_DIAGRAM.md` - Visual diagrams

---

**Last Updated**: When you set this up
**Your Notes**:
- 
- 
- 
