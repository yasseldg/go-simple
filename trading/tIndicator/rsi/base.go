package rsi

import (
	"fmt"
	"sync"

	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/types/sFloats"
)

type Base struct {
	mu sync.Mutex

	loss sFloats.SmoothedAverage
	gain sFloats.SmoothedAverage

	close float64
}

func New(period int) *Base {
	return &Base{
		mu: sync.Mutex{},

		loss: *sFloats.NewSmoothedAverage(period),
		gain: *sFloats.NewSmoothedAverage(period),
	}
}

func (rsi *Base) String() string {
	return fmt.Sprintf("calc: %f  ..  gain: %f  ..  loss: %f", rsi.calc(), rsi.gain.Value(), rsi.loss.Value())
}

func (rsi *Base) Log() {
	rsi.mu.Lock()
	defer rsi.mu.Unlock()

	sLog.Info("RSI: %s", rsi.String())
}

func (rsi *Base) Periods() int {
	return rsi.gain.Periods()
}

func (rsi *Base) Filled() bool {
	return rsi.gain.Filled()
}

func (rsi *Base) Get() float64 {
	rsi.mu.Lock()
	defer rsi.mu.Unlock()

	if !rsi.gain.Filled() {
		return 0
	}

	return rsi.calc()
}

func (rsi *Base) Add(close float64) {
	rsi.mu.Lock()
	defer rsi.mu.Unlock()

	if close == 0 {
		return
	}

	rsi.add(close)
}

func (rsi *Base) add(close float64) {

	if rsi.close == 0 {
		rsi.close = close
		return
	}

	delta := close - rsi.close

	if delta > 0 {
		rsi.gain.AddPos(delta)
		rsi.loss.AddNeg(0)
	} else {
		rsi.gain.AddPos(0)
		rsi.loss.AddNeg(delta)
	}

	rsi.close = close
}

func (rsi *Base) calc() float64 {

	temp := rsi.gain.Value() + rsi.loss.Value()
	if !((-0.00000000000001 < temp) && (temp < 0.00000000000001)) {
		return 100 * (rsi.gain.Value() / temp)
	}

	return 0
}
