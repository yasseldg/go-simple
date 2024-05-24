package sFloats

import (
	"sync"

	"github.com/yasseldg/go-simple/logs/sLog"

	"gonum.org/v1/gonum/stat"
)

// ---- PeriodValues

type InterPeriodValues interface {
	Log(qty int)

	Filled() bool
	ToFill() int
	Values() []float64
	Periods() int

	Add(value float64)

	Fill(value float64) bool
	// DeepCopy() InterPeriodValues
	Mean() float64
	MeanStdDev() (mean, std float64)
	Sum() float64
}

type PeriodValues struct {
	mu sync.Mutex

	periods int
	values  []float64

	filled bool
}

func NewPeriodValues(periods int) *PeriodValues {
	return &PeriodValues{
		mu: sync.Mutex{},

		periods: periods,
		values:  make([]float64, 0, periods),
	}
}

func (pv *PeriodValues) Periods() int {
	return pv.periods
}

func (pv *PeriodValues) Values() []float64 {
	pv.mu.Lock()
	defer pv.mu.Unlock()

	return pv.values
}

func (pv *PeriodValues) Filled() bool {
	pv.mu.Lock()
	defer pv.mu.Unlock()

	return pv.filled
}

func (pv *PeriodValues) ToFill() int {
	pv.mu.Lock()
	defer pv.mu.Unlock()

	return pv.periods - len(pv.values)
}

func (pv *PeriodValues) Log(qty int) {
	pv.mu.Lock()
	defer pv.mu.Unlock()

	if len(pv.values) < 1 {
		sLog.Info("PeriodValues.Log: periods: %-5d  ..  len: %-5d", pv.periods, len(pv.values))
		return
	}
	if len(pv.values) < qty {
		sLog.Info("PeriodValues.Log: periods: %-5d  ..  len: %-5d  ..  values: %v ", pv.periods, len(pv.values), pv.values)
		return
	}
	sLog.Info("PeriodValues.Log: periods: %-5d  ..  len: %-5d  ..  firsts: %v  ..  lasts: %v", pv.periods, len(pv.values), pv.values[:3], pv.values[len(pv.values)-3:])
}

func (pv *PeriodValues) Add(value float64) {
	pv.mu.Lock()
	defer pv.mu.Unlock()

	if !pv.filled {
		pv.add(value)
		return
	}

	pv.values = append(pv.values[1:], value)
}

func (pv *PeriodValues) add(value float64) {
	pv.values = append(pv.values, value)

	if len(pv.values) >= pv.periods {
		pv.filled = true
	}
}

// Fill adds a value to the PeriodValues if it is not filled yet
// and returns true. If the PeriodValues is already filled, it returns
// false. Fill inverse order of Add.
func (pv *PeriodValues) Fill(value float64) bool {
	pv.mu.Lock()
	defer pv.mu.Unlock()

	return pv.fill(value)
}

func (pv *PeriodValues) fill(value float64) bool {
	if pv.filled {
		return false
	}

	if len(pv.values) >= pv.periods {
		pv.filled = true
		return false
	}

	pv.values = append([]float64{value}, pv.values...)

	return true
}

func (pv *PeriodValues) DeepCopy() *PeriodValues {
	pv.mu.Lock()
	defer pv.mu.Unlock()

	return pv.deepCopy()
}

func (pv *PeriodValues) deepCopy() *PeriodValues {
	newValues := make([]float64, len(pv.values))
	copy(newValues, pv.values)

	return &PeriodValues{
		mu:      sync.Mutex{},
		periods: pv.periods, // Los valores de tipo int se copian por valor
		values:  newValues,
	}
}

func (pv *PeriodValues) Mean() float64 {
	pv.mu.Lock()
	defer pv.mu.Unlock()

	if pv.filled {
		return stat.Mean(pv.values, nil)
	}
	return 0
}

func (pv *PeriodValues) MeanStdDev() (mean, std float64) {
	pv.mu.Lock()
	defer pv.mu.Unlock()

	if pv.filled {
		return stat.MeanStdDev(pv.values, nil)
	}
	return 0, 0
}

func (pv *PeriodValues) Sum() float64 {
	sum := 0.0
	for _, v := range pv.values {
		sum += v
	}
	return sum
}
