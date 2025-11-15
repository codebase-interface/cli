package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Validation ValidationConfig `yaml:"validation"`
}

type ValidationConfig struct {
	Agents AgentsConfig `yaml:"agents"`
	Output OutputConfig `yaml:"output"`
}

type AgentsConfig struct {
	EssentialFiles       EssentialFilesConfig       `yaml:"essential-files"`
	GitConfiguration     GitConfigurationConfig     `yaml:"git-configuration"`
	DevelopmentStandards DevelopmentStandardsConfig `yaml:"development-standards"`
}

type EssentialFilesConfig struct {
	Enabled             bool `yaml:"enabled"`
	RequireReadme       bool `yaml:"require_readme"`
	RequireContributing bool `yaml:"require_contributing"`
}

type GitConfigurationConfig struct {
	Enabled              bool `yaml:"enabled"`
	RequireGitignore     bool `yaml:"require_gitignore"`
	RequireGitattributes bool `yaml:"require_gitattributes"`
	RequireEditorconfig  bool `yaml:"require_editorconfig"`
}

type DevelopmentStandardsConfig struct {
	Enabled                    bool `yaml:"enabled"`
	CheckCommitHistory         bool `yaml:"check_commit_history"`
	CommitHistoryDepth         int  `yaml:"commit_history_depth"`
	RequireConventionalCommits bool `yaml:"require_conventional_commits"`
}

type OutputConfig struct {
	Format  string `yaml:"format"`
	Verbose bool   `yaml:"verbose"`
}

func DefaultConfig() *Config {
	return &Config{
		Validation: ValidationConfig{
			Agents: AgentsConfig{
				EssentialFiles: EssentialFilesConfig{
					Enabled:             true,
					RequireReadme:       true,
					RequireContributing: true,
				},
				GitConfiguration: GitConfigurationConfig{
					Enabled:              true,
					RequireGitignore:     true,
					RequireGitattributes: false,
					RequireEditorconfig:  true,
				},
				DevelopmentStandards: DevelopmentStandardsConfig{
					Enabled:                    true,
					CheckCommitHistory:         true,
					CommitHistoryDepth:         10,
					RequireConventionalCommits: true,
				},
			},
			Output: OutputConfig{
				Format:  "table",
				Verbose: false,
			},
		},
	}
}

func Load(targetPath string) (*Config, error) {
	cfg := DefaultConfig()

	configPath := filepath.Join(targetPath, ".codebase-validation.yml")

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return cfg, nil
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	return cfg, nil
}

func (c *Config) IsAgentEnabled(agentName string) bool {
	switch agentName {
	case "essential-files":
		return c.Validation.Agents.EssentialFiles.Enabled
	case "git-configuration":
		return c.Validation.Agents.GitConfiguration.Enabled
	case "development-standards":
		return c.Validation.Agents.DevelopmentStandards.Enabled
	default:
		return false
	}
}
