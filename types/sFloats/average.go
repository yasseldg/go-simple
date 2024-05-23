package sFloats

import (
	"fmt"
	"sync"
)

type InterAverage interface {
	String() string

	Add(float64)
	Calc() float64
	Value() float64
	Qty() int
	Reset()
}

type Average struct {
	mu sync.Mutex

	qty   int
	value float64
}

func NewAverage() *Average {
	return &Average{
		mu: sync.Mutex{},
	}
}

func (a *Average) String() string {
	a.mu.Lock()
	defer a.mu.Unlock()

	return fmt.Sprintf("Average ( %f ): qty: %d  ..  value: %f", a.calc(), a.qty, a.value)
}

func (a *Average) calc() float64 {
	return GetValid(a.value / float64(a.qty))
}

func (a *Average) Add(v float64) {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.qty++
	a.value += v
}

func (a *Average) Calc() float64 {
	a.mu.Lock()
	defer a.mu.Unlock()

	return a.calc()
}

func (a *Average) Value() float64 {
	a.mu.Lock()
	defer a.mu.Unlock()

	return a.value
}

func (a *Average) Qty() int {
	a.mu.Lock()
	defer a.mu.Unlock()

	return a.qty
}

func (a *Average) Reset() {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.qty = 0
	a.value = 0
}
