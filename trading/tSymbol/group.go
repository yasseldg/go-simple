package tSymbol

import (
	"github.com/yasseldg/go-simple/types/sStrings"
)

type Group struct {
	Common `bson:",inline"`

	M_symbols []string `bson:"syms" json:"syms"`
}

func NewGroup(name, exchange string, symbols ...string) (*Group, error) {

	common, err := newCommon(name, exchange)
	if err != nil {
		return nil, err
	}

	return &Group{
		Common:    *common,
		M_symbols: sStrings.ToUpper(symbols),
	}, nil
}

func (g *Group) Names() []string {
	return g.M_symbols
}
