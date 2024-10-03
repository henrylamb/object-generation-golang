package testingLite

type CodeTest struct {
	CodeQuality          int       `json:"codeQuality"`
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
	Feedback string `json:"feedback"`
}
