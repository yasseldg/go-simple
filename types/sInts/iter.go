package sInts

import (
	"fmt"

	"github.com/yasseldg/go-simple/data/dIter"
	"github.com/yasseldg/go-simple/logs/sLog"
)

type InterIter interface {
	dIter.Inter

	Value() int64
	Count() int
	Reset()
}

type Iter struct {
	dIter.Inter

	from       int64
	to         int64
	step       int64
	values     []int64
	values_map map[int64]bool

	index   int
	current int64
}

func NewIter(from, to, step int64, values ...int64) *Iter {
	return &Iter{
		Inter: dIter.New(),

		from:       from,
		to:         to,
		step:       step,
		values:     values,
		values_map: make(map[int64]bool),

		index: -1,
	}
}

func (b *Iter) Reset() {
	b.index = -1
	b.current = 0
	b.values_map = make(map[int64]bool)
}

func (b *Iter) String(name string) string {
	return fmt.Sprintf("%s %d  ..  value: %d", b.Inter.String(name), b.Count(), b.current)
}

func (b *Iter) Log(name string) {
	sLog.Info(b.String(name))
}

func (b *Iter) Value() int64 {
	return b.current
}

func (b *Iter) Count() int {
	return b.index + 1
}

func (b *Iter) Next() bool {

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

	if b.step == 0 && b.current == b.from {

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
