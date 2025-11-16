# 📥 Installation Guide

**Get up and running with Codebase Interface CLI in minutes!**

This guide covers all the ways to install the CLI, from one-line installers to building from source.

## 🚀 Recommended Installation

### One-Line Installation

**Linux/macOS:**
```bash
curl -sSL https://raw.githubusercontent.com/codebase-interface/cli/main/install.sh | bash
```

**Windows (PowerShell):**
```powershell
Invoke-WebRequest -Uri "https://github.com/codebase-interface/cli/releases/latest/download/codebase-interface-windows-amd64.exe" -OutFile "codebase-interface.exe"
# Add to PATH or move to a directory in your PATH
```

### ✅ Verify Installation

```bash
# Check if installed correctly
codebase-interface version
# Or use the short alias
cbi version
```

You should see output like:
```
Codebase Interface CLI v1.0.0
```

## 📦 Package Managers

### Homebrew (macOS/Linux)

```bash
# Add the tap
brew tap codebase-interface/cli

# Install the CLI
brew install codebase-interface

# Verify installation
cbi version
```

**Benefits:**
- ✅ Automatic updates with `brew upgrade`
- ✅ Easy uninstall with `brew uninstall codebase-interface`
- ✅ Manages dependencies automatically

### Chocolatey (Windows)

```powershell
# Install the CLI
choco install codebase-interface

# Verify installation
cbi version
```

**Benefits:**
- ✅ Windows-native package management
- ✅ Automatic PATH configuration
- ✅ Easy updates with `choco upgrade codebase-interface`

### Go Install

If you have Go installed:

```bash
go install github.com/codebase-interface/cli/cmd/codebase-interface@latest

# Verify installation (ensure $GOPATH/bin is in your PATH)
cbi version
```

**Benefits:**
- ✅ Always gets the latest version
- ✅ Works on any platform with Go
- ✅ Minimal dependencies

## 🐳 Container Usage

### Docker

**One-time validation:**
```bash
# Validate current directory
docker run --rm -v $(pwd):/workspace ghcr.io/codebase-interface/cli:latest validate

# Create configuration file
docker run --rm -v $(pwd):/workspace ghcr.io/codebase-interface/cli:latest init-config basic
```

**Interactive session:**
```bash
# Launch interactive shell with CLI available
docker run --rm -it -v $(pwd):/workspace ghcr.io/codebase-interface/cli:latest bash

# Inside container, use CLI normally
cbi validate
cbi init-config open-source
```

**CI/CD Integration:**
```yaml
# GitHub Actions example
- name: Validate Codebase
  run: |
    docker run --rm -v ${{ github.workspace }}:/workspace \
      ghcr.io/codebase-interface/cli:latest validate --output json
```

**Benefits:**
- ✅ No local installation required
- ✅ Consistent environment across teams
- ✅ Perfect for CI/CD pipelines
- ✅ Isolated from host system

## 💾 Pre-built Binaries

