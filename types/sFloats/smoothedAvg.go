package sFloats

import (
	"fmt"
	"sync"

	"github.com/yasseldg/go-simple/logs/sLog"
)

type SmoothedAverage struct {
	mu sync.Mutex

	period int
	value  float64

	filled int
}

func NewSmoothedAverage(period int) *SmoothedAverage {
	return &SmoothedAverage{
		mu: sync.Mutex{},

		period: period,
		filled: period,
	}
}

func (sa *SmoothedAverage) String() string {
	return fmt.Sprintf("period: %d  ..  value: %f  ..  filled: %d", sa.period, sa.value, sa.filled)
}

func (sa *SmoothedAverage) Log() {
	sa.mu.Lock()
	defer sa.mu.Unlock()

	sLog.Info("SmoothedAverage: %s", sa.String())
}

// AddPos adds a value to the average
func (sa *SmoothedAverage) AddPos(v float64) {
	sa.mu.Lock()
	defer sa.mu.Unlock()

	if sa.filled > 0 {
		sa.fillPos(v)
		return
	}

	sa.value *= float64(sa.period - 1)
	sa.value += v
	sa.value /= float64(sa.period)
}

// AddNeg subtracts a value from the average
func (sa *SmoothedAverage) AddNeg(v float64) {
	sa.mu.Lock()
	defer sa.mu.Unlock()

	if sa.filled > 0 {
		sa.fillNeg(v)
		return
	}

	sa.value *= float64(sa.period - 1)
	sa.value -= v
	sa.value /= float64(sa.period)
}

func (sa *SmoothedAverage) fillPos(v float64) {
	sa.value += v
	sa.fill()
}

func (sa *SmoothedAverage) fillNeg(v float64) {
	sa.value -= v
	sa.fill()
}

func (sa *SmoothedAverage) fill() {
	sa.filled--

	if sa.filled == 0 {
		sa.value /= float64(sa.period)
	}
}

func (sa *SmoothedAverage) Filled() bool {
	sa.mu.Lock()
	defer sa.mu.Unlock()

	return sa.filled == 0
}

func (sa *SmoothedAverage) Value() float64 {
	sa.mu.Lock()
	defer sa.mu.Unlock()

	if sa.filled > 0 {
		return 0
	}

	return sa.value
}

func (sa *SmoothedAverage) Period() int {
	sa.mu.Lock()
	defer sa.mu.Unlock()

	return sa.period
}
