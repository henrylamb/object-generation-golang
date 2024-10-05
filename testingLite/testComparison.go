package testingLite

func TestComparison(yourTest *CodeTest, defaultTest *CodeTest) bool {
	// Check code quality
	if yourTest.CodeQuality < defaultTest.CodeQuality {
		return false
	}

	// Check probability of success
	if yourTest.ProbabilityOfSuccess < defaultTest.ProbabilityOfSuccess {
		return false
	}

	// Check security
	if yourTest.Security != nil {
		if yourTest.Security.LowSeverity > defaultTest.Security.LowSeverity {
			return false
		}
		if yourTest.Security.MediumSeverity > defaultTest.Security.MediumSeverity {
			return false
		}
		if yourTest.Security.HighSeverity > defaultTest.Security.HighSeverity {
			return false
		}
		if yourTest.Security.CriticalSeverity > defaultTest.Security.CriticalSeverity {
			return false
		}
	}

	// Check SOLID principles (if pointer is not nil)
	if yourTest.Solid != nil {
		// Single Responsibility Principle
		if yourTest.Solid.SingleResponsibilityScore < defaultTest.Solid.SingleResponsibilityScore {
			return false
		}
		// Open/Closed Principle
		if yourTest.Solid.OpenClosedScore < defaultTest.Solid.OpenClosedScore {
			return false
		}
		// Liskov Substitution Principle
		if yourTest.Solid.LiskovSubstitutionScore < defaultTest.Solid.LiskovSubstitutionScore {
			return false
		}
		// Interface Segregation Principle
		if yourTest.Solid.InterfaceSegregationScore < defaultTest.Solid.InterfaceSegregationScore {
			return false
		}
		// Dependency Inversion Principle
		if yourTest.Solid.DependencyInversionScore < defaultTest.Solid.DependencyInversionScore {
			return false
		}
	}

	// If all checks pass
	return true
}
