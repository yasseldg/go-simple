package tSymbol

import (
	"fmt"
	"strings"

	"github.com/yasseldg/go-simple/trading/tExchange"
)

type Base struct {
	name     string
	exchange tExchange.Inter

	precision int
}

func New(name string, exchange tExchange.Inter) *Base {
	return &Base{
		name:     strings.ToUpper(name),
		exchange: exchange,
	}
}

func (s *Base) Name() string {
	return s.name
}

func (s *Base) IsValid() bool {
	// TODO: implement
	return s.name != ""
}

func (s *Base) Exchange() tExchange.Inter {
	return s.exchange
}

func (s *Base) String() string {
	return fmt.Sprintf("%s_%s", s.exchange.Name(), s.name)
}

func (s *Base) Precision() int {
	return s.precision
}

func (s *Base) SetPrecision(prec int) {
	s.precision = prec
}
