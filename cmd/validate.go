package cmd

import (
	"fmt"
	"os"

	"github.com/codebase-interface/cli/internal/agents"
	"github.com/codebase-interface/cli/internal/config"
	"github.com/codebase-interface/cli/internal/output"
	"github.com/spf13/cobra"
)

var (
	outputFormat string
	targetPath   string
	agentName    string
)

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate codebase structure and standards",
	Long: `Validate codebase structure and standards including:
- Essential files (README.md, CONTRIBUTING.md)
- Git configuration (.gitignore, .gitattributes, .editorconfig)
- Development standards (conventional commits, branch naming)`,
	RunE: runValidate,
}

func runValidate(cmd *cobra.Command, args []string) error {
	cfg, err := config.Load(targetPath)
	if err != nil {
		return fmt.Errorf("failed to load configuration: %w", err)
	}

	agentRegistry := agents.NewRegistry()
	agentRegistry.Register("essential-files", agents.NewEssentialFilesAgent())
	agentRegistry.Register("git-configuration", agents.NewGitConfigurationAgent())
	agentRegistry.Register("development-standards", agents.NewDevelopmentStandardsAgent())

	var results []agents.ValidationResult

	if agentName != "" {
		agent, exists := agentRegistry.Get(agentName)
		if !exists {
			return fmt.Errorf("agent '%s' not found", agentName)
		}

		result, err := agent.Validate(targetPath, cfg)
		if err != nil {
			return fmt.Errorf("validation failed: %w", err)
		}
		results = append(results, result)
	} else {
		for name, agent := range agentRegistry.All() {
			if cfg.IsAgentEnabled(name) {
				result, err := agent.Validate(targetPath, cfg)
				if err != nil {
					return fmt.Errorf("validation failed for agent %s: %w", name, err)
				}
				results = append(results, result)
			}
		}
	}

	formatter, err := output.NewFormatter(outputFormat)
	if err != nil {
		return fmt.Errorf("invalid output format: %w", err)
	}

	if err := formatter.Format(results, os.Stdout); err != nil {
		return fmt.Errorf("failed to format output: %w", err)
	}

	for _, result := range results {
		if result.Status == "fail" {
			os.Exit(1)
		}
	}

	return nil
}

func init() {
	rootCmd.AddCommand(validateCmd)

	validateCmd.Flags().StringVarP(&outputFormat, "output", "o", "table", "Output format (json, table)")
	validateCmd.Flags().StringVarP(&targetPath, "path", "p", ".", "Path to validate")
	validateCmd.Flags().StringVarP(&agentName, "agent", "a", "", "Run specific agent (essential-files, git-configuration, development-standards)")
}
