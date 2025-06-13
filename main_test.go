package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestExtensionMap_GetDisplayName(t *testing.T) {
	extMap := NewExtensionMap()

	tests := []struct {
		extension string
		expected  string
	}{
		{".ts", "TypeScript"},
		{".tsx", "TypeScript"},
		{".js", "JavaScript"},
		{".jsx", "JavaScript"},
		{".go", "Go"},
		{".py", "Python"},
		{".unknown", ".unknown"},
		{"", ""},
	}

	for _, test := range tests {
		result := extMap.GetDisplayName(test.extension)
		if result != test.expected {
			t.Errorf("GetDisplayName(%s) = %s; expected %s", test.extension, result, test.expected)
		}
	}
}

func TestCountLinesInFile(t *testing.T) {
	tempDir := t.TempDir()

	testFile := filepath.Join(tempDir, "test.txt")
	content := "line1\nline2\nline3\n"

	err := os.WriteFile(testFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("テストファイルの作成に失敗: %v", err)
	}

	lines, err := countLinesInFile(testFile)
	if err != nil {
		t.Fatalf("countLinesInFile() failed: %v", err)
	}

	if lines != 3 {
		t.Errorf("countLinesInFile() = %d; expected 3", lines)
	}
}

func TestCountLinesInEmptyFile(t *testing.T) {
	tempDir := t.TempDir()

	testFile := filepath.Join(tempDir, "empty.txt")

	err := os.WriteFile(testFile, []byte(""), 0644)
	if err != nil {
		t.Fatalf("テストファイルの作成に失敗: %v", err)
	}

	lines, err := countLinesInFile(testFile)
	if err != nil {
		t.Fatalf("countLinesInFile() failed: %v", err)
	}

	if lines != 0 {
		t.Errorf("countLinesInFile() = %d; expected 0", lines)
	}
}

func TestAnalyzeDirectory(t *testing.T) {
	tempDir := t.TempDir()

	files := map[string]string{
		"test.go":   "package main\n\nfunc main() {\n}\n",
		"script.js": "console.log('hello');\nconsole.log('world');\n",
		"style.css": "body { margin: 0; }\n",
		"README.md": "# Title\n\nContent here.\n",
	}

	for filename, content := range files {
		filePath := filepath.Join(tempDir, filename)
		err := os.WriteFile(filePath, []byte(content), 0644)
		if err != nil {
			t.Fatalf("テストファイル %s の作成に失敗: %v", filename, err)
		}
	}

	stats, err := analyzeDirectory(tempDir)
	if err != nil {
		t.Fatalf("analyzeDirectory() failed: %v", err)
	}

	expectedStats := map[string]int{
		"Go":         4,
		"JavaScript": 2,
		"CSS":        1,
		"Markdown":   3,
	}

	for fileType, expectedLines := range expectedStats {
		if stat, exists := stats[fileType]; !exists {
			t.Errorf("Expected file type %s not found", fileType)
		} else if stat.Lines != expectedLines {
			t.Errorf("File type %s: expected %d lines, got %d", fileType, expectedLines, stat.Lines)
		}
	}
}

func TestAnalyzeDirectoryNonExistent(t *testing.T) {
	_, err := analyzeDirectory("/non/existent/path")
	if err == nil {
		t.Error("Expected error for non-existent directory, got nil")
	}
}

func TestFileStatsStructure(t *testing.T) {
	stats := &FileStats{
		Extension: "TypeScript",
		Lines:     100,
		Files:     5,
	}

	if stats.Extension != "TypeScript" {
		t.Errorf("Extension = %s; expected TypeScript", stats.Extension)
	}
	if stats.Lines != 100 {
		t.Errorf("Lines = %d; expected 100", stats.Lines)
	}
	if stats.Files != 5 {
		t.Errorf("Files = %d; expected 5", stats.Files)
	}
}

func TestExtensionHandling(t *testing.T) {
	tempDir := t.TempDir()

	noExtFile := filepath.Join(tempDir, "Dockerfile")
	err := os.WriteFile(noExtFile, []byte("FROM ubuntu\nRUN apt-get update\n"), 0644)
	if err != nil {
		t.Fatalf("拡張子なしファイルの作成に失敗: %v", err)
	}

	stats, err := analyzeDirectory(tempDir)
	if err != nil {
		t.Fatalf("analyzeDirectory() failed: %v", err)
	}

	if _, exists := stats["Dockerfile"]; !exists {
		t.Error("拡張子なしファイル 'Dockerfile' が見つかりません")
	}
}

func TestCaseInsensitiveExtensions(t *testing.T) {
	extMap := NewExtensionMap()

	result1 := extMap.GetDisplayName(".JS")
	result2 := extMap.GetDisplayName(".js")

	if result1 != "JavaScript" || result2 != "JavaScript" {
		t.Errorf("Case insensitive test failed: .JS=%s, .js=%s", result1, result2)
	}
}

func TestSortingByLines(t *testing.T) {
	tempDir := t.TempDir()

	files := map[string]string{
		"small.txt": "line1\n",
		"medium.go": "package main\n\nfunc main() {\n  fmt.Println(\"hello\")\n}\n",
		"large.js":  strings.Repeat("console.log('test');\n", 10),
	}

	for filename, content := range files {
		filePath := filepath.Join(tempDir, filename)
		err := os.WriteFile(filePath, []byte(content), 0644)
		if err != nil {
			t.Fatalf("テストファイル %s の作成に失敗: %v", filename, err)
		}
	}

	stats, err := analyzeDirectory(tempDir)
	if err != nil {
		t.Fatalf("analyzeDirectory() failed: %v", err)
	}

	if len(stats) != 3 {
		t.Fatalf("Expected 3 file types, got %d", len(stats))
	}

	jsLines := stats["JavaScript"].Lines
	goLines := stats["Go"].Lines
	textLines := stats["Text"].Lines

	if jsLines <= goLines || goLines <= textLines {
		t.Errorf("Sorting verification failed: JS=%d, Go=%d, Text=%d", jsLines, goLines, textLines)
	}
}
