# Configuration Examples

This directory contains real-world examples of codebase validation configurations for different project types and team requirements.

## Available Examples

### Project Types

- **[Beginner-Friendly](beginner.yml)** - Perfect starting point for new developers
- **[Basic Configuration](basic.yml)** - Simple setup for small projects
- **[Go Project](go-project.yml)** - Complete Go project configuration
- **[Open Source](open-source.yml)** - Configuration for open source projects
- **[Strict Standards](strict.yml)** - High validation standards for professional teams
- **[Node.js Project](nodejs-project.yml)** - JavaScript/TypeScript project setup (coming soon)
- **[Python Project](python-project.yml)** - Python package configuration (coming soon)
- **[Enterprise Setup](enterprise.yml)** - Large organization standards (coming soon)

### Team Configurations

- **[Strict Standards](strict.yml)** - High validation standards
- **[Lenient Standards](lenient.yml)** - Relaxed validation for legacy projects
- **[CI/CD Optimized](cicd.yml)** - Configuration optimized for automated pipelines

### Special Cases

- **[Open Source](open-source.yml)** - Configuration for open source projects
- **[Legacy Migration](legacy-migration.yml)** - Gradual migration from legacy codebases
- **[Multi-Language](multi-language.yml)** - Projects with multiple programming languages

## Usage

Copy any example configuration to your project root as `.codebase-validation.yml` and customize as needed:

```bash
# Copy basic configuration
cp docs/examples/basic.yml .codebase-validation.yml

# Copy and customize for your project
cp docs/examples/go-project.yml .codebase-validation.yml
```

## Contributing Examples

To contribute a new example configuration:

1. Create a new `.yml` file in this directory
2. Follow the naming convention: `descriptive-name.yml`
3. Include comprehensive comments explaining the configuration choices
4. Add an entry to this README with a brief description
5. Test the configuration on a real project

## Example Structure

Each example should include:

```yaml
# Brief description of the use case
# Target project type or team size
# Key features enabled/disabled

validation:
  agents:
    # Agent configurations with comments
    
  output:
    # Output preferences
    
  custom:
    # Any custom settings
```
