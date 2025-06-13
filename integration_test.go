package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestIntegrationBasicUsage(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go", "test")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Command failed: %v\nOutput: %s", err, string(output))
	}

	outputStr := string(output)

	if !strings.Contains(outputStr, "TypeScript") {
		t.Error("Expected 'TypeScript' in output")
	}
	if !strings.Contains(outputStr, "JavaScript") {
		t.Error("Expected 'JavaScript' in output")
	}
	if !strings.Contains(outputStr, "Markdown") {
		t.Error("Expected 'Markdown' in output")
	}
	if !strings.Contains(outputStr, "CSS") {
		t.Error("Expected 'CSS' in output")
	}
	if !strings.Contains(outputStr, "JSON") {
		t.Error("Expected 'JSON' in output")
	}
	if !strings.Contains(outputStr, "Total") {
		t.Error("Expected 'Total' in output")
	}
}

func TestIntegrationAllOption(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go", "--all", "test")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Command with --all failed: %v\nOutput: %s", err, string(output))
	}

	outputStr := string(output)

	if !strings.Contains(outputStr, "TypeScript") {
		t.Error("Expected 'TypeScript' in output with --all option")
	}
	if !strings.Contains(outputStr, "Total") {
		t.Error("Expected 'Total' in output with --all option")
	}
}

func TestIntegrationNonExistentDirectory(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go", "non-existent-dir")
	output, err := cmd.CombinedOutput()
	if err == nil {
		t.Error("Expected command to fail for non-existent directory")
	}

	outputStr := string(output)
	if !strings.Contains(outputStr, "存在しません") {
		t.Error("Expected Japanese error message for non-existent directory")
	}
}

func TestIntegrationNoArguments(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go")
	output, err := cmd.CombinedOutput()
	if err == nil {
		t.Error("Expected command to fail when no arguments provided")
	}

	outputStr := string(output)
	if !strings.Contains(outputStr, "使用方法") {
		t.Error("Expected Japanese usage message when no arguments provided")
	}
}

func TestIntegrationOutputFormat(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go", "test")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Command failed: %v\nOutput: %s", err, string(output))
	}

	outputStr := string(output)
	lines := strings.Split(strings.TrimSpace(outputStr), "\n")

	if len(lines) < 3 {
		t.Error("Expected at least 3 lines of output (header, separator, data)")
	}

	headerLine := lines[0]
	if !strings.Contains(headerLine, "FileType") || !strings.Contains(headerLine, "Lines") || !strings.Contains(headerLine, "Percent") {
		t.Error("Expected proper table header format")
	}

	separatorLine := lines[1]
	if !strings.Contains(separatorLine, "---") {
		t.Error("Expected separator line with dashes")
	}

	totalLine := lines[len(lines)-1]
	if !strings.Contains(totalLine, "Total") {
		t.Error("Expected total line at the end")
	}
}

func TestIntegrationPercentageCalculation(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go", "test")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Command failed: %v\nOutput: %s", err, string(output))
	}

	outputStr := string(output)

	if !strings.Contains(outputStr, "%") {
		t.Error("Expected percentage signs in output")
	}

	if !strings.Contains(outputStr, "100.0%") {
		t.Error("Expected total percentage to be 100.0%")
	}
}
