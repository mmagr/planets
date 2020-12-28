package model

type Tally struct {
	// Day in which the value was observed
	Day int
	// Observed value
	Metric float64
}

func (t *Tally) Observe(day int, metric float64) {
	if metric > t.Metric {
		t.Day = day
		t.Metric = metric
	}
}

type Period struct {
	First int
	Last int

	// Used to tally the observed metric associated with the condition.
	Tally Tally
}

// Structure used to hold state while computing climate conditions in batch
type BatchContext struct{
	// Condition indexed map of predicted conditions
	Conditions map[string][]Period
}

func (bc *BatchContext) Observe(day int, condition string, metric float64) {
	if bc.Conditions == nil {
		bc.Conditions = map[string][]Period{}
	}
	periods, exists := bc.Conditions[condition]
	if exists == false {
		periods = make([]Period, 0, 10)
		bc.Conditions[condition] = periods
	}

	defaultPeriod := Period{day, day, Tally{day, metric}}
	target := defaultPeriod
	if len(periods) > 0 {
		target = periods[len(periods) - 1]
		if target.Last < (day - 1) {
			target = defaultPeriod
		} else {
			target.Last = day
			target.Tally.Observe(day, metric)
			return
		}
	}

	bc.Conditions[condition] = append(periods, target)
}
