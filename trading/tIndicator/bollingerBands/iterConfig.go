package bollingerBands

import (
	"fmt"

	"github.com/yasseldg/go-simple/data/dIter"
	"github.com/yasseldg/go-simple/types/sFloats"
	"github.com/yasseldg/go-simple/types/sInts"
)

type iterConfig struct {
	dIter.InterIterConfig

	name       string
	periods    sInts.InterIter
	deviations sFloats.InterIter
}

func NewIterConfig(name string) *iterConfig {
	return &iterConfig{
		InterIterConfig: dIter.NewIterConfig(fmt.Sprintf("BBands %s", name)),
		name:            name,
	}
}

func (bb *iterConfig) SetPeriods(periods sInts.InterIter) {
	if bb.periods != nil {
		return
	}

	bb.periods = periods
	bb.Add(dIter.NewNameConfig("Periods", bb.periods))
}

func (bb *iterConfig) SetDeviations(deviations sFloats.InterIter) {
	if bb.deviations != nil {
		return
	}

	bb.deviations = deviations
	bb.Add(dIter.NewNameConfig("Deviations", bb.deviations))
}

func (bb *iterConfig) Get() (Inter, error) {
	if bb.periods == nil {
		return nil, fmt.Errorf("periods is required")
	}

	if bb.deviations == nil {
		return nil, fmt.Errorf("deviations is required")
	}

	return New(
		int(bb.periods.Value()),
		bb.deviations.Value()), nil
}

func (bb *iterConfig) GetCandle() InterCandle {
	if bb.periods == nil || bb.deviations == nil {
		return nil
	}

	return NewCandle(
		int(bb.periods.Value()),
		bb.deviations.Value())
}

func (bb *iterConfig) Clone() InterIterConfig {

	clone := NewIterConfig(bb.name)
	clone.SetPeriods(bb.periods.Clone())
	clone.SetDeviations(bb.deviations.Clone())

	return clone
}
