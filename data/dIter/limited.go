package dIter

import (
	"fmt"
	"sync"

	"github.com/yasseldg/go-simple/logs/sLog"
)

type Limited[T any] struct {
	Inter

	mu sync.Mutex

	index int
	items []T

	cloneFunc func(T) T
}

func NewLimited[T any](cloneFunc func(T) T) *Limited[T] {
	return &Limited[T]{
		Inter: New(),

		cloneFunc: cloneFunc,
		index:     -1,
	}
}

func (i *Limited[T]) String(name string) string {
	return fmt.Sprintf("%s %d  ..  index: %d", i.Inter.String(name), i.Count(), i.index)
}

func (i *Limited[T]) Log(name string) {
	sLog.Info(i.String(name))
}

func (i *Limited[T]) Add(items ...T) {
	i.mu.Lock()
	defer i.mu.Unlock()

	i.items = append(i.items, items...)
}

func (i *Limited[T]) Reset() {
	i.mu.Lock()
	defer i.mu.Unlock()

	i.index = -1
}

func (i *Limited[T]) Next() bool {
	i.mu.Lock()
	defer i.mu.Unlock()

	if i.index >= len(i.items)-1 {
		return false
	}

	i.index++
	return true
}

func (i *Limited[T]) Item() T {
	i.mu.Lock()
	defer i.mu.Unlock()

	return i.items[i.index]
}

func (i *Limited[T]) Count() int {
	i.mu.Lock()
	defer i.mu.Unlock()

	return len(i.items)
}

func (i *Limited[T]) Clone() InterLimited[T] {
	i.mu.Lock()
	defer i.mu.Unlock()

	clone := NewLimited[T](i.cloneFunc)
	clone.index = i.index

	if i.cloneFunc == nil {
		i.cloneFunc = func(item T) T { return item }
	}

	for _, item := range i.items {
		clone.Add(i.cloneFunc(item))
	}

	return clone
}
