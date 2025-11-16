package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var initConfigCmd = &cobra.Command{
	Use:   "init-config [type]",
	Short: "Initialize a new .codebase-validation.yml configuration file",
	Long: `Initialize a new .codebase-validation.yml configuration file with sensible defaults.

Available types:
  basic      - Simple configuration for small projects
  strict     - High standards for professional development
  beginner   - Gentle introduction with relaxed standards  
  open-source - Configuration optimized for open source projects
  go-project - Complete configuration for Go projects

If no type is specified, 'basic' will be used.`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		configType := "basic"
		if len(args) > 0 {
			configType = args[0]
		}

		configPath := ".codebase-validation.yml"

		// Check if config already exists
		if _, err := os.Stat(configPath); err == nil {
			overwrite, _ := cmd.Flags().GetBool("force")
			if !overwrite {
				fmt.Printf("❌ Configuration file already exists: %s\n", configPath)
				fmt.Println("   Use --force to overwrite")
				os.Exit(1)
			}
		}

		var configContent string
		switch configType {
		case "basic":
			configContent = getBasicConfig()
		case "strict":
			configContent = getStrictConfig()
		case "beginner":
			configContent = getBeginnerConfig()
		case "open-source":
			configContent = getOpenSourceConfig()
		case "go-project":
			configContent = getGoProjectConfig()
		default:
			fmt.Printf("❌ Unknown configuration type: %s\n", configType)
			fmt.Println("   Available types: basic, strict, beginner, open-source, go-project")
			os.Exit(1)
		}

		if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
			fmt.Printf("❌ Failed to create config file: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("✅ Created %s configuration: %s\n", configType, configPath)
		fmt.Printf("\n🎯 Next steps:\n")
		fmt.Printf("   1. Review and customize the configuration\n")
		fmt.Printf("   2. Run: codebase-interface validate\n")
		fmt.Printf("   3. Validate config: codebase-interface validate-config\n")
		fmt.Printf("\n📖 Documentation: https://cli.codebaseinterface.org/configuration/\n")
	},
}

func getBasicConfig() string {
	return `# Basic Project Configuration
# Simple setup for small projects and teams starting with validation
# Focuses on essential files with minimal complexity

validation:
  agents:
    # Essential files validation
    essential-files:
      enabled: true
      require_readme: true           # README.md or README.rst required
      require_contributing: true     # CONTRIBUTING.md required
      require_docs_directory: false # Docs directory optional for small projects
      
    # Git configuration
    git-configuration:
      enabled: true
      require_gitignore: true        # .gitignore required
      require_gitattributes: false   # .gitattributes optional
      require_editorconfig: true     # .editorconfig for consistency
      
    # Development standards
    development-standards:
      enabled: true
      check_commit_history: true
      commit_history_depth: 5        # Check last 5 commits only
      require_conventional_commits: false  # Relaxed for small teams
      validation_threshold: 0.6      # 60% of commits should be good

  # Output configuration
  output:
    format: "table"                  # Human-readable table format
    verbose: false                   # Concise output

  # Scoring thresholds
  scoring:
    pass_threshold: 0.7              # 70% score needed to pass
    warning_threshold: 0.5           # Below 50% is critical failure
`
}

func getStrictConfig() string {
	return `# Strict Standards Configuration
# High validation standards for professional development teams
# Requires comprehensive documentation and strict adherence to conventions

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
        min_doc_files: 5
      readme_quality:
        min_lines: 30
        require_description: true
        require_installation: true
        require_usage: true
        
    git-configuration:
      enabled: true
      require_gitignore: true
      require_gitattributes: true     # Required for strict standards
      require_editorconfig: true
      validation_rules:
        gitignore_validation: true
        editorconfig_validation: true
        gitattributes_validation: true
        
    development-standards:
      enabled: true
      check_commit_history: true
      commit_history_depth: 20        # Check more commits
      require_conventional_commits: true
      validation_threshold: 0.9       # 90% of commits must be conventional
      branch_validation: true
      conventional_commits:
        require_scope: true           # Scope required in strict mode
        require_breaking_change_footer: true

  output:
    format: "table"
    verbose: true

  scoring:
    pass_threshold: 0.95              # 95% score required to pass
    warning_threshold: 0.85           # 85% for warnings
`
}

