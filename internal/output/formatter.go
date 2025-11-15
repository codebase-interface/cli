package output

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/codebase-interface/cli/internal/agents"
)

type Formatter interface {
	Format(results []agents.ValidationResult, writer io.Writer) error
}

func NewFormatter(format string) (Formatter, error) {
	switch format {
	case "json":
		return &JSONFormatter{}, nil
	case "table":
		return &TableFormatter{}, nil
	default:
		return nil, fmt.Errorf("unsupported format: %s", format)
	}
}

type JSONFormatter struct{}

func (f *JSONFormatter) Format(results []agents.ValidationResult, writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "  ")
	return encoder.Encode(results)
}

type TableFormatter struct{}

func (f *TableFormatter) Format(results []agents.ValidationResult, writer io.Writer) error {
	successStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("2")).Bold(true)
	failStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("1")).Bold(true)
	warningStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("3")).Bold(true)
	infoStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("4"))

	var output strings.Builder
	var totalScore float64
	var criticalFailures bool

	for _, result := range results {
		totalScore += result.Score

		var statusSymbol, statusText string
		var statusStyle lipgloss.Style

		switch result.Status {
		case "pass":
			statusSymbol = "✓"
			statusText = "PASS"
			statusStyle = successStyle
		case "fail":
			statusSymbol = "✗"
			statusText = "FAIL"
			statusStyle = failStyle
			criticalFailures = true
		case "warning":
			statusSymbol = "⚠"
			statusText = "WARN"
			statusStyle = warningStyle
		default:
			statusSymbol = "?"
			statusText = "UNKNOWN"
			statusStyle = infoStyle
		}

		agentTitle := fmt.Sprintf("%s %s Agent - %s (Score: %.1f)",
			statusSymbol,
			strings.Title(strings.ReplaceAll(result.Agent, "-", " ")),
			statusText,
			result.Score,
		)

		output.WriteString(statusStyle.Render(agentTitle))
		output.WriteString("\n")

		for _, finding := range result.Findings {
			var symbol string
			var style lipgloss.Style

			switch finding.Severity {
			case "critical":
				if finding.Type == "missing" || finding.Type == "invalid" {
					symbol = "  ✗"
					style = failStyle
				} else {
					symbol = "  ✓"
					style = successStyle
				}
			case "warning":
				symbol = "  ⚠"
				style = warningStyle
			case "info":
				symbol = "  ✓"
				style = successStyle
			default:
				symbol = "  ℹ"
				style = infoStyle
			}

			findingText := fmt.Sprintf("%s %s", symbol, finding.Message)
			output.WriteString(style.Render(findingText))
			output.WriteString("\n")
		}

		output.WriteString("\n")
	}

	overallScore := totalScore / float64(len(results))
	var overallStatus string
	var overallStyle lipgloss.Style

	if criticalFailures || overallScore < 0.8 {
		overallStatus = "FAIL"
		overallStyle = failStyle
	} else if overallScore < 1.0 {
		overallStatus = "PASS (with warnings)"
		overallStyle = warningStyle
	} else {
		overallStatus = "PASS"
		overallStyle = successStyle
	}

	overallText := fmt.Sprintf("Overall Score: %.2f - %s", overallScore, overallStatus)
	output.WriteString(overallStyle.Render(overallText))
	output.WriteString("\n")

	_, err := writer.Write([]byte(output.String()))
	return err
}
