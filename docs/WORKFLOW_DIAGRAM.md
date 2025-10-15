# GitHub Actions Workflow Diagram

## Build Workflow (build.yml)

### Trigger: Push to `main` or `develop` / PR to `main`

```
┌─────────────────────────────────────────────────────────────────┐
│                         PUSH TO MAIN                             │
└────────────────────────────┬────────────────────────────────────┘
                             │
                             ▼
                    ┌────────────────┐
                    │   Job: Test    │
                    │                │
                    │ • Checkout     │
                    │ • Setup Go     │
                    │ • Download deps│
                    │ • Run tests    │
                    └────────┬───────┘
                             │
                    ┌────────▼───────┐
                    │ Tests Pass? ✓  │
                    └────────┬───────┘
                             │
                             ▼
                    ┌────────────────────────────┐
                    │      Job: Build            │
                    │                            │
                    │ • Checkout (full history) │
                    │ • Setup Go 1.21            │
                    │ • Get version info         │
                    │   - git describe --tags    │
                    │   - build timestamp        │
                    │   - commit hash            │
                    │ • Build binary             │
                    │   CGO_ENABLED=0            │
                    │   -ldflags with version    │
                    │ • Create tar.gz archive    │
                    │   - binary                 │
                    │   - README.md              │
                    │ • Generate checksums       │
                    └────────┬───────────────────┘
                             │
                ┌────────────┴────────────┐
                │                         │
                ▼                         ▼
    ┌───────────────────┐    ┌────────────────────────┐
    │ Upload Artifact 1 │    │  Upload Artifact 2     │
    │                   │    │                        │
    │ Name:             │    │ Name:                  │
    │ websearch-mcp-    │    │ websearch-mcp-{ver}-   │
    │ binary            │    │ linux-amd64            │
    │                   │    │                        │
    │ Contains:         │    │ Contains:              │
    │ • Raw binary      │    │ • Archive (.tar.gz)    │
    │                   │    │ • Checksums.txt        │
    │                   │    │                        │
    │ Retention: 90d    │    │ Retention: 90d         │
    └───────────────────┘    └────────────────────────┘
                             │
                    ┌────────▼───────────┐
                    │ Only on main push? │
                    │      (Yes/No)      │
                    └────────┬───────────┘
                             │ Yes
                             ▼
              ┌──────────────────────────────────┐
              │ Job: Create Snapshot Release     │
              │                                  │
              │ • Download artifacts             │
              │ • Create GitHub Release          │
              │   Tag: snapshot-{sha}            │
              │   Type: Pre-release              │
              │   Assets: tar.gz + checksums     │
              │ • Generate release notes         │
              └──────────────┬───────────────────┘
                             │
                             ▼
                   ┌─────────────────────┐
                   │  SNAPSHOT RELEASE   │
                   │     PUBLISHED!      │
                   │                     │
                   │ Available at:       │
                   │ Releases page       │
                   │                     │
                   │ ✓ Downloadable      │
                   │ ✓ With checksums    │
                   │ ✓ With docs         │
                   └─────────────────────┘
```

## Release Workflow (release.yml)

### Trigger: Push tag `v*` (e.g., v1.0.0)

```
┌──────────────────────────────────────────────────────────────┐
│                    PUSH VERSION TAG (v*.*.*)                 │
│                    git push origin v1.0.0                    │
└────────────────────────────┬─────────────────────────────────┘
                             │
                             ▼
                ┌────────────────────────────┐
                │ Job: Build and Package     │
                │                            │
                │ • Extract version from tag │
                │ • Build binary             │
                │ • Create tar.gz            │
                │ • Upload as artifact       │
                └────────┬───────────────────┘
                         │
                         ▼
                ┌────────────────────────────┐
                │ Job: Create Release        │
                │                            │
                │ • Download artifacts       │
                │ • Create GitHub Release    │
                │   Tag: v1.0.0              │
                │   Type: Official           │
                │   Status: Published        │
                │ • Attach binary assets     │
                │ • Generate changelog       │
                └────────┬───────────────────┘
                         │
                         ▼
                ┌─────────────────────┐
                │  OFFICIAL RELEASE   │
                │    PUBLISHED!       │
                │                     │
                │ Version: v1.0.0     │
                │ Status: Stable      │
                │ ✓ Full release notes│
                │ ✓ Downloadable      │
                └─────────────────────┘
```

## Workflow Comparison

