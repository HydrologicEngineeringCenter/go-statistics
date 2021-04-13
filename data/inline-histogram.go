package data

import (
	"fmt"
	"math"
	"strings"
)

type InlineHistogram struct {
	binWidth float64
	bins     []int64
	minValue float64
	maxValue float64
	pm       *productMoments
}

func Init(binwidth float64, minval float64, maxval float64) *InlineHistogram {
	val := math.Ceil((maxval - minval) / binwidth)
	b := make([]int64, int(val))
	maxval = minval + math.Floor(val*binwidth)
	moments := CreateProductMoments()
	ih := InlineHistogram{bins: b, binWidth: binwidth, minValue: minval, maxValue: maxval, pm: moments}
	return &ih
}

func (ih *InlineHistogram) AddObservation(value float64) {
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
		//add a bin. bins are bin start inclusive, bin end exclusive, therefore we need a new bin for this case.
		tmpBins := make([]int64, 1)
		ih.bins = append(ih.bins, tmpBins...)
		ih.maxValue = ih.maxValue + ih.binWidth
	}
	ih.pm.AddObservation(value) //not critically necessary
	//increment by 1
	ih.bins[int(index)] += 1 //truncation towards zero
}
func (ih *InlineHistogram) AddObservations(values []float64) {
	for _, val := range values {
		ih.AddObservation(val)
	}
}
func (ih *InlineHistogram) GetBins() []int64 {
	return ih.bins
}
func (ih *InlineHistogram) InvCDF(probability float64) float64 {
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
		binOffset := float64(int64(idx+1)) - float64(int64(cobs)-numobs)/float64(obs)
		return ih.minValue + ih.binWidth*(binOffset)
	} else {
		idx := len(ih.bins) - 1
		obs := ih.bins[idx]
		cobs := ih.pm.GetSampleSize() - obs
		for cobs > numobs {
			idx--
			obs = ih.bins[idx]
			cobs -= obs
		}
		binOffset := float64(int64(len(ih.bins)-idx)) + float64(numobs-int64(cobs))/float64(obs)
		return ih.maxValue - ih.binWidth*(binOffset)
	}
}
func (ih *InlineHistogram) CDF(value float64) float64 {
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
func (ih *InlineHistogram) PDF(value float64) float64 {
	index := (value - ih.minValue) / ih.binWidth
	if index < 0 {
		return 0.0
	}
	if int(index) > len(ih.bins) {
		return 0.0
	}
	return float64(ih.bins[int(index)]) / (ih.binWidth * float64(ih.pm.GetSampleSize()))
}
func (ih *InlineHistogram) testTailConvergence(tailValue float64, zAlpha float64, relativeError float64) (bool, int64) {
	qVal := ih.InvCDF(tailValue)
	qSlope := ih.PDF(qVal)
	pOneMinusp := tailValue * (1.0 - tailValue)
	variance := (pOneMinusp) / (float64(ih.pm.GetSampleSize()) * qSlope * qSlope)
	e := math.Abs(zAlpha * math.Sqrt(variance) / qVal)
	converged := (e <= relativeError*0.5)

	if converged {
		return true, 0
	} else {
		iterations := int64((pOneMinusp / (qSlope * qSlope)) * math.Pow((2*zAlpha)/(qVal*relativeError), 2))
		remainingIters := math.Abs(float64(iterations))
		if iterations == math.MinInt64 {
			//cheating. i know.
			return false, 10000
		}
		return false, int64(remainingIters) // calculate this number...
	}
}
func (ih *InlineHistogram) TestForConvergence(minConfidenceLimit float64, maxConfidenceLimit float64, zAlpha float64, relativeError float64) (bool, int64) {
	minConverged, minItersLeft := ih.testTailConvergence(minConfidenceLimit, zAlpha, relativeError)
	//can early exit here.
	if !minConverged {
		return minConverged, minItersLeft
	}
	maxConverged, maxItersLeft := ih.testTailConvergence(maxConfidenceLimit, zAlpha, relativeError)
	if !maxConverged {
		return maxConverged, maxItersLeft //min should be zero if we get to here.
	}
	//converged = true;
	//convergedIteration = _numObs;
	return true, 0
}
func (ih *InlineHistogram) String() string {
	s := fmt.Sprintf("InlineHistogram:\nBinCount: %v\nObservations: %v\nMin: %f\nMax: %f\nMean: %f\n", len(ih.bins), ih.pm.sampleSize, ih.pm.min, ih.pm.max, ih.pm.GetMean())
	s += "Bin Start, Count\n"
	for idx, val := range ih.bins {
		s += fmt.Sprintf("%f, %v\n", ih.minValue+(ih.binWidth*float64(idx)), val)
	}
	return s
}

func (ih *InlineHistogram) HistogramToEmpiricalString() string {
	s := fmt.Sprintf("InlineHistogram:\nBinCount: %v\nObservations: %v\nMin: %f\nMax: %f\nMean: %f\n", len(ih.bins), ih.pm.sampleSize, ih.pm.min, ih.pm.max, ih.pm.GetMean())
	s += "Bin Start"
	for idx, _ := range ih.bins {
		s += fmt.Sprintf("%f, ", ih.minValue+(ih.binWidth*float64(idx)))
	}
	s += "\nBin Count:"
	for _, val := range ih.bins {
		s += fmt.Sprintf(" %v", val)
	}
	return s
}

func (ih *InlineHistogram) StringSparse() string {
	s := fmt.Sprintf("InlineHistogram:\nBinCount: %v\nBinWidth: %v\nObservations: %v\nMin: %f\nMax: %f\nMean: %f\n", len(ih.bins), ih.binWidth, ih.pm.sampleSize, ih.pm.min, ih.pm.max, ih.pm.GetMean())
	s += "Bin Start, Count (bins with zero counts not reported!)\n"
	for idx, val := range ih.bins {
		if val != 0 {
			s += fmt.Sprintf("%f, %v\n", ih.minValue+(ih.binWidth*float64(idx)), val)
		}
	}
	return s
}
func (ih InlineHistogram) MarshalJSON() ([]byte, error) {

	s := fmt.Sprintf("{\"inlinehistogram\":{\"bincount\":%v,\"observations\":%v,\"min\":%f,\"max\":%f,\"mean\":%f,", len(ih.bins), ih.pm.sampleSize, ih.pm.min, ih.pm.max, ih.pm.GetMean())
	s += "\"histogram\":["
	for idx, val := range ih.bins {
		s += fmt.Sprintf("{\"binstart\":%f,\"count\":%v},", ih.minValue+(ih.binWidth*float64(idx)), val)
	}
	s = strings.TrimRight(s, ",")
	s += "]}}"
	return []byte(s), nil
}
