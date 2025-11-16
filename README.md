# Codebase Interface CLI

A command line interface (CLI) for validating codebase structure and development standards. This tool helps ensure your projects follow best practices for essential files, Git configuration, and development workflows.

## Features

- **Essential Files Validation**: Checks for README.md, CONTRIBUTING.md
- **Git Configuration**: Validates .gitignore, .editorconfig, .gitattributes  
- **Development Standards**: Enforces conventional commits and branch naming
- **JSON Schema Validation**: Robust configuration file validation with detailed error reporting
- **Configuration Presets**: Quick setup with 5 built-in templates (basic, strict, beginner, open-source, go-project)
- **Multiple Output Formats**: Table (styled) and JSON output
- **Configurable**: YAML configuration file support with schema validation
- **Extensible**: Modular agent architecture for easy expansion

## Installation

### 🚀 Quick Installation (Recommended)

**Linux/macOS:**
```bash
curl -sSL https://raw.githubusercontent.com/codebase-interface/cli/main/install.sh | bash
```

**Windows (PowerShell):**
```powershell
Invoke-WebRequest -Uri "https://github.com/codebase-interface/cli/releases/latest/download/codebase-interface-windows-amd64.exe" -OutFile "codebase-interface.exe"
```

### 📦 Package Managers

**Homebrew (macOS/Linux):**
```bash
brew tap codebase-interface/cli
brew install codebase-interface
```

**Chocolatey (Windows):**
```powershell
choco install codebase-interface
```

**Go Install:**
```bash
go install github.com/codebase-interface/cli/cmd/codebase-interface@latest
```

### 🐳 Container

**Docker:**
```bash
# Run validation in current directory
docker run --rm -v $(pwd):/workspace ghcr.io/codebase-interface/cli:latest validate

# Interactive mode
docker run --rm -it -v $(pwd):/workspace ghcr.io/codebase-interface/cli:latest
```

### 📥 Pre-built Binaries

