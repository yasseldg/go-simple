package sInts

import "reflect"

type Json struct {
	From   int64   `json:"from"`
	To     int64   `json:"to"`
	Step   int64   `json:"step"`
	Values []int64 `json:"values"`
}

func (jn Json) IsZero() bool {
	return reflect.ValueOf(jn).IsZero()
}

func (jn Json) GetIter() *Iter {
	return NewIter(jn.From, jn.To, jn.Step, jn.Values...)
}
