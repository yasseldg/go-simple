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

func (e *Base) GetSymbols(symbols ...string) (tSymbol.Inters, error) {
	if len(symbols) == 0 {
		return nil, ErrEmptySymbols
	}

	var err error = nil
	var errs string
	var inters tSymbol.Inters

	for _, name := range symbols {
		symbol := tSymbol.New(e.Name(), name)
		if !symbol.IsValid() {
			errs = fmt.Sprintf(" %s %s, ", errs, name)
			continue
		}

		inters = append(inters, symbol)
	}

	if len(errs) > 0 {
		err = fmt.Errorf("%s: %s", tSymbol.ErrInvalidSymbol, errs)
	}

	if len(inters) > 0 {
		return inters, err
	}

	return nil, err
}
