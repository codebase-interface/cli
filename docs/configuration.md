# ⚙️ Configuration Mastery

**Ready to make the CLI work exactly how you want?** This guide will turn you into a configuration wizard! 🧙‍♂️

The Codebase Interface CLI is powerful out of the box, but its real magic happens when you customize it for your specific needs.

## 🎯 Quick Setup

### Where to Put Your Config

Just create a `.codebase-validation.yml` file in your project root. That's it! The CLI will automatically find and use it.

```yaml
# .codebase-validation.yml - Your project's validation rules
validation:
  agents:
    essential-files:
      enabled: true
      require_readme: true
      require_contributing: true
      
  output:
    format: "table"  # Pretty output for humans
```

## 🏗️ The Complete Blueprint

Here's a comprehensive configuration that shows all the possibilities:

```yaml
validation:
  agents:
    # 📋 Essential Files - Make sure the basics are covered
    essential-files:
      enabled: true
      require_readme: true
      require_contributing: true
      require_docs_directory: true
      
    # ⚙️ Git Configuration - Keep your repo clean and consistent
    git-configuration:
      enabled: true
      require_gitignore: true
      require_gitattributes: false  # Optional by default
      require_editorconfig: true
      
    # 📜 Development Standards - Maintain quality practices
    development-standards:
      enabled: true
      check_commit_history: true
      commit_history_depth: 10      # How many commits to check
      require_conventional_commits: true
      branch_naming_patterns:
        - "^(feature|feat)/.+"
        - "^(fix|bugfix)/.+"
        - "^(hotfix|patch)/.+"
        - "^(docs|documentation)/.+"
        - "^(chore|task)/.+"
        - "^(main|master|develop|development)$"

  # 🎨 Output customization
  output:
    format: "table"    # "json" for machines, "table" for humans
    verbose: false     # Set to true for more detailed output
```

## 🎯 Agent Configuration Guide

### 📋 Essential Files Agent

*The foundation checker - makes sure your project has the basics covered.*

**Core Settings:**
```yaml
essential-files:
  enabled: true
  require_readme: true           # README.md or README.rst
  require_contributing: true     # CONTRIBUTING.md  
  require_docs_directory: true   # docs/ folder with good content
```

#### Custom README Files

```yaml
validation:
  agents:
    essential-files:
      enabled: true
      require_readme: true
      custom_readme_files:
        - "README.md"
        - "README.rst"
        - "docs/README.md"
```

#### Documentation Requirements

```yaml
validation:
  agents:
    essential-files:
      require_docs_directory: true
      docs_requirements:
        require_usage_guide: true      # Require docs/usage.md
        require_configuration_guide: true  # Require docs/configuration.md
        require_examples: true         # Require docs/examples/
        min_doc_files: 3              # Minimum number of documentation files
```

### Git Configuration Agent

Validates Git-related configuration files that ensure consistent development environment.

| Setting | Type | Default | Description |
|---------|------|---------|-------------|
| `enabled` | boolean | `true` | Enable/disable the agent |
| `require_gitignore` | boolean | `true` | Require .gitignore file |
| `require_gitattributes` | boolean | `false` | Require .gitattributes file |
| `require_editorconfig` | boolean | `true` | Require .editorconfig file |

#### Language-Specific Gitignore Validation

```yaml
validation:
  agents:
    git-configuration:
      require_gitignore: true
      gitignore_validation:
        check_language_specific: true
        required_patterns:
          go:
            - "*.exe"
            - "*.test"
            - "*.out"
          node:
            - "node_modules/"
            - "*.log"
          python:
            - "__pycache__/"
            - "*.pyc"
```

#### EditorConfig Validation

```yaml
validation:
  agents:
    git-configuration:
      require_editorconfig: true
      editorconfig_validation:
        require_root_declaration: true
        require_charset: true
        require_indent_style: true
```

### Development Standards Agent

Validates development workflow standards including commit messages and branch naming.

| Setting | Type | Default | Description |
|---------|------|---------|-------------|
| `enabled` | boolean | `true` | Enable/disable the agent |
| `check_commit_history` | boolean | `true` | Validate recent commit messages |
| `commit_history_depth` | integer | `10` | Number of recent commits to check |
| `require_conventional_commits` | boolean | `true` | Enforce conventional commit format |

#### Conventional Commits Configuration

```yaml
validation:
  agents:
    development-standards:
      require_conventional_commits: true
      conventional_commits:
        allowed_types:
          - "feat"     # New features
          - "fix"      # Bug fixes
          - "docs"     # Documentation changes
          - "style"    # Formatting changes
          - "refactor" # Code refactoring
          - "test"     # Test changes
          - "chore"    # Build/tooling changes
          - "perf"     # Performance improvements
          - "ci"       # CI/CD changes
          - "build"    # Build system changes
          - "revert"   # Revert commits
        require_scope: false        # Require scope in commits
        require_description: true   # Require description
        min_description_length: 10  # Minimum description length
        max_line_length: 72         # Maximum first line length
```

#### Branch Naming Patterns

```yaml
validation:
  agents:
    development-standards:
      branch_naming:
        patterns:
          - "^(feature|feat)/.+"           # feature/description
          - "^(fix|bugfix)/.+"            # fix/description  
          - "^(hotfix|patch)/.+"          # hotfix/description
          - "^(release|rel)/.+"           # release/version
          - "^(docs|documentation)/.+"    # docs/description
          - "^(chore|task)/.+"            # chore/description
          - "^(main|master|develop|development)$"  # Main branches
        case_sensitive: false
        allow_numbers: true
        min_length: 3
        max_length: 100
```

#### Commit Message Validation

