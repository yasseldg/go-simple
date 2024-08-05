package tIndicator

import (
	"github.com/yasseldg/go-simple/data/dIter"
	"github.com/yasseldg/go-simple/types/sInts"
)

type InterPeriodsConfig interface {
	dIter.Inter

	Periods() int

	Count() int

	Reset()
	Next() bool
}

type PeriodsConfig struct {
	dIter.Inter

	periods sInts.InterIter
}

func NewPeriodsConfig(periods sInts.InterIter) *PeriodsConfig {
	return &PeriodsConfig{
		Inter: dIter.New(),

		periods: periods,
	}
}

func (st *PeriodsConfig) Periods() int {
	return int(st.periods.Value())
}

func (st *PeriodsConfig) Count() int {
	return st.periods.Count()
}

func (st *PeriodsConfig) Reset() {
	st.periods.Reset()
}

func (st *PeriodsConfig) Next() bool {
	return st.periods.Next()
}
