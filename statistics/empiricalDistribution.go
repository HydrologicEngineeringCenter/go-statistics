package statistics

import (
	"errors"
	"math"
)

type EmpiricalDistribution struct {
	binStarts []int64
	binWidth  int64
	binCounts []int64
	minValue  int64
	maxValue  int64
}

func Init(binstarts []int64, bincounts []int64) (*EmpiricalDistribution, error) {
	if binstarts == nil {
		return nil, errors.New("bin starts array cannot be empty")
	} else if bincounts == nil {
		return nil, errors.New("bin counts array cannot be empty")
	} else if len(binstarts) != len(bincounts) {
		return nil, errors.New("the arrays must be the same size")
	} else {
		for bincount := range bincounts {
			if bincount < 0 {
				return nil, errors.New("bin counts can not be negative")
			}
		}
		for i := 1; i < len(binstarts); i++ {
			if binstarts[i-1] >= binstarts[i] {
				return nil, errors.New("bin starts must be monotonically increasing")
			}
		}
		for i := 1; i < len(binstarts)-1; i++ {
			if binstarts[i]-binstarts[i-1] != binstarts[i+1]-binstarts[i] {
				return nil, errors.New("bin width must be constant")
			}
		}
		w := binstarts[1] - binstarts[0]
		b := binstarts
		c := bincounts
		min := binstarts[0]
		max := binstarts[len(binstarts)-1] + w
		e := EmpiricalDistribution{binStarts: b, binWidth: w, binCounts: c, minValue: min, maxValue: max}
		return &e, nil
	}
}

func (e EmpiricalDistribution) GetSampleSize() int64 {
	var sum int64
	sum = 0
	for i := 0; i < len(e.binCounts); i++ {
		sum += e.binCounts[i]
	}
	return sum
}

func (e EmpiricalDistribution) InvCDF(probability float64) float64 {
	if probability <= 0.0 {
		return float64(e.minValue)
	}
	if probability >= 1.0 {
		return float64(e.maxValue)
	}
	numobs := int64(float64(e.GetSampleSize()) * probability)
	if probability <= 0.5 {
		idx := 0
		obs := e.binCounts[idx] // bin counts
		cobs := obs
		for cobs < numobs {
			idx++
			obs = e.binCounts[idx]
			cobs += obs
		}
		binOffSet := float64(idx+1) - float64(cobs-numobs)/float64(obs)
		return float64(e.minValue) + float64(e.binWidth)*binOffSet
	} else {
		idx := len(e.binCounts)
		obs := e.binCounts[idx]
		cobs := e.GetSampleSize() - obs
		for cobs > numobs {
			idx--
			obs = e.binCounts[idx]
			cobs -= obs
		}
		binOffSet := float64(len(e.binCounts)-idx) + float64(numobs-cobs)/float64(obs)
		return float64(e.maxValue) - float64(e.binWidth)*binOffSet
	}
}

func (e EmpiricalDistribution) CDF(value float64) float64 {
	if value <= float64(e.minValue) {
		return 0.0
	}
	if value >= float64(e.maxValue) {
		return 1.0
	}
	dIdx := (value - float64(e.minValue)) / float64(e.binWidth)
	if dIdx <= 0 {
		return 0.0
	}
	if int(dIdx) >= len(e.binCounts) {
		return 1.0
	}
	val := float64(len(e.binCounts)) / 2
	if dIdx <= val {
		idx := int64(math.Floor(dIdx))
		var cobs int64 = 0
		var i int64 = 0
		for i < idx {
			cobs += e.binCounts[i]
			i++
		}
		cobs += (int64(dIdx) - idx) * e.binCounts[idx]
		return float64(cobs) / float64(e.GetSampleSize())
	} else {
		idx := int64(math.Floor(dIdx))
		var cobs int64 = e.GetSampleSize()
		var i int64 = int64(len(e.binCounts) - 1)
		for i > idx {
			cobs -= e.binCounts[i]
			i--
		}
		cobs -= (idx + 1 - int64(dIdx)) * e.binCounts[idx]
		return float64(cobs) / float64(e.GetSampleSize())

	}
}

func (e EmpiricalDistribution) PDF(value float64) float64 {
	// binwidth returns 0, even when hard-wired to be 1
	idx := (value - float64(e.minValue)) / float64(e.binWidth)
	if idx < 0 {
		return 0.0
	}
	if int(idx) > len(e.binCounts) {
		return 0.0
	}
	return float64(e.binCounts[int64(idx)]) / float64(e.binWidth*e.GetSampleSize())

}

func (e EmpiricalDistribution) CentralTendency() float64 {
	return e.InvCDF(0.5)
}
