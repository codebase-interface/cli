package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/xeipuuv/gojsonschema"
	"gopkg.in/yaml.v3"
)

var validateConfigCmd = &cobra.Command{
	Use:   "validate-config [path]",
	Short: "Validate a .codebase-validation.yml configuration file",
	Long: `Validate a .codebase-validation.yml configuration file against the JSON schema.
This helps catch configuration errors and shows the expected format.`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var configPath string
		if len(args) > 0 {
			// Check if argument is a file or directory
			if filepath.Ext(args[0]) == ".yml" || filepath.Ext(args[0]) == ".yaml" {
				configPath = args[0]
			} else {
				configPath = filepath.Join(args[0], ".codebase-validation.yml")
			}
		} else {
			configPath = ".codebase-validation.yml"
		}

		// Check if config file exists
		if _, err := os.Stat(configPath); os.IsNotExist(err) {
			fmt.Printf("❌ Configuration file not found: %s\n", configPath)
			fmt.Println("\n💡 To create a configuration file, try:")
			fmt.Println("   codebase-interface init-config")
			os.Exit(1)
		}

		// Read and parse the YAML file
		data, err := os.ReadFile(configPath)
		if err != nil {
			fmt.Printf("❌ Failed to read config file: %v\n", err)
			os.Exit(1)
		}

		// Validate YAML syntax
		var config interface{}
		if err := yaml.Unmarshal(data, &config); err != nil {
			fmt.Printf("❌ Invalid YAML syntax: %v\n", err)
			os.Exit(1)
		}

		// Convert to JSON for schema validation
		jsonData, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			fmt.Printf("❌ Failed to convert to JSON: %v\n", err)
			os.Exit(1)
		}

		// Load embedded schema
		schemaContent, err := schemaFS.ReadFile("schema/codebase-validation.schema.json")
		if err != nil {
			fmt.Printf("❌ Failed to load schema: %v\n", err)
			os.Exit(1)
		}

		// Validate against JSON Schema
		schemaLoader := gojsonschema.NewBytesLoader(schemaContent)
		documentLoader := gojsonschema.NewBytesLoader(jsonData)

		result, err := gojsonschema.Validate(schemaLoader, documentLoader)
		if err != nil {
			fmt.Printf("❌ Schema validation error: %v\n", err)
			os.Exit(1)
		}

		if !result.Valid() {
			fmt.Printf("❌ Configuration file has validation errors:\n\n")
			for i, desc := range result.Errors() {
				fmt.Printf("%d. %s\n", i+1, desc)
			}
			fmt.Printf("\n💡 Use 'codebase-interface schema -o schema.json' to get the full schema\n")
			os.Exit(1)
		}

		fmt.Printf("✅ Configuration file is valid: %s\n", configPath)
		fmt.Printf("\n📋 Configuration preview (JSON format):\n")
		fmt.Printf("%s\n", string(jsonData))

		fmt.Printf("\n💡 Schema documentation: https://cli.codebaseinterface.org/schema/\n")
	},
}

func init() {
	rootCmd.AddCommand(validateConfigCmd)
}
