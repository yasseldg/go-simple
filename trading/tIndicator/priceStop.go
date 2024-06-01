package tIndicator

import (
	"sync"

	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/trading/tCandle"
	"github.com/yasseldg/go-simple/trading/tSide"
)

// El Índice de Movimiento Direccional Promedio (ADX, por sus siglas en inglés)

type PriceStop struct {
	mu sync.Mutex

	prev tCandle.Candle

	side tSide.Side
	low  float64
	high float64
	stop float64
}

func NewPriceStop(side tSide.Side) *PriceStop {
	if side == tSide.DEFAULT {
		return nil
	}

	return &PriceStop{
		mu: sync.Mutex{},

		side: side,
	}
}

func (ps *PriceStop) Log() {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	sLog.Info("PriceStop: ")
}

func (ps *PriceStop) Add(candle tCandle.Candle) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	switch ps.side {
	case tSide.Buy:
		ps.long(candle)

	case tSide.Sell:
		ps.short(candle)
	}

	ps.prev = candle
}

func (ps *PriceStop) long(candle tCandle.Candle) {
	if ps.prev.Low() > candle.Low() {
		ps.low = candle.Low()
	}

	if candle.Close() > ps.high && ps.low > 0 {
		ps.stop = ps.low
	}

	if candle.High() > ps.high {
		ps.high = candle.High()
	}
}

func (ps *PriceStop) short(candle tCandle.Candle) {
	if candle.Low() < ps.low {
		ps.low = candle.Low()

		if ps.stop > ps.high {
			ps.stop = ps.high
		}
	}

	if ps.prev.High() < candle.High() {
		ps.high = candle.High()
	}
}

func (ps *PriceStop) SetStop(stop float64) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	ps.stop = stop
}

func (ps *PriceStop) Stop() float64 {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	return ps.stop
}

func (ps *PriceStop) Reset() {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	ps.prev = tCandle.Candle{}

	ps.low = 0
	ps.high = 0
	ps.stop = 0
}
