package data

import (
	"math"
)

type inlineHistogram struct {
	binWidth float64
	bins     []int64
	minValue float64
	maxValue float64
	pm       *productMoments
}

func Init(binwidth float64, minval float64, maxval float64) *inlineHistogram {
	val := math.Ceil((maxval - minval) / binwidth)
	b := make([]int64, int(val))
	maxval = minval + math.Floor(val*binwidth)
	moments := CreateProductMoments()
	ih := inlineHistogram{bins: b, binWidth: binwidth, minValue: minval, maxValue: maxval, pm: moments}
	return &ih
}
func (ih *inlineHistogram) AddObservation(value float64) {
	//check if value is a valid number?
	if ih.maxValue < value {
		//add bins to end of array
		additionalBins := math.Ceil((value - ih.maxValue) / ih.binWidth)
		tmpBins := make([]int64, int(additionalBins))
		ih.bins = append(ih.bins, tmpBins...)
		ih.maxValue = ih.maxValue + additionalBins*ih.binWidth
	} else if ih.minValue > value {
		//add bins to beginning of array
		additionalBins := math.Ceil((ih.minValue - value) / ih.binWidth)
		tmpBins := make([]int64, int(additionalBins))
		ih.bins = append(tmpBins, ih.bins...)
		ih.minValue = ih.minValue - additionalBins*ih.binWidth
	}
	//index it to the right location
	index := (value - ih.minValue) / (ih.binWidth)
	if value == ih.maxValue {
		index -= 1 //edge case
	}
	ih.pm.AddObservation(value) //not critically necessary
	//increment by 1
	ih.bins[int(index)] += 1 //truncation towards zero
}
func (ih *inlineHistogram) AddObservations(values []float64) {
	for _, val := range values {
		ih.AddObservation(val)
	}
}
func (ih *inlineHistogram) GetBins() []int64 {
	return ih.bins
}
func (ih *inlineHistogram) InvCDF(probability float64) float64 {
	if probability <= 0.0 {
		return ih.minValue
	}
	if probability >= 1.0 {
		return ih.maxValue
	}

	return 0.0
}
func (ih *inlineHistogram) CDF(value float64) float64 {
	return 0.0
}
func (ih *inlineHistogram) PDF(value float64) float64 {
	return 0.0
}
