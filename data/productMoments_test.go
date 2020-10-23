package data

import (
	"testing"
)

func TestAddObservationVariance(t *testing.T) {
	pm := CreateProductMoments()
	vals := []float64{.5, 1.5, 2.5, 3.5, 4.5}
	expected := 2.5
	for idx := range vals {
		pm.AddObservation(vals[idx])
	}
	if pm.GetSampleVariance() != expected {
		t.Errorf("GetSampleVariance() = %f; expected %f", pm.GetSampleVariance(), expected)
	}
}
func TestAddObservationMean(t *testing.T) {
	pm := CreateProductMoments()
	vals := []float64{.5, 1.5, 2.5, 3.5, 4.5}
	expected := 2.5
	for idx := range vals {
		pm.AddObservation(vals[idx])
	}
	if pm.GetMean() != expected {
		t.Errorf("GetMean() = %f; expected %f", pm.GetMean(), expected)
	}
}
func TestAddObservationMin(t *testing.T) {
	pm := CreateProductMoments()
	vals := []float64{.5, 1.5, 2.5, 3.5, 4.5}
	expected := .5
	for idx := range vals {
		pm.AddObservation(vals[idx])
	}
	if pm.GetMin() != expected {
		t.Errorf("GetMin() = %f; expected %f", pm.GetMin(), expected)
	}
}
func TestAddObservationMax(t *testing.T) {
	pm := CreateProductMoments()
	vals := []float64{.5, 1.5, 2.5, 3.5, 4.5}
	expected := 4.5
	for idx := range vals {
		pm.AddObservation(vals[idx])
	}
	if pm.GetMax() != expected {
		t.Errorf("GetMax() = %f; expected %f", pm.GetMax(), expected)
	}
}
func TestAddObservationSampleSize(t *testing.T) {
	pm := CreateProductMoments()
	vals := []float64{.5, 1.5, 2.5, 3.5, 4.5}
	expected := 5
	for idx := range vals {
		pm.AddObservation(vals[idx])
	}
	if pm.GetSampleSize() != int64(expected) {
		t.Errorf("GetSampleSize() = %v; expected %d", pm.GetSampleSize(), expected)
	}
}
