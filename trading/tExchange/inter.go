package tExchange

import (
	"github.com/yasseldg/go-simple/trading/tSymbol"

	"github.com/yasseldg/mgm/v4"
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
	mgm.ModelDateState

	Name() string
}
