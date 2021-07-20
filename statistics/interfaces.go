package statistics

import (
	"fmt"
	"reflect"
)

type ContinuousDistribution interface {
	InvCDF(probability float64) float64
	CDF(value float64) float64
	PDF(value float64) float64
	CentralTendency() float64
}
type FittableDistribution interface {
	Fit(inputData []float64) // could have an interface FittableDistribution or something like that
}

func Marshal(c ContinuousDistribution) {
	ref := reflect.Indirect(reflect.ValueOf(c))
	fmt.Print("{\"type\":\"")
	fmt.Print(reflect.TypeOf(c))
	fmt.Print("\",\"parameters\":[")
	for i := 0; i < ref.NumField(); i++ {
		fmt.Print("{\"")
		field := ref.Type().Field(i)
		name := field.Name
		fmt.Print(name)
		fmt.Print("\":")
		value := ref.Field(i)
		fmt.Print(fmt.Sprintf("%v", value))
		if i < ref.NumField()-1 {
			fmt.Print("},")
		} else {
			fmt.Print("}")
		}
	}
	fmt.Println("]}")

}