Download directly from [GitHub Releases](https://github.com/codebase-interface/cli/releases/latest):

### Download Commands

**Linux (x64):**
```bash
curl -L https://github.com/codebase-interface/cli/releases/latest/download/codebase-interface-linux-amd64 -o codebase-interface
chmod +x codebase-interface
sudo mv codebase-interface /usr/local/bin/
```

**macOS (Intel):**
```bash
curl -L https://github.com/codebase-interface/cli/releases/latest/download/codebase-interface-darwin-amd64 -o codebase-interface
chmod +x codebase-interface
sudo mv codebase-interface /usr/local/bin/
```

**macOS (Apple Silicon):**
```bash
curl -L https://github.com/codebase-interface/cli/releases/latest/download/codebase-interface-darwin-arm64 -o codebase-interface
chmod +x codebase-interface
sudo mv codebase-interface /usr/local/bin/
```

**Windows:**
- Download `codebase-interface-windows-amd64.exe`
- Rename to `codebase-interface.exe`
- Add to your PATH or place in a directory that's already in PATH

### Create Short Alias

```bash
# Linux/macOS - Add to ~/.bashrc or ~/.zshrc
echo 'alias cbi="codebase-interface"' >> ~/.bashrc
source ~/.bashrc

# Or create a symlink
sudo ln -sf /usr/local/bin/codebase-interface /usr/local/bin/cbi
```

## 🛠️ Build from Source

### Prerequisites

- **Go 1.21+** - [Install Go](https://golang.org/doc/install)
- **Git** - For cloning the repository
- **Task** (optional) - [Install Task](https://taskfile.dev/installation/) for easier building

### Build Steps

```bash
# Clone the repository
git clone https://github.com/codebase-interface/cli.git
cd cli

# Option 1: Using Task (recommended)
task build

# Option 2: Using Go directly
go build -o bin/codebase-interface ./cmd/codebase-interface

# Option 3: Install globally
task install
# Or: go install ./cmd/codebase-interface
```

### Development Build

```bash
# Build for development with debug info
task build

# Run tests
task test

# Build for all platforms
task build:all

# Package for distribution
task package
```

## 🔧 Post-Installation Setup

### 1. Verify Installation

```bash
# Check version and available commands
cbi --help
cbi version
```

### 2. Create Your First Configuration

```bash
# Navigate to your project
cd /path/to/your/project

# Create a basic configuration
cbi init-config basic

# Validate the configuration
cbi validate-config

# Run your first validation
cbi validate
```

### 3. Customize for Your Project

```bash
# Try different presets
cbi init-config open-source  # For open source projects
cbi init-config go-project   # For Go projects
cbi init-config strict       # For production codebases

# Edit the configuration file
nano .codebase-validation.yml
```

## 🚨 Troubleshooting

### Command Not Found

**Issue:** `cbi: command not found`

**Solutions:**
1. **Check PATH:** Ensure the installation directory is in your PATH
   ```bash
   echo $PATH
   which cbi
   ```

2. **Reinstall using package manager:**
   ```bash
   # Homebrew
   brew reinstall codebase-interface
   
   # Or try the installation script again
   curl -sSL https://raw.githubusercontent.com/codebase-interface/cli/main/install.sh | bash
   ```

3. **Manual PATH addition:**
   ```bash
   # Add to ~/.bashrc or ~/.zshrc
   export PATH="/usr/local/bin:$PATH"
   source ~/.bashrc
   ```

### Permission Denied

**Issue:** `Permission denied` when running installation

**Solution:**
```bash
# Use sudo for system-wide installation
sudo curl -sSL https://raw.githubusercontent.com/codebase-interface/cli/main/install.sh | bash

# Or install to user directory
curl -sSL https://raw.githubusercontent.com/codebase-interface/cli/main/install.sh | bash -s -- ~/bin
```

### Docker Issues

**Issue:** Docker permission errors

**Solution:**
```bash
# Add user to docker group (Linux)
sudo usermod -aG docker $USER
# Then log out and back in

# Or use sudo
sudo docker run --rm -v $(pwd):/workspace ghcr.io/codebase-interface/cli:latest validate
```

## 🔄 CI/CD Integration

### GitHub Actions

Add validation to your GitHub workflow:

```yaml
name: Validate Codebase

on:
  push:
    branches: [ main, develop ]
  pull_request:
    branches: [ main ]

jobs:
  validate:
    runs-on: ubuntu-latest
    
    steps:
    - name: Checkout code
      uses: actions/checkout@v4
      
    - name: Validate codebase structure
      run: |
        docker run --rm -v ${{ github.workspace }}:/workspace \
          ghcr.io/codebase-interface/cli:latest validate --output json
```

### GitLab CI

Add to your `.gitlab-ci.yml`:

```yaml
validate_codebase:
  image: ghcr.io/codebase-interface/cli:latest
  stage: test
  script:
    - codebase-interface validate --output json
  rules:
    - if: $CI_PIPELINE_SOURCE == "push"
    - if: $CI_PIPELINE_SOURCE == "merge_request_event"
```

### Jenkins Pipeline

```groovy
pipeline {
    agent any
    
    stages {
        stage('Validate Codebase') {
            steps {
                script {
                    docker.image('ghcr.io/codebase-interface/cli:latest').inside('-v ${WORKSPACE}:/workspace') {
                        sh 'codebase-interface validate --output json'
                    }
                }
            }
        }
    }
}
```

### Azure Pipelines

```yaml
trigger:
- main
- develop

pool:
  vmImage: 'ubuntu-latest'

steps:
- task: Docker@2
  displayName: 'Validate Codebase'
  inputs:
    command: 'run'
    arguments: '--rm -v $(Build.SourcesDirectory):/workspace ghcr.io/codebase-interface/cli:latest validate --output json'
```

## 📱 Platform-Specific Notes

### macOS
- **Gatekeeper:** First run might require going to Security & Privacy settings to allow the binary
- **Homebrew recommended** for easiest installation and updates
- **Apple Silicon** users should use the ARM64 binary for better performance

### Windows
- **PowerShell recommended** for best experience
- **Windows Defender** might flag the binary initially - this is normal for new binaries
- **Chocolatey** provides the smoothest Windows experience
- **WSL users** can use the Linux installation methods

### Linux
- **Package managers vary** - the installation script auto-detects your system
- **Snap package** coming soon for Ubuntu users
- **AppImage** available for distribution-agnostic installation

## 🎯 Next Steps

After installation, check out:

1. **[Usage Guide](usage.md)** - Learn the core commands and workflows
2. **[Configuration Guide](configuration.md)** - Customize validation rules
3. **[Examples](examples/)** - Copy-paste configurations for your project type

---

*💡 **Tip:** Bookmark this page! Installation is just the beginning of your journey toward better codebase organization.*