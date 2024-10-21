package priceAction

import (
	"fmt"
	"sync"

	"github.com/yasseldg/go-simple/trading/tIndicator/indicator"

	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/trading/tCandle"
	"github.com/yasseldg/go-simple/types/sTime"
)

type Base struct {
	indicator.Inter

	mu sync.Mutex

	new_low   float64
	new_high  float64
	last_low  float64
	last_high float64

	state   state
	trigger bool
	values  []ind

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

		state:  neutral_up,
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
		pa.Count(), len(pa.values), pa.state.String(), pa.last_high, pa.last_low, sTime.ForLog(last_ts, 0))
}

func (pa *Base) Log() {
	sLog.Info(pa.String())
}

func (pa *Base) Values() []ind {
	return pa.values
}

func (pa *Base) State() state {
	return pa.state
}

func (pa *Base) IsTrigger() bool {
	return pa.trigger
}
