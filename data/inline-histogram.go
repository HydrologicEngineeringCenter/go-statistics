package data

import "math"

type InlineHistogram struct {
	BinWidth float64
	Bins     []int64
	MinValue float64
	MaxValue float64
}

func (ih *InlineHistogram) AddObservation(value float64) {
	if ih.MaxValue > value && value > ih.MinValue {
		//index it to the right location
		index := float64(len(ih.Bins)) * (value - ih.MinValue) / (ih.MaxValue - ih.MinValue)
		//increment by 1
		ih.Bins[int(index)] += 1 //truncation towards zero
	} else {
		//bins need to change.
		if ih.MaxValue < value {
			//add to end of array
			additionalBins := math.Ceil((ih.MaxValue - value) / ih.BinWidth) //should take ceiling
			newCount := len(ih.Bins) + int(additionalBins)
			tmpBins := make([]int64, int(additionalBins))
			ih.Bins = append(ih.Bins, tmpBins...)
			ih.MaxValue = value      //should go to the end of the next whole bin
			ih.Bins[newCount-1] += 1 //i think this is right.
		} else {
			//add to beginning of array
			additionalBins := math.Ceil((value - ih.MinValue) / ih.BinWidth) //should take ceiling
			tmpBins := make([]int64, int(additionalBins))
			ih.Bins = append(tmpBins, ih.Bins...)
			ih.MinValue = value //should go to the start of the first whole bin
			ih.Bins[0] += 1     //i think this is right.
		}
	}
}
