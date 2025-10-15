# WebSearch MCP Server Documentation

Welcome to the WebSearch MCP Server documentation! This directory contains comprehensive guides for using, building, and deploying the server.

## 📚 Documentation Index

### Getting Started

1. **[Quick Start Guide](QUICK_START.md)** ⚡
   - Platform-specific installation instructions
   - Download and run in 5 minutes
   - Basic configuration for Tabnine
   - **Start here if you're new!**

2. **[Platform Support](PLATFORM_SUPPORT.md)** 🖥️
   - Supported platforms and architectures
   - Platform-specific requirements and installation
   - Running as a service (systemd, launchd, Windows Service)
   - Platform-specific troubleshooting

3. **[Usage Guide](USAGE.md)** 📖
   - How to use the server
   - API examples and endpoints
   - Configuration options
   - Advanced usage patterns

### Building and Development

4. **[Building from Source](BUILDING.md)** 🔨
   - Prerequisites and dependencies
   - Platform-specific build instructions
   - Building for all platforms
   - Build configuration and customization
   - Troubleshooting build issues

5. **[Multi-Platform Update](MULTI_PLATFORM_UPDATE.md)** 🌐
   - Summary of multi-platform support
   - What changed in the build system
   - Migration guide
   - Binary naming conventions
   - Quick reference

### Integration

6. **[Tabnine Setup Guide](TABNINE_SETUP.md)** 🤖
   - Complete Tabnine integration guide
   - Configuration examples
   - Testing and verification
   - Common issues and solutions

7. **[Tabnine Quick Reference](TABNINE_QUICK_REFERENCE.md)** 📝
   - Quick configuration snippets
   - Common commands
   - Troubleshooting cheatsheet

8. **[MCP Introduction](mcp-introduction.md)** 🔌
   - Understanding Model Context Protocol
   - How MCP works
   - Protocol specifications

9. **[MCP Examples](mcp-examples.md)** 💡
   - Example MCP requests and responses
   - Use cases
   - Integration patterns

### CI/CD and Releases

10. **[Workflows Documentation](WORKFLOWS.md)** ⚙️
    - GitHub Actions workflows
    - Build and release processes
    - Matrix build strategy
    - Platform-specific build details
    - Troubleshooting CI/CD

11. **[Release Guide](RELEASE_GUIDE.md)** 🚀
    - Creating releases
    - Pre-release checklist
    - Release types (stable, pre-release, snapshot)
    - Post-release tasks
    - Rollback procedures

12. **[Setup Checklist](SETUP_CHECKLIST.md)** ✅
    - Repository setup checklist
    - CI/CD configuration
    - Release preparation

### Project Information

13. **[Changes / Changelog](CHANGES.md)** 📋
    - Version history
    - What's new in each release
    - Breaking changes

14. **[GitHub Setup](GITHUB_SETUP.md)** 🔧
    - Repository configuration
    - GitHub Actions setup
    - Permissions and secrets

15. **[Workflow Diagrams](WORKFLOW_DIAGRAM.md)** 📊
    - Visual workflow representations
    - Architecture diagrams

16. **[Workflow Summary](WORKFLOW_SUMMARY.md)** 📑
    - Quick workflow overview
    - Key processes summary

## 🗺️ Documentation Map

### By User Type

#### 👤 End Users
1. [Quick Start Guide](QUICK_START.md) - Get up and running
2. [Platform Support](PLATFORM_SUPPORT.md) - Choose your platform
3. [Usage Guide](USAGE.md) - Learn to use the server
4. [Tabnine Setup](TABNINE_SETUP.md) - Integrate with Tabnine

#### 👨‍💻 Developers
1. [Building from Source](BUILDING.md) - Build your own binaries
2. [MCP Introduction](mcp-introduction.md) - Understand the protocol
3. [MCP Examples](mcp-examples.md) - See example code
4. [Usage Guide](USAGE.md) - Advanced usage

#### 🔧 Maintainers
1. [Release Guide](RELEASE_GUIDE.md) - Create releases
2. [Workflows Documentation](WORKFLOWS.md) - Manage CI/CD
3. [Multi-Platform Update](MULTI_PLATFORM_UPDATE.md) - Build system details
4. [GitHub Setup](GITHUB_SETUP.md) - Configure repository

### By Task