```yaml
validation:
  agents:
    development-standards:
      commit_validation:
        check_commit_history: true
        commit_history_depth: 20
        validation_threshold: 0.8  # 80% of commits must be valid
        ignore_merge_commits: true
        ignore_revert_commits: false
```

## Output Configuration

Controls how validation results are displayed.

| Setting | Type | Default | Description |
|---------|------|---------|-------------|
| `format` | string | `"table"` | Output format: `table` or `json` |
| `verbose` | boolean | `false` | Include detailed validation information |

### Output Customization

```yaml
validation:
  output:
    format: "table"
    verbose: true
    styling:
      colors: true           # Enable colored output
      unicode_symbols: true  # Use Unicode symbols (✓, ✗, ⚠)
      show_score: true      # Show numerical scores
      show_summary: true    # Show overall summary
```

### JSON Output Configuration

```yaml
validation:
  output:
    format: "json"
    json_options:
      pretty_print: true    # Format JSON with indentation
      include_metadata: true # Include validation metadata
      timestamp: true       # Add timestamp to output
```

## Environment-Specific Configuration

### Development Environment

```yaml
# .codebase-validation.yml for development
validation:
  agents:
    essential-files:
      enabled: true
      require_docs_directory: true
    git-configuration:
      enabled: true
    development-standards:
      enabled: true
      check_commit_history: true
      commit_history_depth: 5  # Check fewer commits during development

  output:
    format: "table"
    verbose: true
```

### CI/CD Environment

```yaml
# .codebase-validation.yml for CI/CD
validation:
  agents:
    essential-files:
      enabled: true
    git-configuration:
      enabled: true
    development-standards:
      enabled: true
      check_commit_history: true
      commit_history_depth: 50  # Check more commits in CI

  output:
    format: "json"
    json_options:
      pretty_print: false
      include_metadata: true
      timestamp: true
```

### Production Validation

```yaml
# .codebase-validation.yml for production validation
validation:
  agents:
    essential-files:
      enabled: true
      require_readme: true
      require_contributing: true
      require_docs_directory: true
    git-configuration:
      enabled: true
      require_gitignore: true
      require_editorconfig: true
    development-standards:
      enabled: false  # Skip commit validation for releases

  output:
    format: "json"
    verbose: false
```

## Configuration Examples

### Minimal Configuration

```yaml
validation:
  agents:
    essential-files:
      enabled: true
    git-configuration:
      enabled: false
    development-standards:
      enabled: false
```

### Strict Configuration

```yaml
validation:
  agents:
    essential-files:
      enabled: true
      require_readme: true
      require_contributing: true
      require_docs_directory: true
    git-configuration:
      enabled: true
      require_gitignore: true
      require_gitattributes: true
      require_editorconfig: true
    development-standards:
      enabled: true
      check_commit_history: true
      commit_history_depth: 20
      require_conventional_commits: true

  output:
    format: "table"
    verbose: true
```

### Language-Specific Configuration

#### Go Project

```yaml
validation:
  agents:
    essential-files:
      enabled: true
    git-configuration:
      enabled: true
      gitignore_validation:
        required_patterns:
          - "*.exe"
          - "*.test"
          - "*.out"
          - "vendor/"
    development-standards:
      enabled: true
```

#### Node.js Project

```yaml
validation:
  agents:
    essential-files:
      enabled: true
    git-configuration:
      enabled: true
      gitignore_validation:
        required_patterns:
          - "node_modules/"
          - "*.log"
          - ".env"
          - "dist/"
    development-standards:
      enabled: true
```

## Validation Overrides

### Temporary Disabling

```yaml
validation:
  agents:
    development-standards:
      enabled: false  # Temporarily disable for legacy projects
```

### File-Specific Overrides

```yaml
validation:
  agents:
    essential-files:
      file_overrides:
        "projects/legacy/":
          require_contributing: false  # Legacy projects may not have CONTRIBUTING.md
        "tools/":
          require_docs_directory: false  # Tool projects may not need extensive docs
```

## Configuration Validation

The CLI validates the configuration file itself:

- YAML syntax validation
- Required field validation  
- Type checking
- Range validation for numeric values

Invalid configuration will result in an error with specific details about the issue.

## Best Practices

1. **Version Control**: Always commit your `.codebase-validation.yml` file
2. **Documentation**: Document any custom configuration decisions
3. **Environment-Specific**: Use different configurations for different environments
4. **Gradual Adoption**: Start with minimal validation and gradually increase strictness
5. **Team Alignment**: Ensure all team members agree on validation rules

## Migration Guide

### From Default to Custom Configuration

1. Run with default settings to see current validation status
2. Create `.codebase-validation.yml` with default values
3. Gradually customize settings based on project needs
4. Test configuration changes thoroughly

### Updating Configuration

When updating configuration:

1. Test changes locally first
2. Update documentation if validation behavior changes
3. Communicate changes to team members
4. Consider backward compatibility

## Troubleshooting Configuration

### Common Configuration Errors

#### Invalid YAML Syntax

```bash
Error: failed to parse config file: yaml: line 5: found character that cannot start any token
```

**Solution**: Validate YAML syntax using an online YAML validator or CLI tool.

#### Unknown Configuration Keys

```bash
Warning: unknown configuration key 'validation.agents.unknown-agent'
```

**Solution**: Remove unknown keys or check for typos in agent names.

#### Invalid Values

```bash
Error: commit_history_depth must be a positive integer, got: -5
```

**Solution**: Ensure numeric values are within valid ranges.

## Further Reading

- [Agent Documentation](agents.md) - Detailed agent behavior and customization
- [Usage Guide](usage.md) - CLI usage instructions
- [Examples](examples/) - Real-world configuration examples

