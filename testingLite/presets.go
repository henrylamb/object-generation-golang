package testingLite

// / These values are the thresholds which can be used as a default for your code to pass their tests
var (
	// LenientTesting represents the most relaxed evaluation thresholds
	KindTesting = CodeTest{
		CodeQuality:          50,
		ProbabilityOfSuccess: 40,
		Security: &Security{
			LowSeverity:      15,
			MediumSeverity:   10,
			HighSeverity:     5,
			CriticalSeverity: 2,
		},
		Solid: &Solid{
			SingleResponsibilityScore: 30,
			OpenClosedScore:           30,
			LiskovSubstitutionScore:   30,
			InterfaceSegregationScore: 30,
			DependencyInversionScore:  30,
		},
	}
	LenientTesting = CodeTest{
		CodeQuality:          60,
		ProbabilityOfSuccess: 50,
		Security: &Security{
			LowSeverity:      10,
			MediumSeverity:   5,
			HighSeverity:     3,
			CriticalSeverity: 1,
		},
		Solid: &Solid{
			SingleResponsibilityScore: 50,
			OpenClosedScore:           50,
			LiskovSubstitutionScore:   50,
			InterfaceSegregationScore: 50,
			DependencyInversionScore:  50,
		},
	}

	// ModerateTesting represents a balanced evaluation with moderate thresholds
	ModerateTesting = CodeTest{
		CodeQuality:          70,
		ProbabilityOfSuccess: 70,
		Security: &Security{
			LowSeverity:      5,
			MediumSeverity:   3,
			HighSeverity:     1,
			CriticalSeverity: 0,
		},
		Solid: &Solid{
			SingleResponsibilityScore: 70,
			OpenClosedScore:           70,
			LiskovSubstitutionScore:   70,
			InterfaceSegregationScore: 70,
			DependencyInversionScore:  70,
		},
	}

	// StrictTesting represents the most stringent evaluation with strict thresholds
	StrictTesting = CodeTest{
		CodeQuality:          80,
		ProbabilityOfSuccess: 85,
		Security: &Security{
			LowSeverity:      3,
			MediumSeverity:   1,
			HighSeverity:     0,
			CriticalSeverity: 0,
		},
		Solid: &Solid{
			SingleResponsibilityScore: 80,
			OpenClosedScore:           80,
			LiskovSubstitutionScore:   80,
			InterfaceSegregationScore: 80,
			DependencyInversionScore:  80,
		},
	}

	// StrictTesting represents the most stringent evaluation with strict thresholds
	ExtremeTesting = CodeTest{
		CodeQuality:          90,
		ProbabilityOfSuccess: 95,
		Security: &Security{
			LowSeverity:      1,
			MediumSeverity:   0,
			HighSeverity:     0,
			CriticalSeverity: 0,
		},
		Solid: &Solid{
			SingleResponsibilityScore: 95,
			OpenClosedScore:           95,
			LiskovSubstitutionScore:   95,
			InterfaceSegregationScore: 95,
			DependencyInversionScore:  95,
		},
	}
)
