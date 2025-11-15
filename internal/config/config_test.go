package config

import (
	"os"
	"testing"
)

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig()

	if cfg == nil {
		t.Fatal("DefaultConfig returned nil")
	}

	if !cfg.Validation.Agents.EssentialFiles.Enabled {
		t.Error("Essential files agent should be enabled by default")
	}

	if !cfg.Validation.Agents.EssentialFiles.RequireReadme {
		t.Error("README should be required by default")
	}

	if !cfg.Validation.Agents.EssentialFiles.RequireContributing {
		t.Error("CONTRIBUTING should be required by default")
	}

	if !cfg.Validation.Agents.GitConfiguration.Enabled {
		t.Error("Git configuration agent should be enabled by default")
	}

	if !cfg.Validation.Agents.GitConfiguration.RequireGitignore {
		t.Error("Gitignore should be required by default")
	}

	if !cfg.Validation.Agents.GitConfiguration.RequireEditorconfig {
		t.Error("Editorconfig should be required by default")
	}

	if cfg.Validation.Agents.GitConfiguration.RequireGitattributes {
		t.Error("Gitattributes should not be required by default")
	}

	if !cfg.Validation.Agents.DevelopmentStandards.Enabled {
		t.Error("Development standards agent should be enabled by default")
	}

	if !cfg.Validation.Agents.DevelopmentStandards.CheckCommitHistory {
		t.Error("Commit history check should be enabled by default")
	}

	if cfg.Validation.Agents.DevelopmentStandards.CommitHistoryDepth != 10 {
		t.Errorf("Expected commit history depth 10, got %d", cfg.Validation.Agents.DevelopmentStandards.CommitHistoryDepth)
	}

	if cfg.Validation.Output.Format != "table" {
		t.Errorf("Expected default format 'table', got %s", cfg.Validation.Output.Format)
	}

	if cfg.Validation.Output.Verbose {
		t.Error("Verbose should be false by default")
	}
}

func TestLoad_NoConfigFile(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "config-test-")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	cfg, err := Load(tmpDir)
	if err != nil {
		t.Errorf("Load failed: %v", err)
	}

	defaultCfg := DefaultConfig()
	if cfg.Validation.Agents.EssentialFiles.Enabled != defaultCfg.Validation.Agents.EssentialFiles.Enabled {
		t.Error("Loaded config doesn't match default when no file exists")
	}
}

func TestIsAgentEnabled(t *testing.T) {
	cfg := DefaultConfig()

	tests := []struct {
		agentName string
		expected  bool
	}{
		{"essential-files", true},
		{"git-configuration", true},
		{"development-standards", true},
		{"unknown-agent", false},
	}

	for _, tt := range tests {
		t.Run(tt.agentName, func(t *testing.T) {
			result := cfg.IsAgentEnabled(tt.agentName)
			if result != tt.expected {
				t.Errorf("IsAgentEnabled(%s) = %v, expected %v", tt.agentName, result, tt.expected)
			}
		})
	}
}
