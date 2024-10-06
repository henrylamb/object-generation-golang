package testingLite

import "path/filepath"

// LanguageExtractor defines an interface for extracting the programming language from a file extension.
type LanguageExtractor interface {
	ExtractLanguage(fileName string) string
}

// DefaultLanguageExtractor holds the map of file extensions to languages and implements the LanguageExtractor interface.
type DefaultLanguageExtractor struct {
	languageMap map[string]string
}

// NewDefaultLanguageExtractor creates a new instance of DefaultLanguageExtractor with a predefined language map.
func NewDefaultLanguageExtractor() *DefaultLanguageExtractor {
	return &DefaultLanguageExtractor{
		languageMap: map[string]string{
			".go":   "Go",
			".py":   "Python",
			".js":   "JavaScript",
			".java": "Java",
			".cpp":  "C++",
			".rb":   "Ruby",
			".ts":   "TypeScript",
			".rs":   "Rust",
			// Additional languages can be added here without modifying the function
		},
	}
}

// AddLanguage allows adding a new language mapping to the extractor
func (e *DefaultLanguageExtractor) AddLanguage(extension, language string) {
	e.languageMap[extension] = language
}

// ExtractLanguage extracts the programming language from the file extension
func (e *DefaultLanguageExtractor) ExtractLanguage(fileName string) string {
	ext := filepath.Ext(fileName)
	if language, exists := e.languageMap[ext]; exists {
		return language
	}
	return "Unknown"
}
