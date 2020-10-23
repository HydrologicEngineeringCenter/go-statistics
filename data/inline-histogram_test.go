package data

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/HenryGeorgist/go-statistics/statistics"
)

func TestAddObservationBinData(t *testing.T) {
	ih := Init(1.0, 0.0, 5.0)
	vals := []float64{.5, 1.5, 2.5, 3.5, 4.5}
	expected := []int64{1, 1, 1, 1, 1}
	for idx := range vals {
		ih.AddObservation(vals[idx])
	}
	for idx, val := range expected {
		if ih.GetBins()[idx] != val {
			t.Errorf("Bin(%d) = %d; expected %d", idx, ih.GetBins()[idx], val)
		}
	}
}
func TestAddObservationExceedUpper(t *testing.T) {
	ih := Init(1.0, 0.0, 5.0)
	vals := []float64{.5, 1.5, 2.5, 3.5, 4.5, 5.5}
	expected := []int64{1, 1, 1, 1, 1, 1}
	for idx := range vals {
		ih.AddObservation(vals[idx])
	}
	for idx, val := range expected {
		if ih.GetBins()[idx] != val {
			t.Errorf("Bin(%d) = %d; expected %d", idx, ih.GetBins()[idx], val)
		}
	}
}
func TestAddObservationBelowLower(t *testing.T) {
	ih := Init(1.0, 0.0, 5.0)
	vals := []float64{-.5, .5, 1.5, 2.5, 3.5, 4.5}
	expected := []int64{1, 1, 1, 1, 1, 1}
	for idx := range vals {
		ih.AddObservation(vals[idx])
	}
	for idx, val := range expected {
		if ih.GetBins()[idx] != val {
			t.Errorf("Bin(%d) = %d; expected %d", idx, ih.GetBins()[idx], val)
		}
	}
}
func TestAddObservationNormalDist(t *testing.T) {
	ih := Init(.01, -1.0, 1.0)
	n := statistics.NormalDistribution{Mean: 0, StandardDeviation: 1}
	for i := 0; i < 100000; i++ {
		ih.AddObservation(n.InvCDF(rand.Float64()))
	}
	for _, val := range ih.GetBins() {
		fmt.Println(val)
	}
}
