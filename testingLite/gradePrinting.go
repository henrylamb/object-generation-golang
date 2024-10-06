package testingLite

import "fmt"

// evaluateGrade assigns a grade (A, B, C, D) based on the thresholds and test scores
func EvaluateGrade(testVal *CodeTest) string {
	if TestComparison(testVal, &ExtremeTesting) {
		return "A*"
	} else if TestComparison(testVal, &StrictTesting) {
		return "A"
	} else if TestComparison(testVal, &ModerateTesting) {
		return "B"
	} else if TestComparison(testVal, &LenientTesting) {
		return "C"
	} else if TestComparison(testVal, &KindTesting) {
		return "D"
	} else {
		return "F"
	}
}

// PrintResult prints the result based on verbose mode
func PrintResult(grade string, testVal *CodeTest, verbose bool) {
	if verbose {
		// Verbose output: includes detailed scores and grade
		fmt.Printf("Grade: %s\n", grade)
		fmt.Printf("Code Quality: %d\n", testVal.CodeQuality)
		fmt.Printf("Probability of Success: %d\n", testVal.ProbabilityOfSuccess)
		fmt.Printf("Security: Low Severity: %d, Medium Severity: %d, High Severity: %d, Critical Severity: %d\n",
			testVal.Security.LowSeverity, testVal.Security.MediumSeverity, testVal.Security.HighSeverity, testVal.Security.CriticalSeverity)
		fmt.Printf("SOLID Scores: SRP: %d, OCP: %d, LSP: %d, ISP: %d, DIP: %d\n",
			testVal.Solid.SingleResponsibilityScore, testVal.Solid.OpenClosedScore, testVal.Solid.LiskovSubstitutionScore,
			testVal.Solid.InterfaceSegregationScore, testVal.Solid.DependencyInversionScore)
	} else {
		// Non-verbose output: only shows grade
		fmt.Printf("Grade: %s\n", grade)
	}
}
