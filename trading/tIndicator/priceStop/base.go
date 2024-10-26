package priceStop

import (
	"sync"

	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/trading/tCandle"
	"github.com/yasseldg/go-simple/trading/tSide"
)

type Base struct {
	mu sync.Mutex

	prev tCandle.Inter

	side tSide.Side
	low  float64
	high float64
	stop float64
}

func New(side tSide.Side) *Base {
	if side == tSide.DEFAULT {
		return nil
	}

	return &Base{
		mu: sync.Mutex{},

		prev: new(tCandle.Candle),

		side: side,
	}
}

func (ps *Base) String() string {
	return ""
}

func (ps *Base) Log() {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	sLog.Info("PriceStop: %s", ps.String())
}

func (ps *Base) SetStop(stop float64) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	ps.stop = stop
}

func (ps *Base) Stop() float64 {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	return ps.stop
}

func (ps *Base) Reset() {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	ps.prev = new(tCandle.Candle)

	ps.low = 0
	ps.high = 0
	ps.stop = 0
}

func (ps *Base) Add(candle tCandle.Inter) {
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

func (ps *Base) long(candle tCandle.Inter) {
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

func (ps *Base) short(candle tCandle.Inter) {
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
