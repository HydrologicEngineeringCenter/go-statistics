package data

import "math"

type productMoments struct {
	min            float64
	max            float64
	sampleSize     int64
	mean           float64
	sampleVariance float64
}

func CreateProductMoments() *productMoments {
	pm := productMoments{}
	return &pm
}
func (pm *productMoments) GetMin() float64 {
	return pm.min
}
func (pm *productMoments) GetMean() float64 {
	return pm.mean
}
func (pm *productMoments) GetMax() float64 {
	return pm.max
}
func (pm *productMoments) GetSampleSize() int64 {
	return pm.sampleSize
}
func (pm *productMoments) GetSampleVariance() float64 {
	return pm.sampleVariance
}
func (pm *productMoments) AddObservation(value float64) {
	if pm.sampleSize == 0 {
		pm.max = value
		pm.min = value
		pm.mean = value
		pm.sampleVariance = 0
		pm.sampleSize = 1
	} else {
		if value > pm.max {
			pm.max = value
		} else if value < pm.min {
			pm.min = value
		}
		pm.sampleSize += 1
		pm.sampleVariance = (((float64(pm.sampleSize-2) / float64(pm.sampleSize-1)) * pm.sampleVariance) + (math.Pow((value-pm.mean), 2))/float64(pm.sampleSize))
		pm.mean = pm.mean + ((value - pm.mean) / float64(pm.sampleSize))
	}
}
func (pm *productMoments) AddObservations(values []float64) {
	for _, val := range values {
		pm.AddObservation(val)
	}
}
