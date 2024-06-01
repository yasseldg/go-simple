package tSymbol

import (
	"fmt"

	"github.com/yasseldg/go-simple/trading/tExchange"
)

func Get(exchange tExchange.Inter, symbols ...string) (Inters, error) {
	if len(symbols) == 0 {
		return nil, ErrEmptySymbols
	}

	var err error
	var errs string
	var inters Inters

	for _, name := range symbols {
		symbol := New(exchange, name)
		if !symbol.IsValid() {
			errs = fmt.Sprintf(" %s %s, ", errs, name)
			continue
		}

		inters = append(inters, symbol)
	}

	if len(errs) > 0 {
		err = fmt.Errorf("%s: %s", ErrInvalidSymbol, errs)
	} else {
		err = nil
	}

	if len(inters) > 0 {
		return inters, err
	}

	return nil, err
}
