package testingLite

import "log"

func TestComparison(yourTest *CodeTest, defaultTest *CodeTest) bool {
	// Check code quality
	if yourTest.CodeQuality < defaultTest.CodeQuality {
		log.Println("Code quality: ", yourTest.CodeQuality, " < ", defaultTest.CodeQuality)
		return false
	}

	// Check probability of success
	if yourTest.ProbabilityOfSuccess < defaultTest.ProbabilityOfSuccess {
		log.Println("Probability of success: ", yourTest.ProbabilityOfSuccess, " < ", defaultTest.ProbabilityOfSuccess)
		return false
	}

	// Check security
	if yourTest.Security != nil {
		if yourTest.Security.LowSeverity > defaultTest.Security.LowSeverity {
			log.Println("Low severity: ", yourTest.Security.LowSeverity, " > ", defaultTest.Security.LowSeverity)
			return false
		}
		if yourTest.Security.MediumSeverity > defaultTest.Security.MediumSeverity {
			log.Println("Medium severity: ", yourTest.Security.MediumSeverity, " > ", defaultTest.Security.MediumSeverity)
			return false
		}
		if yourTest.Security.HighSeverity > defaultTest.Security.HighSeverity {
			log.Println("High severity: ", yourTest.Security.HighSeverity, " > ", defaultTest.Security.HighSeverity)
			return false
		}
		if yourTest.Security.CriticalSeverity > defaultTest.Security.CriticalSeverity {
			log.Println("Critical severity: ", yourTest.Security.CriticalSeverity, " > ", defaultTest.Security.CriticalSeverity)
			return false
		}
	}

	// Check SOLID principles (if pointer is not nil)
	if yourTest.Solid != nil {
		// Single Responsibility Principle
		if yourTest.Solid.SingleResponsibilityScore < defaultTest.Solid.SingleResponsibilityScore {
			log.Println("Single Responsibility Score: ", yourTest.Solid.SingleResponsibilityScore, " < ", defaultTest.Solid.SingleResponsibilityScore)
			return false
		}
		// Open/Closed Principle
		if yourTest.Solid.OpenClosedScore < defaultTest.Solid.OpenClosedScore {
			log.Println("Open/Closed Score: ", yourTest.Solid.OpenClosedScore, " < ", defaultTest.Solid.OpenClosedScore)
			return false
		}
		// Liskov Substitution Principle
		if yourTest.Solid.LiskovSubstitutionScore < defaultTest.Solid.LiskovSubstitutionScore {
			log.Println("Liskov Substitution Score: ", yourTest.Solid.LiskovSubstitutionScore, " < ", defaultTest.Solid.LiskovSubstitutionScore)
			return false
		}
		// Interface Segregation Principle
		if yourTest.Solid.InterfaceSegregationScore < defaultTest.Solid.InterfaceSegregationScore {
			log.Println("Interface Segregation Score: ", yourTest.Solid.InterfaceSegregationScore, " < ", defaultTest.Solid.InterfaceSegregationScore)
			return false
		}
		// Dependency Inversion Principle
		if yourTest.Solid.DependencyInversionScore < defaultTest.Solid.DependencyInversionScore {
			log.Println("Dependency Inversion Score: ", yourTest.Solid.DependencyInversionScore, " < ", defaultTest.Solid.DependencyInversionScore)
			return false
		}
	}

	// If all checks pass
	return true
}
