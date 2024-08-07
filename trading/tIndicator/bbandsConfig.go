package tIndicator

import (
	"github.com/yasseldg/go-simple/types/sFloats"
	"github.com/yasseldg/go-simple/types/sInts"
)

type InterBBConfig interface {
	InterPeriodsConfig

	Deviations() float64
	AtClose() bool

	Count() int

	Reset()
	Next() bool
}

type BBConfig struct {
	InterPeriodsConfig

	deviations sFloats.InterIter
	at_close   bool
}

func NewBBConfig(periods sInts.InterIter, deviations sFloats.InterIter, at_close bool) *BBConfig {
	return &BBConfig{
		InterPeriodsConfig: NewPeriodsConfig(periods),

		deviations: deviations,
		at_close:   at_close,
	}
}

func (bb *BBConfig) Deviations() float64 {
	return bb.deviations.Value()
}

func (bb *BBConfig) AtClose() bool {
	return bb.at_close
}

func (bb *BBConfig) Count() int {
	return bb.InterPeriodsConfig.Count() * bb.deviations.Count()
}

func (st *BBConfig) Reset() {
	st.deviations.Reset()

	st.InterPeriodsConfig.Reset()

	st.InterPeriodsConfig.Next()
}

func (st *BBConfig) Next() bool {

	if st.deviations.Next() {
		return true
	}

	if st.InterPeriodsConfig.Next() {
		st.deviations.Reset()

		return st.Next()
	}

	return false
}
