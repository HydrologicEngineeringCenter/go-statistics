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
	//var wg sync.WaitGroup
	iterations := 1000
	//wg.Add(iterations)
	for i := 0; i < iterations; i++ {
		//defer wg.Done()
		//go ih.AddObservation(n.InvCDF(rand.Float64())) //i think the appending of slices in the add observation method on inlinehistogram causes the problems
		ih.AddObservation(n.InvCDF(rand.Float64()))
	}
	//wg.Wait()
	for _, val := range ih.GetBins() {
		fmt.Println(val)
	}
	fmt.Println(fmt.Sprintf("numbins %d", len(ih.GetBins())))
	fmt.Println(fmt.Sprintf("min %f", ih.pm.GetMin()))
	fmt.Println(fmt.Sprintf("max %f", ih.pm.GetMax()))
	fmt.Println(fmt.Sprintf("mean %f", ih.pm.GetMean()))
	fmt.Println(fmt.Sprintf("variance %f", ih.pm.GetSampleVariance()))
	fmt.Println(fmt.Sprintf("sample size %d", ih.pm.GetSampleSize()))
}
func TestAddObservationNormalDistConvergence(t *testing.T) {
	ih := Init(.01, -1.0, 1.0)
	n := statistics.NormalDistribution{Mean: 0, StandardDeviation: 1}
	var convergence bool = false
	var iterations int64 = 1000
	for convergence != true {
		var i int64 = 0
		for i < iterations {
			ih.AddObservation(n.InvCDF(rand.Float64()))
			i++
		}
		convergence, iterations = ih.TestForConvergence(.01, .99, .95, .001) //upper confidence limit test, lower confidence limit test, confidenece, error tolerance
		fmt.Println(fmt.Sprintf("Computed %d estimated to need %d more iterations", ih.pm.GetSampleSize(), iterations))
	}
	fmt.Println("****Converged******")
	fmt.Println(fmt.Sprintf("numbins %d", len(ih.GetBins())))
	fmt.Println(fmt.Sprintf("min %f", ih.pm.GetMin()))
	fmt.Println(fmt.Sprintf("max %f", ih.pm.GetMax()))
	fmt.Println(fmt.Sprintf("mean %f", ih.pm.GetMean()))
	fmt.Println(fmt.Sprintf("variance %f", ih.pm.GetSampleVariance()))
	fmt.Println(fmt.Sprintf("sample size %d", ih.pm.GetSampleSize()))
}
