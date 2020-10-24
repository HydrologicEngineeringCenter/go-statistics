package data

type InlineStats interface {
	AddObservation(value float64)
	AddObservations(value []float64)
}
type ConvergenceTestable interface {
	TestForConvergence(minConfidenceLimit float64, maxConfidenceLimit float64, zAlpha float64, relativeError float64) (bool, int64)
}
