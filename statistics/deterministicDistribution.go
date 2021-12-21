package statistics

type DeterministicDistribution struct {
	Value float64 `json:"value"`
}

func InitDeterministic(val float64) (*DeterministicDistribution, error) {
	d := DeterministicDistribution{Value: val} // what errors to check for
	return &d, nil
}

func (d DeterministicDistribution) InvCDF(probability float64) float64 {
	return d.Value
}
func (d DeterministicDistribution) PDF(val float64) float64 {
	if val == d.Value {
		return 1.0
	} else {
		return 0.0
	}
}
func (d DeterministicDistribution) CDF(val float64) float64 {
	if val >= d.Value {
		return 1.0
	} else {
		return 0.0
	}
}
func (d DeterministicDistribution) CentralTendency() float64 {
	return d.Value
}
