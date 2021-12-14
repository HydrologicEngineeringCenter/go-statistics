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
func (d ShiftedGammaDistribution) CDF(value float64) float64 {
	g := distuv.Gamma{Alpha: d.Alpha, Beta: 1.0 / d.Beta}
	v := value - d.Shift
	return g.CDF(v)
}
func (d ShiftedGammaDistribution) PDF(value float64) float64 {
	g := distuv.Gamma{Alpha: d.Alpha, Beta: 1.0 / d.Beta}
	return g.Prob(value - d.Shift)
}
func (n ShiftedGammaDistribution) CentralTendency() float64 {
	return .5
}
