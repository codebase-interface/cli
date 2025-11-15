# Codebase Validation Agents

This document defines the validation agents for the codebase interface CLI. These agents validate essential files and configurations for proper codebase setup.

## Technical Implementation

### Technology Stack

- **Language:** Go (latest stable version)
- **CLI Framework:** [Cobra](https://github.com/spf13/cobra) - Standard Go CLI framework with command structure and flags
- **TUI Framework:** [Bubble Tea](https://github.com/charmbracelet/bubbletea) - Modern terminal UI for interactive validation reports
- **Task Runner:** [Taskfile](https://taskfile.dev/) - Task runner for development workflows
- **Testing:** Test-driven development (TDD) approach with Go's built-in testing framework

### Development Principles

- **Test-Driven Development (TDD)** - All features must be developed with tests first
- **Go Best Practices** - Follow Go idioms, effective Go guidelines, and community standards
- **Documentation-Driven** - Maintain comprehensive README.md and CONTRIBUTING.md
- **Task Automation** - Use Taskfile for all development interactions (build, test, lint, release)
- **Conventional Commits** - Follow [Conventional Commits](https://www.conventionalcommits.org/) specification for commit messages

### Architecture Requirements

- Modular agent system for extensibility
- Clean separation between validation logic and UI presentation
- Configuration-driven validation rules
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
codebase-cli/
├── cmd/                    # Cobra CLI commands
│   ├── root.go            # Root command
│   ├── validate.go        # Validation command
│   └── version.go         # Version command
├── internal/              # Internal packages
│   ├── agents/           # Validation agents
│   ├── config/           # Configuration handling
│   ├── output/           # Output formatters
│   └── ui/               # Bubble Tea TUI components
├── pkg/                   # Public API packages
├── test/                  # Test files and fixtures
├── Taskfile.yml          # Development tasks
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
      - go build -o bin/codebase-cli ./cmd/codebase-cli

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

  tidy:
    desc: Tidy Go modules
    cmds:
      - go mod tidy

  install:
    desc: Install CLI locally
    cmds:
      - go install ./cmd/codebase-cli
```

### Test-Driven Development Requirements

- **Unit Tests:** Each agent must have comprehensive unit tests
- **Integration Tests:** End-to-end validation scenarios
- **Table-Driven Tests:** Use Go's table-driven test patterns
- **Test Coverage:** Minimum 80% code coverage
- **Test Fixtures:** Provide sample project structures for testing

---

## Validation Execution

### Command Structure

```bash
# Validate all essential files
codebase-cli validate

# Validate specific agent
codebase-cli validate --agent essential-files
codebase-cli validate --agent git-configuration
codebase-cli validate --agent development-standards

# Output formats
codebase-cli validate --output json
codebase-cli validate --output table
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
