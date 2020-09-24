package data

type ProductMoments struct {
	Min            float64
	Max            float64
	SampleSize     int64
	Mean           float64
	SampleVariance float64
}

/*
func (pm *ProductMoments) AddObservation(value float64) {
	if value != nil {
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
			pm.SampleVariance = ((((pm.SampleSize - 2) / (pm.SampleSize - 1)) * pm.SampleVariance) + ((value-pm.Mean)^2)/pm.SampleSize)
			pm.Mean = pm.Mean + ((value - pm.Mean) / pm.SampleSize)
		}
	}
}
*/
