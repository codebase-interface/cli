package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "codebase-cli",
	Short: "A CLI for validating codebase structure and standards",
	Long: `Codebase CLI validates essential files and configurations for proper codebase setup.
It checks for README.md, CONTRIBUTING.md, Git configuration files, and development standards.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Add global flags here if needed
}
