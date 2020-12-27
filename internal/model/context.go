package model

type Tally struct {
	// Day in which the value was observed
	Day int
	// Observed value
	Metric float64
}

type Period struct {
	First int
	Last int

	// Used to tally the observed metric associated with the condition.
	MetricMax float64
}

// Structure used to hold state while computing climate conditions in batch
type BatchContext struct{
	// Condition indexed map of predicted conditions
	Conditions map[string]struct {
		Periods []Period
	}
}