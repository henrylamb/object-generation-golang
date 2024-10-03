package testingLite

// these values are the thresholds which can be used as a defualt for your code to pass their tests
var (
	// LenientTesting represents the most relaxed evaluation thresholds
	LenientTesting = CodeTest{
		CodeQuality:          60,
		ProbabilityOfSuccess: 50,
		Security: &Security{
			LowSeverity:      10,
			MediumSeverity:   5,
			HighSeverity:     3,
			CriticalSeverity: 1,
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
	}
)
