package tIndicator

import (
	"sync"

	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/trading/tCandle"
	"github.com/yasseldg/go-simple/types/sFloats"
	"github.com/yasseldg/go-simple/types/sTime"
)

//  BBands (Bollinger Bands)

type BBands struct {
	mu sync.Mutex

	deviations float64
	closes     sFloats.PeriodValues

	mean float64
	std  float64
}

func NewBBands(period int, deviations float64) *BBands {
	return &BBands{
		mu: sync.Mutex{},

		deviations: deviations,
		closes:     *sFloats.NewPeriodValues(period),
	}
}

func (bb *BBands) Add(close float64) {
	bb.mu.Lock()
	defer bb.mu.Unlock()

	if close == 0 {
		return
	}

	bb.add(close)
}

func (bb *BBands) Filled() bool {
	bb.mu.Lock()
	defer bb.mu.Unlock()

	return bb.closes.Filled()
}

func (bb *BBands) Get() (mean, upper, lower float64) {
	bb.mu.Lock()
	defer bb.mu.Unlock()

	return bb.calc(bb.deviations)
}

func (bb *BBands) Calc(deviations float64) (mean, upper, lower float64) {
	bb.mu.Lock()
	defer bb.mu.Unlock()

	return bb.calc(deviations)
}

func (bb *BBands) add(close float64) {
	bb.closes.Add(close)

	if !bb.closes.Filled() {
		return
	}

	bb.mean, bb.std = bb.closes.MeanStdDev()
}

func (bb *BBands) calc(deviations float64) (mean, upper, lower float64) {
	if !bb.closes.Filled() {
		return 0, 0, 0
	}

	mean = bb.mean
	upper = bb.mean + (bb.std * deviations)
	lower = bb.mean - (bb.std * deviations)
	return
}

//  BBcandle is a BBands indicator for candles

type BBcandle struct {
	BBands

	c    int
	prev tCandle.Candle
}

func NewBBcandle(period int, deviations float64) *BBcandle {
	return &BBcandle{
		BBands: *NewBBands(period, deviations),
	}
}

func (bb *BBcandle) Log() {
	bb.mu.Lock()
	defer bb.mu.Unlock()

	sLog.Info("BBands: %d: %s  .. mean: %f  .. std: %f ", bb.c, sTime.ForLog(bb.prev.Ts, 0), bb.mean, bb.std)
}

func (bb *BBcandle) Add(candle tCandle.Candle) {
	bb.mu.Lock()
	defer bb.mu.Unlock()

	if candle.Close == 0 {
		return
	}

	bb.add(candle.Close)

	bb.c++
	bb.prev = candle
}

func (bb *BBcandle) Filled() bool {
	return bb.BBands.Filled()
}
