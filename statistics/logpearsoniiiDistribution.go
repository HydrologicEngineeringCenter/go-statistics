package statistics

import (
	"math"

	"github.com/HydrologicEngineeringCenter/go-statistics/data"
)

type LogPearsonIIIDistribution struct {
	Mean              float64 `json:"mean"`
	StandardDeviation float64 `json:"standarddeviation"`
	Skew              float64 `json:"skew"`
}

func (n *LogPearsonIIIDistribution) Fit(inputData []float64) {
	pm := data.CreateProductMoments()
	//TO DO: log data.
	pm.AddObservations(inputData)
	n.Mean = pm.GetMean()
	n.StandardDeviation = pm.GetSampleVariance() //check to see if this is right
	n.Skew = .4                                  //why not. TO DO: compute skew.

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
	z := PearsonIIIDistribution{Mean: n.Mean, StandardDeviation: n.StandardDeviation, Skew: n.Skew}
	return z.CDF(math.Log10(value))
}
func (n LogPearsonIIIDistribution) PDF(value float64) float64 {
	z := PearsonIIIDistribution{Mean: n.Mean, StandardDeviation: n.StandardDeviation, Skew: n.Skew}
	return z.PDF(math.Log10(value)) / value / math.Log(10)
}
func (n LogPearsonIIIDistribution) CentralTendency() float64 {
	return n.Mean
}
