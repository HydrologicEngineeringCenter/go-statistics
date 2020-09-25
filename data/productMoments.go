package data

import "math"

type ProductMoments struct {
	Min            float64
	Max            float64
	SampleSize     int64
	Mean           float64
	SampleVariance float64
}

func (pm *ProductMoments) AddObservation(value float64) {
	if pm.SampleSize == 0 {
		pm.Max = value
		pm.Min = value
		pm.Mean = value
		pm.SampleVariance = 0
		pm.SampleSize = 1
	} else {
		if value > pm.Max {
			pm.Max = value
		} else if value < pm.Min {
			pm.Min = value
		}
		pm.SampleSize += 1
		pm.SampleVariance = ((float64((pm.SampleSize-2)/(pm.SampleSize-1)) * pm.SampleVariance) + (math.Pow((value-pm.Mean), 2))/float64(pm.SampleSize))
		pm.Mean = pm.Mean + ((value - pm.Mean) / float64(pm.SampleSize))
	}
}
func (pm *ProductMoments) AddObservations(values []float64) {
	for _, val := range values {
		pm.AddObservation(val)
	}
}
