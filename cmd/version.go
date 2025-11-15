package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of codebase-cli",
	Long:  "Print the version number of codebase-cli",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("codebase-cli v0.1.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
