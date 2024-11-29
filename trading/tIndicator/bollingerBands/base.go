package bollingerBands

import (
	"fmt"
	"sync"

	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/types/sFloats"
)

type base struct {
	mu sync.Mutex

	deviations float64
	closes     sFloats.InterPeriodValues

	mean float64
	std  float64
}

func New(periods int, deviations float64) *base {
	return &base{
		mu: sync.Mutex{},

		deviations: deviations,
		closes:     sFloats.NewPeriodValues(periods),
	}
}

func (bb *base) String() string {

	mean, upper, lower := bb.get()

	return fmt.Sprintf("deviations: %f  ..  periods: %d  ..  std: %f  ..  upper: %f  ..  mean: %f  ..  lower: %f",
		bb.deviations, bb.closes.Periods(), bb.std, upper, mean, lower)
}

func (bb *base) Log() {
	bb.mu.Lock()
	defer bb.mu.Unlock()

	sLog.Info("BBands: %s", bb.String())
}

func (bb *base) Deviations() float64 {
	return bb.deviations
}

func (bb *base) Periods() int {
	return bb.closes.Periods()
}

func (bb *base) Filled() bool {
	bb.mu.Lock()
	defer bb.mu.Unlock()

	return bb.closes.Filled()
}

func (bb *base) Get() (mean, upper, lower float64) {
	bb.mu.Lock()
	defer bb.mu.Unlock()

	return bb.get()
}

func (bb *base) Calc(deviations float64) (mean, upper, lower float64) {
	bb.mu.Lock()
	defer bb.mu.Unlock()

	return bb.calc(deviations)
}

func (bb *base) Add(close float64) {
	bb.mu.Lock()
	defer bb.mu.Unlock()

	if close == 0 {
		return
	}

	bb.add(close)
}

func (bb *base) add(close float64) {
	bb.closes.Add(close)

	if !bb.closes.Filled() {
		return
	}

	bb.mean, bb.std = bb.closes.MeanStdDev()
}

func (bb *base) get() (mean, upper, lower float64) {
	return bb.calc(bb.deviations)
}

func (bb *base) calc(deviations float64) (mean, upper, lower float64) {
	if !bb.closes.Filled() {
		return 0, 0, 0
	}

	mean = bb.mean
	upper = bb.mean + (bb.std * deviations)
	lower = bb.mean - (bb.std * deviations)
	return
}
