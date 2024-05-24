package tIndicator

import (
	"fmt"

	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/trading/tCandle"
	"github.com/yasseldg/go-simple/types/sTime"
)

//  BBcandle is a BBands indicator for candles

type InterBBcandle interface {
	interBBands

	Candle() tCandle.Inter
	Add(candle tCandle.Inter)
}

type BBcandle struct {
	BBands

	c    int
	prev tCandle.Inter
}

func NewBBcandle(period int, deviations float64) *BBcandle {
	return &BBcandle{
		BBands: *NewBBands(period, deviations),
		prev:   new(tCandle.Candle),
	}
}

func (bb *BBcandle) String() string {
	return fmt.Sprintf("c: %d: %s  ..  %s", bb.c, sTime.ForLog(bb.prev.Ts(), 0), bb.BBands.String())
}

func (bb *BBcandle) Log() {
	bb.mu.Lock()
	defer bb.mu.Unlock()

	sLog.Info("BBands: %s", bb.String())
}

func (bb *BBcandle) Add(candle tCandle.Inter) {
	bb.mu.Lock()
	defer bb.mu.Unlock()

	if candle.Close() == 0 {
		return
	}

	bb.add(candle.Close())

	bb.c++
	bb.prev = candle
}
