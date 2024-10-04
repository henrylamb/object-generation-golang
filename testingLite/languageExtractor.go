package testingLite

import "path/filepath"

func extractLanguage(fileName string) string {
	ext := filepath.Ext(fileName)
	languageMap := map[string]string{
		".go":   "Go",
		".py":   "Python",
		".js":   "JavaScript",
		".java": "Java",
		".cpp":  "C++",
		".rb":   "Ruby",
		".ts":   "TypeScript",
		".rs":   "Rust",
		// Add more languages here...
	}
	if language, exists := languageMap[ext]; exists {
		return language
	}
	return "Unknown"
}
