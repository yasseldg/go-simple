package bollingerBands

import (
	"fmt"
	"sync"

	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/types/sFloats"
)

type Base struct {
	mu sync.Mutex

	deviations float64
	closes     sFloats.InterPeriodValues

	mean float64
	std  float64
}

func New(periods int, deviations float64) *Base {
	return &Base{
		mu: sync.Mutex{},

		deviations: deviations,
		closes:     sFloats.NewPeriodValues(periods),
	}
}

func (bb *Base) String() string {

	mean, upper, lower := bb.get()

	return fmt.Sprintf("deviations: %f  ..  periods: %d  ..  std: %f  ..  upper: %f  ..  mean: %f  ..  lower: %f",
		bb.deviations, bb.closes.Periods(), bb.std, upper, mean, lower)
}

func (bb *Base) Log() {
	bb.mu.Lock()
	defer bb.mu.Unlock()

	sLog.Info("BBands: %s", bb.String())
}

func (bb *Base) Deviations() float64 {
	return bb.deviations
}

func (bb *Base) Periods() int {
	return bb.closes.Periods()
}

func (bb *Base) Filled() bool {
	bb.mu.Lock()
	defer bb.mu.Unlock()

	return bb.closes.Filled()
}

func (bb *Base) Get() (mean, upper, lower float64) {
	bb.mu.Lock()
	defer bb.mu.Unlock()

	return bb.get()
}

func (bb *Base) Calc(deviations float64) (mean, upper, lower float64) {
	bb.mu.Lock()
	defer bb.mu.Unlock()

	return bb.calc(deviations)
}

func (bb *Base) Add(close float64) {
	bb.mu.Lock()
	defer bb.mu.Unlock()

	if close == 0 {
		return
	}

	bb.add(close)
}

func (bb *Base) add(close float64) {
	bb.closes.Add(close)

	if !bb.closes.Filled() {
		return
	}

	bb.mean, bb.std = bb.closes.MeanStdDev()
}

func (bb *Base) get() (mean, upper, lower float64) {
	return bb.calc(bb.deviations)
}

func (bb *Base) calc(deviations float64) (mean, upper, lower float64) {
	if !bb.closes.Filled() {
		return 0, 0, 0
	}

	mean = bb.mean
	upper = bb.mean + (bb.std * deviations)
	lower = bb.mean - (bb.std * deviations)
	return
}