#### 📥 Installing
- New user? → [Quick Start Guide](QUICK_START.md)
- Specific platform? → [Platform Support](PLATFORM_SUPPORT.md)
- Running as service? → [Platform Support](PLATFORM_SUPPORT.md#running-as-a-service)

#### 🔨 Building
- First time building? → [Building from Source](BUILDING.md)
- Build for all platforms? → [Building - Build for All Platforms](BUILDING.md#build-for-all-platforms)
- Build issues? → [Building - Troubleshooting](BUILDING.md#troubleshooting)

#### 🤝 Integrating
- Using Tabnine? → [Tabnine Setup Guide](TABNINE_SETUP.md)
- Understanding MCP? → [MCP Introduction](mcp-introduction.md)
- Need examples? → [MCP Examples](mcp-examples.md)

#### 🚀 Releasing
- Creating a release? → [Release Guide](RELEASE_GUIDE.md)
- CI/CD issues? → [Workflows Documentation](WORKFLOWS.md)
- Understanding builds? → [Multi-Platform Update](MULTI_PLATFORM_UPDATE.md)

## 🆕 New to WebSearch MCP?

**Recommended reading order:**

1. **[Quick Start Guide](QUICK_START.md)** - Get it running (5 minutes)
2. **[Platform Support](PLATFORM_SUPPORT.md)** - Understand your platform (10 minutes)
3. **[Tabnine Setup Guide](TABNINE_SETUP.md)** - Integrate with Tabnine (15 minutes)
4. **[Usage Guide](USAGE.md)** - Learn advanced features (as needed)

## 🔍 Quick Find

### Common Questions

**How do I download and run the server?**
→ [Quick Start Guide](QUICK_START.md)

**Which binary should I download?**
→ [Platform Support](PLATFORM_SUPPORT.md#supported-platforms)

**How do I build from source?**
→ [Building from Source](BUILDING.md)

**How do I integrate with Tabnine?**
→ [Tabnine Setup Guide](TABNINE_SETUP.md)

**How do I create a release?**
→ [Release Guide](RELEASE_GUIDE.md)

**What platforms are supported?**
→ [Platform Support](PLATFORM_SUPPORT.md#supported-platforms)

**What changed in the multi-platform update?**
→ [Multi-Platform Update](MULTI_PLATFORM_UPDATE.md)

**How do GitHub Actions workflows work?**
→ [Workflows Documentation](WORKFLOWS.md)

**What's the MCP protocol?**
→ [MCP Introduction](mcp-introduction.md)

**Where are the API examples?**
→ [MCP Examples](mcp-examples.md)

## 📊 Documentation Statistics

- **Total Documents**: 17
- **Getting Started Guides**: 4
- **Technical Guides**: 7
- **Reference Docs**: 6
- **Total Pages**: ~150+ (estimated)

## 🤝 Contributing to Documentation

Found an issue or want to improve the docs?

1. **Typos/Errors**: Open an issue or PR
2. **New Content**: Discuss in an issue first
3. **Examples**: Always welcome! Open a PR
4. **Clarifications**: Open an issue to discuss

### Documentation Standards

- Use clear, concise language
- Include code examples where helpful
- Provide platform-specific instructions when relevant
- Add links to related documentation
- Keep tables of contents updated

## 📞 Getting Help

If you can't find what you're looking for:

1. **Search the docs** - Use your browser's search (Cmd/Ctrl+F)
2. **Check the FAQ** - Each guide has common questions
3. **GitHub Issues** - Search existing issues
4. **Open an Issue** - Describe what you're trying to do

## 🔄 Documentation Updates

Documentation is updated with each release. Check the [Changes](CHANGES.md) file for documentation updates.

**Last Updated**: 2024-01-15
**Version**: 1.0.0 (multi-platform support)

## 📝 Document Maintenance

| Document | Last Updated | Status | Maintainer |
|----------|--------------|--------|------------|
| Quick Start | 2024-01-15 | Current | Core team |
| Platform Support | 2024-01-15 | Current | Core team |
| Building | 2024-01-15 | Current | Core team |
| Usage Guide | 2024-01-15 | Current | Core team |
| Tabnine Setup | 2024-01-15 | Current | Core team |
| Workflows | 2024-01-15 | Current | Core team |
| Release Guide | 2024-01-15 | Current | Core team |
| Multi-Platform Update | 2024-01-15 | Current | Core team |
| MCP Introduction | Earlier | Current | Core team |
| MCP Examples | Earlier | Current | Core team |

## 🌟 Featured Documentation

### Must-Read for New Users
1. [Quick Start Guide](QUICK_START.md)
2. [Platform Support](PLATFORM_SUPPORT.md)
3. [Tabnine Setup Guide](TABNINE_SETUP.md)

### Must-Read for Developers
1. [Building from Source](BUILDING.md)
2. [MCP Introduction](mcp-introduction.md)
3. [MCP Examples](mcp-examples.md)

### Must-Read for Maintainers
1. [Release Guide](RELEASE_GUIDE.md)
2. [Workflows Documentation](WORKFLOWS.md)
3. [Multi-Platform Update](MULTI_PLATFORM_UPDATE.md)

---

**Navigate**: [Back to Main README](../README.md) | [Project Repository](https://github.com/youruser/websearch-mcp)
