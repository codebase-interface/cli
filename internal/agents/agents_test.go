package agents

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/codebase-interface/cli/internal/config"
)

func TestEssentialFilesAgent_Validate(t *testing.T) {
	tests := []struct {
		name           string
		setupFiles     []string
		expectedScore  float64
		expectedStatus string
		expectedFiles  []string
	}{
		{
			name:           "all files present",
			setupFiles:     []string{"README.md", "CONTRIBUTING.md"},
			expectedScore:  1.0,
			expectedStatus: "pass",
			expectedFiles:  []string{"README.md", "CONTRIBUTING.md"},
		},
		{
			name:           "only README present",
			setupFiles:     []string{"README.md"},
			expectedScore:  0.5,
			expectedStatus: "fail",
			expectedFiles:  []string{"README.md"},
		},
		{
			name:           "no files present",
			setupFiles:     []string{},
			expectedScore:  0.0,
			expectedStatus: "fail",
			expectedFiles:  []string{},
		},
		{
			name:           "README.rst variant",
			setupFiles:     []string{"README.rst", "CONTRIBUTING.md"},
			expectedScore:  1.0,
			expectedStatus: "pass",
			expectedFiles:  []string{"README.rst", "CONTRIBUTING.md"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create temporary directory
			tmpDir, err := os.MkdirTemp("", "codebase-test-")
			if err != nil {
				t.Fatalf("Failed to create temp dir: %v", err)
			}
			defer os.RemoveAll(tmpDir)

			// Setup files
			for _, filename := range tt.setupFiles {
				filePath := filepath.Join(tmpDir, filename)
				if err := os.WriteFile(filePath, []byte("test content"), 0644); err != nil {
					t.Fatalf("Failed to create test file %s: %v", filename, err)
				}
			}

			// Run validation
			agent := NewEssentialFilesAgent()
			cfg := config.DefaultConfig()
			result, err := agent.Validate(tmpDir, cfg)

			// Check results
			if err != nil {
				t.Errorf("Validation failed: %v", err)
			}

			if result.Agent != "essential-files" {
				t.Errorf("Expected agent 'essential-files', got %s", result.Agent)
			}

			if result.Score != tt.expectedScore {
				t.Errorf("Expected score %f, got %f", tt.expectedScore, result.Score)
			}

			if result.Status != tt.expectedStatus {
				t.Errorf("Expected status %s, got %s", tt.expectedStatus, result.Status)
			}
		})
	}
}

func TestGitConfigurationAgent_Validate(t *testing.T) {
	tests := []struct {
		name           string
		setupFiles     []string
		expectedScore  float64
		expectedStatus string
	}{
		{
			name:           "all required files present",
			setupFiles:     []string{".gitignore", ".editorconfig"},
			expectedScore:  1.0,
			expectedStatus: "pass",
		},
		{
			name:           "only gitignore present",
			setupFiles:     []string{".gitignore"},
			expectedScore:  0.5,
			expectedStatus: "fail",
		},
		{
			name:           "no files present",
			setupFiles:     []string{},
			expectedScore:  0.0,
			expectedStatus: "fail",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpDir, err := os.MkdirTemp("", "codebase-test-")
			if err != nil {
				t.Fatalf("Failed to create temp dir: %v", err)
			}
			defer os.RemoveAll(tmpDir)

			for _, filename := range tt.setupFiles {
				filePath := filepath.Join(tmpDir, filename)
				if err := os.WriteFile(filePath, []byte("test content"), 0644); err != nil {
					t.Fatalf("Failed to create test file %s: %v", filename, err)
				}
			}

			agent := NewGitConfigurationAgent()
			cfg := config.DefaultConfig()
			result, err := agent.Validate(tmpDir, cfg)

			if err != nil {
				t.Errorf("Validation failed: %v", err)
			}

			if result.Score != tt.expectedScore {
				t.Errorf("Expected score %f, got %f", tt.expectedScore, result.Score)
			}

			if result.Status != tt.expectedStatus {
				t.Errorf("Expected status %s, got %s", tt.expectedStatus, result.Status)
			}
		})
	}
}

func TestRegistry(t *testing.T) {
	registry := NewRegistry()

	// Test registration
	agent := NewEssentialFilesAgent()
	registry.Register("test-agent", agent)

	// Test get
	retrievedAgent, exists := registry.Get("test-agent")
	if !exists {
		t.Error("Expected agent to exist in registry")
	}
	if retrievedAgent != agent {
		t.Error("Retrieved agent doesn't match registered agent")
	}

	// Test get non-existent
	_, exists = registry.Get("non-existent")
	if exists {
		t.Error("Expected non-existent agent to not exist")
	}

	// Test All
	all := registry.All()
	if len(all) != 1 {
		t.Errorf("Expected 1 agent in registry, got %d", len(all))
	}
}
