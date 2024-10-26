package bollingerBands

import (
	"fmt"

	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/trading/tCandle"
	"github.com/yasseldg/go-simple/types/sTime"
)

type Candle struct {
	Base

	c    int
	prev tCandle.Inter
}

func NewCandle(periods int, deviations float64) *Candle {
	return &Candle{
		Base: *New(periods, deviations),
		prev: new(tCandle.Candle),
	}
}

func (bb *Candle) String() string {
	return fmt.Sprintf("c: %d: %s  ..  %s", bb.c, sTime.ForLog(bb.prev.Ts(), 0), bb.Base.String())
}

func (bb *Candle) Log() {
	bb.mu.Lock()
	defer bb.mu.Unlock()

	sLog.Info("BBands: %s", bb.String())
}

func (bb *Candle) Candle() tCandle.Inter {
	return bb.prev
}

func (bb *Candle) Add(candle tCandle.Inter) {
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
