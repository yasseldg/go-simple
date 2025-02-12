package tExchange

import (
	"fmt"
	"strings"

	"github.com/yasseldg/go-simple/trading/tSymbol"
)

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

func (e *Base) GetSymbols(symbols ...string) (tSymbol.InterIterLimited, error) {
	if len(symbols) == 0 {
		return nil, ErrEmptySymbols
	}

	var err error = nil
	var errs string

	iter := tSymbol.NewIterLimited()

	for _, name := range symbols {
		symbol, err := tSymbol.New(name, e.Name())
		if err != nil {
			errs = fmt.Sprintf(" %s %s, ", errs, name)
			continue
		}

		iter.Add(symbol)
	}

	if len(errs) > 0 {
		err = fmt.Errorf("%s: %s", tSymbol.ErrInvalidSymbol, errs)
	}

	return iter, err
}
