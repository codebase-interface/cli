# Package Manager Setup Guide

This guide explains how to set up the automated package publishing for Homebrew and Chocolatey.

## 🍺 Homebrew Setup

### 1. Create Homebrew Tap Repository

First, you need to create a separate repository for the Homebrew tap:

```bash
# Create a new repository named 'homebrew-cli'
# This should be done through GitHub UI or API
# Repository should be: codebase-interface/homebrew-cli
```

The repository structure will look like:
```
homebrew-cli/
└── Formula/
    └── codebase-interface.rb
```

### 2. Set Up GitHub Token

Create a GitHub token with repository access:

1. Go to GitHub Settings → Developer settings → Personal access tokens
2. Generate a new token (classic) with `repo` permissions
3. Add it as a repository secret named `HOMEBREW_TAP_GITHUB_TOKEN`

### 3. How It Works

When you create a release:

1. **GoReleaser** automatically creates a Homebrew formula
2. **Formula** gets pushed to `codebase-interface/homebrew-cli` 
3. **Users** can then install with:
   ```bash
   brew tap codebase-interface/cli
   brew install codebase-interface
   ```

The formula will be generated automatically and look like:
```ruby
class CodebaseInterface < Formula
  desc "A CLI tool for validating codebase structure and development standards"
  homepage "https://github.com/codebase-interface/cli"
  url "https://github.com/codebase-interface/cli/releases/download/v1.0.0/codebase-interface-Darwin-x86_64.tar.gz"
  sha256 "..."
  license "MIT"

  def install
    bin.install "codebase-interface"
    bin.install_symlink bin/"codebase-interface" => "cbi"
  end

  test do
    system "#{bin}/codebase-interface", "version"
    system "#{bin}/cbi", "version"
  end
end
```

## 🍫 Chocolatey Setup

### 1. Create Chocolatey Account

1. Create account at [chocolatey.org](https://chocolatey.org/)
2. Get your API key from your account profile
3. Add it as a repository secret named `CHOCOLATEY_API_KEY`

### 2. How It Works

When you create a release:

1. **GoReleaser** creates a Chocolatey package (.nupkg)
2. **Package** gets automatically uploaded to Chocolatey repository
3. **Users** can then install with:
   ```powershell
   choco install codebase-interface
   ```

The package will include:
- Windows executable
- PowerShell install/uninstall scripts
- Automatic PATH configuration
- Both `codebase-interface.exe` and `cbi.exe` commands

## 🔧 Required Repository Secrets

Add these secrets to your GitHub repository:

```bash
# For Homebrew tap publishing
HOMEBREW_TAP_GITHUB_TOKEN=github_pat_...

# For Chocolatey package publishing  
CHOCOLATEY_API_KEY=...

# Already configured for GitHub releases
GITHUB_TOKEN=... (automatic)
```

## 🎯 Testing Package Managers

### Test Homebrew Locally

```bash
# Install from your tap
brew tap codebase-interface/cli
brew install codebase-interface

# Verify installation
codebase-interface version
cbi version

# Test upgrade path
brew upgrade codebase-interface
```

### Test Chocolatey Locally

```powershell
# Install from Chocolatey
choco install codebase-interface

# Verify installation  
codebase-interface version
cbi version

# Test upgrade path
choco upgrade codebase-interface
```

## 📋 Release Checklist

Before releasing with package manager support:

- [ ] Create `codebase-interface/homebrew-cli` repository
- [ ] Add `HOMEBREW_TAP_GITHUB_TOKEN` secret 
- [ ] Add `CHOCOLATEY_API_KEY` secret
- [ ] Test GoReleaser configuration locally:
  ```bash
  goreleaser check
  goreleaser release --snapshot --clean
  ```
- [ ] Create a test release with a pre-release tag
- [ ] Verify packages are published correctly

## 🚀 First Release

For your first release with package managers:

1. **Create the required repositories and secrets**
2. **Tag your release:**
   ```bash
   git tag v1.0.0
   git push --tags
   ```
3. **Monitor the GitHub Action** to ensure all steps complete
4. **Test installation** from both package managers
5. **Update documentation** with installation instructions

The release process will:
- ✅ Build binaries for all platforms  
- ✅ Create GitHub release with assets
- ✅ Publish Docker images to GHCR
- ✅ Create Homebrew formula in tap repository
- ✅ Publish Chocolatey package to chocolatey.org

After the first successful release, users will be able to install your CLI using standard package managers!