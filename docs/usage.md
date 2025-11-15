# CLI Usage Guide

Comprehensive guide for using the Codebase Interface CLI tool.

## Installation

### Prerequisites

- Go 1.21 or later
- Git (for development standards validation)

### Building from Source

```bash
git clone https://github.com/codebase-interface/cli.git
cd cli
go build -o bin/codebase-cli ./cmd/codebase-cli
```

### Installing Locally

```bash
# Build and install to $GOPATH/bin
go install ./cmd/codebase-cli

# Or using Taskfile (if available)
task install
```

## Commands Overview

### `codebase-cli validate`

The primary command for validating codebase structure and standards.

#### Basic Usage

```bash
# Validate current directory
codebase-cli validate

# Validate specific directory
codebase-cli validate --path /path/to/project

# Validate with JSON output
codebase-cli validate --output json

# Validate specific agent only
codebase-cli validate --agent essential-files
```

#### Flags

| Flag | Short | Description | Default |
|------|-------|-------------|---------|
| `--path` | `-p` | Path to the project directory to validate | `.` (current directory) |
| `--output` | `-o` | Output format (`table` or `json`) | `table` |
| `--agent` | `-a` | Run specific validation agent | (all enabled agents) |
| `--help` | `-h` | Show help for the command | |

#### Available Agents

- `essential-files` - Validates README.md, CONTRIBUTING.md, and docs/ structure
- `git-configuration` - Validates .gitignore, .editorconfig, .gitattributes
- `development-standards` - Validates conventional commits and branch naming

### `codebase-cli version`

Display version information.

```bash
codebase-cli version
# Output: codebase-cli v0.1.0
```

### `codebase-cli help`

Display help information for any command.

```bash
# General help
codebase-cli help

# Help for specific command
codebase-cli help validate
```

## Output Formats

### Table Format (Default)

Provides a human-readable, styled output with colors and symbols:

```text
✓ Essential Files Agent - PASS (Score: 1.0)
  ✓ README.md present
  ✓ CONTRIBUTING.md present
  ✓ docs/ directory structure present

✗ Git Configuration Agent - FAIL (Score: 0.5)
  ✓ .gitignore present
  ✗ .editorconfig missing

Overall Score: 0.75 - PASS (with warnings)
```

### JSON Format

Provides structured output suitable for programmatic processing:

```json
[
  {
    "agent": "essential-files",
    "status": "pass",
    "score": 1.0,
    "findings": [
      {
        "type": "present",
        "file": "README.md",
        "message": "README.md present",
        "severity": "info"
      }
    ]
  }
]
```

## Exit Codes

The CLI uses standard exit codes to indicate validation results:

- **0** - All validations passed successfully
- **1** - Critical validations failed (missing required files, etc.)
- **2** - Warnings present but no critical failures

## Usage Examples

### Basic Project Validation

```bash
# Navigate to your project
cd /path/to/my-project

# Run validation
codebase-cli validate

# Expected output for a well-structured project:
✓ Essential Files Agent - PASS (Score: 1.0)
  ✓ README.md present
  ✓ CONTRIBUTING.md present

✓ Git Configuration Agent - PASS (Score: 1.0)
  ✓ .gitignore present
  ✓ .editorconfig present

✓ Development Standards Agent - PASS (Score: 1.0)
  ✓ Recent commits follow conventional format
  ✓ Branch naming follows conventions: feature/new-validation

Overall Score: 1.00 - PASS
```

### CI/CD Integration

Use JSON output for automated processing in CI/CD pipelines:

```bash
# Generate JSON report
codebase-cli validate --output json > validation-report.json

# Example CI script
#!/bin/bash
if codebase-cli validate --output json > /dev/null 2>&1; then
    echo "✅ Codebase validation passed"
    exit 0
else
    echo "❌ Codebase validation failed"
    codebase-cli validate --output json
    exit 1
fi
```

### Specific Agent Validation

```bash
# Check only essential files
codebase-cli validate --agent essential-files

# Check only Git configuration
codebase-cli validate --agent git-configuration

# Check only development standards
codebase-cli validate --agent development-standards
```

### Validation with Custom Configuration

```bash
# With custom config file in project root
codebase-cli validate

# The tool automatically looks for .codebase-validation.yml
# See configuration.md for details on customizing validation rules
```

## Troubleshooting

### Common Issues

#### "Failed to check commit history"

This warning appears when the tool cannot access Git history:

- Ensure you're in a Git repository
- Verify Git is installed and accessible
- Check that the repository has commits

#### "Branch name doesn't follow conventions"

The development standards agent validates branch names against common patterns:

- `feature/description` or `feat/description`
- `fix/description` or `bugfix/description`
- `hotfix/description`
- `docs/description`
- `chore/description`
- `main`, `master`, `develop`, `development`

#### Missing Configuration File

If `.codebase-validation.yml` is not found, the tool uses default settings. Create a configuration file to customize validation rules (see [configuration.md](configuration.md)).

### Debug Information

For detailed output, use verbose mode (when available) or JSON output:

```bash
# JSON output shows all validation details
codebase-cli validate --output json | jq '.'
```

## Documentation

### Viewing Documentation Locally

For easy browsing of the complete documentation, you can serve it locally:

```bash
# Serve documentation at http://localhost:8000
task docs:serve

# Serve and automatically open in browser
task docs:open

# Check documentation for common issues
task docs:check
```

This provides a better reading experience with proper formatting, working links, and easy navigation between documentation files.

## Integration with Development Workflow

### Pre-commit Hooks

Add validation as a pre-commit hook:

```bash
# .git/hooks/pre-commit
#!/bin/bash
echo "Running codebase validation..."
if ! codebase-cli validate; then
    echo "❌ Codebase validation failed. Fix issues before committing."
    exit 1
fi
echo "✅ Codebase validation passed"
```

### IDE Integration

Many IDEs can be configured to run external tools. Configure your IDE to run:

```bash
codebase-cli validate --path ${PROJECT_DIR}
```

### GitHub Actions

```yaml
name: Validate Codebase
on: [push, pull_request]

jobs:
  validate:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v4
        with:
          go-version: '1.21'
      - name: Build CLI
        run: go build -o codebase-cli ./cmd/codebase-cli
      - name: Validate Codebase
        run: ./codebase-cli validate
```

## Advanced Usage

### Custom Project Paths

Validate multiple projects or specific subdirectories:

```bash
# Validate multiple projects
for project in project1 project2 project3; do
    echo "Validating $project..."
    codebase-cli validate --path "$project"
done

# Validate with custom config per project
codebase-cli validate --path ./backend
codebase-cli validate --path ./frontend
```

### Automated Reporting

Generate reports for compliance tracking:

```bash
#!/bin/bash
# Generate validation report with timestamp
timestamp=$(date +"%Y-%m-%d_%H-%M-%S")
report_file="validation-report-${timestamp}.json"

codebase-cli validate --output json > "$report_file"
echo "Validation report saved to: $report_file"

# Upload to reporting system or send notifications
# Example: send to webhook, save to database, etc.
```

## Further Reading

- [Configuration Reference](configuration.md) - Detailed configuration options
- [Agent Documentation](agents.md) - Understanding validation agents
- [Examples](examples/) - Real-world configuration examples

