# 🎯 Getting Started with Codebase CLI

**Ready to transform your project into a well-organized, professional codebase?** You're in the right place! This guide will have you validating and improving your projects in just a few minutes.

## 📥 Installation First

If you haven't installed the CLI yet, check out our **[Installation Guide](installation.md)** for multiple easy options:

- **One-line installer** for Linux/macOS
- **Package managers** (Homebrew, Chocolatey)  
- **Docker containers** for CI/CD
- **Pre-built binaries** for all platforms
- **Build from source** for developers

**Quick install:**
```bash
# Linux/macOS one-liner
curl -sSL https://raw.githubusercontent.com/codebase-interface/cli/main/install.sh | bash

# Verify installation
cbi version
```

## ⚡ Your First Validation

Let's validate your first project! Navigate to any project directory and run:

```bash
codebase-interface validate
# or use the short alias
cbi validate
```

**What happens next?** The CLI will scan your project and show you exactly what's missing or could be improved:

```text
✓ Essential Files Agent - PASS (Score: 1.0)
  ✓ README.md present and well-structured
  ✓ CONTRIBUTING.md found

⚠️  Git Configuration Agent - WARNING (Score: 0.7)
  ✓ .gitignore present
  ⚠️  .editorconfig missing (recommended for consistent formatting)

❌ Development Standards Agent - FAIL (Score: 0.3)
  ❌ Recent commits don't follow conventional format
  ✓ Branch naming follows conventions

Overall Score: 0.67 - NEEDS IMPROVEMENT
```

## 🎮 Command Mastery

### The Main Event: `validate`

This is your primary tool for checking project quality. Here's how to wield it:

```bash
# Validate current directory (most common)
codebase-interface validate
# or: cbi validate

# Check a specific project
codebase-interface validate --path /path/to/my-awesome-project
# or: cbi validate --path /path/to/my-awesome-project

# Get machine-readable output
codebase-interface validate --output json
# or: cbi validate --output json

# Focus on specific areas
codebase-interface validate --agent essential-files
codebase-interface validate --agent git-configuration  
codebase-interface validate --agent development-standards
# or use the short alias:
# cbi validate --agent essential-files
```

### 🛠️ Command Options Reference

| Flag | Short | What It Does | Default |
|------|-------|-------------|----------|
| `--path` | `-p` | 📁 Which project to validate | `.` (current directory) |
| `--output` | `-o` | 📊 How to show results (`table` or `json`) | `table` |
| `--agent` | `-a` | 🎯 Focus on one validator only | (all enabled) |
| `--help` | `-h` | 📚 Show help for the command | |

### 🤖 Meet Your Validation Agents

- **📋 essential-files** - Ensures you have README.md, CONTRIBUTING.md, and proper docs
- **⚙️ git-configuration** - Checks for .gitignore, .editorconfig, .gitattributes
- **📜 development-standards** - Validates commit messages and branch naming

### 🔍 Other Handy Commands

**Configuration Management:**
```bash
# Create a new configuration file
codebase-interface init-config basic
cbi init-config go-project

# Validate your configuration file
codebase-interface validate-config
cbi validate-config

# Get the JSON schema for editor support
codebase-interface schema -o schema.json
cbi schema
```

**Version and Help:**
```bash
# Check what version you're running
codebase-interface version
# or: cbi version

# Get help anytime
codebase-interface help
codebase-interface help validate
# or: cbi help
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

## 🚦 Understanding Exit Codes

When the CLI finishes, it tells you exactly how things went:

- **🟢 Exit 0** - Perfect! All validations passed
- **🔴 Exit 1** - Issues found that need attention
- **🟡 Exit 2** - Minor warnings, but nothing critical

*This is especially useful for automation and CI/CD pipelines!*

## 🎆 Real-World Examples

### 🎉 Your First Success Story

```bash
# Navigate to your project
cd /path/to/my-awesome-project

# Run the magic
codebase-cli validate

# Celebrate when you see this:
✓ Essential Files Agent - PASS (Score: 1.0)
  ✓ README.md present and informative
  ✓ CONTRIBUTING.md guides contributors well

✓ Git Configuration Agent - PASS (Score: 1.0)
  ✓ .gitignore keeps unwanted files out
  ✓ .editorconfig ensures consistent formatting

✓ Development Standards Agent - PASS (Score: 1.0)
  ✓ Recent commits follow conventional format
  ✓ Branch naming follows conventions: feature/amazing-feature

🎆 Overall Score: 1.00 - EXCELLENT PROJECT!
```

### 🤖 Automation Made Simple

Integrate with your CI/CD pipeline effortlessly:

```bash
# Generate a detailed JSON report
codebase-interface validate --output json > validation-report.json
# or: cbi validate --output json > validation-report.json

# Simple CI/CD script that works everywhere
#!/bin/bash
if codebase-interface validate --output json > /dev/null 2>&1; then
    echo "✅ Your codebase rocks! Validation passed"
    exit 0
else
    echo "⚠️  Time for some improvements!"
    codebase-interface validate
    exit 1
fi
```

### 🎯 Laser-Focused Validation

Sometimes you want to check just one thing:

```bash
# Just the essentials
codebase-interface validate --agent essential-files
# or: cbi validate --agent essential-files

