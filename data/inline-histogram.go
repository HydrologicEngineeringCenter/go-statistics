package data

import (
	"math"
)

type inlineHistogram struct {
	binWidth float64
	bins     []int64
	minValue float64
	maxValue float64
}

func Init(binwidth float64, minval float64, maxval float64) *inlineHistogram {
	val := math.Ceil((maxval - minval) / binwidth)
	b := make([]int64, int(val))
	maxval = minval + math.Floor(val*binwidth)
	ih := inlineHistogram{bins: b, binWidth: binwidth, minValue: minval, maxValue: maxval}
	return &ih
}
func (ih *inlineHistogram) AddObservation(value float64) {
	if ih.maxValue > value && value > ih.minValue {
		//index it to the right location
		index := float64(len(ih.bins)) * (value - ih.minValue) / (ih.maxValue - ih.minValue)
		//increment by 1
		ih.bins[int(index)] += 1 //truncation towards zero
	} else {
		//bins need to change.
		if ih.maxValue < value {
			//add to end of array
			additionalBins := math.Ceil((value - ih.maxValue) / ih.binWidth) //should take ceiling
			newCount := len(ih.bins) + int(additionalBins)
			tmpBins := make([]int64, int(additionalBins))
			ih.bins = append(ih.bins, tmpBins...)
			ih.maxValue = value      //should go to the end of the next whole bin
			ih.bins[newCount-1] += 1 //i think this is right.
		} else {
			//add to beginning of array
			additionalBins := math.Ceil((ih.minValue - value) / ih.binWidth) //should take ceiling
			tmpBins := make([]int64, int(additionalBins))
			ih.bins = append(tmpBins, ih.bins...)
			ih.minValue = value //should go to the start of the first whole bin
			ih.bins[0] += 1     //i think this is right.
		}
	}
}
func (ih *inlineHistogram) GetBins() []int64 {
	return ih.bins
}
