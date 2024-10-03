package testingLite

import (
	"fmt"
	"os"
	"path/filepath"
)

func extractLanguagesAndContents(codePath string) ([]string, map[string]string, error) {
	// Map of file extensions to programming languages
	extensionToLanguage := map[string]string{
		".go":   "Go",
		".py":   "Python",
		".js":   "JavaScript",
		".java": "Java",
		".cpp":  "C++",
		".rb":   "Ruby",
		".ts":   "TypeScript",
		".rs":   "Rust",
		// Add more languages as needed
	}

	// Map to store the contents of each file
	fileContents := map[string]string{}

	// Set to store detected languages (to ensure uniqueness)
	languages := map[string]struct{}{}

	// Check if the path is a directory
	fileInfo, err := os.Stat(codePath)
	if err != nil {
		return nil, nil, fmt.Errorf("Error accessing path: %s", err.Error())
	}

	// If it's a directory, read all files from the directory
	if fileInfo.IsDir() {
		files, err := os.ReadDir(codePath)
		if err != nil {
			return nil, nil, fmt.Errorf("Error reading directory: %s", err.Error())
		}

		// Iterate through the files and process them
		for _, file := range files {
			if !file.IsDir() { // Ignore subdirectories
				ext := filepath.Ext(file.Name())
				filePath := filepath.Join(codePath, file.Name())

				// Read the content of the file
				content, err := os.ReadFile(filePath)
				if err != nil {
					return nil, nil, fmt.Errorf("Error reading file %s: %s", file.Name(), err.Error())
				}

				// Store the content
				fileContents[file.Name()] = string(content)

				// Identify the language based on file extension
				if lang, ok := extensionToLanguage[ext]; ok {
					languages[lang] = struct{}{}
				}
			}
		}
	} else {
		// If it's a single file, read and process that file
		ext := filepath.Ext(codePath)

		// Read the content of the file
		content, err := os.ReadFile(codePath)
		if err != nil {
			return nil, nil, fmt.Errorf("Error reading file: %s", err.Error())
		}

		// Store the content of the file
		fileContents[filepath.Base(codePath)] = string(content)

		// Identify the language based on file extension
		if lang, ok := extensionToLanguage[ext]; ok {
			languages[lang] = struct{}{}
		}
	}

	// Convert map keys (languages) to a slice to return the list of unique languages
	var languageList []string
	for lang := range languages {
		languageList = append(languageList, lang)
	}

	return languageList, fileContents, nil
}
