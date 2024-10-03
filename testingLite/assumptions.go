package testingLite

const (
	// Core assumptions
	WorkingAssumption = "Assume that all the functions/methods or requests will return as intended."
	BrokenAssumption  = "Assume that all the functions/methods or requests will NOT return as intended."
	// Additional assumptions
	PartialSuccessAssumption         = "Assume that only some functions/methods will return as intended, while others will fail."
	BoundaryConditionAssumption      = "Assume that input values are at or near their boundary limits (e.g., minimum, maximum, or extreme values)."
	ConcurrencyAssumption            = "Assume that multiple threads or processes will call the same function concurrently, potentially causing race conditions."
	PerformanceDegradationAssumption = "Assume that the system is under heavy load, causing performance degradation, such as slow database queries or high CPU usage."
	NetworkInstabilityAssumption     = "Assume that the network is unreliable, leading to intermittent timeouts or packet loss."
	InvalidInputAssumption           = "Assume that the input provided to functions or methods is invalid, such as malformed data or out-of-range values."
	LatencyAndDelayAssumption        = "Assume that certain operations, such as database queries or network requests, will take significantly longer than expected."
	DataCorruptionAssumption         = "Assume that some data will be corrupted or incomplete (e.g., truncated or missing fields)."
	SecurityBreachAssumption         = "Assume that security vulnerabilities exist, such as unauthorized access or injection attacks."
	StateManagementAssumption        = "Assume that state or session management may be inconsistent, such as stale cache or race conditions."
	ResourceLimitAssumption          = "Assume that system resources, such as memory, CPU, or disk space, are limited or exhausted."
	CompatibilityAssumption          = "Assume that the code is running in different environments or platforms, such as different OS versions or browser types."
	ExternalDependencyAssumption     = "Assume that external dependencies, such as third-party APIs or libraries, may be unavailable or malfunctioning."
	LegacySystemAssumption           = "Assume that certain components or dependencies are outdated or incompatible with newer code."
	UserBehaviorAssumption           = "Assume that users will interact with the system in unexpected ways, such as rapid clicks or entering unexpected input."
	NoAssumptions                    = ""
)
