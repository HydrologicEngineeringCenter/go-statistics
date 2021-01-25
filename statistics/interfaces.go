package statistics

type ContinuousDistribution interface {
	InvCDF(probability float64) float64
	CDF(value float64) float64
	PDF(value float64) float64
	CentralTendency() float64
}
type FittableDistribution interface {
	Fit(inputData []float64) // could have an interface FittableDistribution or something like that
}
