package atr

import (
	"fmt"
	"math"

	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/trading/tCandle"
	"github.com/yasseldg/go-simple/types/sFloats"
)

type Smoothed struct {
	Base

	trs sFloats.SmoothedAverage
}

// NewSmoothed create a new Smoothed ATR
func NewSmoothed(periods int) *Smoothed {
	return &Smoothed{
		Base: *newBase(),
		trs:  *sFloats.NewSmoothedAverage(periods),
	}
}

func (atr *Smoothed) String() string {
	return fmt.Sprintf("%s  ..  %f", atr.Base.String(), atr.get())
}

func (atr *Smoothed) Log() {
	atr.mu.Lock()
	defer atr.mu.Unlock()

	sLog.Info("SmATR: %s", atr.String())
}

func (atr *Smoothed) Periods() int {
	return atr.trs.Periods()
}

func (atr *Smoothed) Filled() bool {
	return atr.trs.Filled()
}

// Add adds a candle to the SmATR
func (atr *Smoothed) Add(candle tCandle.Inter) {
	atr.mu.Lock()
	defer atr.mu.Unlock()

	atr.add(candle)
}

func (atr *Smoothed) Get() float64 {
	atr.mu.Lock()
	defer atr.mu.Unlock()

	return atr.get()
}

// private methods

func (atr *Smoothed) add(candle tCandle.Inter) {
	atr.c++

	if atr.prev.Close() > 0 {
		atr.trs.AddPos(max(candle.High()-candle.Low(), math.Abs(candle.High()-atr.prev.Close()), math.Abs(candle.Low()-atr.prev.Close())))
	}

	atr.prev = candle
}

func (atr *Smoothed) get() float64 {
	return atr.trs.Value()
}
