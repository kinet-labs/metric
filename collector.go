package metric

// Collector is a marker interface for compatibility with Registerer.
type Collector interface{}

// AsCollector returns a metric as a Collector for registration.
// Registration is a no-op for high-perf metrics, so this function simply
// returns the input value.
func AsCollector(v interface{}) Collector {
	return v
}