Download from [GitHub Releases](https://github.com/codebase-interface/cli/releases/latest):

- **Linux**: `codebase-interface-linux-amd64`
- **macOS Intel**: `codebase-interface-darwin-amd64`  
- **macOS Apple Silicon**: `codebase-interface-darwin-arm64`
- **Windows**: `codebase-interface-windows-amd64.exe`

### 🛠️ Build from Source

```bash
git clone https://github.com/codebase-interface/cli.git
cd cli
task build  # Creates binary in bin/

# Install globally
task install  # Installs to $GOPATH/bin
```

## Quick Start

### ✅ Verify Installation
```bash
# Check if installed correctly
codebase-interface version
# Or use the short alias
cbi version
```

### 🎯 Initial Setup
```bash
# Create a configuration file for your project
cbi init-config basic

# Validate your configuration
cbi validate-config

# Run validation
cbi validate
```

### 🔧 Common Usage
```bash
# Validate current directory
cbi validate

# Validate specific path  
cbi validate --path /path/to/project

# JSON output for CI/CD
cbi validate --output json

# Run specific validation agent
cbi validate --agent essential-files

# Get help
cbi --help
```

## Validation Agents

### 1. Essential Files Agent

Validates presence of fundamental project files:

- **README.md** or **README.rst** - Project documentation
- **CONTRIBUTING.md** - Contribution guidelines

### 2. Git Configuration Agent  

Validates Git configuration files:

- **.gitignore** - Files to ignore in Git
- **.editorconfig** - Consistent coding styles
- **.gitattributes** - Git attributes (optional)

### 3. Development Standards Agent

Validates development workflow standards:

- **Conventional Commits** - Commit message format
- **Branch Naming** - Branch naming conventions

## Configuration

### Quick Setup

Create a configuration file with preset templates:

```bash
# Create basic configuration
codebase-interface init-config basic

# Or choose a specific template
codebase-interface init-config go-project
codebase-interface init-config open-source
codebase-interface init-config strict
```

### Manual Configuration

Create `.codebase-validation.yml` in your project root:

```yaml
validation:
  agents:
    essential-files:
      enabled: true
      require_readme: true
      require_contributing: true
      
    git-configuration:
      enabled: true
      require_gitignore: true
      require_gitattributes: false  # Optional
      require_editorconfig: true
      
    development-standards:
      enabled: true
      check_commit_history: true
      commit_history_depth: 10
      require_conventional_commits: true

  output:
    format: "table"  # json, table
    verbose: false
```

## Commands

### validate-config

Validates configuration files against the JSON schema.

```bash
# Validate default configuration file
codebase-interface validate-config

# Validate specific file
codebase-interface validate-config my-config.yml

# Validate configuration in a specific directory
codebase-interface validate-config /path/to/project
```

### schema

Get the JSON schema for configuration validation and editor integration.

```bash
# Display schema to stdout
codebase-interface schema

# Save schema to file for editor integration
codebase-interface schema --output codebase-validation.schema.json

# Use shorter alias
cbi schema -o schema.json
```

### init-config

Creates a new configuration file with preset templates.

```bash
codebase-interface init-config [type]

# Available types: basic, strict, beginner, open-source, go-project
codebase-interface init-config basic
codebase-interface init-config --force  # Overwrite existing
```

### schema

Display or save the JSON schema for configuration validation.

```bash
codebase-interface schema                    # Display schema
codebase-interface schema -o schema.json    # Save to file
```

### validate

Validates codebase structure and standards.

```bash
codebase-interface validate [flags]

Flags:
  -a, --agent string    Run specific agent (essential-files, git-configuration, development-standards)
  -o, --output string   Output format (json, table) (default "table")
  -p, --path string     Path to validate (default ".")
```

### version

Print version information.

```bash
codebase-interface version
```

## Development

### Prerequisites

- Go 1.21 or later
- [Task](https://taskfile.dev/) for task automation

### Development Tasks

```bash
# Setup development environment
task setup

# Build the CLI
task build

# Run tests
task test

# Run tests in watch mode
task test:watch

# Run linting
task lint

# Install locally
task install

# Validate the CLI project itself
task validate-self

# Serve documentation locally
task docs:serve

# Open documentation in browser
task docs:open

# Check documentation for issues
# Check documentation for issues\ntask docs:check\n```bash
```

### Architecture

```txt
codebase-cli/
├── cmd/                    # Cobra CLI commands
│   ├── root.go            # Root command
│   ├── validate.go        # Validation command
│   └── version.go         # Version command
├── internal/              # Internal packages
│   ├── agents/           # Validation agents
│   ├── config/           # Configuration handling
│   ├── output/           # Output formatters
│   └── ui/               # TUI components
├── pkg/                   # Public API packages
├── test/                  # Test files and fixtures
├── Taskfile.yml          # Development tasks
└── .codebase-validation.yml  # Self-validation config
```

## Technology Stack

- **Language**: Go (latest stable)
- **CLI Framework**: [Cobra](https://github.com/spf13/cobra)
- **Styling**: [Lipgloss](https://github.com/charmbracelet/lipgloss)
- **Configuration**: YAML with [yaml.v3](https://gopkg.in/yaml.v3)
- **Testing**: Go's built-in testing with table-driven tests
- **Development**: Test-driven development (TDD)

## Contributing

We welcome contributions! See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines on:

- Setting up the development environment
- Code style and conventions
- Testing requirements
- Pull request process

## Examples

### Successful Validation

```text
✓ Essential Files Agent - PASS (Score: 1.0)
  ✓ README.md present
  ✓ CONTRIBUTING.md present

✓ Git Configuration Agent - PASS (Score: 1.0)
  ✓ .gitignore present
  ✓ .editorconfig present

✓ Development Standards Agent - PASS (Score: 1.0)
  ✓ Recent commits follow conventional format
  ✓ Branch naming follows conventions: main

Overall Score: 1.00 - PASS
```

### Failed Validation

```text
✗ Essential Files Agent - FAIL (Score: 0.5)
  ✓ README.md present
  ✗ CONTRIBUTING.md missing

Overall Score: 0.50 - FAIL
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
