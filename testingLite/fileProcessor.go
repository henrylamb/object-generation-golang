package testingLite

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// FileProcessor handles file operations, including checking for test files.
type FileProcessor struct{}

// isTestFile is now a method of FileProcessor that checks if a file is a test file
func (fp *FileProcessor) isTestFile(fileName string) bool {
	return strings.HasSuffix(fileName, "_test.go") || // Go
		strings.Contains(fileName, ".test.") || // JS/TS
		strings.Contains(fileName, ".spec.") || // JS/TS
		strings.HasSuffix(fileName, "Test.java") || // Java
		strings.HasSuffix(fileName, "_test.py") || // Python
		strings.HasSuffix(fileName, "_test.rb") || // Ruby
		strings.HasSuffix(fileName, "Test.php") || // PHP
		strings.HasSuffix(fileName, "_test.cpp") || // C++
		strings.HasSuffix(fileName, "Test.swift") || // Swift
		strings.HasSuffix(fileName, "Test.kt") || // Kotlin
		strings.HasSuffix(fileName, "_test.exs") || // Elixir
		strings.HasSuffix(fileName, "_test.cs") || // C#
		strings.HasSuffix(fileName, "Test.m") || // Objective-C
		strings.HasSuffix(fileName, "Spec.groovy") || // Groovy
		strings.HasSuffix(fileName, "Test.groovy") || // Groovy
		strings.HasSuffix(fileName, "Test.scala") || // Scala
		strings.Contains(fileName, "test_") || // Python (alternate)
		strings.HasSuffix(fileName, "Test.rb") // Ruby (alternate)
}

// extractFileContents processes a file or directory and returns the file contents (excluding test files)
func (fp *FileProcessor) extractFileContents(codePath string) (map[string]string, error) {
	fileContents := map[string]string{}

	fileInfo, err := os.Stat(codePath)
	if err != nil {
		return nil, fmt.Errorf("Error accessing path: %s", err.Error())
	}

	if fileInfo.IsDir() {
		err = fp.processDirectory(codePath, fileContents)
		if err != nil {
			return nil, err
		}
	} else {
		err = fp.processFile(codePath, fileContents)
		if err != nil {
			return nil, err
		}
	}

	return fileContents, nil
}

// processDirectory processes all files in a directory and skips test files
func (fp *FileProcessor) processDirectory(directoryPath string, fileContents map[string]string) error {
	files, err := os.ReadDir(directoryPath)
	if err != nil {
		return fmt.Errorf("Error reading directory: %s", err.Error())
	}

	for _, file := range files {
		if !file.IsDir() {
			filePath := filepath.Join(directoryPath, file.Name())
			err = fp.processFile(filePath, fileContents)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// processFile processes a single file and updates fileContents
func (fp *FileProcessor) processFile(filePath string, fileContents map[string]string) error {
	fileName := filepath.Base(filePath)

	// Skip test files
	if fp.isTestFile(fileName) {
		return nil
	}

	// Read the file content
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("Error reading file %s: %s", fileName, err.Error())
	}

	fileContents[fileName] = string(content)
	return nil
}
