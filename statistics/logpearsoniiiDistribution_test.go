package statistics

import (
	"math"
	"testing"
)

func TestLogPearsonIIIInvCDF(t *testing.T) {
	n := LogPearsonIIIDistribution{Mean: 3.368, StandardDeviation: .246, Skew: .668}
	probs := []float64{.998, .995, .99, .98, .95, .9, .8, .5, .2, .1, .05, .01}
	expected := []float64{18878.87515053270180942491, 14246.58825980164874636102, 11408.83966308754315832630, 9043.72657283687294693664, 6511.95816420457322237780, 4961.12702987368902540766, 3656.87315507261564562214, 2191.79779904862152761780, 1435.93911608508096833248, 1189.92079576230275961279, 1035.43101823480742496031, 827.66401592971760692308}
	for idx := range probs {
		got := n.InvCDF(probs[idx])
		diff := expected[idx] - got
		if math.Abs(diff) > .0000000001 {
			t.Errorf("InvCDF(%f) = %f; expected %f", probs[idx], got, expected[idx])
		}
	}
}
