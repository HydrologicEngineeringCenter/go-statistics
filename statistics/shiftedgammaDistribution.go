package statistics

import (
	"gonum.org/v1/gonum/stat/distuv"
)

type ShiftedGammaDistribution struct {
	Alpha float64 `json:"alpha"`
	Beta  float64 `json:"beta"`
	Shift float64 `json:"shift"`
}

func (d ShiftedGammaDistribution) InvCDF(probability float64) float64 {
	if probability > 1 {
		panic("nope")
	}
	if probability <= 0 {
		panic("nope")
	}
	g := distuv.Gamma{Alpha: d.Alpha, Beta: 1.0 / d.Beta}
	return g.Quantile(probability) + d.Shift
}
func (n ShiftedGammaDistribution) CDF(value float64) float64 {
	return 0.0
}
func (n ShiftedGammaDistribution) PDF(value float64) float64 {

	return 0.0
}
func (n ShiftedGammaDistribution) CentralTendency() float64 {
	return .5
}
