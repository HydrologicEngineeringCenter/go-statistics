package statistics

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
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

func Marshal(c ContinuousDistribution) (string, error) {
	ref := reflect.Indirect(reflect.ValueOf(c))
	s := "{\"type\":\""
	s += fmt.Sprintf("%v", reflect.TypeOf(c))
	s += "\",\"parameters\":["
	for i := 0; i < ref.NumField(); i++ {
		s += "{\""
		field := ref.Type().Field(i)
		name := field.Name
		s += name
		s += "\":"
		value := ref.Field(i)
		s += fmt.Sprintf("%v", value)
		if i < ref.NumField()-1 {
			s += "},"
		} else {
			s += "}"
		}
	}
	s += "]}"
	return s, nil
}
func Unmarshal(distribution string) (ContinuousDistribution, error) {
	distribution = strings.Replace(distribution, "{\"type\":\"", "", 1)
	structname := strings.Split(distribution, "\"")[0]
	bracketIdx := strings.Index(distribution, "[")
	parameterString := distribution[bracketIdx+1 : len(distribution)-2]
	params := strings.Split(parameterString, ",")
	var c ContinuousDistribution
	var err error
	switch structname {
	case "statistics.NormalDistribution":
		c, err = parseDistribution(&NormalDistribution{}, params)
	case "statistics.TriangularDistribution":
		c, err = parseDistribution(&TriangularDistribution{}, params)
	case "statistics.UniformDistribution":
		c, err = parseDistribution(&UniformDistribution{}, params)
	case "statistics.LogNormalDistribution":
		c, err = parseDistribution(&LogNormalDistribution{}, params)
	case "*statistics.DeterministicDistribution":
		c, err = parseDistribution(&DeterministicDistribution{}, params)
	case "*statistics.EmpiricalDistribution":
		c, err = parseDistribution(&EmpiricalDistribution{}, params)
	default:
		return NormalDistribution{}, errors.New(structname + " not found")
	}
	return c, err
}
func parseDistribution(c interface{}, params []string) (ContinuousDistribution, error) {
	for i, s := range params {
		parts := strings.Split(s, ":")
		fname := parts[0][2 : len(parts[0])-1]
		vstring := parts[1][:len(parts[1])-1]
		t := reflect.ValueOf(c)
		v := t.Elem().FieldByName(fname)
		switch v.Kind() {
		case reflect.Float64:
			val, _ := strconv.ParseFloat(vstring, 64)
			rf := reflect.ValueOf(c).Elem().Field(i)
			if rf.CanSet() {
				rf.SetFloat(val)
			} else {
				rf := reflect.NewAt(rf.Type(), unsafe.Pointer(rf.UnsafeAddr())).Elem()
				rf.SetFloat(val)
			}
		case reflect.Int64:
			val, _ := strconv.ParseInt(vstring, 10, 64)
			rf := reflect.ValueOf(c).Elem().Field(i)
			if rf.CanSet() {
				rf.SetInt(val)
			} else {
				rf := reflect.NewAt(rf.Type(), unsafe.Pointer(rf.UnsafeAddr())).Elem()
				rf.SetInt(val)
			}
		case reflect.Slice:
			fmt.Println("we gotta slice...")
		default:
			fmt.Println(fname)
			fmt.Println(vstring)
		}
	}
	d, ok := c.(ContinuousDistribution)
	if ok {
		return d, nil
	}
	return NormalDistribution{}, errors.New("stupid reflection and interfaces are dumb.")
}
