package atr

import (
	"fmt"
	"sync"

	"github.com/yasseldg/go-simple/trading/tCandle"
	"github.com/yasseldg/go-simple/types/sTime"
)

type Base struct {
	mu sync.Mutex

	prev tCandle.Inter

	c int
}

func newBase() *Base {
	return &Base{
		mu: sync.Mutex{},

		prev: new(tCandle.Candle),
	}
}

func (atr *Base) String() string {
	return fmt.Sprintf("%d: %s", atr.c, sTime.ForLog(atr.prev.Ts(), 0))
}

func (atr *Base) Count() int {
	atr.mu.Lock()
	defer atr.mu.Unlock()

	return atr.c
}

func (atr *Base) Prev() tCandle.Inter {
	atr.mu.Lock()
	defer atr.mu.Unlock()

	return atr.prev
}
