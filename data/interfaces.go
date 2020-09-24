package data

type InlineStats interface{
	AddObservation(value float64)
	AddObservations(value []float64)
}