package bollingerBands

import (
	"fmt"

	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/trading/tCandle"
	"github.com/yasseldg/go-simple/types/sTime"
)

type candle struct {
	base

	c    int
	prev tCandle.Inter
}

func NewCandle(periods int, deviations float64) *candle {
	return &candle{
		base: *New(periods, deviations),
		prev: new(tCandle.Candle),
	}
}

func (bb *candle) String() string {
	return fmt.Sprintf("c: %d: %s  ..  %s", bb.c, sTime.ForLog(bb.prev.Ts(), 0), bb.base.String())
}

func (bb *candle) Log() {
	bb.mu.Lock()
	defer bb.mu.Unlock()

	sLog.Info("BBands: %s", bb.String())
}

func (bb *candle) Candle() tCandle.Inter {
	return bb.prev
}

func (bb *candle) Add(candle tCandle.Inter) {
	bb.mu.Lock()
	defer bb.mu.Unlock()

	if candle.Close() == 0 {
		return
	}

	if bb.prev.Ts() >= candle.Ts() {
		return
	}

	bb.add(candle.Close())

	bb.c++
	bb.prev = candle
}
