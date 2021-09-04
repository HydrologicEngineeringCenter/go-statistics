package statistics

import (
	"math"
)

type LogPearsonIIIDistribution struct {
	Mean              float64 `json:"mean"`
	StandardDeviation float64 `json:"standarddeviation"`
	Skew              float64 `json:"skew"`
}

func (n LogPearsonIIIDistribution) InvCDF(probability float64) float64 {
	if probability > 1 {
		panic("nope")
	}
	if probability <= 0 {
		panic("nope")
	}

	z := PearsonIIIDistribution{Mean: n.Mean, StandardDeviation: n.StandardDeviation, Skew: n.Skew}
	return math.Pow(10, z.InvCDF(probability))
}
func (n LogPearsonIIIDistribution) CDF(value float64) float64 {

	return 0.0
}
func (n LogPearsonIIIDistribution) PDF(value float64) float64 {

	return 0.0
}
func (n LogPearsonIIIDistribution) CentralTendency() float64 {
	return 0.0
}