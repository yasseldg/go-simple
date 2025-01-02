package rsi

import (
	"fmt"

	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/trading/tCandle"
	"github.com/yasseldg/go-simple/types/sTime"
)

// RSIcandle is a RSI indicator for candles

type Candle struct {
	base

	c      int
	candle tCandle.Inter
}

func NewCandle(period int) *Candle {
	return &Candle{
		base:   *New(period),
		candle: new(tCandle.Candle),
	}
}

func (rsi *Candle) String() string {
	return fmt.Sprintf("c: %d: %s  ..  %s",
		rsi.c, sTime.ForLog(rsi.candle.Ts(), 0), rsi.base.String())
}

func (rsi *Candle) Log() {
	rsi.mu.Lock()
	defer rsi.mu.Unlock()

	sLog.Info("RSI: %s ", rsi.String())
}

func (rsi *Candle) Candle() tCandle.Inter {
	rsi.mu.Lock()
	defer rsi.mu.Unlock()

	return rsi.candle
}

func (rsi *Candle) Add(candle tCandle.Inter) {
	rsi.mu.Lock()
	defer rsi.mu.Unlock()

	if candle.Close() == 0 {
		return
	}

	rsi.add(candle.Close())

	rsi.c++
	rsi.candle = candle
}
