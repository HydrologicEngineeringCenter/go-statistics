package statistics

import (
	"fmt"
	"testing"
)

func TestNormalEncoding(t *testing.T) {
	var c ContinuousDistribution
	n := NormalDistribution{Mean: 0, StandardDeviation: 1}
	sn, _ := Marshal(n)
	fmt.Printf("%v\n", sn)
	c, _ = Unmarshal(sn)
	fmt.Printf("%v\n", DistributionToString(c))
	ln := LogNormalDistribution{Mean: 0, StandardDeviation: 1}
	sln, _ := Marshal(ln)
	fmt.Printf("%v\n", sln)
	c, _ = Unmarshal(sln)
	fmt.Printf("%v\n", DistributionToString(c))
	tri := TriangularDistribution{Min: 1, MostLikely: 2, Max: 3}
	stri, _ := Marshal(tri)
	fmt.Printf("%v\n", stri)
	c, _ = Unmarshal(stri)
	fmt.Printf("%v\n", DistributionToString(c))
	u := UniformDistribution{Min: 1, Max: 5}
	su, _ := Marshal(u)
	fmt.Printf("%v\n", su)
	c, _ = Unmarshal(su)
	fmt.Printf("%v\n", DistributionToString(c))
	d, _ := InitDeterministic(2.3)
	sd, _ := Marshal(d)
	fmt.Printf("%v\n", sd)
	c, _ = Unmarshal(sd)
	fmt.Printf("%v\n", DistributionToString(c))
	/*
		binstarts := []float64{0, 1, 2, 3, 4}
		bincounts := []int64{2, 3, 4, 3, 2}
		e, _ := Init(binstarts, bincounts)
		es, _ := Marshal(e)
		c, _ = Unmarshal(es)
		fmt.Printf("%v", c)
	*/
}