```
┌─────────────────┬──────────────────────┬─────────────────────┐
│     Trigger     │   Workflow File      │      Result         │
├─────────────────┼──────────────────────┼─────────────────────┤
│ Push to main    │ build.yml            │ Snapshot Release    │
│                 │                      │ (Pre-release)       │
├─────────────────┼──────────────────────┼─────────────────────┤
│ Push to develop │ build.yml            │ Artifacts only      │
│                 │                      │ (No release)        │
├─────────────────┼──────────────────────┼─────────────────────┤
│ Pull Request    │ build.yml            │ Tests + Build       │
│                 │                      │ (No artifacts/      │
│                 │                      │  release)           │
├─────────────────┼──────────────────────┼─────────────────────┤
│ Push tag v*     │ release.yml          │ Official Release    │
│                 │                      │ (Stable)            │
└─────────────────┴──────────────────────┴─────────────────────┘
```

## Access Points

```
┌────────────────────────────────────────────────────────────┐
│                    HOW TO GET BINARIES                     │
└────────────────────────────────────────────────────────────┘

Option 1: GitHub Actions Artifacts (Team/Internal)
   └─► Actions Tab
       └─► Select Workflow Run
           └─► Scroll to Artifacts
               └─► Download
                   • Available: 90 days
                   • Access: Repo collaborators

Option 2: Snapshot Releases (Latest main)
   └─► Releases Page
       └─► Find "snapshot-{sha}" (marked Pre-release)
           └─► Download from Assets
               • Available: Until deleted
               • Access: Public (if repo public)
               • Use for: Testing latest changes

Option 3: Official Releases (Stable)
   └─► Releases Page
       └─► Find version tag (v1.0.0, etc.)
           └─► Download from Assets
               • Available: Permanent
               • Access: Public (if repo public)
               • Use for: Production
```

## Version Resolution

```
┌────────────────────────────────────────────────────────────┐
│                  HOW VERSION IS DETERMINED                 │
└────────────────────────────────────────────────────────────┘

git describe --tags --always --dirty
     │
     ├─► Has tags? ──Yes──► v1.2.3 or v1.2.3-5-gabc123
     │
     └─► No tags? ──Yes──► abc1234 (short commit hash)
                             │
                             └─► Fallback: "dev-abc1234"

Examples:
• v1.0.0           → Clean tag
• v1.0.0-3-gabc123 → 3 commits after v1.0.0, hash abc123
• abc1234          → No tags, just commit hash
• dev-abc1234      → Fallback format
```

## File Flow

```
Source Code
    │
    ▼
┌─────────────┐
│   go build  │ ──────► websearch-mcp (binary)
└─────────────┘            │
                           ▼
                    ┌──────────────┐
                    │  tar + gzip  │ ──► websearch-mcp-{ver}-linux-amd64.tar.gz
                    └──────────────┘       │
                           │                ├─► websearch-mcp (binary)
                           │                └─► README.md
                           ▼
                    ┌──────────────┐
                    │  sha256sum   │ ──► checksums.txt
                    └──────────────┘       │
                           │                ├─► {archive}.tar.gz checksum
                           │                └─► websearch-mcp checksum
                           ▼
                    ┌──────────────────┐
                    │ Upload to GitHub │
                    └──────────────────┘
                           │
                    ┌──────┴───────┐
                    │              │
                    ▼              ▼
            ┌────────────┐  ┌──────────────┐
            │ Artifacts  │  │   Releases   │
            └────────────┘  └──────────────┘
```

## Quick Reference Commands

```bash
# Trigger snapshot release
git push origin main

# Trigger official release
git tag -a v1.0.0 -m "Release v1.0.0"
git push origin v1.0.0

# Build locally (current platform)
make build

# Build locally (all platforms)
make build-all-platforms

# Clean up
make clean
```

## Permissions Required

```
GitHub Settings → Actions → General → Workflow permissions

┌─────────────────────────────────────────────────┐
│ ● Read and write permissions                   │ ← SELECT THIS
│ ○ Read repository contents and packages        │
│   permissions                                   │
│                                                 │
│ ☑ Allow GitHub Actions to create and approve   │
│   pull requests                                 │
└─────────────────────────────────────────────────┘
```

## Success Indicators

```
✅ Workflow Run
   ├─► All jobs green ✓
   ├─► Artifacts present (2 items)
   └─► No errors in logs

✅ Snapshot Release (for main pushes)
   ├─► Appears in Releases page
   ├─► Tagged as "Pre-release"
   ├─► Has downloadable assets
   └─► Tag format: snapshot-{commit-sha}

✅ Downloaded Binary
   ├─► File size > 0
   ├─► Checksum matches
   ├─► Executable permission works
   └─► Runs without errors
```
