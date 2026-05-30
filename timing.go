package metric

// TimingMetric measures durations and records them in a histogram.
type TimingMetric = timingMetric

// NewTimingMetric creates a timing metric bound to the provided histogram.
func NewTimingMetric(histogram Histogram) *TimingMetric {
	return newTimingMetric(histogram)
}
