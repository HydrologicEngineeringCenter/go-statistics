package statistics

import "math"

type TriangularDistribution struct {
	Min        float64 `json:"min"`
	MostLikely float64 `json:"mostlikely"`
	Max        float64 `json:"max"`
}

func (t TriangularDistribution) InvCDF(probability float64) float64 {
	a := t.MostLikely - t.Min
	b := t.Max - t.MostLikely
	if probability <= 0 {
		return t.Min
	} else if probability < a/(t.Max-t.Min) {
		return t.Min + math.Sqrt(probability*(t.Max-t.Min)*a)
	} else if probability < 1 {
		return t.Max - math.Sqrt((1-probability)*(t.Max-t.Min)*b)
	}
	return t.Max
}
func (t TriangularDistribution) CDF(value float64) float64 {
	if value < t.Min {
		return 0
	} else if value <= t.MostLikely {
		return math.Pow((value-t.Min), 2) / ((t.Max - t.Min) * (t.MostLikely - t.Min))
	} else if value <= t.Max {
		return 1 - math.Pow((t.Max-value), 2)/((t.Max-t.Min)*(t.Max-t.MostLikely))
	}
	return 1
}
func (t TriangularDistribution) PDF(value float64) float64 {
	if value < t.Min {
		return 0
	} else if value <= t.MostLikely {
		return 2 * (value - t.Min) / ((t.Max - t.Min) * (t.MostLikely - t.Min))
	} else if value <= t.Max {
		return 2 * (t.Max - value) / ((t.Max - t.Min) * (t.Max - t.MostLikely))
	} else {
		return 0
	}
}
func (t TriangularDistribution) CentralTendency() float64 {
	return t.MostLikely
}
