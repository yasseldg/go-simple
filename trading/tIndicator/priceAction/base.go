package priceAction

import (
	"fmt"
	"sync"

	"github.com/yasseldg/go-simple/trading/tIndicator/indicator"

	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/trading/tCandle"
	"github.com/yasseldg/go-simple/types/sTime"
)

const (
	downtrend = iota - 1
	neutral_down
	neutral_up
	uptrend
)

type Base struct {
	indicator.Inter

	mu sync.Mutex

	new_low   float64
	new_high  float64
	last_low  float64
	last_high float64

	state  int
	values []ind

	last tCandle.Inter

	swing    bool
	high_low bool
	closes   bool

	funcHigh funcHighLow
	funcLow  funcHighLow
}

type funcHighLow func(candle tCandle.Inter) float64

func New() *Base {
	pa := &Base{
		Inter: indicator.NewBase(),

		state:  neutral_down,
		values: make([]ind, 0),
	}
	pa.SetCloses(false)

	return pa
}

func (pa *Base) String() string {
	last_ts := int64(0)
	if pa.last != nil {
		last_ts = pa.last.Ts()
	}
	return fmt.Sprintf("PriceAction: c: %d  ..  values: %d  ..  state: %s  ..  high: %f  ..  low: %f  ..  %s",
		pa.Count(), len(pa.values), pa.State(), pa.last_high, pa.last_low, sTime.ForLog(last_ts, 0))
}

func (pa *Base) Log() {
	sLog.Info(pa.String())
}

func (pa *Base) Values() []ind {
	return pa.values
}

func (pa *Base) State() string {
	switch pa.state {
	case downtrend:
		return "downtrend"

	case neutral_down:
		return "neutral down"

	case neutral_up:
		return "neutral up"

	case uptrend:
		return "uptrend"

	default:
		return "unknown"
	}
}

func (pa *Base) IsUptrend() bool {
	return pa.state == uptrend
}

// private types

type ind struct {
	ts    int64
	state int
}

func newInd(ts int64, state int) ind {
	return ind{ts: ts, state: state}
}

func (i *ind) Ts() int64 {
	return i.ts
}

func (i *ind) State() int {
	return i.state
}
