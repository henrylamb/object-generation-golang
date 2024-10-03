package testingLite

import (
	"fmt"
	"github.com/henrylamb/object-generation-golang/jsonSchema"
	"path/filepath"
)

func SingleUnitTestWrapper(assumption string, codePath string, model jsonSchema.ModelType) (*jsonSchema.Definition, string, error) {
	definitions, code, err := UnitTestConstructor(assumption, codePath, model)
	if err != nil {
		return nil, "", err
	}

	return definitions[0], code[0], nil
}

// TestConstructor constructs a definition based on the assumption and codePath
// this will return the definition to be sent and the prompt to GoR
func UnitTestConstructor(assumption string, codePath string, model jsonSchema.ModelType) ([]*jsonSchema.Definition, []string, error) {
	// Extract languages and file contents using the function created earlier
	_, fileContents, err := extractLanguagesAndContents(codePath)
	if err != nil {
		return nil, nil, fmt.Errorf("Error processing code path: %s", err.Error())
	}

	// Prepare to return a list of definitions and code texts
	var definitions []*jsonSchema.Definition
	var codeTexts []string

	// Iterate through each file's content and create a corresponding Definition
	for fileName, codeText := range fileContents {
		// Get the file extension to determine the language
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
		// Add more languages as needed
		default:
			language = "Unknown"
		}

		// Construct the system prompts for each section
		systemPromptGeneral := fmt.Sprintf("You are an expert in %s programming language, and you are tasked with reviewing code for overall quality, maintainability, and best practices.", language)
		systemPromptReview := fmt.Sprintf("You are reviewing the code written in %s, and you should provide detailed feedback on its structure, readability, and maintainability.", language)
		systemPromptSecurity := fmt.Sprintf("You are evaluating the security aspects of the code written in %s, identifying vulnerabilities and potential risks.", language)
		systemPromptSuccess := fmt.Sprintf("You are estimating the probability of successful deployment for the code written in %s without significant issues.", language)

		// Construct the jsonSchema.Definition for the current file
		codeQA := &jsonSchema.Definition{
			Type:         jsonSchema.Object,
			Instruction:  fmt.Sprintf("Analyze the overall quality of the provided code and provide a summary evaluation. \n\n %s", assumption),
			SystemPrompt: &systemPromptGeneral,
			Model:        model,
			Properties: map[string]jsonSchema.Definition{
				"codeQuality": {
					Type:        jsonSchema.Number,
					Instruction: "Rate the overall quality of the code on a scale from 1 to 100, where 1 is poor and 100 is excellent.",
				},
				"review": {
					Type:         jsonSchema.Object,
					SystemPrompt: &systemPromptReview,
					Instruction:  fmt.Sprintf("Provide a detailed review of the code, including key strengths and weaknesses. \n\n %s", assumption),
					Properties: map[string]jsonSchema.Definition{
						"feedback": {
							Type:        jsonSchema.String,
							Instruction: "Give specific feedback on the code’s structure, readability, maintainability, and any issues detected.",
						},
					},
				},
				"security": {
					Type:         jsonSchema.Object,
					Instruction:  fmt.Sprintf("Evaluate the security of the code by analyzing vulnerabilities across severity levels. Each severity level (low, medium, high, and critical) must be reported based on actual, observable issues. If no vulnerabilities of a certain severity level are present, report a value of 0. Make sure to differentiate clearly between theoretical risks and actual vulnerabilities present in the code.\n\n%s", assumption),
					SystemPrompt: &systemPromptSecurity,
					Model:        model,
					SelectFields: []string{"review"},
					Properties: map[string]jsonSchema.Definition{
						"lowSeverity": {
							Type:        jsonSchema.Number,
							Instruction: "Identify the number of low-severity security vulnerabilities found in the code. Low-severity vulnerabilities are minor issues that do not pose a significant risk but should be addressed. Only report actual low-severity vulnerabilities found in the code. For example, minor information exposure that doesn't lead to direct threats.",
						},
						"mediumSeverity": {
							Type:        jsonSchema.Number,
							Instruction: "Identify the number of medium-severity security vulnerabilities found in the code. Medium-severity vulnerabilities may allow limited access or leakage of data, but typically require specific conditions to be exploited. Provide examples such as hardcoded sensitive information or minor access control issues. Report a 0 if none are found.",
						},
						"highSeverity": {
							Type:        jsonSchema.Number,
							Instruction: "Identify the number of high-severity security vulnerabilities found in the code. High-severity vulnerabilities can lead to significant data loss, unauthorized access, or system compromise if exploited. Examples include SQL injections or lack of input validation for critical operations. Ensure this number is based only on real, observable issues, and report 0 if none are found.",
						},
						"criticalSeverity": {
							Type:        jsonSchema.Number,
							Voters:      true,
							Instruction: "Identify the number of critical-severity security vulnerabilities found in the code. Critical-severity vulnerabilities are the most severe, allowing immediate or total system compromise. Examples include unrestricted remote code execution or major authentication bypasses. Only report these vulnerabilities if explicitly found in the code, and return 0 if none exist.",
						},
					},
				},
				"probabilityOfSuccess": {
					Type:         jsonSchema.Number,
					SystemPrompt: &systemPromptSuccess,
					Instruction:  "Estimate the probability of the code’s successfully running given the context of the code provided, on a scale from 0 to 100.",
				},
			},
		}

		// Add the generated codeQA and the file content to their respective slices
		definitions = append(definitions, codeQA)
		codeTexts = append(codeTexts, codeText)
	}

	// Return the list of jsonSchema.Definition and the corresponding code texts
	return definitions, codeTexts, nil
}
