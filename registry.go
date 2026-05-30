//go:build metrics

package metric

// NewRegistry returns a new in-process registry.
func NewRegistry() Registry {
	return newRegistry()
}
