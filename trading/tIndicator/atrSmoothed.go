package tIndicator

import (
	"fmt"
	"math"

	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/trading/tCandle"
	"github.com/yasseldg/go-simple/types/sFloats"
)

type SmATR struct {
	BaseATR

	trs sFloats.SmoothedAverage
}

func NewSmATR(periods int) *SmATR {
	return &SmATR{
		BaseATR: *newBaseATR(),
		trs:     *sFloats.NewSmoothedAverage(periods),
	}
}

func (atr *SmATR) String() string {
	return fmt.Sprintf("%s  ..  %f", atr.BaseATR.String(), atr.get())
}

func (atr *SmATR) Log() {
	atr.mu.Lock()
	defer atr.mu.Unlock()

	sLog.Info("SmATR: %s", atr.String())
}

func (atr *SmATR) Periods() int {
	return atr.trs.Periods()
}

func (atr *SmATR) Filled() bool {
	return atr.trs.Filled()
}

// Add adds a candle to the SmATR
func (atr *SmATR) Add(candle tCandle.Inter) {
	atr.mu.Lock()
	defer atr.mu.Unlock()

	atr.add(candle)
}

func (atr *SmATR) Get() float64 {
	atr.mu.Lock()
	defer atr.mu.Unlock()

	return atr.get()
}

// private methods

func (atr *SmATR) add(candle tCandle.Inter) {
	atr.c++

	if atr.prev.Close() > 0 {
		atr.trs.AddPos(max(candle.High()-candle.Low(), math.Abs(candle.High()-atr.prev.Close()), math.Abs(candle.Low()-atr.prev.Close())))
	}

	atr.prev = candle
}

func (atr *SmATR) get() float64 {
	return atr.trs.Value()
}
