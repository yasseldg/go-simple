package sStrings

import "reflect"

type Json struct {
	Values []string `json:"values"`
}

func (jn Json) IsZero() bool {
	return reflect.ValueOf(jn).IsZero()
}

func (jn Json) GetIter() InterIterLimited {
	return NewIterLimited(jn.Values...)
}
