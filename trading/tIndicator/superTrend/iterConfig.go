package superTrend

import (
	"fmt"

	"github.com/yasseldg/go-simple/data/dIter"
	"github.com/yasseldg/go-simple/types/sFloats"
	"github.com/yasseldg/go-simple/types/sInts"
)

type iterConfig struct {
	dIter.InterIterConfig

	name        string
	periods     sInts.InterIter
	multipliers sFloats.InterIter
	smoothed    sInts.InterIter
}

func NewIterConfig(name string) *iterConfig {
	return &iterConfig{
		InterIterConfig: dIter.NewIterConfig(fmt.Sprintf("SuperTrend %s", name)),
		name:            name,
	}
}

func (st *iterConfig) SetPeriods(periods sInts.InterIter) {
	if st.periods != nil {
		return
	}

	st.periods = periods
	st.Add(dIter.NewNameConfig("Periods", st.periods))
}

func (st *iterConfig) SetMultiplier(multipliers sFloats.InterIter) {
	if st.multipliers != nil {
		return
	}

	st.multipliers = multipliers
	st.Add(dIter.NewNameConfig("Multipliers", st.multipliers))
}

func (st *iterConfig) SetSmoothed(smoothed sInts.InterIter) {
	if st.smoothed != nil {
		return
	}

	st.smoothed = smoothed
	st.Add(dIter.NewNameConfig("Smoothed", st.smoothed))
}

func (st *iterConfig) Get() (Inter, error) {
	if st.periods == nil {
		return nil, fmt.Errorf("periods is required")
	}

	if st.multipliers == nil {
		return nil, fmt.Errorf("multipliers is required")
	}

	if st.smoothed == nil {
		return nil, fmt.Errorf("smoothed is required")
	}

	return New(
		int(st.periods.Value()),
		st.multipliers.Value(),
		st.smoothed.Value() == 1), nil
}

func (st *iterConfig) Clone() InterIterConfig {

	clone := NewIterConfig(st.name)
	clone.SetPeriods(st.periods.Clone())
	clone.SetMultiplier(st.multipliers.Clone())
	clone.SetSmoothed(st.smoothed.Clone())

	return clone
}
