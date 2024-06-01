package tIndicator

import (
	"github.com/yasseldg/go-simple/types/sFloats"
	"github.com/yasseldg/go-simple/types/sInts"
)

type InterSTConfig interface {
	InterPeriodsConfig

	Multiplier() float64
	AtClose() bool
	Smoothed() bool

	Reset()
	Next() bool
}

type STConfig struct {
	InterPeriodsConfig

	multipliers sFloats.Iter
	at_close    bool
	smoothed    bool
}

func NewSTConfig(periods sInts.Iter, multipliers sFloats.Iter, at_close, smoothed bool) *STConfig {
	return &STConfig{
		InterPeriodsConfig: NewPeriodsConfig(periods),

		multipliers: multipliers,
		at_close:    at_close,
		smoothed:    smoothed,
	}
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

func (st *STConfig) Next() bool {

	if st.multipliers.Next() {
		return true
	}

	if st.InterPeriodsConfig.Next() {
		st.multipliers.Reset()

		return st.Next()
	}

	return false
}

func (st *STConfig) Reset() {
	st.multipliers.Reset()

	st.InterPeriodsConfig.Reset()

	st.InterPeriodsConfig.Next()
}
