package tIndicator

import (
	"github.com/yasseldg/go-simple/data/dIter"
	"github.com/yasseldg/go-simple/types/sInts"
)

type InterPeriodsConfig interface {
	dIter.Inter

	Periods() int

	Reset()
	Next() bool
}

type PeriodsConfig struct {
	dIter.Inter

	periods sInts.Iter
}

func NewPeriodsConfig(periods sInts.Iter) *PeriodsConfig {
	return &PeriodsConfig{
		Inter: dIter.New(),

		periods: periods,
	}
}

func (st *PeriodsConfig) Periods() int {
	return int(st.periods.Value())
}

func (st *PeriodsConfig) Next() bool {

	return st.periods.Next()
}

func (st *PeriodsConfig) Reset() {
	st.periods.Reset()
}
