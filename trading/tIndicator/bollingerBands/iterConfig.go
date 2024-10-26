package bollingerBands

import (
	"fmt"

	"github.com/yasseldg/go-simple/data/dIter"
	"github.com/yasseldg/go-simple/types/sFloats"
	"github.com/yasseldg/go-simple/types/sInts"
)

type IterConfig struct {
	dIter.InterIterConfig

	name       string
	periods    sInts.InterIter
	deviations sFloats.InterIter
}

func NewIterConfig(name string) *IterConfig {
	return &IterConfig{
		InterIterConfig: dIter.NewIterConfig(fmt.Sprintf("SuperTrend %s", name)),
		name:            name,
	}
}

func (bb *IterConfig) SetPeriods(periods sInts.InterIter) {
	if bb.periods != nil {
		return
	}

	bb.periods = periods
	bb.Add(dIter.NewNameConfig("Periods", bb.periods))
}

func (bb *IterConfig) SetDeviations(deviations sFloats.InterIter) {
	if bb.deviations != nil {
		return
	}

	bb.deviations = deviations
	bb.Add(dIter.NewNameConfig("Deviations", bb.deviations))
}

func (bb *IterConfig) Get() Inter {
	if bb.periods == nil || bb.deviations == nil {
		return nil
	}

	return New(
		int(bb.periods.Value()),
		bb.deviations.Value())
}

func (bb *IterConfig) GetCandle() InterCandle {
	if bb.periods == nil || bb.deviations == nil {
		return nil
	}

	return NewCandle(
		int(bb.periods.Value()),
		bb.deviations.Value())
}

func (bb *IterConfig) Clone() InterIterConfig {

	clone := NewIterConfig(bb.name)
	clone.SetPeriods(bb.periods.Clone())
	clone.SetDeviations(bb.deviations.Clone())

	return clone
}
