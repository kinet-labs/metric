package metric

// Registerer is the minimal interface required to register and create metrics.
type Registerer interface {
	Metrics
	Register(Collector) error
	MustRegister(...Collector)
}

// Registry is a registerer that can also gather metric families.
type Registry interface {
	Registerer
	Gatherer
}
