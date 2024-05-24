package tIndicator

import (
	"fmt"

	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/trading/tCandle"
	"github.com/yasseldg/go-simple/types/sTime"
)

type InterRSIcandle interface {
	interRSI

	Candle() tCandle.Inter
	Add(candle tCandle.Inter)
}

// RSIcandle is a RSI indicator for candles

type RSIcandle struct {
	RSI

	c      int
	candle tCandle.Inter
}

func NewRSIcandle(period int) *RSIcandle {
	return &RSIcandle{
		RSI:    *NewRSI(period),
		candle: new(tCandle.Candle),
	}
}

func (rsi *RSIcandle) String() string {
	return fmt.Sprintf("c: %d: %s  ..  %s", rsi.c, sTime.ForLog(rsi.candle.Ts(), 0), rsi.RSI.String())
}

func (rsi *RSIcandle) Log() {
	rsi.mu.Lock()
	defer rsi.mu.Unlock()

	sLog.Info("RSI: %s ", rsi.String())
}

func (rsi *RSIcandle) Candle() tCandle.Inter {
	rsi.mu.Lock()
	defer rsi.mu.Unlock()

	return rsi.candle
}

func (rsi *RSIcandle) Add(candle tCandle.Inter) {
	rsi.mu.Lock()
	defer rsi.mu.Unlock()

	if candle.Close() == 0 {
		return
	}

	rsi.add(candle.Close())

	rsi.c++
	rsi.candle = candle
}
