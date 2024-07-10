package tSymbol

import (
	"fmt"
	"strings"

	"github.com/yasseldg/go-simple/logs/sLog"
)

type Base struct {
	exchange  string
	name      string
	precision int
}

func New(exchange, name string) *Base {
	return &Base{
		name:     strings.ToUpper(name),
		exchange: exchange,
	}
}

func (s *Base) String() string {
	return fmt.Sprintf("%s_%s", s.exchange, s.name)
}

func (s *Base) Log() {
	sLog.Info("%s .. prec: %d", s.String(), s.precision)
}

func (s *Base) Exchange() string {
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
		exchange:  s.exchange,
		name:      s.name,
		precision: s.precision,
	}
}

func (s *Base) Model() InterModel {
	return NewModel(
		s.exchange,
		s.name,
		s.precision,
	)
}
