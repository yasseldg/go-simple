package tIndicator

import (
	"fmt"

	"github.com/yasseldg/go-simple/data/dIter"
	"github.com/yasseldg/go-simple/types/sFloats"
	"github.com/yasseldg/go-simple/types/sInts"
)

type InterSTConfig interface {
	dIter.InterIterConfig

	Periods() int
	Multiplier() float64
	AtClose() bool
	Smoothed() bool
}

type STConfig struct {
	dIter.InterIterConfig

	periods     InterPeriodsConfig
	multipliers sFloats.InterIter
	at_close    bool
	smoothed    bool
}

func NewSTConfig(periods sInts.InterIter, multipliers sFloats.InterIter, at_close, smoothed bool, name string) *STConfig {
	name = fmt.Sprintf("SuperTrend %s", name)
	st_config := STConfig{
		InterIterConfig: dIter.NewIterConfig(name),

		periods:     NewPeriodsConfig(periods),
		multipliers: multipliers,
		at_close:    at_close,
		smoothed:    smoothed,
	}

	st_config.Add(dIter.NewNameConfig("Periods", st_config.periods),
		dIter.NewNameConfig("Multipliers", st_config.multipliers))

	return &st_config
}

func (st *STConfig) Periods() int {
	return st.periods.Periods()
}

func (st *STConfig) Multiplier() float64 {
	return st.multipliers.Value()
}

func (st *STConfig) AtClose() bool {
	return st.at_close
}

func (st *STConfig) Smoothed() bool {
	return st.smoothed
}
