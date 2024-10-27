package atr

import (
	"fmt"

	"github.com/yasseldg/go-simple/data/dIter"
	"github.com/yasseldg/go-simple/types/sInts"
)

type IterConfig struct {
	dIter.InterIterConfig

	name     string
	periods  sInts.InterIter
	smoothed sInts.InterIter
}

func NewIterConfig(name string) *IterConfig {
	return &IterConfig{
		InterIterConfig: dIter.NewIterConfig(fmt.Sprintf("SuperTrend %s", name)),
		name:            name,
	}
}

func (st *IterConfig) SetPeriods(periods sInts.InterIter) {
	if st.periods != nil {
		return
	}

	st.periods = periods
	st.Add(dIter.NewNameConfig("Periods", st.periods))
}

func (st *IterConfig) SetSmoothed(smoothed sInts.InterIter) {
	if st.smoothed != nil {
		return
	}

	st.smoothed = smoothed
	st.Add(dIter.NewNameConfig("Smoothed", st.smoothed))
}

func (st *IterConfig) Smoothed() bool {
	return st.smoothed.Value() == 1
}

func (st *IterConfig) Get() Inter {
	if st.periods == nil || st.smoothed == nil {
		return nil
	}

	if st.Smoothed() {
		return NewSmoothed(int(st.periods.Value()))
	}
	return NewAvg(int(st.periods.Value()))
}

func (st *IterConfig) Clone() InterIterConfig {

	clone := NewIterConfig(st.name)
	clone.SetPeriods(st.periods.Clone())
	clone.SetSmoothed(st.smoothed.Clone())

	return clone
}
