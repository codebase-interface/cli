package agents

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/codebase-interface/cli/internal/config"
)

type ValidationResult struct {
	Agent    string    `json:"agent"`
	Status   string    `json:"status"` // pass, fail, warning
	Score    float64   `json:"score"`  // 0.0-1.0
	Findings []Finding `json:"findings"`
}

type Finding struct {
	Type     string `json:"type"` // missing, present, invalid
	File     string `json:"file"`
	Message  string `json:"message"`
	Severity string `json:"severity"` // critical, warning, info
}

type Agent interface {
	Validate(targetPath string, cfg *config.Config) (ValidationResult, error)
}

type Registry struct {
	agents map[string]Agent
}

func NewRegistry() *Registry {
	return &Registry{
		agents: make(map[string]Agent),
	}
}

func (r *Registry) Register(name string, agent Agent) {
	r.agents[name] = agent
}

func (r *Registry) Get(name string) (Agent, bool) {
	agent, exists := r.agents[name]
	return agent, exists
}

func (r *Registry) All() map[string]Agent {
	return r.agents
}

type EssentialFilesAgent struct{}

func NewEssentialFilesAgent() *EssentialFilesAgent {
	return &EssentialFilesAgent{}
}

func (a *EssentialFilesAgent) Validate(targetPath string, cfg *config.Config) (ValidationResult, error) {
	result := ValidationResult{
		Agent:    "essential-files",
		Status:   "pass",
		Score:    1.0,
		Findings: []Finding{},
	}

	agentCfg := cfg.Validation.Agents.EssentialFiles
	totalChecks := 0
	passedChecks := 0

	if agentCfg.RequireReadme {
		totalChecks++
		readmePaths := []string{"README.md", "README.rst", "readme.md", "readme.rst"}
		found := false

		for _, readmePath := range readmePaths {
			if _, err := os.Stat(filepath.Join(targetPath, readmePath)); err == nil {
				found = true
				result.Findings = append(result.Findings, Finding{
					Type:     "present",
					File:     readmePath,
					Message:  fmt.Sprintf("%s present", readmePath),
					Severity: "info",
				})
				break
			}
		}

		if found {
			passedChecks++
		} else {
			result.Findings = append(result.Findings, Finding{
				Type:     "missing",
				File:     "README.md",
				Message:  "README.md or README.rst missing",
				Severity: "critical",
			})
		}
	}

	if agentCfg.RequireContributing {
		totalChecks++
		contributingPath := filepath.Join(targetPath, "CONTRIBUTING.md")

		if _, err := os.Stat(contributingPath); err == nil {
			passedChecks++
			result.Findings = append(result.Findings, Finding{
				Type:     "present",
				File:     "CONTRIBUTING.md",
				Message:  "CONTRIBUTING.md present",
				Severity: "info",
			})
		} else {
			result.Findings = append(result.Findings, Finding{
				Type:     "missing",
				File:     "CONTRIBUTING.md",
				Message:  "CONTRIBUTING.md missing",
				Severity: "critical",
			})
		}
	}

	if totalChecks > 0 {
		result.Score = float64(passedChecks) / float64(totalChecks)
	}

	if result.Score < 1.0 {
		result.Status = "fail"
	}

	return result, nil
}

type GitConfigurationAgent struct{}

func NewGitConfigurationAgent() *GitConfigurationAgent {
	return &GitConfigurationAgent{}
}

func (a *GitConfigurationAgent) Validate(targetPath string, cfg *config.Config) (ValidationResult, error) {
	result := ValidationResult{
		Agent:    "git-configuration",
		Status:   "pass",
		Score:    1.0,
		Findings: []Finding{},
	}

	agentCfg := cfg.Validation.Agents.GitConfiguration
	totalChecks := 0
	passedChecks := 0

	if agentCfg.RequireGitignore {
		totalChecks++
		gitignorePath := filepath.Join(targetPath, ".gitignore")

		if _, err := os.Stat(gitignorePath); err == nil {
			passedChecks++
			result.Findings = append(result.Findings, Finding{
				Type:     "present",
				File:     ".gitignore",
				Message:  ".gitignore present",
				Severity: "info",
			})
		} else {
			result.Findings = append(result.Findings, Finding{
				Type:     "missing",
				File:     ".gitignore",
				Message:  ".gitignore missing",
				Severity: "critical",
			})
		}
	}

	if agentCfg.RequireGitattributes {
		totalChecks++
		gitattributesPath := filepath.Join(targetPath, ".gitattributes")

		if _, err := os.Stat(gitattributesPath); err == nil {
			passedChecks++
			result.Findings = append(result.Findings, Finding{
				Type:     "present",
				File:     ".gitattributes",
				Message:  ".gitattributes present",
				Severity: "info",
			})
		} else {
			result.Findings = append(result.Findings, Finding{
				Type:     "missing",
				File:     ".gitattributes",
				Message:  ".gitattributes missing (optional)",
				Severity: "warning",
			})
		}
	}

	if agentCfg.RequireEditorconfig {
		totalChecks++
		editorconfigPath := filepath.Join(targetPath, ".editorconfig")

		if _, err := os.Stat(editorconfigPath); err == nil {
			passedChecks++
			result.Findings = append(result.Findings, Finding{
				Type:     "present",
				File:     ".editorconfig",
				Message:  ".editorconfig present",
				Severity: "info",
			})
		} else {
			result.Findings = append(result.Findings, Finding{
				Type:     "missing",
				File:     ".editorconfig",
				Message:  ".editorconfig missing",
				Severity: "critical",
			})
		}
	}

	if totalChecks > 0 {
		result.Score = float64(passedChecks) / float64(totalChecks)
	}

	if result.Score < 1.0 {
		result.Status = "fail"
	}

	return result, nil
}

