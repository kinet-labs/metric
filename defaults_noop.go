//go:build !metrics

package metric

func init() {
	DefaultRegistry = NewNoOpRegistry()
	defaultFactory = NewNoOpFactory()
}
