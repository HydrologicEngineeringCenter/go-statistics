package statistics

import (
	"math"
	"testing"
)

func TestBetaDistribution1_InvCDF(t *testing.T) {
	bd := BetaDistribution{Alpha: 0.5, Beta: 0.5}
	probs := []float64{.25, .5, .75, 1}
	expected := []float64{0.1464466, 0.5000000, 0.8535534, 1.0000000} //R Stats
	for idx := range probs {
		got := bd.InvCDF(probs[idx])
		diff := expected[idx] - got
		if math.Abs(diff) > .0000001 {
			t.Errorf("InvCDF(%f) = %f; expected %f", probs[idx], got, expected[idx])
		}
	}
}

func TestBetaDistribution2_InvCDF(t *testing.T) {
	bd := BetaDistribution{Alpha: 1.5, Beta: 0.5}
	probs := []float64{.25, .5, .75, 1}
	expected := []float64{0.5971501, 0.8368060, 0.9609369, 1.0000000} //R Stats
	for idx := range probs {
		got := bd.InvCDF(probs[idx])
		diff := expected[idx] - got
		if math.Abs(diff) > .0000001 {
			t.Errorf("InvCDF(%f) = %f; expected %f", probs[idx], got, expected[idx])
		}
	}
}

func TestBetaDistribution3_InvCDF(t *testing.T) {
	bd := BetaDistribution{Alpha: 0.5, Beta: 1.5}
	probs := []float64{.25, .5, .75, 1}
	expected := []float64{0.03906313, 0.16319399, 0.40284992, 1.00000000} //R Stats
	for idx := range probs {
		got := bd.InvCDF(probs[idx])
		diff := expected[idx] - got
		if math.Abs(diff) > .0000001 {
			t.Errorf("InvCDF(%f) = %f; expected %f", probs[idx], got, expected[idx])
		}
	}
}

func TestBetaDistribution_CDF(t *testing.T) {
	bd := BetaDistribution{Alpha: 0.5, Beta: 1.5}
	probs := []float64{0, .25, .5, .75, 1}
	expected := []float64{0.0000000, 0.6089978, 0.8183099, 0.9423311, 1.0000000} //R Stats
	for idx := range probs {
		got := bd.CDF(probs[idx])
		diff := expected[idx] - got
		if math.Abs(diff) > .0000001 {
			t.Errorf("CDF(%f) = %f; expected %f", probs[idx], got, expected[idx])
		}
	}
}

func TestBetaDistribution_PDF(t *testing.T) {
	bd := BetaDistribution{Alpha: 0.5, Beta: 1.5}
	probs := []float64{0.00000001, .25, .5, .75, 1}
	expected := []float64{6366.1976918, 1.1026578, 0.6366198, 0.3675526, 0.0000000} //R Stats
	for idx := range probs {
		got := bd.PDF(probs[idx])
		diff := expected[idx] - got
		if math.Abs(diff) > .0000001 {
			t.Errorf("InvCDF(%f) = %f; expected %f", probs[idx], got, expected[idx])
		}
	}
}
