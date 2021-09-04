package statistics

import (
	"math"
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
		z := zeroSkewDistribution(d)
		return z.InvCDF(probability)
	} else {
		if d.Skew > 0 {

			g := positiveSkewDistribution(d)
			return g.InvCDF(probability)
		} else {

			g := negativeSkewDistribution(d)
			return -g.InvCDF(1 - probability)
		}
	}
}
func (d PearsonIIIDistribution) CDF(value float64) float64 {
	noSkew := .00001
	if math.Abs(d.Skew) < noSkew {
		z := zeroSkewDistribution(d)
		return z.CDF(value)
	} else {
		if d.Skew > 0 {

			g := positiveSkewDistribution(d)
			return g.CDF(value)
		} else {

			g := negativeSkewDistribution(d)
			return 1 - g.CDF(-value)
		}
	}
}
func (d PearsonIIIDistribution) PDF(value float64) float64 {
	noSkew := .00001
	if math.Abs(d.Skew) < noSkew {
		z := zeroSkewDistribution(d)
		return z.PDF(value)
	} else {
		if d.Skew > 0 {

			g := positiveSkewDistribution(d)
			return g.PDF(value)
		} else {

			g := negativeSkewDistribution(d)
			return -g.PDF(value)
		}
	}
}
func (n PearsonIIIDistribution) CentralTendency() float64 {
	return n.Mean
}
func zeroSkewDistribution(d PearsonIIIDistribution) ContinuousDistribution {
	return NormalDistribution{Mean: d.Mean, StandardDeviation: d.StandardDeviation}
}
func negativeSkewDistribution(d PearsonIIIDistribution) ContinuousDistribution {
	shift := -d.Mean + 2.0*d.StandardDeviation/d.Skew
	alpha := 4.0 / (d.Skew * d.Skew)
	beta := .5 * d.StandardDeviation * d.Skew
	beta = -beta
	return ShiftedGammaDistribution{Alpha: alpha, Beta: beta, Shift: shift}
}
func positiveSkewDistribution(d PearsonIIIDistribution) ContinuousDistribution {
	alpha := 4.0 / (d.Skew * d.Skew)
	beta := .5 * d.StandardDeviation * d.Skew
	shift := d.Mean - 2.0*d.StandardDeviation/d.Skew
	return ShiftedGammaDistribution{Alpha: alpha, Beta: beta, Shift: shift}
}