func getBeginnerConfig() string {
	return `# Beginner-Friendly Configuration
# Perfect for developers just starting with codebase validation
# Gentle introduction with helpful explanations and relaxed standards

validation:
  agents:
    essential-files:
      enabled: true
      require_readme: true           # Every project needs a welcoming README
      require_contributing: false    # Optional for personal projects
      require_docs_directory: false # Can add documentation later
      
    git-configuration:
      enabled: true
      require_gitignore: true        # Essential - keeps junk out of your repo
      require_gitattributes: false   # Advanced feature, skip for now
      require_editorconfig: false    # Nice to have, not required initially
      
    development-standards:
      enabled: true
      check_commit_history: true
      commit_history_depth: 5        # Check only recent commits
      require_conventional_commits: false  # Learn this later
      validation_threshold: 0.3      # Very forgiving - 30% is enough to start

  output:
    format: "table"                  # Pretty, human-readable output
    verbose: false                   # Don't overwhelm with details

  scoring:
    pass_threshold: 0.5              # 50% score is passing - you've got this!
    warning_threshold: 0.3           # Only warn if really struggling
`
}

func getOpenSourceConfig() string {
	return `# Open Source Project Configuration
# Optimized for open source projects with community contributions
# Emphasizes documentation, contribution guidelines, and welcoming setup

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
        min_doc_files: 5
      readme_quality:
        min_lines: 50                # Comprehensive README
        require_description: true
        require_installation: true
        require_usage: true
        check_badges: true           # Status badges important for OSS
      custom_files:
        - pattern: "LICENSE*"
          required: true
          description: "License file mandatory for open source"
        - pattern: "CODE_OF_CONDUCT*"
          required: true
          description: "Code of conduct for community"
        - pattern: "SECURITY*"
          required: true
          description: "Security policy"
        
    git-configuration:
      enabled: true
      require_gitignore: true
      require_gitattributes: true
      require_editorconfig: true
      
    development-standards:
      enabled: true
      check_commit_history: true
      commit_history_depth: 10
      require_conventional_commits: true
      validation_threshold: 0.7       # 70% - accommodates new contributors

  output:
    format: "table"
    verbose: true

  scoring:
    pass_threshold: 0.8
    warning_threshold: 0.6
`
}

func getGoProjectConfig() string {
	return `# Go Project Configuration
# Comprehensive configuration for Go projects with modern development practices
# Includes Go-specific validation rules and strict quality standards

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
      custom_files:
        - pattern: "LICENSE*"
          required: true
          description: "License file required"
        - pattern: "go.mod"
          required: true
          description: "Go module file"
        - pattern: "Taskfile.yml"
          required: true
          description: "Task automation file"
          
    git-configuration:
      enabled: true
      require_gitignore: true
      require_gitattributes: true
      require_editorconfig: true
      validation_rules:
        gitignore_validation: true
        gitattributes_validation: true
      gitignore_validation:
        check_language_specific: true
        detect_project_type: true
        required_patterns:
          go: 
            - "*.exe"
            - "*.test"
            - "*.out"
            - "vendor/"
            - "bin/"
            
    development-standards:
      enabled: true
      check_commit_history: true
      commit_history_depth: 15
      require_conventional_commits: true
      validation_threshold: 0.8
      conventional_commits:
        allowed_types:
          - "feat"
          - "fix"
          - "docs"
          - "style"
          - "refactor"
          - "test"
          - "chore"
          - "perf"
          - "ci"
          - "build"

  output:
    format: "table"
    verbose: true

  scoring:
    pass_threshold: 0.85
    warning_threshold: 0.7
`
}

func init() {
	rootCmd.AddCommand(initConfigCmd)
	initConfigCmd.Flags().Bool("force", false, "Overwrite existing configuration file")
}
