package tSymbol

import (
	"fmt"
	"strings"

	"github.com/yasseldg/go-simple/types/sStrings"
)

type Group struct {
	Common `bson:",inline"`

	M_symbols []string `bson:"syms" json:"syms"`
}

func NewGroup(name, exchange string, symbols ...string) (*Group, error) {

	if len(name) == 0 {
		return nil, fmt.Errorf("name is empty")
	}

	if len(exchange) == 0 {
		return nil, fmt.Errorf("exchange is empty")
	}

	return &Group{
		Common: Common{
			M_name:     strings.ToUpper(name),
			M_exchange: strings.ToUpper(exchange),
		},
		M_symbols: sStrings.ToUpper(symbols),
	}, nil
}

func (g *Group) Names() []string {
	return g.M_symbols
}
