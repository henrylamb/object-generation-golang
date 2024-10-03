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
						Type: jsonSchema.Number,
						Instruction: `Identify the number of low-severity security vulnerabilities found in the integrated code. 
	Low-severity vulnerabilities are typically minor issues that have a limited impact on the system's overall security posture. 
	These issues often represent small misconfigurations, or minor information exposures that do not pose an immediate threat but may 
	indicate areas for improvement.

	Examples of low-severity vulnerabilities:
	- Minor input validation issues that do not lead to security bypass (related to OWASP A07:2021 - Identification and Authentication Failures).
	- Weak or overly verbose error messages that could give attackers minor information (related to OWASP A09:2021 - Security Logging and Monitoring Failures).
	- Insufficiently strict Content Security Policies (CSP) that don't create immediate risks but could be exploited under certain conditions (related to OWASP A05:2021 - Security Misconfiguration).

	Only report observable, low-risk issues. Return a value of 0 if none are found.`,
					},
					"mediumSeverity": {
						Type: jsonSchema.Number,
						Instruction: `Identify the number of medium-severity security vulnerabilities found in the integrated code. 
	Medium-severity vulnerabilities represent issues that could be exploited in more specific or controlled circumstances, 
	and may allow an attacker to gain some unauthorized access or leakage of information. These vulnerabilities often require 
	a combination of flaws to be exploited or depend on certain configurations.

	Examples of medium-severity vulnerabilities:
	- Insecure Direct Object References (related to OWASP A01:2021 - Broken Access Control) that allow limited access to unauthorized resources.
	- Missing or improper authentication methods for non-critical functionality (related to OWASP A07:2021 - Identification and Authentication Failures).
	- Storing sensitive information in non-secure ways that is hard to access but still could be leaked in certain conditions (related to OWASP A03:2021 - Injection).

	Only report actual, exploitable medium-severity issues. Return 0 if no such vulnerabilities are found.`,
					},
					"highSeverity": {
						Type: jsonSchema.Number,
						Instruction: `Identify the number of high-severity security vulnerabilities found in the integrated code. 
	High-severity vulnerabilities can result in significant impacts, such as unauthorized access to critical resources or the ability to 
	manipulate sensitive data. These vulnerabilities typically require immediate attention as they can lead to data breaches or other severe 
	consequences if exploited.

	Examples of high-severity vulnerabilities:
	- SQL injection vulnerabilities that allow attackers to execute arbitrary queries on the database (related to OWASP A03:2021 - Injection).
	- Missing access control checks on sensitive operations (related to OWASP A01:2021 - Broken Access Control).
	- Exposing sensitive data, such as passwords or personal information, in unencrypted storage or over insecure channels (related to OWASP A02:2021 - Cryptographic Failures).

	These vulnerabilities pose a clear risk to data integrity or confidentiality. Only report actual high-severity vulnerabilities, and return a value of 0 if none are present.`,
					},
					"criticalSeverity": {
						Type: jsonSchema.Number,
						Instruction: `Identify the number of critical-severity security vulnerabilities found in the integrated code. 
	Critical-severity vulnerabilities are the most severe and represent issues that, if exploited, could allow an attacker to completely compromise the system, 
	execute arbitrary code, or perform any action within the system with full privileges. These are typically vulnerabilities that must be resolved immediately 
	to prevent catastrophic consequences.

	Examples of critical-severity vulnerabilities:
	- Remote code execution vulnerabilities that allow attackers to run arbitrary commands on the system (related to OWASP A06:2021 - Vulnerable and Outdated Components or A10:2021 - Server-Side Request Forgery).
	- Authentication bypasses that allow attackers full access to the system without valid credentials (related to OWASP A01:2021 - Broken Access Control).
	- Unrestricted file uploads that could lead to server compromise (related to OWASP A08:2021 - Software and Data Integrity Failures).

	Only report vulnerabilities of this severity if they are directly observable in the code. Return 0 if no critical vulnerabilities are found.`,
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
