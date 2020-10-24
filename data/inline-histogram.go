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
	numobs := int64(float64(ih.pm.GetSampleSize()) * probability)
	if probability <= 0.5 {
		idx := 0
		obs := ih.bins[idx]
		cobs := obs
		for cobs < numobs {
			idx++
			obs = ih.bins[idx]
			cobs += obs
		}
		return ih.minValue + ih.binWidth*(float64(int64(idx+1)-(int64(cobs)-numobs))/float64(obs))
	} else {
		idx := len(ih.bins) - 1
		obs := ih.bins[idx]
		cobs := ih.pm.GetSampleSize() - obs
		for cobs > numobs {
			idx--
			obs = ih.bins[idx]
			cobs -= obs
		}
		return ih.maxValue + ih.binWidth*(float64(int64(len(ih.bins)-1)-(int64(cobs)-numobs))/float64(obs))
	}
}
func (ih *inlineHistogram) CDF(value float64) float64 {
	if value <= ih.minValue {
		return 0.0
	}
	if value >= ih.maxValue {
		return 1.0
	}
	dIdx := (value - ih.minValue) / ih.binWidth
	if dIdx <= 0 {
		return 0.0
	}
	if int(dIdx) >= len(ih.bins) {
		return 1.0
	}
	val := float64(len(ih.bins)) / 2
	if dIdx <= val {
		idx := int64(math.Floor(dIdx))
		var cobs int64 = 0
		var i int64 = 0
		for i < idx {
			cobs += ih.bins[i]
			i++
		}
		cobs += (int64(dIdx) - idx) * ih.bins[idx]
		return float64(cobs) / float64(ih.pm.GetSampleSize())
	} else {
		idx := int64(math.Floor(dIdx))
		var cobs int64 = ih.pm.GetSampleSize()
		var i int64 = int64(len(ih.bins) - 1)
		for i > idx {
			cobs -= ih.bins[i]
			i--
		}
		cobs -= (idx + 1 - int64(dIdx)) * ih.bins[idx]
		return float64(cobs) / float64(ih.pm.GetSampleSize())
	}
}
func (ih *inlineHistogram) PDF(value float64) float64 {
	index := (value - ih.minValue) / (ih.binWidth)
	if index < 0 {
		return 0.0
	}
	if index > ih.maxValue {
		return 0.0
	}
	return float64(ih.bins[int(index)] / int64(ih.binWidth*float64(ih.pm.GetSampleSize())))
}
