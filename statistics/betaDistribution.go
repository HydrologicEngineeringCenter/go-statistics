package statistics

import (
	"gonum.org/v1/gonum/stat/distuv"
)

type BetaDistribution struct {
	Alpha float64 `json:"alpha"`
	Beta  float64 `json:"beta"`
}

func (d BetaDistribution) InvCDF(probability float64) float64 {
	if probability > 1 {
		panic("nope")
	}
	if probability <= 0 {
		panic("nope")
	}
	b := distuv.Beta{Alpha: d.Alpha, Beta: d.Beta}
	return b.Quantile(probability)
}

func (d BetaDistribution) CDF(value float64) float64 {
	b := distuv.Beta{Alpha: d.Alpha, Beta: d.Beta}
	return b.CDF(value)
}

func (d BetaDistribution) PDF(value float64) float64 {
	b := distuv.Beta{Alpha: d.Alpha, Beta: d.Beta}
	return b.Prob(value)
}

func (d BetaDistribution) CentralTendency() float64 {
	b := distuv.Beta{Alpha: d.Alpha, Beta: d.Beta}
	return b.Mean()
}
