package testingLite

import (
	"fmt"
	"github.com/henrylamb/object-generation-golang/jsonSchema"
	"strings"
)

// CodeProcessor handles the responsibility of processing code and generating prompts
type CodeProcessor interface {
	ProcessCode(assumption string, fileContents map[string]string, model jsonSchema.ModelType) (*jsonSchema.Definition, string, error)
}

type DefaultCodeProcessor struct{}

func NewDefaultCodeProcessor() *DefaultCodeProcessor {
	return &DefaultCodeProcessor{}
}

func (p *DefaultCodeProcessor) ProcessCode(assumption string, fileContents map[string]string, model jsonSchema.ModelType) (*jsonSchema.Definition, string, error) {
	languageSet := map[string]struct{}{}
	var combinedCodeText strings.Builder

	for fileName, codeText := range fileContents {
		language := extractLanguage(fileName)

		if language != "Unknown" {
			languageSet[language] = struct{}{}
		}

		combinedCodeText.WriteString(fmt.Sprintf("\n\n// Start of File: %s\n", fileName))
		combinedCodeText.WriteString(codeText)
		combinedCodeText.WriteString("\n// End of File\n")
	}

	languageList := make([]string, 0, len(languageSet))
	for lang := range languageSet {
		languageList = append(languageList, lang)
	}

	language := strings.Join(languageList, ", ")
	systemPromptGeneral := fmt.Sprintf(
		"You are an expert in the following programming languages: %s. You are tasked with reviewing code for overall quality, maintainability, best practices, and integration between these languages. The separate files are all merged together, do not fault the code for being structured like this.",
		language,
	)

	// Construct the system prompts for each section
	systemPromptReview := fmt.Sprintf("You are reviewing the code written in %s, and you should provide detailed feedback on its structure, readability, and maintainability.", language)
	systemPromptSecurity := fmt.Sprintf("You are evaluating the security aspects of the code written in %s, identifying vulnerabilities and potential risks.", language)
	systemPromptSuccess := fmt.Sprintf("You are estimating the probability of successful deployment for the code written in %s without significant issues.", language)

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
				Type:         jsonSchema.Number,
				SystemPrompt: &systemPromptSuccess,
				Instruction:  "Estimate the probability of the code’s successfully running given the context of the code provided, on a scale from 0 to 100.",
			},
			"solid": {
				Type:         jsonSchema.Object,
				Instruction:  fmt.Sprintf("Evaluate the code based on the five SOLID principles (Single Responsibility, Open/Closed, Liskov Substitution, Interface Segregation, Dependency Inversion). Rate each principle from 0 to 100, where 10 is poor adherence and 90 is excellent adherence. \n\n %s", assumption),
				SystemPrompt: &systemPromptGeneral,
				Properties: map[string]jsonSchema.Definition{
					"singleResponsibilityScore": {
						Type:        jsonSchema.Number,
						Instruction: "Rate the code’s adherence to the Single Responsibility Principle (SRP). This principle states that a class should have only one reason to change, meaning it should have only one job or responsibility.",
					},
					"openClosedScore": {
						Type:        jsonSchema.Number,
						Instruction: "Rate the code’s adherence to the Open/Closed Principle (OCP). This principle states that a class should be open for extension but closed for modification, meaning it should allow new functionality to be added without changing existing code.",
					},
					"liskovSubstitutionScore": {
						Type:        jsonSchema.Number,
						Instruction: "Rate the code’s adherence to the Liskov Substitution Principle (LSP). This principle states that objects of a superclass should be replaceable with objects of its subclasses without affecting the correctness of the program.",
					},
					"interfaceSegregationScore": {
						Type:        jsonSchema.Number,
						Instruction: "Rate the code’s adherence to the Interface Segregation Principle (ISP). This principle states that a client should not be forced to implement an interface that it does not use, meaning interfaces should be specific to the client’s needs.",
					},
					"dependencyInversionScore": {
						Type:        jsonSchema.Number,
						Instruction: "Rate the code’s adherence to the Dependency Inversion Principle (DIP). This principle states that high-level modules should not depend on low-level modules, but both should depend on abstractions, meaning the code should rely on interfaces rather than concrete implementations.",
					},
				},
			},
		},
	}

	return codeQA, combinedCodeText.String(), nil
}
