package tSymbol

import (
	"fmt"
	"strings"

	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/trading/tExchange"
)

type Base struct {
	exchange tExchange.Inter

	name      string
	precision int
}

func New(exchange tExchange.Inter, name string) *Base {
	return &Base{
		name:     strings.ToUpper(name),
		exchange: exchange,
	}
}

func (s *Base) String() string {
	return fmt.Sprintf("%s_%s", s.exchange.Name(), s.name)
}

func (s *Base) Log() {
	sLog.Info("%s .. prec: %d", s.String(), s.precision)
}

func (s *Base) Exchange() tExchange.Inter {
	return s.exchange
}

func (s *Base) Name() string {
	return s.name
}

func (s *Base) Precision() int {
	return s.precision
}

func (s *Base) SetPrecision(prec int) {
	s.precision = prec
}

func (s *Base) IsValid() bool {
	// TODO: implement
	return s.name != ""
}

func (s *Base) Clone() Inter {
	return &Base{
		exchange:  s.Exchange().Clone(),
		name:      s.name,
		precision: s.precision,
	}
}

func (s *Base) Model() InterModel {
	return NewModel(
		s.exchange.Model().StringID(),
		s.name,
		s.precision,
	)
}
