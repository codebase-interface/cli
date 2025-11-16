# 🤖 Understanding Your Validation Agents

**Ever wonder what the CLI is actually checking?** Meet your three validation assistants - each one has a specific job to make your project awesome!

Think of these agents as your personal code quality team, each with their own specialty and expertise.

## 🧠 How Validation Works

When you run `codebase-cli validate`, each agent:

1. 🔍 **Scans** your project for specific files and patterns
2. 📊 **Scores** what it finds (0.0 = needs work, 1.0 = perfect!)
3. 📝 **Reports** exactly what's missing or could be improved
4. ✅ **Passes or Fails** based on your configuration

Each agent gives you actionable feedback, not just complaints!

## 📋 Essential Files Agent - "The Welcomer"

**Job:** *"Make sure visitors to your project feel welcome and know what's going on!"*

### What It's Looking For

This agent ensures your project has the fundamental files that make it professional and approachable:

**🎯 README File** - Your project's front door

- Looks for `README.md`, `README.rst`, or similar
- This is the first thing people see!

**🤝 Contributing Guidelines** - How others can help

- Searches for `CONTRIBUTING.md`
- Shows people how to get involved

**📚 Documentation Directory** - The knowledge base

- Validates that you have a `docs/` folder
- Checks that it actually contains useful content

### How It Scores

```txt
Your Score = (Files Found) ÷ (Files Required)
```

- Missing README? Major problem! 😱
- Missing CONTRIBUTING? Room for improvement 📈
- All files present? You're a documentation champion! 🏆

- Each required file contributes equally to the final score
- Missing critical files result in "fail" status
- All files present results in "pass" status

### Configuration Options

```yaml
validation:
  agents:
    essential-files:
      enabled: true
      require_readme: true
      require_contributing: true
      require_docs_directory: true
      docs_requirements:
        require_usage_guide: true
        require_examples: true
        min_doc_files: 3
```

### File Detection Rules

#### README Files

The agent searches for README files in this order:

1. `README.md` (preferred)
2. `README.rst` (reStructuredText)
3. `readme.md` (lowercase variant)
4. `readme.rst` (lowercase reStructuredText)

#### Documentation Directory

When `require_docs_directory` is enabled, the agent checks for:

- `docs/` directory exists
- `docs/README.md` or `docs/index.md` (documentation index)
- `docs/usage.md` (usage instructions) - if `require_usage_guide: true`
- `docs/examples/` directory - if `require_examples: true`

### Quality Checks

#### README Quality Assessment

When a README file is found, the agent can perform quality checks:

```yaml
validation:
  agents:
    essential-files:
      readme_quality:
        min_lines: 10          # Minimum number of lines
        require_description: true    # Must have description section
        require_installation: true   # Must have installation section
        require_usage: true         # Must have usage section
        check_badges: false         # Check for status badges
```

#### Documentation Completeness

```yaml
validation:
  agents:
    essential-files:
      docs_completeness:
        check_toc: true            # Table of contents in docs/README.md
        check_cross_references: true # Links between documentation files
        require_api_docs: true     # API documentation for libraries
```

### Custom File Patterns

```yaml
validation:
  agents:
    essential-files:
      custom_files:
        - pattern: "LICENSE*"
          required: true
          description: "License file"
        - pattern: "CHANGELOG*"
          required: false
          description: "Changelog file"
```

### Examples

#### Minimal Project Structure

```txt
project/
├── README.md
├── CONTRIBUTING.md
└── docs/
    └── README.md
```

**Result**: ✅ PASS (Score: 1.0) - All essential files present

#### Incomplete Project Structure

```txt
project/
└── README.md
```

**Result**: ❌ FAIL (Score: 0.33) - Missing CONTRIBUTING.md and docs/

## Git Configuration Agent

**Purpose**: Validates Git configuration files that ensure consistent development environment.
**Priority**: High

### Validation Logic

The Git Configuration Agent checks for:

1. **`.gitignore`**: Prevents committing unwanted files
2. **`.editorconfig`**: Ensures consistent coding styles
3. **`.gitattributes`**: Defines file handling attributes (optional by default)

### Scoring Algorithm

