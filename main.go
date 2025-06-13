package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type FileStats struct {
	Extension string
	Lines     int
	Files     int
}

type ExtensionMap struct {
	extensions map[string]string
}

func NewExtensionMap() *ExtensionMap {
	return &ExtensionMap{
		extensions: map[string]string{
			".ts":    "TypeScript",
			".tsx":   "TypeScript",
			".js":    "JavaScript",
			".jsx":   "JavaScript",
			".py":    "Python",
			".go":    "Go",
			".java":  "Java",
			".c":     "C",
			".cpp":   "C++",
			".cc":    "C++",
			".cxx":   "C++",
			".h":     "C Header",
			".hpp":   "C++ Header",
			".cs":    "C#",
			".php":   "PHP",
			".rb":    "Ruby",
			".rs":    "Rust",
			".kt":    "Kotlin",
			".swift": "Swift",
			".dart":  "Dart",
			".scala": "Scala",
			".r":     "R",
			".m":     "Objective-C",
			".mm":    "Objective-C++",
			".sql":   "SQL",
			".html":  "HTML",
			".htm":   "HTML",
			".css":   "CSS",
			".scss":  "SCSS",
			".sass":  "Sass",
			".less":  "Less",
			".vue":   "Vue",
			".xml":   "XML",
			".json":  "JSON",
			".yaml":  "YAML",
			".yml":   "YAML",
			".toml":  "TOML",
			".ini":   "INI",
			".cfg":   "Config",
			".conf":  "Config",
			".md":    "Markdown",
			".txt":   "Text",
			".sh":    "Shell",
			".bash":  "Bash",
			".zsh":   "Zsh",
			".fish":  "Fish",
			".ps1":   "PowerShell",
			".bat":   "Batch",
			".cmd":   "Batch",
		},
	}
}

func (em *ExtensionMap) GetDisplayName(ext string) string {
	if displayName, exists := em.extensions[strings.ToLower(ext)]; exists {
		return displayName
	}
	return ext
}

func countLinesInFile(filePath string) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	// bufio.NewScannerはデフォルトで64KB制限があり、1行が長すぎると
	// "token too long" エラーが発生するため、バイト単位での読み取りを使用
	lineCount := 0
	reader := bufio.NewReader(file)
	
	for {
		_, err := reader.ReadBytes('\n')
		if err != nil {
			if err.Error() == "EOF" {
				// ファイル末尾に到達
				return lineCount, nil
			}
			return lineCount, err
		}
		lineCount++
	}
}

func analyzeDirectory(dirPath string) (map[string]*FileStats, error) {
	stats := make(map[string]*FileStats)
	extMap := NewExtensionMap()

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		ext := filepath.Ext(path)
		if ext == "" {
			ext = filepath.Base(path)
		}

		displayName := extMap.GetDisplayName(ext)

		lines, err := countLinesInFile(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "警告: %s の行数をカウントできませんでした: %v\n", path, err)
			return nil
		}

		if stat, exists := stats[displayName]; exists {
			stat.Lines += lines
			stat.Files++
		} else {
			stats[displayName] = &FileStats{
				Extension: displayName,
				Lines:     lines,
				Files:     1,
			}
		}

		return nil
	})

	return stats, err
}

func printTable(stats map[string]*FileStats, showAll bool) {
	if len(stats) == 0 {
		fmt.Println("ファイルが見つかりませんでした。")
		return
	}

	statsList := make([]*FileStats, 0, len(stats))
	totalLines := 0

	for _, stat := range stats {
		statsList = append(statsList, stat)
		totalLines += stat.Lines
	}

	sort.Slice(statsList, func(i, j int) bool {
		return statsList[i].Lines > statsList[j].Lines
	})

	displayStats := statsList
	if !showAll && len(statsList) > 5 {
		displayStats = statsList[:5]
	}

	fmt.Printf("| %-12s | %5s | %7s |\n", "FileType", "Lines", "Percent")
	fmt.Printf("| %s | %s | %s |\n", strings.Repeat("-", 12), strings.Repeat("-", 5), strings.Repeat("-", 7))

	for _, stat := range displayStats {
		percent := float64(stat.Lines) / float64(totalLines) * 100
		fmt.Printf("| %-12s | %5d | %6.1f%% |\n", stat.Extension, stat.Lines, percent)
	}

	fmt.Printf("| %s | %s | %s |\n", strings.Repeat("=", 12), strings.Repeat("=", 5), strings.Repeat("=", 7))
	fmt.Printf("| %-12s | %5d | %6.1f%% |\n", "Total", totalLines, 100.0)
}

func main() {
	var showAll bool
	flag.BoolVar(&showAll, "all", false, "全てのファイルタイプを表示")
	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "使用方法: %s [--all] <ディレクトリパス>\n", os.Args[0])
		os.Exit(1)
	}

	dirPath := args[0]

	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "エラー: ディレクトリ '%s' が存在しません\n", dirPath)
		os.Exit(1)
	}

	stats, err := analyzeDirectory(dirPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "エラー: ディレクトリの解析に失敗しました: %v\n", err)
		os.Exit(1)
	}

	printTable(stats, showAll)
}
