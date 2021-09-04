package statistics

import (
	"math"
	"testing"
)

func TestLogPearsonIIIInvCDF_PositiveSkew(t *testing.T) {
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
func TestLogPearsonIIIInvCDF_NegativeSkew(t *testing.T) {
	n := LogPearsonIIIDistribution{Mean: 2.966, StandardDeviation: .668, Skew: -.473}
	probs := []float64{.998, .995, .99, .98, .95, .9, .8, .5, .2, .1, .05, .01}
	expected := []float64{32506.82384690305480035022, 24602.16720938941944041289, 19284.35566198145897942595, 14566.43229235407852684148, 9288.03872218876131228171, 6041.78073029362531087827, 3451.04415806516362863476, 1043.47618679165134381037, 265.80156381679307742161, 121.13255664011704482164, 60.95393539477918665170, 15.29554262000421083201}
	for idx := range probs {
		got := n.InvCDF(probs[idx])
		diff := expected[idx] - got
		if math.Abs(diff) > .000000001 {
			t.Errorf("InvCDF(%f) = %f; expected %f", probs[idx], got, expected[idx])
		}
	}
}
