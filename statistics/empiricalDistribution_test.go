package statistics

import (
	"math"
	"testing"
)

func TestEmpiricalDistribution_GetSampleSize(t *testing.T) {
	binstarts := []float64{0, 1, 2, 3, 4}
	bincounts := []float64{2, 3, 4, 3, 2}
	e, _ := Init(binstarts, bincounts)
	var expected int64 = 14
	got := e.GetSampleSize()
	if expected != got {
		t.Errorf("e.GetSampleSize(): %v; expected: %v", got, expected)
	}
}
func TestEmpiricalDistribution_InvCDF(t *testing.T) {
	binstarts := []float64{0, 1, 2, 3, 4}
	bincounts := []float64{2, 3, 4, 3, 2}
	e, _ := Init(binstarts, bincounts)
	var expected float64 = 2
	got := e.InvCDF(0.4)
	diff := expected - got
	if math.Abs(diff) > 0.00001 {
		t.Errorf("InvCDF(0.4): %v; expected: %v", got, expected)
	}
}
func TestEmpiricalDistribution_CDF(t *testing.T) {
	binstarts := []float64{0, 1, 2, 3, 4}
	bincounts := []float64{2, 3, 4, 3, 2}
	e, _ := Init(binstarts, bincounts)
	expected := 0.642857
	got := e.CDF(3.5)
	diff := expected - got
	if math.Abs(diff) > 0.00001 {
		t.Errorf("CDF(3.5): %f; expected: %f", got, expected)
	}
}

func TestEmpiricalDistribution_PDF(t *testing.T) {
	binstarts := []float64{0, 1, 2, 3, 4}
	bincounts := []float64{2, 3, 4, 3, 2}
	e, _ := Init(binstarts, bincounts)
	var expected float64 = 0.21428571
	got := e.PDF(3.5)
	diff := expected - got
	if math.Abs(diff) > 0.0000001 {
		t.Errorf("PDF(3.5): %f; expected: %f", got, expected)
	}
}
