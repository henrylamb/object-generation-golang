package testingLite

import (
	"fmt"
	"strings"
)

type CodeTest struct {
	CodeQuality          int       `json:"codeQuality"`
	Solid                *Solid    `json:"solid"`
	Review               *Review   `json:"review"`
	Security             *Security `json:"security"`
	ProbabilityOfSuccess int       `json:"probabilityOfSuccess"`
}

type Security struct {
	LowSeverity      int `json:"lowSeverity"`
	MediumSeverity   int `json:"mediumSeverity"`
	HighSeverity     int `json:"highSeverity"`
	CriticalSeverity int `json:"criticalSeverity"`
}

type Review struct {
	Feedback string `json:"feedback"` //this is the feedback from a review that an LLM will be providing as to tell the developer what they need to improve
}

// Solid struct representing the five SOLID principles.
type Solid struct {
	SingleResponsibilityScore int `json:"singleResponsibilityScore"` // SRP - Single Responsibility Principle
	OpenClosedScore           int `json:"openClosedScore"`           // OCP - Open/Closed Principle
	LiskovSubstitutionScore   int `json:"liskovSubstitutionScore"`   // LSP - Liskov Substitution Principle
	InterfaceSegregationScore int `json:"interfaceSegregationScore"` // ISP - Interface Segregation Principle
	DependencyInversionScore  int `json:"dependencyInversionScore"`  // DIP - Dependency Inversion Principle
}

// Print method for CodeTest struct (returns a string)
func (ct *CodeTest) Print() string {
	var sb strings.Builder

	// Build the string
	sb.WriteString("CodeTest Details:\n")
	sb.WriteString(fmt.Sprintf("  Code Quality: %d\n", ct.CodeQuality))
	sb.WriteString(fmt.Sprintf("  Probability of Success: %d%%\n", ct.ProbabilityOfSuccess))

	if ct.Review != nil {
		sb.WriteString("  Review:\n")
		sb.WriteString(fmt.Sprintf("    Feedback: %s\n", ct.Review.Feedback))
	} else {
		sb.WriteString("  Review: <nil>\n")
	}

	if ct.Security != nil {
		sb.WriteString("  Security:\n")
		sb.WriteString(fmt.Sprintf("    Low Severity: %d\n", ct.Security.LowSeverity))
		sb.WriteString(fmt.Sprintf("    Medium Severity: %d\n", ct.Security.MediumSeverity))
		sb.WriteString(fmt.Sprintf("    High Severity: %d\n", ct.Security.HighSeverity))
		sb.WriteString(fmt.Sprintf("    Critical Severity: %d\n", ct.Security.CriticalSeverity))
	} else {
		sb.WriteString("  Security: <nil>\n")
	}

	if ct.Solid != nil {
		sb.WriteString("  SOLID Scores:\n")
		sb.WriteString(fmt.Sprintf("    Single Responsibility Principle: %d\n", ct.Solid.SingleResponsibilityScore))
		sb.WriteString(fmt.Sprintf("    Open/Closed Principle: %d\n", ct.Solid.OpenClosedScore))
		sb.WriteString(fmt.Sprintf("    Liskov Substitution Principle: %d\n", ct.Solid.LiskovSubstitutionScore))
		sb.WriteString(fmt.Sprintf("    Interface Segregation Principle: %d\n", ct.Solid.InterfaceSegregationScore))
		sb.WriteString(fmt.Sprintf("    Dependency Inversion Principle: %d\n", ct.Solid.DependencyInversionScore))
	} else {
		sb.WriteString("  SOLID Scores: <nil>\n")
	}

	// Return the built string
	return sb.String()
}
