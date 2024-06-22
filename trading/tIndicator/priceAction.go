package tIndicator

import (
	"sync"

	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/trading/tCandle"
)

type InterPriceAction interface {
	String() string
	Log()

	Add(candle tCandle.Inter)
}

type PriceAction struct {
	mu sync.Mutex

	lows  []float64
	highs []float64
}

func NewPriceAction() *PriceAction {
	return &PriceAction{
		mu: sync.Mutex{},

		lows:  make([]float64, 0),
		highs: make([]float64, 0),
	}
}

func (pa *PriceAction) String() string {
	pa.mu.Lock()
	defer pa.mu.Unlock()

	return ""
}

func (pa *PriceAction) Log() {
	pa.mu.Lock()
	defer pa.mu.Unlock()

	sLog.Info("PriceAction: %s", pa.String())
}

func (pa *PriceAction) Add(candle tCandle.Inter) {
	pa.mu.Lock()
	defer pa.mu.Unlock()

}
