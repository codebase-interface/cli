# Contributing to Codebase CLI

Thank you for your interest in contributing to the Codebase CLI! This document provides guidelines for contributing to the project.

## Development Setup

### Prerequisites

- Go 1.21 or later
- [Task](https://taskfile.dev/) for running development tasks
- Git for version control

### Getting Started

1. Clone the repository:

```bash
git clone https://github.com/codebase-interface/cli.git
cd cli
```

1. Install dependencies:

```bash
task setup
```

1. Build the CLI:

```bash
task build
```

**📝 Note:** The `bin/` directory is git-ignored. Built binaries are not committed to the repository - each developer builds locally.

1. Run tests:

```bash
task test
```

## Development Workflow

### Code Style

- Follow Go best practices and idioms
- Use `gofmt` for formatting
- Run linting with `task lint`
- Maintain test coverage above 80%

### Commit Convention

This project follows [Conventional Commits](https://www.conventionalcommits.org/) specification:

- `feat:` for new features
- `fix:` for bug fixes
- `docs:` for documentation updates
- `refactor:` for code refactoring
- `test:` for adding or updating tests
- `chore:` for maintenance tasks

Examples:

```txt
feat: add support for custom validation rules
fix: resolve issue with git branch detection
docs: update README with installation instructions
```

### Branch Naming

Use the following patterns for branch names:

- `feature/description` for new features
- `fix/description` for bug fixes
- `docs/description` for documentation updates
- `chore/description` for maintenance tasks

### Testing

All code must be thoroughly tested:

1. Write tests first (TDD approach)
2. Ensure all tests pass: `task test`
3. Maintain high test coverage
4. Use table-driven tests where appropriate

### Adding New Validation Agents

To add a new validation agent:

1. Create the agent struct in `internal/agents/`
2. Implement the `Agent` interface
3. Add configuration options to `internal/config/`
4. Write comprehensive tests
5. Update documentation

## Pull Request Process

1. Fork the repository
2. Create a feature branch with a descriptive name
3. Make your changes following the coding standards
4. Write or update tests for your changes
5. Ensure all tests pass and linting is clean
6. Commit your changes using conventional commits
7. Push your branch and create a pull request
8. Provide a clear description of your changes

## Code Review Guidelines

- Code must be reviewed by at least one maintainer
- All tests must pass
- Code coverage must be maintained
- Documentation must be updated if applicable
- Follow the existing code style and patterns

## Development Tasks

Use the following Task commands for development:

```bash
# Build the CLI
task build

# Run all tests
task test

# Run tests in watch mode
task test:watch

# Run linting
task lint

# Clean build artifacts
task clean

# Install CLI locally
task install

# Run the CLI on itself
task validate-self

# Serve documentation locally
task docs:serve

# Open documentation in browser
task docs:open

# Check documentation for issues
task docs:check
```

## Architecture Overview

The CLI follows a modular architecture:

```txt
cmd/                    # CLI commands (Cobra)
├── root.go            # Root command
├── validate.go        # Validation command
└── version.go         # Version command

internal/              # Internal packages
├── agents/           # Validation agents
├── config/           # Configuration handling
├── output/           # Output formatters
└── ui/               # User interface components

pkg/                   # Public API (future use)
test/                  # Test fixtures and helpers
```

### Key Components

- **Agents**: Modular validation logic for different aspects
- **Config**: YAML configuration system with sensible defaults
- **Output**: Multiple output formats (JSON, table) with styled output
- **Registry**: Plugin-like system for registering and running agents

## Questions or Issues?

If you have questions or encounter issues:

1. Check existing GitHub issues
2. Create a new issue with detailed information
3. Join our discussions in GitHub Discussions
4. Follow our code of conduct

Thank you for contributing to make codebase validation better for everyone!
