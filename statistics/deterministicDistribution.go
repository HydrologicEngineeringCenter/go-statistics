package statistics

type DeterministicDistribution struct {
	value float64
}

func InitDeterministic(val float64) (*DeterministicDistribution, error) {
	d := DeterministicDistribution{value: val} // what errors to check for
	return &d, nil
}

func (d *DeterministicDistribution) InvCDF(probability float64) float64 {
	return d.value
}

func (d *DeterministicDistribution) PDF(val float64) float64 {
	if val == d.value {
		return 1.0
	} else {
		return 0.0
	}
}

func (d *DeterministicDistribution) CDF(val float64) float64 {
	if val >= d.value {
		return 1.0
	} else {
		return 0.0
	}
}

func (d *DeterministicDistribution) CentralTendency() float64 {
	return d.value
}
