package atr

import (
	"fmt"
	"math"

	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/trading/tCandle"
	"github.com/yasseldg/go-simple/types/sFloats"
)

type Avg struct {
	Base

	trs sFloats.InterPeriodValues
}

func NewAvg(periods int) *Avg {
	return &Avg{
		Base: *newBase(),
		trs:  sFloats.NewPeriodValues(periods),
	}
}

func (atr *Avg) String() string {
	return fmt.Sprintf("%s  ..  atr: %f", atr.Base.String(), atr.get())
}

func (atr *Avg) Log() {
	atr.mu.Lock()
	defer atr.mu.Unlock()

	sLog.Info("AvgATR: %s", atr.String())
}

func (atr *Avg) Periods() int {
	return atr.trs.Periods()
}

func (atr *Avg) Filled() bool {
	return atr.trs.Filled()
}

// Add adds a candle to the AvgATR
func (atr *Avg) Add(candle tCandle.Inter) {
	atr.mu.Lock()
	defer atr.mu.Unlock()

	atr.add(candle)
}

func (atr *Avg) Get() float64 {
	atr.mu.Lock()
	defer atr.mu.Unlock()

	return atr.get()
}

// private methods

func (atr *Avg) add(candle tCandle.Inter) {
	atr.c++

	if atr.prev.Close() > 0 {
		atr.trs.Add(max(candle.High()-candle.Low(), math.Abs(candle.High()-atr.prev.Close()), math.Abs(candle.Low()-atr.prev.Close())))
	}

	atr.prev = candle
}

func (atr *Avg) get() float64 {
	return atr.trs.Mean()
}
