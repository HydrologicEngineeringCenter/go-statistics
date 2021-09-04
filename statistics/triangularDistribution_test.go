package statistics

import (
	"math"
	"testing"
)

func TestTriangularInvCDF(t *testing.T) {
	td := TriangularDistribution{Min: 0, MostLikely: 5.0, Max: 10}
	probs := []float64{0, .25, .5, .75, 1}
	expected := []float64{0, 3.535534, 5, 6.464466, 10.0}
	for idx := range probs {
		got := td.InvCDF(probs[idx])
		diff := expected[idx] - got
		if math.Abs(diff) > .0000001 {
			t.Errorf("InvCDF(%f) = %f; expected %f", probs[idx], got, expected[idx])
		}
	}
}
func TestTriangularCDF(t *testing.T) {
	td := TriangularDistribution{Min: 0, MostLikely: 5.0, Max: 10}
	expected := []float64{0, .25, .5, .75, 1}
	values := []float64{0, 3.535534, 5, 6.464466, 10.0}
	for idx := range values {
		got := td.CDF(values[idx])
		diff := expected[idx] - got
		if math.Abs(diff) > .0000001 {
			t.Errorf("CDF(%f) = %f; expected %f", values[idx], got, expected[idx])
		}
	}
}
func TestTriangularPDF(t *testing.T) {
	td := TriangularDistribution{Min: 0, MostLikely: 5.0, Max: 10}
	vals := []float64{-1, 0, 2.5, 5, 7.5, 10.0, 11}
	expected := []float64{0.0, 0.0, 0.1, 0.2, 0.1, 0.0, 0.0}
	for idx := range vals {
		got := td.PDF(vals[idx])
		if got != expected[idx] {
			t.Errorf("PDF(%f) = %f; expected %f", vals[idx], got, expected[idx])
		}
	}
}
