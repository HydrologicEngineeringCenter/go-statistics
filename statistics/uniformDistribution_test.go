package statistics

import (
	"testing"
)

func TestUniformFit(t *testing.T) {
	values := []float64{5.32, 4.54, 76.434, 32.34, -5.34}
	min := -5.34
	max := 76.434
	u := UniformDistribution{}
	u.Fit(values)
	if u.Min != min {
		t.Errorf("Uniform Distribution Fit test yeilded min of %f; expected %f", u.Min, min)
	}
	if u.Max != max {
		t.Errorf("Uniform Distribution Fit test yeilded max of %f; expected %f", u.Max, max)
	}
}
func TestUniformInvCDF(t *testing.T) {
	u := UniformDistribution{Min: 0, Max: 10}
	probs := []float64{0, .25, .5, .75, 1}
	expected := []float64{0, 2.5, 5, 7.5, 10.0}
	for idx := range probs {
		got := u.InvCDF(probs[idx])
		if got != expected[idx] {
			t.Errorf("InvCDF(%f) = %f; expected %f", probs[idx], got, expected[idx])
		}
	}
}
func TestUniformCDF(t *testing.T) {
	u := UniformDistribution{Min: 0, Max: 10}
	expected := []float64{0, .25, .5, .75, 1}
	values := []float64{0, 2.5, 5, 7.5, 10.0}
	for idx := range values {
		got := u.CDF(values[idx])
		if got != expected[idx] {
			t.Errorf("CDF(%f) = %f; expected %f", values[idx], got, expected[idx])
		}
	}
}
func TestUniformPDF(t *testing.T) {
	u := UniformDistribution{Min: 0, Max: 10}
	vals := []float64{-1, 0, 2.5, 5, 7.5, 10.0, 11}
	expected := []float64{0.0, 0.1, 0.1, 0.1, 0.1, 0.1, 0.0}
	for idx := range vals {
		got := u.PDF(vals[idx])
		if got != expected[idx] {
			t.Errorf("PDF(%f) = %f; expected %f", vals[idx], got, expected[idx])
		}
	}
}
