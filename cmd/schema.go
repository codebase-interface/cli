package cmd

import (
	"embed"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

//go:embed schema/codebase-validation.schema.json
var schemaFS embed.FS

var schemaCmd = &cobra.Command{
	Use:   "schema",
	Short: "Display or save the JSON schema for configuration files",
	Long: `Display or save the JSON schema that defines the structure and validation rules
for .codebase-validation.yml configuration files.

The schema can be used with editors that support JSON Schema for autocompletion
and validation while editing YAML configuration files.`,
	Run: func(cmd *cobra.Command, args []string) {
		outputFile, _ := cmd.Flags().GetString("output")

		schemaContent, err := schemaFS.ReadFile("schema/codebase-validation.schema.json")
		if err != nil {
			fmt.Printf("❌ Failed to read embedded schema: %v\n", err)
			os.Exit(1)
		}

		if outputFile != "" {
			// Save to file
			if err := os.WriteFile(outputFile, schemaContent, 0644); err != nil {
				fmt.Printf("❌ Failed to write schema file: %v\n", err)
				os.Exit(1)
			}
			fmt.Printf("✅ Schema saved to: %s\n", outputFile)
			fmt.Printf("\n💡 You can now reference this schema in your YAML files:\n")
			fmt.Printf("# yaml-language-server: $schema=%s\n", outputFile)
		} else {
			// Display to stdout
			fmt.Print(string(schemaContent))
		}
	},
}

func init() {
	rootCmd.AddCommand(schemaCmd)
	schemaCmd.Flags().StringP("output", "o", "", "Save schema to file instead of displaying")
}
