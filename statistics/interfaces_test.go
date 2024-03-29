package statistics

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestNormalMarshal(t *testing.T) {
	n := NormalDistribution{Mean: 0, StandardDeviation: 1}
	sn, _ := json.Marshal(n)
	fmt.Println(string(sn))
}
func TestNormalEncoding(t *testing.T) {
	n := NormalDistribution{Mean: 0, StandardDeviation: 1}
	sn, _ := Marshal(n)
	fmt.Printf("%v\n", sn)
	var cc ContinuousDistributionContainer
	err := json.Unmarshal([]byte(sn), &cc)
	if err != nil {
		fmt.Println(err)
	}
	s, _ := Marshal(cc.Value)
	fmt.Printf("%v\n", s)
	ln := LogNormalDistribution{Mean: 0, StandardDeviation: 1}
	sln, _ := Marshal(ln)
	fmt.Printf("%v\n", sln)
	var ccln ContinuousDistributionContainer
	err = json.Unmarshal([]byte(sln), &ccln)
	if err != nil {
		fmt.Println(err)
	}
	s, _ = Marshal(ccln.Value)
	fmt.Printf("%v\n", s)
	tri := TriangularDistribution{Min: 1, MostLikely: 2, Max: 3}
	stri, _ := Marshal(tri)
	fmt.Printf("%v\n", stri)
	var cctri ContinuousDistributionContainer
	err = json.Unmarshal([]byte(stri), &cctri)
	if err != nil {
		fmt.Println(err)
	}
	s, _ = Marshal(cctri.Value)
	fmt.Printf("%v\n", s)
	u := UniformDistribution{Min: 1, Max: 5}
	su, _ := Marshal(u)
	fmt.Printf("%v\n", su)
	var ccu ContinuousDistributionContainer
	err = json.Unmarshal([]byte(su), &ccu)
	if err != nil {
		fmt.Println(err)
	}
	s, _ = Marshal(ccu.Value)
	fmt.Printf("%v\n", s)
	d := DeterministicDistribution{Value: 2.3}
	sd, _ := Marshal(d)
	fmt.Printf("%v\n", sd)
	var ccd ContinuousDistributionContainer
	err = json.Unmarshal([]byte(sd), &ccd)
	if err != nil {
		fmt.Println(err)
	}
	s, _ = Marshal(ccd.Value)
	fmt.Printf("%v\n", s)

	binstarts := []float64{0, 1, 2, 3, 4}
	bincounts := []int64{2, 3, 4, 3, 2}
	e, _ := Init(binstarts, bincounts)
	es, _ := Marshal(e)
	fmt.Printf("%v\n", es)
	var cce ContinuousDistributionContainer
	err = json.Unmarshal([]byte(es), &cce)
	if err != nil {
		fmt.Println(err)
	}
	s, _ = Marshal(cce.Value)
	fmt.Printf("%v\n", s)

	lp3 := LogPearsonIIIDistribution{Mean: 2, StandardDeviation: 1, Skew: 3}
	slp3, _ := Marshal(lp3)
	fmt.Printf("%v\n", slp3)
	var cclp3 ContinuousDistributionContainer
	err = json.Unmarshal([]byte(slp3), &cclp3)
	if err != nil {
		fmt.Println(err)
	}
	s, _ = Marshal(cclp3.Value)
	fmt.Printf("%v\n", s)

	b := BetaDistribution{Alpha: .5, Beta: .5}
	sb, _ := Marshal(b)
	fmt.Printf("%v\n", sb)
	var ccb ContinuousDistributionContainer
	err = json.Unmarshal([]byte(sb), &ccb)
	if err != nil {
		fmt.Println(err)
	}
	s, _ = Marshal(ccb.Value)
	fmt.Printf("%v\n", s)

	g := ShiftedGammaDistribution{Shift: .2, Alpha: .3, Beta: .5}
	sg, _ := Marshal(g)
	fmt.Printf("%v\n", sg)
	var ccg ContinuousDistributionContainer
	err = json.Unmarshal([]byte(sg), &ccg)
	if err != nil {
		fmt.Println(err)
	}
	s, _ = Marshal(ccg.Value)
	fmt.Printf("%v\n", s)

}
