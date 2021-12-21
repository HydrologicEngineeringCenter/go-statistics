package statistics

import (
	"errors"
	"fmt"
	"math"
)

type EmpiricalDistribution struct {
	BinStarts []float64 `json:"binstarts"`
	BinWidth  float64   `json:"binwidth"`
	BinCounts []int64   `json:"bincounts"`
	MinValue  float64   `json:"minvalue"`
	MaxValue  float64   `json:"maxvalue"`
}

func Init(binstarts []float64, bincounts []int64) (EmpiricalDistribution, error) {
	if len(binstarts) < 2 {
		return EmpiricalDistribution{}, errors.New("there must be more than 1 bin hence more than 1 bin start")
	} else if len(bincounts) < 2 {
		return EmpiricalDistribution{}, errors.New("there must be more than 1 bin hence more than 1 bin count")
	} else if len(binstarts) != len(bincounts) {
		return EmpiricalDistribution{}, errors.New("the arrays must be the same size")
	} else {
		for bincount := range bincounts {
			if bincount < 0 {
				return EmpiricalDistribution{}, errors.New("bin counts can not be negative")
			}
		}
		for i := 1; i < len(binstarts); i++ {
			if binstarts[i-1] >= binstarts[i] {
				return EmpiricalDistribution{}, errors.New("bin starts must be monotonically increasing")
			}
		}
		for i := 1; i < len(binstarts)-1; i++ {
			if binstarts[i]-binstarts[i-1] != binstarts[i+1]-binstarts[i] {
				return EmpiricalDistribution{}, errors.New("bin width must be constant")
			}
		}
		w := binstarts[1] - binstarts[0]
		b := binstarts
		c := bincounts
		min := binstarts[0]
		max := binstarts[len(binstarts)-1] + float64(w)
		e := EmpiricalDistribution{BinStarts: b, BinWidth: w, BinCounts: c, MinValue: min, MaxValue: max}
		return e, nil
	}
}

func (e EmpiricalDistribution) GetSampleSize() int64 {
	var sum int64
	sum = 0
	for i := 0; i < len(e.BinCounts); i++ {
		sum += int64(e.BinCounts[i])
	}
	return sum
}

func (e EmpiricalDistribution) InvCDF(probability float64) float64 {
	if probability <= 0.0 {
		return float64(e.MinValue)
	}
	if probability >= 1.0 {
		return float64(e.MaxValue)
	}
	numobs := int64(float64(e.GetSampleSize()) * probability)
	if probability <= 0.5 {
		idx := 0
		obs := e.BinCounts[idx] // bin counts
		cobs := obs
		for cobs < numobs {
			idx++
			obs = e.BinCounts[idx]
			cobs += obs
		}
		binOffSet := float64(idx+1) - float64(cobs-numobs)/float64(obs)
		return float64(e.MinValue) + float64(e.BinWidth)*binOffSet
	} else {
		idx := len(e.BinCounts) - 1
		obs := e.BinCounts[idx]
		cobs := e.GetSampleSize() - obs
		for cobs > numobs {
			idx--
			obs = e.BinCounts[idx]
			cobs -= obs
		}
		fraction := float64(numobs-int64(cobs)) / float64(obs)
		binOffset := float64(int64((len(e.BinCounts)) - idx))
		return e.MaxValue - e.BinWidth*(binOffset) + e.BinWidth*(fraction)
	}
}

func (e EmpiricalDistribution) CDF(value float64) float64 {
	if value <= e.MinValue {
		return 0.0
	}
	if value >= e.MaxValue {
		return 1.0
	}
	dIdx := (value - e.MinValue) / float64(e.BinWidth)
	if dIdx <= 0 {
		return 0.0
	}
	if int(dIdx) >= len(e.BinCounts) {
		return 1.0
	}
	val := float64(len(e.BinCounts)) / 2
	if dIdx <= val {
		idx := int64(math.Floor(dIdx))
		var cobs int64 = 0
		var i int64 = 0
		for i < idx {
			cobs += e.BinCounts[i]
			i++
		}
		cobs += (int64(dIdx) - idx) * e.BinCounts[idx]
		return float64(cobs) / float64(e.GetSampleSize())
	} else {
		idx := int64(math.Floor(dIdx))
		var cobs int64 = e.GetSampleSize()
		var i int64 = int64(len(e.BinCounts) - 1)
		for i > idx {
			cobs -= e.BinCounts[i]
			i--
		}
		cobs -= (idx + 1 - int64(dIdx)) * e.BinCounts[idx]
		return float64(cobs) / float64(e.GetSampleSize())

	}
}

func (e EmpiricalDistribution) PDF(value float64) float64 {
	idx := (value - e.MinValue) / float64(e.BinWidth)
	if idx < 0 {
		return 0.0
	}
	if int(idx) > len(e.BinCounts) {
		return 0.0
	}
	return float64(e.BinCounts[int64(idx)]) / float64(e.BinWidth*float64(e.GetSampleSize()))

}

func (e EmpiricalDistribution) CentralTendency() float64 {
	return e.InvCDF(0.5)
}

func (e EmpiricalDistribution) String() string {
	s := fmt.Sprintf("Empirical Distribution:\nBinCount: %v\nObservations: %v\nMin: %v\nMax: %v\nMean: %f\n", len(e.BinStarts), e.GetSampleSize(), e.MinValue, e.MaxValue, e.CentralTendency())
	s += "Bin Start, Count\n"
	for idx, val := range e.BinCounts {
		s += fmt.Sprintf("%v, %v\n", e.MinValue+float64((e.BinWidth*float64(idx))), val)
	}
	return s
}
