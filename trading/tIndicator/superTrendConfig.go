package tIndicator

import (
	"fmt"

	"github.com/yasseldg/go-simple/data/dIter"
	"github.com/yasseldg/go-simple/trading/tInterval"
	"github.com/yasseldg/go-simple/types/sFloats"
	"github.com/yasseldg/go-simple/types/sInts"
)

type InterSTConfig interface {
	dIter.InterIterConfig

	Get() InterSuperTrend

	Periods() int
	Multiplier() float64
	Interval() tInterval.Inter
	AtClose() bool
	Smoothed() bool
}

type STConfig struct {
	dIter.InterIterConfig

	intervals   tInterval.InterIterLimited
	periods     InterPeriodsConfig
	multipliers sFloats.InterIter
	at_close    sInts.InterIter
	smoothed    sInts.InterIter
}

func NewSTConfig(periods sInts.InterIter, multipliers sFloats.InterIter, intervals tInterval.InterIterLimited, at_close, smoothed sInts.InterIter, name string) *STConfig {

	name = fmt.Sprintf("SuperTrend %s", name)

	st_config := STConfig{
		InterIterConfig: dIter.NewIterConfig(name),

		periods:     NewPeriodsConfig(periods),
		multipliers: multipliers,
		intervals:   intervals,
		at_close:    at_close,
		smoothed:    smoothed,
	}

	st_config.Add(dIter.NewNameConfig("Periods", st_config.periods),
		dIter.NewNameConfig("Multipliers", st_config.multipliers),
		dIter.NewNameConfig("Intervals", st_config.intervals),
		dIter.NewNameConfig("AtClose", st_config.at_close),
		dIter.NewNameConfig("Smoothed", st_config.smoothed))

	return &st_config
}

func (st *STConfig) Get() InterSuperTrend {
	s := NewSuperTrend(st.Periods(), st.Multiplier(), st.Smoothed())
	return s
}

func (st *STConfig) Periods() int {
	return st.periods.Periods()
}

func (st *STConfig) Multiplier() float64 {
	return st.multipliers.Value()
}

func (st *STConfig) Interval() tInterval.Inter {
	return st.intervals.Item()
}

func (st *STConfig) AtClose() bool {
	return st.at_close.Value() == 1
}

func (st *STConfig) Smoothed() bool {
	return st.smoothed.Value() == 1
}
