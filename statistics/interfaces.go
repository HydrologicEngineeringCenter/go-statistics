package statistics

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type ContinuousDistribution interface {
	InvCDF(probability float64) float64
	CDF(value float64) float64
	PDF(value float64) float64
	CentralTendency() float64
}
type FittableDistribution interface {
	Fit(inputData []float64) ContinuousDistribution // could have an interface FittableDistribution or something like that
}
type BootstrappableDistribution interface {
	FittableDistribution
	ContinuousDistribution
	Bootstrap(seed int64) ContinuousDistribution
}
type ContinuousDistributionContainer struct {
	Type  string                 `json:"type"`
	Value ContinuousDistribution `json:"parameters"`
}

func Marshal(c ContinuousDistribution) (string, error) {
	s := "{\"type\":\""
	stype := reflect.TypeOf(c)
	typestringparts := strings.Split(stype.String(), ".")
	s += fmt.Sprintf("%v", typestringparts[len(typestringparts)-1])
	s += "\",\"parameters\":"
	dist, err := json.Marshal(c)
	if err != nil {
		return s, err
	}
	s += string(dist)
	s += "}"
	return s, nil
}
func (c *ContinuousDistributionContainer) UnmarshalJSON(data []byte) error {
	value, err := UnmarshalCustomValue(data, "type", "parameters",
		map[string]reflect.Type{
			"NormalDistribution":        reflect.TypeOf(NormalDistribution{}),
			"LogNormalDistribution":     reflect.TypeOf(LogNormalDistribution{}),
			"TriangularDistribution":    reflect.TypeOf(TriangularDistribution{}),
			"UniformDistribution":       reflect.TypeOf(UniformDistribution{}),
			"DeterministicDistribution": reflect.TypeOf(DeterministicDistribution{}),
			"EmpiricalDistribution":     reflect.TypeOf(EmpiricalDistribution{}),
			"LogPearsonIIIDistribution": reflect.TypeOf(LogPearsonIIIDistribution{}),
			"BetaDistribution":          reflect.TypeOf(BetaDistribution{}),
			"ShiftedGammaDistribution":  reflect.TypeOf(ShiftedGammaDistribution{}),
		})
	if err != nil {
		return err
	}

	c.Value = value.(ContinuousDistribution)

	return nil
}

func UnmarshalCustomValue(data []byte, typeJsonField, valueJsonField string, customTypes map[string]reflect.Type) (interface{}, error) {
	m := map[string]interface{}{}
	if err := json.Unmarshal(data, &m); err != nil {
		return nil, err
	}
	typeName := m[typeJsonField].(string)
	var value ContinuousDistribution
	if ty, found := customTypes[typeName]; found {
		value = reflect.New(ty).Interface().(ContinuousDistribution)
	} else {
		return nil, errors.New("statistics: unmarshaling distribution " + typeName + " but it is not in the customTypes map.")
	}
	valueBytes, err := json.Marshal(m[valueJsonField])
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(valueBytes, &value); err != nil {
		return nil, err
	}
	return value, nil
}

var _ json.Unmarshaler = (*ContinuousDistributionContainer)(nil)
