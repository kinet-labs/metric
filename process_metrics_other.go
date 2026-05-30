//go:build !unix && !windows

package metric

func processCPUSeconds() (float64, bool) {
	return 0, false
}

func processResidentBytes() (float64, bool) {
	return 0, false
}
