package tIndicator

import (
	"fmt"
	"sync"

	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/types/sFloats"
)

//  BBands (Bollinger Bands)

type interBBands interface {
	String() string
	Log()

	Periods() int
	Deviations() float64

	Filled() bool
	Get() (mean, upper, lower float64)
	Calc(deviations float64) (mean, upper, lower float64)
}

type InterBBands interface {
	interBBands

	Add(close float64)
}

type BBands struct {
	mu sync.Mutex

	deviations float64
	closes     sFloats.InterPeriodValues

	mean float64
	std  float64
}

func NewBBands(periods int, deviations float64) *BBands {
	return &BBands{
		mu: sync.Mutex{},

		deviations: deviations,
		closes:     sFloats.NewPeriodValues(periods),
	}
}

func (bb *BBands) String() string {

	mean, upper, lower := bb.get()

	return fmt.Sprintf("deviations: %f  ..  periods: %d  ..  std: %f  ..  upper: %f  ..  mean: %f  ..  lower: %f",
		bb.deviations, bb.closes.Periods(), bb.std, upper, mean, lower)
}

func (bb *BBands) Log() {
	bb.mu.Lock()
	defer bb.mu.Unlock()

	sLog.Info("BBands: %s", bb.String())
}

func (bb *BBands) Deviations() float64 {
	return bb.deviations
}

func (bb *BBands) Periods() int {
	return bb.closes.Periods()
}

func (bb *BBands) Filled() bool {
	bb.mu.Lock()
	defer bb.mu.Unlock()

	return bb.closes.Filled()
}

func (bb *BBands) Get() (mean, upper, lower float64) {
	bb.mu.Lock()
	defer bb.mu.Unlock()

	return bb.get()
}

func (bb *BBands) Calc(deviations float64) (mean, upper, lower float64) {
	bb.mu.Lock()
	defer bb.mu.Unlock()

	return bb.calc(deviations)
}

func (bb *BBands) Add(close float64) {
	bb.mu.Lock()
	defer bb.mu.Unlock()

	if close == 0 {
		return
	}

	bb.add(close)
}

func (bb *BBands) add(close float64) {
	bb.closes.Add(close)

	if !bb.closes.Filled() {
		return
	}

	bb.mean, bb.std = bb.closes.MeanStdDev()
}

func (bb *BBands) get() (mean, upper, lower float64) {
	return bb.calc(bb.deviations)
}

func (bb *BBands) calc(deviations float64) (mean, upper, lower float64) {
	if !bb.closes.Filled() {
		return 0, 0, 0
	}

	mean = bb.mean
	upper = bb.mean + (bb.std * deviations)
	lower = bb.mean - (bb.std * deviations)
	return
}