```txt
Score = (Present Files + Weighted Optional Files) / (Total Required Files + Weighted Optional Files)
```

- Required files have weight 1.0
- Optional files have weight 0.5
- Missing required files significantly impact score

### Configuration Options

```yaml
validation:
  agents:
    git-configuration:
      enabled: true
      require_gitignore: true
      require_gitattributes: false  # Optional by default
      require_editorconfig: true
      validation_rules:
        gitignore_validation: true
        editorconfig_validation: true
```

### File Validation Rules

#### .gitignore Validation

```yaml
validation:
  agents:
    git-configuration:
      gitignore_validation:
        check_language_specific: true
        detect_project_type: true    # Auto-detect from files
        required_patterns:
          go: ["*.exe", "*.test", "*.out", "vendor/"]
          node: ["node_modules/", "*.log", ".env"]
          python: ["__pycache__/", "*.pyc", ".venv/"]
          general: [".DS_Store", "Thumbs.db", "*.swp"]
```

#### .editorconfig Validation

```yaml
validation:
  agents:
    git-configuration:
      editorconfig_validation:
        require_root: true           # root = true declaration
        require_charset: true        # charset specification
        require_indent_style: true   # indent_style specification
        check_file_types: true       # File type sections
```

#### .gitattributes Validation

```yaml
validation:
  agents:
    git-configuration:
      gitattributes_validation:
        check_line_endings: true     # Text file line endings
        check_binary_files: true     # Binary file handling
        check_export_ignore: true    # Export ignore patterns
```

### Language-Specific Validation

The agent can detect project type and apply appropriate validation:

```go
// Project type detection
func detectProjectType(targetPath string) []string {
    // Detects based on presence of:
    // - go.mod -> Go project
    // - package.json -> Node.js project  
    // - pyproject.toml/setup.py -> Python project
    // - etc.
}
```

### Examples

#### Well-Configured Git Setup

```txt
project/
├── .gitignore        # Language-appropriate ignore rules
├── .editorconfig     # Consistent code formatting
└── .gitattributes    # Line ending and binary file handling
```

**Result**: ✅ PASS (Score: 1.0) - All configuration files present and valid

#### Minimal Git Setup

```txt
project/
├── .gitignore        # Basic ignore rules
└── .editorconfig     # Basic formatting rules
```

**Result**: ✅ PASS (Score: 0.83) - Required files present, optional .gitattributes missing

## Development Standards Agent

**Purpose**: Validates development workflow standards including commit messages and branch naming.
**Priority**: Medium

### Validation Logic

The Development Standards Agent performs these checks:

1. **Conventional Commits**: Validates recent commit messages
2. **Branch Naming**: Ensures current branch follows naming conventions
3. **Commit Quality**: Assesses commit history quality

### Scoring Algorithm

```txt
Score = (Valid Commits / Total Commits Checked) * 0.7 + (Branch Naming Valid ? 0.3 : 0)
```

- Commit history validation contributes 70% to score
- Branch naming contributes 30% to score
- Configurable thresholds for pass/fail determination

### Configuration Options

```yaml
validation:
  agents:
    development-standards:
      enabled: true
      check_commit_history: true
      commit_history_depth: 10
      require_conventional_commits: true
      validation_threshold: 0.8      # 80% of commits must be valid
      branch_validation: true
```

### Conventional Commits Validation

#### Standard Format

```
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

#### Supported Types

Default conventional commit types:

- `feat`: New features
- `fix`: Bug fixes  
- `docs`: Documentation changes
- `style`: Code formatting changes
- `refactor`: Code refactoring
- `test`: Test additions/modifications
- `chore`: Build process or auxiliary tool changes
- `perf`: Performance improvements
- `ci`: Continuous integration changes
- `build`: Build system changes
- `revert`: Revert previous commits

#### Custom Type Configuration

```yaml
validation:
  agents:
    development-standards:
      conventional_commits:
        allowed_types:
          - "feat"
          - "fix"
          - "docs"
          - "custom"     # Project-specific type
        require_scope: true
        scopes:
          - "api"
          - "ui"
          - "backend"
        require_breaking_change_footer: true
