//go:build !metrics

package metric

// NewRegistry returns a no-op registry when metrics are disabled.
func NewRegistry() Registry {
	return NewNoOpRegistry()
}
