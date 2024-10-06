package testingLite

// Comparison defines the interface for a component that can compare itself with another component.
type Comparison interface {
	IsBetterThan(yourTest *CodeTest, defaultTest *CodeTest) bool
}

type QualityComparator struct{}

func (qc *QualityComparator) IsBetterThan(yourTest *CodeTest, defaultTest *CodeTest) bool {
	return yourTest.CodeQuality >= defaultTest.CodeQuality
}

type SuccessProbabilityComparator struct{}

func (spc *SuccessProbabilityComparator) IsBetterThan(yourTest *CodeTest, defaultTest *CodeTest) bool {
	return yourTest.ProbabilityOfSuccess >= defaultTest.ProbabilityOfSuccess
}

type SecurityComparator struct{}

func (sc *SecurityComparator) IsBetterThan(yourTest *CodeTest, defaultTest *CodeTest) bool {
	if yourTest.Security == nil || defaultTest.Security == nil {
		return true // If security is nil, don't fail the comparison
	}

	return yourTest.Security.LowSeverity <= defaultTest.Security.LowSeverity &&
		yourTest.Security.MediumSeverity <= defaultTest.Security.MediumSeverity &&
		yourTest.Security.HighSeverity <= defaultTest.Security.HighSeverity &&
		yourTest.Security.CriticalSeverity <= defaultTest.Security.CriticalSeverity
}

type SolidPrinciplesComparator struct{}

func (spc *SolidPrinciplesComparator) IsBetterThan(yourTest *CodeTest, defaultTest *CodeTest) bool {
	if yourTest.Solid == nil || defaultTest.Solid == nil {
		return true // If Solid principles are not applicable, pass comparison
	}

	return yourTest.Solid.SingleResponsibilityScore >= defaultTest.Solid.SingleResponsibilityScore &&
		yourTest.Solid.OpenClosedScore >= defaultTest.Solid.OpenClosedScore &&
		yourTest.Solid.LiskovSubstitutionScore >= defaultTest.Solid.LiskovSubstitutionScore &&
		yourTest.Solid.InterfaceSegregationScore >= defaultTest.Solid.InterfaceSegregationScore &&
		yourTest.Solid.DependencyInversionScore >= defaultTest.Solid.DependencyInversionScore
}

// TestComparison function uses comparators to compare different aspects
func TestComparison(yourTest *CodeTest, defaultTest *CodeTest) bool {
	// List of comparators
	comparators := []Comparison{
		&QualityComparator{},
		&SuccessProbabilityComparator{},
		&SecurityComparator{},
		&SolidPrinciplesComparator{},
	}

	// Iterate through each comparator and check if the comparison fails
	for _, comparator := range comparators {
		if !comparator.IsBetterThan(yourTest, defaultTest) {
			return false
		}
	}

	// All checks passed
	return true
}
