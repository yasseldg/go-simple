package tExchange

import (
	"github.com/yasseldg/go-simple/repos/rMongo"
	"github.com/yasseldg/go-simple/trading/tSymbol"
)

// Inter

type Inter interface {
	// Name returns the exchange name
	Name() string
	IsValid() bool

	GetSymbols(symbols ...string) (tSymbol.InterIterLimited, error)

	Clone() Inter
	Model() InterModel
}

// InterModel

type InterModel interface {
	rMongo.InterModelDateState

	Name() string
}
