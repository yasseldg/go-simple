package tIndicator

import (
	"fmt"
	"sync"

	"github.com/yasseldg/go-simple/trading/tCandle"
	"github.com/yasseldg/go-simple/types/sTime"
)

// ATR (Average True Range)

type InterATR interface {
	String() string
	Log()

	Count() int
	Periods() int
	Filled() bool
	Get() float64
	Prev() tCandle.Inter

	Add(candle tCandle.Inter)
}

type BaseATR struct {
	mu sync.Mutex

	prev tCandle.Inter

	c int
}

func newBaseATR() *BaseATR {
	return &BaseATR{
		mu: sync.Mutex{},

		prev: new(tCandle.Candle),
	}
}

func (atr *BaseATR) String() string {
	return fmt.Sprintf("%d: %s", atr.c, sTime.ForLog(atr.prev.Ts(), 0))
}

func (atr *BaseATR) Count() int {
	atr.mu.Lock()
	defer atr.mu.Unlock()

	return atr.c
}

func (atr *BaseATR) Prev() tCandle.Inter {
	atr.mu.Lock()
	defer atr.mu.Unlock()

	return atr.prev
}