type DevelopmentStandardsAgent struct{}

func NewDevelopmentStandardsAgent() *DevelopmentStandardsAgent {
	return &DevelopmentStandardsAgent{}
}

func (a *DevelopmentStandardsAgent) Validate(targetPath string, cfg *config.Config) (ValidationResult, error) {
	result := ValidationResult{
		Agent:    "development-standards",
		Status:   "pass",
		Score:    1.0,
		Findings: []Finding{},
	}

	agentCfg := cfg.Validation.Agents.DevelopmentStandards
	totalChecks := 0
	passedChecks := 0

	if agentCfg.CheckCommitHistory && agentCfg.RequireConventionalCommits {
		totalChecks++

		if hasConventionalCommits, err := a.checkConventionalCommits(targetPath, agentCfg.CommitHistoryDepth); err != nil {
			result.Findings = append(result.Findings, Finding{
				Type:     "invalid",
				File:     "git-history",
				Message:  fmt.Sprintf("Failed to check commit history: %v", err),
				Severity: "warning",
			})
		} else if hasConventionalCommits {
			passedChecks++
			result.Findings = append(result.Findings, Finding{
				Type:     "present",
				File:     "git-history",
				Message:  "Recent commits follow conventional format",
				Severity: "info",
			})
		} else {
			result.Findings = append(result.Findings, Finding{
				Type:     "invalid",
				File:     "git-history",
				Message:  "Recent commits don't follow conventional format",
				Severity: "critical",
			})
		}
	}

	totalChecks++
	if branchValid, branchName, err := a.checkBranchNaming(targetPath); err != nil {
		result.Findings = append(result.Findings, Finding{
			Type:     "invalid",
			File:     "git-branch",
			Message:  fmt.Sprintf("Failed to check branch naming: %v", err),
			Severity: "warning",
		})
	} else if branchValid {
		passedChecks++
		result.Findings = append(result.Findings, Finding{
			Type:     "present",
			File:     "git-branch",
			Message:  fmt.Sprintf("Branch naming follows conventions: %s", branchName),
			Severity: "info",
		})
	} else {
		result.Findings = append(result.Findings, Finding{
			Type:     "invalid",
			File:     "git-branch",
			Message:  fmt.Sprintf("Branch name doesn't follow conventions: %s", branchName),
			Severity: "warning",
		})
	}

	if totalChecks > 0 {
		result.Score = float64(passedChecks) / float64(totalChecks)
	}

	if result.Score < 1.0 {
		result.Status = "fail"
	}

	return result, nil
}

func (a *DevelopmentStandardsAgent) checkConventionalCommits(targetPath string, depth int) (bool, error) {
	cmd := exec.Command("git", "log", "--oneline", fmt.Sprintf("-%d", depth))
	cmd.Dir = targetPath

	output, err := cmd.Output()
	if err != nil {
		return false, fmt.Errorf("git log failed: %w", err)
	}

	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	if len(lines) == 0 {
		return true, nil
	}

	conventionalPattern := regexp.MustCompile(`^[a-f0-9]+ (feat|fix|docs|style|refactor|test|chore|perf|ci|build|revert)(\(.+\))?!?: .+`)

	validCommits := 0
	for _, line := range lines {
		if conventionalPattern.MatchString(line) {
			validCommits++
		}
	}

	threshold := float64(len(lines)) * 0.8
	return float64(validCommits) >= threshold, nil
}

func (a *DevelopmentStandardsAgent) checkBranchNaming(targetPath string) (bool, string, error) {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	cmd.Dir = targetPath

	output, err := cmd.Output()
	if err != nil {
		return false, "", fmt.Errorf("git branch check failed: %w", err)
	}

	branchName := strings.TrimSpace(string(output))

	patterns := []*regexp.Regexp{
		regexp.MustCompile(`^(feature|feat)/.+`),
		regexp.MustCompile(`^(fix|bugfix)/.+`),
		regexp.MustCompile(`^(hotfix|patch)/.+`),
		regexp.MustCompile(`^(release|rel)/.+`),
		regexp.MustCompile(`^(docs|documentation)/.+`),
		regexp.MustCompile(`^(chore|task)/.+`),
		regexp.MustCompile(`^(main|master|develop|development)$`),
	}

	for _, pattern := range patterns {
		if pattern.MatchString(branchName) {
			return true, branchName, nil
		}
	}

	return false, branchName, nil
}
