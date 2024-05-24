package tCandle

import (
	"fmt"
	"time"

	"github.com/yasseldg/go-simple/logs/sLog"
)

type FanOut struct {
	in   <-chan Inter
	outs map[string]chan<- Inter

	active bool
	done   chan bool
}
type FanOuts []*FanOut

func NewFanOut(in <-chan Inter) FanOut {
	return FanOut{
		in:   in,
		outs: make(map[string]chan<- Inter),
		done: make(chan bool),
	}
}

func (f *FanOut) Log() {
	sLog.Info("FanOut: Log: %d\n", len(f.outs))
	for k := range f.outs {
		sLog.Info("FanOut: Log: %s: %v\n", k)
	}
}

func (f *FanOut) Add(topic string, out chan<- Inter) error {
	if _, ok := f.outs[topic]; ok {
		return fmt.Errorf("topic %s already exists", topic)
	}

	f.outs[topic] = out
	return nil
}

func (f *FanOut) Remove(topic string) error {
	if _, ok := f.outs[topic]; !ok {
		return fmt.Errorf("topic %s does not exist", topic)
	}

	if len(f.outs) > 1 {
		delete(f.outs, topic)
		return nil
	}

	f.Stop()
	return nil
}

func (f *FanOut) Start() {
	if f.active {
		return
	}

	f.active = true

	go func() {
		select {
		case <-f.done:
			return

		default:
			for msg := range f.in {
				for _, out := range f.outs {
					out <- msg
				}
			}
		}
	}()
}

func (f *FanOut) Stop() {
	f.done <- true

	time.Sleep(time.Second)

	for _, out := range f.outs {
		close(out)
	}

	close(f.done)

	f.active = false
}

func (f *FanOut) Outs() int {
	return len(f.outs)
}
