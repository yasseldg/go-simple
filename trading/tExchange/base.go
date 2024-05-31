package tExchange

import "strings"

type Base struct {
	name string
}

func New(name string) *Base {
	return &Base{
		name: strings.ToUpper(name),
	}
}

func (e *Base) Name() string {
	return e.name
}

func (e *Base) IsValid() bool {
	// TODO: implement
	return e.name != ""
}

func (e *Base) Clone() Inter {
	return New(e.name)
}

func (e *Base) Model() InterModel {
	return NewModel(e.name)
}
