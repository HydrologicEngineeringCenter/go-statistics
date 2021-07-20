package statistics

import (
	"math"
)

type LogNormalDistribution struct {
	Mean              float64 `json:mean`
	StandardDeviation float64 `json:standarddeviation`
}

func (n LogNormalDistribution) InvCDF(probability float64) float64 {
	z := NormalDistribution{Mean: 0, StandardDeviation: 1}
	return math.Exp(n.Mean + z.InvCDF(probability)*n.StandardDeviation)
}
func (n LogNormalDistribution) CDF(value float64) float64 {
	z := NormalDistribution{Mean: n.Mean, StandardDeviation: n.StandardDeviation}
	return z.CDF(math.Log(value))
}
func (n LogNormalDistribution) PDF(value float64) float64 {
	z := NormalDistribution{Mean: n.Mean, StandardDeviation: n.StandardDeviation}
	return z.PDF(math.Log(value))
}
func (n LogNormalDistribution) CentralTendency() float64 {
	return n.Mean
}
