package statistics

import (
	"math"

	"gonum.org/v1/gonum/stat/distuv"
)

type PearsonIIIDistribution struct {
	Mean              float64 `json:"mean"`
	StandardDeviation float64 `json:"standarddeviation"`
	Skew              float64 `json:"skew"`
}

func (d PearsonIIIDistribution) InvCDF(probability float64) float64 {
	if probability > 1 {
		panic("nope")
	}
	if probability <= 0 {
		panic("nope")
	}
	noSkew := .00001
	if math.Abs(d.Skew) < noSkew {
		z := NormalDistribution{Mean: d.Mean, StandardDeviation: d.StandardDeviation}
		return z.InvCDF(probability)
	} else {
		shift := 0.0
		alpha := 4.0 / (d.Skew * d.Skew)
		beta := .5 * d.StandardDeviation * d.Skew
		if d.Skew > 0 {
			shift = d.Mean - 2.0*d.StandardDeviation/d.Skew
			g := distuv.Gamma{Alpha: alpha, Beta: 1.0 / beta}
			return g.Quantile(probability) + shift
		} else {
			beta = -beta
			shift = -d.Mean + 2.0*d.StandardDeviation/d.Skew
			g := distuv.Gamma{Alpha: alpha, Beta: 1.0 / beta}
			return -(g.Quantile(1-probability) + shift)
		}
	}

}
func (n PearsonIIIDistribution) CDF(value float64) float64 {

	return 0.0
}
func (n PearsonIIIDistribution) PDF(value float64) float64 {

	return 0.0
}
func (n PearsonIIIDistribution) CentralTendency() float64 {
	return 0.0
}
