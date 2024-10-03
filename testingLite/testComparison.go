package testingLite

import "log"

func TestComparison(yourTest *CodeTest, defaultTest *CodeTest) bool {
	if yourTest.CodeQuality < defaultTest.CodeQuality {
		log.Println("Code quality: ", yourTest.CodeQuality, " < ", defaultTest.CodeQuality)
		return false
	}
	if yourTest.ProbabilityOfSuccess < defaultTest.ProbabilityOfSuccess {
		log.Println("Probability of success: ", yourTest.ProbabilityOfSuccess, " < ", defaultTest.ProbabilityOfSuccess)
		return false
	}
	if yourTest.Security != nil {
		if yourTest.Security.LowSeverity > defaultTest.Security.LowSeverity {
			log.Println("Low severity: ", yourTest.Security.LowSeverity, " < ", defaultTest.Security.LowSeverity)
			return false
		}
		if yourTest.Security.MediumSeverity > defaultTest.Security.MediumSeverity {
			log.Println("Medium severity: ", yourTest.Security.MediumSeverity, " < ", defaultTest.Security.MediumSeverity)
			return false
		}
		if yourTest.Security.HighSeverity > defaultTest.Security.HighSeverity {
			log.Println("High severity: ", yourTest.Security.HighSeverity, " < ", defaultTest.Security.HighSeverity)
			return false
		}
		if yourTest.Security.CriticalSeverity > defaultTest.Security.CriticalSeverity {
			log.Println("Critical severity: ", yourTest.Security.CriticalSeverity, " < ", defaultTest.Security.CriticalSeverity)
			return false
		}
	}
	return true
}
