package tIndicator

import (
	"sync"

	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/trading/tCandle"
)

// El Índice de Movimiento Direccional Promedio (ADX, por sus siglas en inglés)

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

func (pa *PriceAction) Log() {
	pa.mu.Lock()
	defer pa.mu.Unlock()

	sLog.Info("PriceAction: ")
}

func (pa *PriceAction) Add(candle tCandle.Candle) {
	pa.mu.Lock()
	defer pa.mu.Unlock()

}
