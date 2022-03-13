package statistics

import (
	"math"
	"math/rand"

	"github.com/HydrologicEngineeringCenter/go-statistics/data"
)

type LogPearsonIIIDistribution struct {
	Mean                    float64 `json:"mean"`
	StandardDeviation       float64 `json:"standarddeviation"`
	Skew                    float64 `json:"skew"`
	EquivalentYearsOfRecord int     `json:"equivalent_years_of_record"`
}

func (n LogPearsonIIIDistribution) Bootstrap(seed int64) ContinuousDistribution {
	rnd := rand.New(rand.NewSource(seed))
	samples := make([]float64, n.EquivalentYearsOfRecord)
	for i := 0; i < n.EquivalentYearsOfRecord; i++ {
		samples[i] = n.InvCDF(rnd.Float64())
	}
	ret := n.Fit(samples)
	return ret
}
func (n LogPearsonIIIDistribution) Fit(inputData []float64) ContinuousDistribution {
	pm := data.CreateProductMoments()
	//log data.
	logData := make([]float64, len(inputData))
	for idx, v := range inputData {
		logData[idx] = math.Log10(v)
	}
	pm.AddObservations(logData)
	ret := LogPearsonIIIDistribution{
		Mean:              pm.GetMean(),
		StandardDeviation: pm.GetSampleVariance(),
		Skew:              n.Skew, //TO DO: compute skew.
	}
	return ret
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
