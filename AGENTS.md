# Codebase Validation Agents

This document defines the validation agents for the codebase interface CLI. These agents validate essential files and configurations for proper codebase setup.

## Technical Implementation

### Technology Stack

- **Language:** Go (latest stable version)
- **CLI Framework:** [Cobra](https://github.com/spf13/cobra) - Standard Go CLI framework with command structure and flags
- **TUI Framework:** [Bubble Tea](https://github.com/charmbracelet/bubbletea) - Modern terminal UI for interactive validation reports
- **Schema Validation:** [JSON Schema](https://json-schema.org/draft/2020-12/schema) with [gojsonschema](https://github.com/xeipuuv/gojsonschema) - Comprehensive configuration validation
- **Task Runner:** [Taskfile](https://taskfile.dev/) - Task runner for development workflows
- **Testing:** Test-driven development (TDD) approach with Go's built-in testing framework

### Development Principles

- **Test-Driven Development (TDD)** - All features must be developed with tests first
- **Go Best Practices** - Follow Go idioms, effective Go guidelines, and community standards
- **Documentation-Driven** - Maintain comprehensive README.md, CONTRIBUTING.md, AGENTS.md and detailed docs/ directory
- **Task Automation** - Use Taskfile for all development interactions (build, test, lint, release)
- **Conventional Commits** - Follow [Conventional Commits](https://www.conventionalcommits.org/) specification for commit messages

### Architecture Requirements

- Modular agent system for extensibility
- Clean separation between validation logic and UI presentation
- Configuration-driven validation rules with JSON Schema validation
- Embedded schema system for offline validation and editor integration
- Configuration presets and templates for quick setup
- Structured JSON output for programmatic use
- Interactive TUI for human-friendly reports

## Validation Category

### MVP (Minimal Viable Project)

The minimal requirements for any codebase to be considered properly structured.

---

## Agent Definitions

### 1. Essential Files Agent

**Purpose:** Validates presence of fundamental project files  
**Priority:** Critical

#### Validation Rules

- **README.md** or **README.rst** - Project documentation root
- **CONTRIBUTING.md** - Contribution guidelines
- **docs/** directory - Comprehensive documentation structure with CLI usage guides

#### Output Format

```json
{
  "agent": "essential-files",
  "status": "pass|fail|warning",
  "score": 0.0-1.0,
  "findings": [
    {
      "type": "missing|present|invalid",
      "file": "README.md",
      "message": "Description of finding",
      "severity": "critical|warning|info"
    }
  ]
}
```

### 2. Git Configuration Agent

**Purpose:** Validates Git configuration files  
**Priority:** High

#### Git Configuration Files

- **.gitignore** - Specifies files Git should ignore
- **.gitattributes** - Defines attributes for paths (line endings, merge strategies)
- **.editorconfig** - Maintains consistent coding styles across editors

#### Git Configuration Output Format

```json
{
  "agent": "git-configuration",
  "status": "pass|fail|warning",
  "score": 0.0-1.0,
  "findings": [
    {
      "type": "missing|present|invalid",
      "file": ".gitignore",
      "message": "Description of finding",
      "severity": "critical|warning|info"
    }
  ]
}
```

### 3. Development Standards Agent

**Purpose:** Validates development workflow standards  
**Priority:** Medium

#### Development Standards

- **Conventional Commits** - Commit messages follow conventional commits specification
- **Commit History** - Recent commits adhere to conventional format (feat:, fix:, docs:, etc.)
- **Branch Naming** - Branch names follow conventional patterns (feature/, fix/, hotfix/)

---

## Development Workflow

### Project Structure

```text
codebase-interface-cli/
├── cmd/                    # Cobra CLI commands
│   ├── codebase-interface/ # Main CLI entry point
│   ├── root.go            # Root command with alias support
│   ├── validate.go        # Validation command
│   ├── validate-config.go # Configuration validation command
│   ├── init-config.go     # Configuration preset generator
│   ├── schema.go          # Schema export command
│   ├── schema/            # Embedded JSON schema
│   │   └── codebase-validation.schema.json
│   └── version.go         # Version command
├── internal/              # Internal packages
│   ├── agents/           # Validation agents
│   ├── config/           # Configuration handling with schema validation
│   ├── output/           # Output formatters
│   └── ui/               # Bubble Tea TUI components
├── pkg/                   # Public API packages
├── test/                  # Test files and fixtures
├── docs/                  # Documentation directory (MkDocs)
│   ├── README.md         # Documentation overview
│   ├── usage.md          # Detailed CLI usage instructions
│   ├── configuration.md  # Configuration reference with schema docs
│   ├── agents.md         # Agent documentation
│   └── examples/         # Example configurations and use cases
├── .vscode/               # VS Code workspace settings
│   └── settings.json     # YAML schema integration
├── bin/                   # Compiled binaries (git-ignored)
├── Taskfile.yml          # Development tasks with MkDocs integration
├── mkdocs.yml            # Documentation site configuration
├── README.md             # Project documentation
├── CONTRIBUTING.md       # Development guidelines
└── go.mod                # Go module definition
```

### Required Development Tasks (Taskfile.yml)

```yaml
tasks:
  build:
    desc: Build the CLI binary
    cmds:
      - go build -o bin/codebase-interface ./cmd/codebase-interface

  test:
    desc: Run all tests
    cmds:
      - go test ./...

  test:watch:
    desc: Run tests in watch mode
    cmds:
      - gotestsum --watch

  lint:
    desc: Run linting
    cmds:
      - golangci-lint run

  validate-schema:
    desc: Validate JSON schema and test configuration validation
    cmds:
      - ./bin/codebase-interface validate-config
      - ./bin/codebase-interface schema --output /tmp/schema-test.json

  docs:serve:
    desc: Serve documentation locally
    cmds:
      - mkdocs serve

  docs:build:
    desc: Build documentation site
    cmds:
      - mkdocs build

  tidy:
    desc: Tidy Go modules
    cmds:
      - go mod tidy

  install:
    desc: Install CLI locally
    cmds:
      - go install ./cmd/codebase-interface
```

### Test-Driven Development Requirements

- **Unit Tests:** Each agent must have comprehensive unit tests
- **Integration Tests:** End-to-end validation scenarios
- **Table-Driven Tests:** Use Go's table-driven test patterns
- **Test Coverage:** Minimum 80% code coverage
- **Test Fixtures:** Provide sample project structures for testing

### Documentation Requirements

- **docs/README.md:** Overview of the documentation structure and how to navigate it
- **docs/usage.md:** Comprehensive CLI usage guide with examples for all commands and flags
- **docs/configuration.md:** Complete reference for .codebase-validation.yml configuration options
- **docs/agents.md:** Detailed documentation of each validation agent, their rules, and customization options
- **docs/examples/:** Real-world examples showing different project setups and configurations
- **API Documentation:** Generated documentation for public packages in pkg/

---

## Validation Execution

### Command Structure

```bash
# Validate all essential files
codebase-interface validate

# Validate specific agent
codebase-interface validate --agent essential-files
codebase-interface validate --agent git-configuration
codebase-interface validate --agent development-standards

# Output formats
codebase-interface validate --output json
codebase-interface validate --output table

# Configuration management
codebase-interface init-config basic           # Create configuration from preset
codebase-interface validate-config            # Validate configuration against schema
codebase-interface schema --output schema.json # Export schema for editor integration

# Short alias available for all commands
cbi validate
cbi init-config go-project
cbi validate-config
```

### Exit Codes

- **0** - All validations passed
- **1** - Critical validations failed
- **2** - Warnings present but no critical failures

### Scoring System

- Each agent returns a score between 0.0 and 1.0
- Overall score is averaged across all agents

---

## Configuration

### Configuration Validation & Management

The CLI provides comprehensive configuration management with JSON Schema validation:

#### Schema Validation Features

- **Type Safety:** Validates data types (boolean, string, number, array)
- **Range Validation:** Ensures numeric values are within valid ranges (0.0-1.0)
- **Enum Validation:** Checks string values against allowed options
- **Property Validation:** Detects typos and unknown configuration keys
- **Required Field Validation:** Ensures all mandatory fields are present
- **Editor Integration:** Provides autocomplete and real-time validation in IDEs

#### Configuration Presets

```bash
# Available preset configurations
cbi init-config basic        # Simple validation rules
cbi init-config strict       # Comprehensive validation for production
cbi init-config beginner     # Gentle validation for learning projects
cbi init-config open-source  # Optimized for public repositories
cbi init-config go-project   # Go-specific validation rules
```

#### Schema Integration

```yaml
# Enable editor autocomplete and validation
# yaml-language-server: $schema=https://raw.githubusercontent.com/codebase-interface/cli/refs/heads/main/cmd/schema/codebase-validation.schema.json
```

### Default Configuration File: `.codebase-validation.yml`

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
      commit_history_depth: 10  # Check last 10 commits
      require_conventional_commits: true

  output:
    format: "table"  # json, table
    verbose: false
```

---

## Examples

### Successful Validation

```bash
codebase-cli validate
```

Expected output:

```text
✓ Essential Files Agent - PASS (Score: 1.0)
  ✓ README.md present
  ✓ CONTRIBUTING.md present

✓ Git Configuration Agent - PASS (Score: 1.0)
  ✓ .gitignore present
  ✓ .editorconfig present
  ℹ .gitattributes not found (optional)

✓ Development Standards Agent - PASS (Score: 1.0)
  ✓ Recent commits follow conventional format
  ✓ Branch naming follows conventions

Overall Score: 1.0 - PASS
```

### Failed Validation

```bash
codebase-cli validate
```

Expected output:

```text
✗ Essential Files Agent - FAIL (Score: 0.5)
  ✓ README.md present
  ✗ CONTRIBUTING.md missing

✗ Git Configuration Agent - FAIL (Score: 0.5)
  ✓ .gitignore present
  ✗ .editorconfig missing

✗ Development Standards Agent - FAIL (Score: 0.3)
  ✗ Recent commits don't follow conventional format
  ✓ Branch naming follows conventions

Overall Score: 0.43 - FAIL
```

## Distribution and Installation

- The CLI should be available via pre-compiled binaries for major OSes (Linux, macOS, Windows)
- Installation via Go `go install github.com/codebase-interface/cli/cmd/codebase-interface@latest`
- Homebrew formula for macOS users
- Winget package for Windows users
