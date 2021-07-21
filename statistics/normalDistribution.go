package statistics

import "math"

type NormalDistribution struct {
	Mean              float64 `json:"mean"`
	StandardDeviation float64 `json:"standarddeviation`
}

func (n NormalDistribution) InvCDF(probability float64) float64 {
	if probability == .5 {
		return n.Mean
	}
	if probability <= 0 {
		probability = .000000000000001
	}
	if probability >= 1 {
		probability = .999999999999999
	}
	i := -1
	if probability > .5 {
		probability = 1 - probability
		i = 1
	}
	t := math.Sqrt(math.Log(1 / (probability * probability)))
	c0 := 2.515517
	c1 := 0.802853
	c2 := 0.010328
	d1 := 1.432788
	d2 := 0.189269
	d3 := 0.001308
	tsqrd := t * t
	x := t - (c0+c1*t+c2*(tsqrd))/(1+d1*t+d2*(tsqrd)+d3*(t*tsqrd))
	x = float64(i) * x
	return n.Mean + x*n.StandardDeviation
}
func (n NormalDistribution) CDF(value float64) float64 {
	return math.Erfc(-(value-n.Mean)/(n.StandardDeviation*math.Sqrt2)) / 2
}
func (n NormalDistribution) PDF(value float64) float64 {
	z := value - n.Mean
	return math.Exp(-z*z/(2*n.StandardDeviation*n.StandardDeviation)) * ((1 / math.Sqrt(2*math.Pi)) / n.StandardDeviation)
}
func (n NormalDistribution) CentralTendency() float64 {
	return n.Mean
}
