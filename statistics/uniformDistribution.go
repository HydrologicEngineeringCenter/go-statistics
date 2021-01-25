package statistics

type UniformDistribution struct {
	Min float64
	Max float64
}

func (u *UniformDistribution) Fit(inputData []float64) {
	var min = inputData[0]
	var max = inputData[0]
	for _, val := range inputData {
		if val < min {
			min = val
		} else if val > max {
			max = val
		}
	}
	u.Min = min
	u.Max = max
}
func (u UniformDistribution) InvCDF(probability float64) float64 {
	return u.Min + ((u.Max - u.Min) * probability)
}
func (u UniformDistribution) CDF(value float64) float64 {
	if value < u.Min {
		return 0
	} else if value <= u.Max {
		return (value - u.Min) / (u.Max - u.Min)
	} else {
		return 1
	}
}
func (u UniformDistribution) PDF(value float64) float64 {
	if value < u.Min {
		return 0
	} else if value <= u.Max {
		return 1 / (u.Max - u.Min)
	} else {
		return 0
	}
}
func (u UniformDistribution) CentralTendency() float64 {
	return (u.Min + u.Max)/2
}
