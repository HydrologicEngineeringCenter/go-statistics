package data

import (
	"testing"
)

func TestAddObservationBinData(t *testing.T) {
	bins := make([]int64, 5)
	ih := InlineHistogram{Bins: bins, BinWidth: 1.0, MinValue: 0.0, MaxValue: 5.0}
	vals := []float64{.5, 1.5, 2.5, 3.5, 4.5}
	expected := []int64{1, 1, 1, 1, 1}
	for idx := range vals {
		ih.AddObservation(vals[idx])
	}
	for idx, val := range expected {
		if ih.Bins[idx] != val {
			t.Errorf("Bin(%d) = %d; expected %d", idx, ih.Bins[idx], val)
		}
	}
}
func TestAddObservationExceedUpper(t *testing.T) {
	bins := make([]int64, 5)
	ih := InlineHistogram{Bins: bins, BinWidth: 1.0, MinValue: 0.0, MaxValue: 5.0}
	vals := []float64{.5, 1.5, 2.5, 3.5, 4.5, 6.5}
	expected := []int64{1, 1, 1, 1, 1, 1}
	for idx := range vals {
		ih.AddObservation(vals[idx])
	}
	for idx, val := range expected {
		if ih.Bins[idx] != val {
			t.Errorf("Bin(%d) = %d; expected %d", idx, ih.Bins[idx], val)
		}
	}
}
func TestAddObservationBelowLower(t *testing.T) {
	bins := make([]int64, 5)
	ih := InlineHistogram{Bins: bins, BinWidth: 1.0, MinValue: 0.0, MaxValue: 5.0}
	vals := []float64{-.5, .5, 1.5, 2.5, 3.5, 4.5}
	expected := []int64{1, 1, 1, 1, 1, 1}
	for idx := range vals {
		ih.AddObservation(vals[idx])
	}
	for idx, val := range expected {
		if ih.Bins[idx] != val {
			t.Errorf("Bin(%d) = %d; expected %d", idx, ih.Bins[idx], val)
		}
	}
}
