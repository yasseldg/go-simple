package tIndicator

import (
	"fmt"
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

	close float64
}

func NewRSI(period int) *RSI {
	return &RSI{
		mu: sync.Mutex{},

		loss: *sFloats.NewSmoothedAverage(period),
		gain: *sFloats.NewSmoothedAverage(period),
	}
}

func (rsi *RSI) Period() int {
	return rsi.gain.Period()
}

func (rsi *RSI) Filled() bool {
	return rsi.gain.Filled()
}

func (rsi *RSI) String() string {
	return fmt.Sprintf("calc: %f  ..  gain: %f  ..  loss: %f", rsi.calc(), rsi.gain.Value(), rsi.loss.Value())
}

func (rsi *RSI) Log() {
	rsi.mu.Lock()
	defer rsi.mu.Unlock()

	sLog.Info("RSI: %s", rsi.String())
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

	rsi.close = close
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

	c      int
	candle tCandle.Candle
}

func NewRSIcandle(period int) *RSIcandle {
	return &RSIcandle{
		RSI: *NewRSI(period),
	}
}

func (rsi *RSIcandle) String() string {
	return fmt.Sprintf("c: %d: %s  ..  %s", rsi.c, sTime.ForLog(rsi.candle.Ts, 0), rsi.RSI.String())
}

func (rsi *RSIcandle) Log() {
	rsi.mu.Lock()
	defer rsi.mu.Unlock()

	sLog.Info("RSI: %s", rsi.String())
}

func (rsi *RSIcandle) Candle() tCandle.Candle {
	rsi.mu.Lock()
	defer rsi.mu.Unlock()

	return rsi.candle
}

func (rsi *RSIcandle) Add(candle tCandle.Candle) {
	rsi.mu.Lock()
	defer rsi.mu.Unlock()

	if candle.Close == 0 {
		return
	}

	rsi.add(candle.Close)

	rsi.c++
	rsi.candle = candle
}
