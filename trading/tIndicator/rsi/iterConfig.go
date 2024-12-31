package rsi

import (
	"fmt"

	"github.com/yasseldg/go-simple/data/dIter"
	"github.com/yasseldg/go-simple/types/sInts"
)

type iterConfig struct {
	dIter.InterIterConfig

	name    string
	periods sInts.InterIter
}

func NewIterConfig(name string) *iterConfig {
	return &iterConfig{
		InterIterConfig: dIter.NewIterConfig(fmt.Sprintf("RSI %s", name)),
		name:            name,
	}
}

func (ic *iterConfig) SetPeriods(periods sInts.InterIter) InterIterConfig {
	if ic.periods != nil {
		return ic
	}

	ic.periods = periods
	ic.Add(dIter.NewNameConfig("Periods", ic.periods))
	return ic
}

func (ic *iterConfig) Periods() int64 {
	return ic.periods.Value()
}

func (ic *iterConfig) Get() (Inter, error) {

	if ic.periods == nil {
		return nil, fmt.Errorf("periods is required")
	}

	return New(int(ic.Periods())), nil
}

func (ic *iterConfig) Clone() InterIterConfig {

	clone := NewIterConfig(ic.name)
	clone.SetPeriods(ic.periods.Clone())

	return clone
}