```

### Branch Naming Validation

#### Default Patterns

```regex
^(feature|feat)/.+          # feature/description
^(fix|bugfix)/.+           # fix/description
^(hotfix|patch)/.+         # hotfix/description
^(release|rel)/.+          # release/version
^(docs|documentation)/.+   # docs/description
^(chore|task)/.+           # chore/description
^(main|master|develop|development)$  # Main branches
```

#### Custom Pattern Configuration

```yaml
validation:
  agents:
    development-standards:
      branch_naming:
        patterns:
          - "^epic/.+"           # Epic branches
          - "^spike/.+"          # Research spikes
          - "^prototype/.+"      # Prototypes
        case_sensitivity: false
        min_length: 5
        max_length: 80
```

### Commit History Analysis

#### Quality Metrics

The agent analyzes commit history for:

- **Conventional format compliance**
- **Commit message length**
- **Frequency of fix vs feature commits**
- **Breaking change indicators**

#### Configuration

```yaml
validation:
  agents:
    development-standards:
      commit_analysis:
        min_message_length: 10
        max_message_length: 72      # First line
        check_breaking_changes: true
        ignore_merge_commits: true
        ignore_fixup_commits: true
```

### Git Integration

The agent uses Git commands to gather information:

```bash
# Get commit history
git log --oneline -<depth>

# Get current branch
git rev-parse --abbrev-ref HEAD

# Check if in git repository
git rev-parse --git-dir
```

### Error Handling

Common scenarios and handling:

- **Not a Git repository**: Warning with graceful degradation
- **No commits**: Skip commit validation
- **Git not available**: Warning and skip git-based checks

### Examples

#### Well-Managed Repository

```bash
$ git log --oneline -5
abc1234 feat(api): add user authentication endpoint
def5678 fix(ui): resolve login form validation
ghi9012 docs: update API documentation
jkl3456 test(auth): add integration tests
mno7890 chore(deps): update dependencies
```

**Current branch**: `feature/user-authentication`

**Result**: ✅ PASS (Score: 1.0) - All commits follow conventions, branch naming correct

#### Legacy Repository

```bash
$ git log --oneline -5
abc1234 Fixed bug in login
def5678 Added new feature
ghi9012 Update README
jkl3456 WIP
mno7890 Cleanup
```

**Current branch**: `develop`

**Result**: ❌ FAIL (Score: 0.3) - No commits follow conventional format, branch naming acceptable

## Agent Extension

### Custom Agents

Implement the Agent interface to create custom validation:

```go
type CustomAgent struct{}

func (a *CustomAgent) Validate(targetPath string, cfg *config.Config) (ValidationResult, error) {
    // Custom validation logic
    return ValidationResult{
        Agent:    "custom-agent",
        Status:   "pass",
        Score:    1.0,
        Findings: []Finding{},
    }, nil
}
```

### Agent Registration

```go
registry := agents.NewRegistry()
registry.Register("custom-agent", &CustomAgent{})
```

### Plugin System (Future)

Future versions may support plugin-based agents:

```yaml
validation:
  plugins:
    - name: "security-agent"
      path: "./plugins/security.so"
      config:
        scan_dependencies: true
        check_secrets: true
```

## Best Practices

### Agent Configuration

1. **Start Simple**: Begin with default configurations
2. **Iterative Improvement**: Gradually increase validation strictness  
3. **Team Consensus**: Ensure all team members agree on standards
4. **Environment-Specific**: Use different configurations for different environments

### Performance Optimization

1. **Limit Commit History Depth**: Use reasonable `commit_history_depth` values
2. **Disable Expensive Checks**: Turn off quality checks for large repositories
3. **Cache Results**: Consider caching validation results for CI/CD

### Troubleshooting

1. **Use JSON Output**: For detailed debugging information
2. **Enable Verbose Mode**: When available, for additional context
3. **Check Git Configuration**: Ensure Git is properly configured
4. **Validate Configuration**: Test configuration changes incrementally

## Further Reading

- [Configuration Reference](configuration.md) - Detailed configuration options
- [Usage Guide](usage.md) - CLI usage instructions
- [Examples](examples/) - Real-world configuration examples
