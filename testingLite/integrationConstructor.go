package testingLite

import (
	"fmt"
	"github.com/henrylamb/object-generation-golang/jsonSchema"
	"path/filepath"
	"strings"
)

// IntegrationTestConstructor constructs a definition based on the assumption and codePath
// Ensure that the model being used for this kind of testing has a very high token input count
// the multiple team recommends that the Gemini or Anthropic models are used
func IntegrationTestConstructor(assumption string, codePath string, model jsonSchema.ModelType) (*jsonSchema.Definition, string, error) {
	// Extract languages and file contents using the function created earlier
	_, fileContents, err := extractLanguagesAndContents(codePath)
	if err != nil {
		return nil, "", fmt.Errorf("Error processing code path: %s", err.Error())
	}

	// Prepare to compile the languages used into a single system prompt
	languageSet := map[string]struct{}{}
	var combinedCodeText strings.Builder

	// Iterate through each file's content and compile a unified system prompt
	for fileName, codeText := range fileContents {
		// Determine the language from the file extension
		ext := filepath.Ext(fileName)
		var language string
		switch ext {
		case ".go":
			language = "Go"
		case ".py":
			language = "Python"
		case ".js":
			language = "JavaScript"
		case ".java":
			language = "Java"
		case ".cpp":
			language = "C++"
		case ".rb":
			language = "Ruby"
		case ".ts":
			language = "TypeScript"
		case ".rs":
			language = "Rust"
		case ".cs":
			language = "C#"
		case ".php":
			language = "PHP"
		case ".swift":
			language = "Swift"
		case ".kt", ".kts":
			language = "Kotlin"
		case ".m", ".mm":
			language = "Objective-C"
		case ".scala":
			language = "Scala"
		case ".sh":
			language = "Shell"
		case ".pl":
			language = "Perl"
		case ".r":
			language = "R"
		case ".lua":
			language = "Lua"
		case ".dart":
			language = "Dart"
		case ".hs":
			language = "Haskell"
		case ".clj":
			language = "Clojure"
		case ".erl":
			language = "Erlang"
		case ".ex", ".exs":
			language = "Elixir"
		case ".sql":
			language = "SQL"
		case ".xml", ".html":
			language = "Markup"
		default:
			language = "Unknown"
		}

		// Add the language to the set if it is recognized
		if language != "Unknown" {
			languageSet[language] = struct{}{}
		}

		// Add a marker and file name to the combined code text
		combinedCodeText.WriteString(fmt.Sprintf("\n\n// Start of File: %s\n", fileName))
		combinedCodeText.WriteString(codeText)
		combinedCodeText.WriteString("\n// End of File\n")
	}

	// Create a list of languages used
	var languageList []string
	for lang := range languageSet {
		languageList = append(languageList, lang)
	}

	// Create a combined system prompt mentioning all languages involved
	systemPromptGeneral := fmt.Sprintf(
		"You are an expert in the following programming languages: %s. You are tasked with reviewing code for overall quality, maintainability, best practices, and integration between these languages.",
		strings.Join(languageList, ", "),
	)

	// Create a jsonSchema.Definition for the integration test
	integrationQA := &jsonSchema.Definition{
		Type:         jsonSchema.Object,
		Model:        model,
		Instruction:  fmt.Sprintf("Analyze the overall quality of the provided code for integration and multiple languages, and provide a summary evaluation. \n\n%s", assumption),
		SystemPrompt: &systemPromptGeneral,
		Properties: map[string]jsonSchema.Definition{
			"codeQuality": {
				Type:        jsonSchema.Number,
				Instruction: "Rate the overall quality of the integrated code on a scale from 1 to 10, where 1 is poor and 10 is excellent.",
			},
			"review": {
				Type:        jsonSchema.Object,
				Instruction: fmt.Sprintf("Provide a detailed review of the integrated code, including key strengths and weaknesses for integration across multiple languages. \n\n%s", assumption),
				Model:       model,
				Properties: map[string]jsonSchema.Definition{
					"feedback": {
						Type:        jsonSchema.String,
						Instruction: "Give specific feedback on the structure, readability, maintainability, and integration of the code across the detected languages.",
					},
				},
			},
			"security": {
				Type:        jsonSchema.Object,
				Model:       model,
				Instruction: fmt.Sprintf("Evaluate the security of the integrated code by analyzing vulnerabilities categorized by severity level across the different languages. \n\n%s", assumption),
				Properties: map[string]jsonSchema.Definition{
					"lowSeverity": {
						Type:        jsonSchema.Number,
						Instruction: "Identify the number of low-severity security vulnerabilities found in the integrated code.",
					},
					"mediumSeverity": {
						Type:        jsonSchema.Number,
						Instruction: "Identify the number of medium-severity security vulnerabilities found in the integrated code.",
					},
					"highSeverity": {
						Type:        jsonSchema.Number,
						Instruction: "Identify the number of high-severity security vulnerabilities found in the integrated code.",
					},
					"criticalSeverity": {
						Type:        jsonSchema.Number,
						Instruction: "Identify the number of critical-severity security vulnerabilities found in the integrated code.",
					},
				},
			},
			"probabilityOfSuccess": {
				Type:        jsonSchema.Number,
				Instruction: "Estimate the probability of successful deployment of the integrated code without significant issues, on a scale from 0 to 100%.",
			},
		},
	}

	// Return the integration definition and the combined code text
	return integrationQA, combinedCodeText.String(), nil
}
