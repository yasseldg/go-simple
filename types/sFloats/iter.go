package sFloats

import (
	"fmt"

	"github.com/yasseldg/go-simple/data/dIter"
	"github.com/yasseldg/go-simple/logs/sLog"
)

type Iter interface {
	dIter.Inter

	Value() float64
	Count() int
	Reset()
}

type BaseIter struct {
	dIter.Inter

	from       float64
	to         float64
	step       float64
	values     []float64
	values_map map[float64]bool

	index   int
	current float64
}

func NewBaseIter(from, to, step float64, values ...float64) *BaseIter {
	return &BaseIter{
		Inter: dIter.New(),

		from:       from,
		to:         to,
		step:       step,
		values:     values,
		values_map: make(map[float64]bool),

		index: -1,
	}
}

func (b *BaseIter) Reset() {
	b.index = -1
	b.current = 0
	b.values_map = make(map[float64]bool)
}

func (b *BaseIter) String(name string) string {
	return fmt.Sprintf("%s %d  ..  value: %f", b.Inter.String(name), b.Count(), b.current)
}

func (b *BaseIter) Log(name string) {
	sLog.Info(b.String(name))
}

func (b *BaseIter) Value() float64 {
	return b.current
}

func (b *BaseIter) Count() int {
	return b.index + 1
}

func (b *BaseIter) Next() bool {

	if b.index < len(b.values) {
		if b.nextValue() {
			return true
		}

		if b.from == 0 && b.to == 0 {
			return false
		}

		b.current = b.from

		return b.verify()
	}

	return b.nextRange()
}

// private methods

func (b *BaseIter) nextValue() bool {

	b.index++

	if b.index >= len(b.values) {
		return false
	}

	b.current = b.values[b.index]

	b.values_map[b.current] = true

	return true
}

func (b *BaseIter) nextRange() bool {

	nextValue := b.current + b.step
	if (b.step > 0 && CompareTruncFloat(nextValue, b.to, 3) <= 0) ||
		(b.step < 0 && CompareTruncFloat(nextValue, b.to, 3) >= 0) {

		b.index++
		b.current = nextValue

		return b.verify()
	}

	if b.step == 0 && b.current == b.from {

		b.index++
		b.current = b.to
		return b.verify()
	}

	return false
}

func (b *BaseIter) verify() bool {
	if _, ok := b.values_map[b.current]; ok {
		return b.nextRange()
	}

	return true
}
