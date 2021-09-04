package statistics

import (
	"math"
	"testing"
)

func TestNormalInvCDF(t *testing.T) {
	n := NormalDistribution{Mean: 0, StandardDeviation: 1}
	probs := []float64{0, .25, .5, .75, 1}
	expected := []float64{-7.940906, -0.674189, 0, 0.674189, 7.941005}
	for idx := range probs {
		got := n.InvCDF(probs[idx])
		diff := expected[idx] - got
		if math.Abs(diff) > .000001 {
			t.Errorf("InvCDF(%f) = %f; expected %f", probs[idx], got, expected[idx])
		}
	}
}
func TestNormalCDF(t *testing.T) {
	n := NormalDistribution{Mean: 0, StandardDeviation: 1}
	expected := []float64{0, .25, .5, .75, 1}
	values := []float64{-7.940906, -0.674189, 0, 0.674189, 7.941005}
	for idx := range values {
		got := n.CDF(values[idx])
		diff := expected[idx] - got
		if math.Abs(diff) > .0001 {
			t.Errorf("CDF(%f) = %f; expected %f", values[idx], got, expected[idx])
		}
	}
}
func TestNormalPDF(t *testing.T) {
	n := NormalDistribution{Mean: 0, StandardDeviation: 1}
	vals := []float64{-7.940906, -0.674189, 0, 0.674189, 7.941005}
	expected := []float64{0.0, 0.317841, 0.398942, 0.317841, 0.000000}
	for idx := range vals {
		got := n.PDF(vals[idx])
		diff := expected[idx] - got
		if math.Abs(diff) > .000001 {
			t.Errorf("PDF(%f) = %f; expected %f", vals[idx], got, expected[idx])
		}
	}
}
