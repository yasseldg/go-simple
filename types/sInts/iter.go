package sInts

import (
	"fmt"

	"github.com/yasseldg/go-simple/data/dIter"
	"github.com/yasseldg/go-simple/logs/sLog"
)

type Iter struct {
	dIter.Inter

	from       int64
	to         int64
	step       int64
	values     []int64
	values_map map[int64]bool

	count   int
	index   int
	current int64
}

func NewIter(from, to, step int64, values ...int64) *Iter {

	if len(values) == 0 {
		values = nil
	}

	iter := Iter{
		Inter: dIter.New(),

		from:       from,
		to:         to,
		step:       step,
		values:     values,
		values_map: make(map[int64]bool, len(values)),

		index: -1,
	}

	iter.setCount()

	return &iter
}

func (b *Iter) Reset() {
	b.index = -1
	b.current = 0
	b.values_map = make(map[int64]bool, len(b.values))
}

func (b *Iter) String(name string) string {
	return fmt.Sprintf("%s %d  ..  from: %d  ..  to: %d  ..  step: %d  ..  values: %v  ..  value: %d",
		b.Inter.String(name), b.Count(), b.from, b.to, b.step, b.values, b.current)
}

func (b *Iter) Log(name string) {
	sLog.Info(b.String(name))
}

func (b *Iter) Value() int64 {
	return b.current
}

func (b *Iter) Count() int {
	return b.count
}

func (b *Iter) Next() bool {

	if b.index < len(b.values) {
		if b.nextValue() {
			return true
		}

		if b.from > 0 && b.from == b.to && b.from == b.current {
			return false
		}

		b.current = b.from

		return b.verify()
	}

	return b.nextRange()
}

func (b *Iter) Clone() InterIter {
	return NewIter(b.from, b.to, b.step, b.values...)
}

// private methods

func (b *Iter) setCount() {
	for b.Next() {
		b.count++
	}
	b.Reset()
}

func (b *Iter) nextValue() bool {

	b.index++

	if b.index >= len(b.values) {
		return false
	}

	b.current = b.values[b.index]

	b.values_map[b.current] = true

	return true
}

func (b *Iter) nextRange() bool {

	nextValue := b.current + b.step
	if (b.step > 0 && nextValue <= b.to) || (b.step < 0 && nextValue >= b.to) {

		b.index++
		b.current = nextValue

		return b.verify()
	}

	if b.from == b.to && b.from == b.current {
		return false
	}

	if b.step == 0 && b.from == b.current {

		b.index++
		b.current = b.to
		return b.verify()
	}

	return false
}

func (b *Iter) verify() bool {
	if _, ok := b.values_map[b.current]; ok {
		return b.nextRange()
	}

	return true
}