# Only Git setup
codebase-interface validate --agent git-configuration
# or: cbi validate --agent git-configuration

# Just development practices
codebase-interface validate --agent development-standards
# or: cbi validate --agent development-standards
```

*Perfect for when you're working on specific improvements!*

### 🎨 Custom Configuration Magic

**Quick Start with Templates:**
```bash
# Create a configuration file from templates
codebase-interface init-config basic      # Simple setup
codebase-interface init-config strict     # High standards  
codebase-interface init-config beginner   # Gentle introduction
codebase-interface init-config open-source # OSS projects
codebase-interface init-config go-project  # Go-specific setup
```

**Validate Your Configuration:**
```bash
# Make sure your config file is correct
codebase-interface validate-config
# or: cbi validate-config

# The CLI automatically finds your .codebase-validation.yml file
codebase-interface validate
# or: cbi validate
```

**Editor Integration:**
```bash
# Get the JSON schema for autocomplete and validation in your editor
codebase-interface schema -o .vscode/codebase-validation.schema.json
```

*Then add this to the top of your `.codebase-validation.yml`:*
```yaml
# yaml-language-server: $schema=.vscode/codebase-validation.schema.json
```

*Head over to our [configuration guide](configuration.md) to unlock the full power of customization.*

## 🔧 Troubleshooting Guide

### 🤔 "Hmm, something's not quite right..."

**🚨 "Failed to check commit history"**

Don't worry! This usually means:
- You're not in a Git repository (try `git init` if you want one)
- Git isn't installed (grab it from [git-scm.com](https://git-scm.com))
- Your repository doesn't have any commits yet (make your first commit!)

**🏷️ "Branch name doesn't follow conventions"**

Your branch name needs to follow these friendly patterns:
- `feature/amazing-new-thing` or `feat/cool-feature`
- `fix/annoying-bug` or `bugfix/quick-fix`  
- `hotfix/urgent-patch`
- `docs/better-readme`
- `chore/cleanup-task`
- Or stick with classics: `main`, `master`, `develop`

**📝 "Missing Configuration File"**

No `.codebase-validation.yml`? No problem! The CLI works great with sensible defaults. But if you want to customize, check out our [configuration guide](configuration.md) and [examples](examples/).

### 🔍 Getting More Details

Need to see exactly what's happening?

```bash
# Get all the details in JSON format
codebase-interface validate --output json
# or: cbi validate --output json

# Pretty-print with jq if you have it installed
codebase-interface validate --output json | jq '.'
# or: cbi validate --output json | jq '.'
```

## 🚀 Power User Features

### 🎯 Pre-commit Hooks (Never Break the Build Again!)

Add this to your `.git/hooks/pre-commit` file:

```bash
#!/bin/bash
echo "🔍 Running codebase validation..."
if ! codebase-interface validate; then
    echo "❌ Hold up! Fix these issues before committing."
    exit 1
fi
echo "✅ All good! Committing..."
```

*Make it executable: `chmod +x .git/hooks/pre-commit`*

### 💻 IDE Integration (Code While You Validate)

Most IDEs can run external tools. Set yours up with:

```bash
codebase-interface validate --path ${PROJECT_DIR}
# or: cbi validate --path ${PROJECT_DIR}
```

*Works great with VS Code tasks, IntelliJ run configurations, and more!*

### 🔄 GitHub Actions (Automate All The Things)

```yaml
name: "🤖 Validate Codebase"
on: [push, pull_request]

jobs:
  validate:
    name: "🔍 Check Project Quality"
    runs-on: ubuntu-latest
    steps:
      - name: "📥 Get the code"
        uses: actions/checkout@v4
        
      - name: "🔧 Setup Go"
        uses: actions/setup-go@v4
        with:
          go-version: '1.21'
          
      - name: "📎 Build CLI"
        run: go build -o codebase-interface ./cmd/codebase-interface
        
      - name: "✨ Validate Codebase"
        run: ./codebase-interface validate
```

## 🎮 Advanced Techniques

### 📋 Multi-Project Validation

Got multiple projects? Validate them all:

```bash
# Check all your awesome projects
for project in frontend backend mobile; do
    echo "🔍 Validating $project..."
    codebase-interface validate --path "./$project"
    # or: cbi validate --path "./$project"
done

# Or check different environments
codebase-interface validate --path ./development
codebase-interface validate --path ./staging
codebase-interface validate --path ./production
```

### 📈 Automated Reporting & Analytics

Track your project quality over time:

```bash
#!/bin/bash
# Create timestamped quality reports
timestamp=$(date +"%Y-%m-%d_%H-%M-%S")
report_file="quality-report-${timestamp}.json"

codebase-interface validate --output json > "reports/$report_file"
# or: cbi validate --output json > "reports/$report_file"
echo "📈 Quality report saved: $report_file"

# Send to your monitoring system
# curl -X POST "https://your-api.com/quality-reports" -d @"reports/$report_file"
```

---

## 🎆 What's Next?

**Ready to customize?** → [Configuration Guide](configuration.md) - Tailor validation to your needs

**Want to understand more?** → [Validation Agents](agents.md) - Deep dive into what gets checked

**Looking for inspiration?** → [Examples](examples/) - Copy-paste configurations for any project type

---

*💬 **Questions?** Open an issue on GitHub - we'd love to help make your codebase amazing!*

