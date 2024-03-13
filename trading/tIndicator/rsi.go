package tIndicator

import (
	"sync"

	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/trading/tCandle"
	"github.com/yasseldg/go-simple/types/sFloats"
	"github.com/yasseldg/go-simple/types/sTime"
)

// RSI (Relative Strength Index)

type RSI struct {
	mu sync.Mutex

	loss sFloats.SmoothedAverage
	gain sFloats.SmoothedAverage

	sum   float64
	close float64
}

func NewRSI(period int) *RSI {
	return &RSI{
		mu: sync.Mutex{},

		loss: *sFloats.NewSmoothedAverage(period),
		gain: *sFloats.NewSmoothedAverage(period),
	}
}

func (rsi *RSI) Add(close float64) {
	rsi.mu.Lock()
	defer rsi.mu.Unlock()

	if close == 0 {
		return
	}

	rsi.add(close)
}

func (rsi *RSI) Get() float64 {
	rsi.mu.Lock()
	defer rsi.mu.Unlock()

	if !rsi.gain.Filled() {
		return 0
	}

	return rsi.calc()
}

func (rsi *RSI) add(close float64) {

	if rsi.close == 0 {
		rsi.close = close
		return
	}

	delta := close - rsi.close

	if delta > 0 {
		rsi.gain.AddPos(delta)
		rsi.loss.AddNeg(0)
	} else {
		rsi.gain.AddPos(0)
		rsi.loss.AddNeg(delta)
	}
}

func (rsi *RSI) calc() float64 {

	temp := rsi.gain.Value() + rsi.loss.Value()
	if !((-0.00000000000001 < temp) && (temp < 0.00000000000001)) {
		return 100 * (rsi.gain.Value() / temp)
	}

	return 0
}

// RSIcandle is a RSI indicator for candles

type RSIcandle struct {
	RSI

	c    int
	prev tCandle.Candle
}

func NewRSIcandle(period int) *RSIcandle {
	return &RSIcandle{
		RSI: *NewRSI(period),
	}
}

func (rsi *RSIcandle) Log() {
	rsi.mu.Lock()
	defer rsi.mu.Unlock()

	sLog.Info("RSI: %d: %s  ..  %f  ..  %f  ..  %f  ..  %f", rsi.c, sTime.ForLog(rsi.prev.Ts, 0), rsi.calc(), rsi.gain.Value(), rsi.loss.Value(), rsi.sum)
}

func (rsi *RSIcandle) Add(candle tCandle.Candle) {
	rsi.mu.Lock()
	defer rsi.mu.Unlock()

	if candle.Close == 0 {
		return
	}

	rsi.add(candle.Close)

	rsi.c++
	rsi.prev = candle
}
