package tIndicator

import (
	"math"
	"sync"

	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/trading/tCandle"
	"github.com/yasseldg/go-simple/types/sFloats"
	"github.com/yasseldg/go-simple/types/sTime"
)

// ATR (Average True Range)

type ATR struct {
	mu sync.Mutex

	prev tCandle.Candle

	trs sFloats.SmoothedAverage

	c int
}

func NewATR(period int) *ATR {
	return &ATR{
		mu: sync.Mutex{},

		trs: *sFloats.NewSmoothedAverage(period),
	}
}

func (atr *ATR) Log() {
	atr.mu.Lock()
	defer atr.mu.Unlock()

	sLog.Info("ATR %d: %s  ..  %f", atr.c, sTime.ForLog(atr.prev.Ts, 0), atr.trs.Value())
}

// Add adds a candle to the ATR
func (atr *ATR) Add(candle tCandle.Candle) {
	atr.mu.Lock()
	defer atr.mu.Unlock()

	atr.add(candle)
}

func (atr *ATR) add(candle tCandle.Candle) {
	atr.c++

	if atr.prev.Close > 0 {
		atr.trs.AddPos(max(candle.High-candle.Low, math.Abs(candle.High-atr.prev.Close), math.Abs(candle.Low-atr.prev.Close)))
	}

	atr.prev = candle
}

func (atr *ATR) Get() float64 {
	atr.mu.Lock()
	defer atr.mu.Unlock()

	return atr.trs.Value()
}
