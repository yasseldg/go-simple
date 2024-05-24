package tIndicator

import (
	"fmt"
	"math"

	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/trading/tCandle"
	"github.com/yasseldg/go-simple/types/sFloats"
)

type AvgATR struct {
	BaseATR

	trs sFloats.InterPeriodValues
}

func NewAvgATR(periods int) *AvgATR {
	return &AvgATR{
		BaseATR: *newBaseATR(),
		trs:     sFloats.NewPeriodValues(periods),
	}
}

func (atr *AvgATR) String() string {
	return fmt.Sprintf("%s  ..  atr: %f", atr.BaseATR.String(), atr.get())
}

func (atr *AvgATR) Log() {
	atr.mu.Lock()
	defer atr.mu.Unlock()

	sLog.Info("AvgATR: %s", atr.String())
}

func (atr *AvgATR) Periods() int {
	return atr.trs.Periods()
}

func (atr *AvgATR) Filled() bool {
	return atr.trs.Filled()
}

// Add adds a candle to the AvgATR
func (atr *AvgATR) Add(candle tCandle.Inter) {
	atr.mu.Lock()
	defer atr.mu.Unlock()

	atr.add(candle)
}

func (atr *AvgATR) Get() float64 {
	atr.mu.Lock()
	defer atr.mu.Unlock()

	return atr.get()
}

// private methods

func (atr *AvgATR) add(candle tCandle.Inter) {
	atr.c++

	if atr.prev.Close() > 0 {
		atr.trs.Add(max(candle.High()-candle.Low(), math.Abs(candle.High()-atr.prev.Close()), math.Abs(candle.Low()-atr.prev.Close())))
	}

	atr.prev = candle
}

func (atr *AvgATR) get() float64 {
	return atr.trs.Mean()
}
