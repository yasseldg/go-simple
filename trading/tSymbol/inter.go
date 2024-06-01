package tSymbol

import (
	"github.com/yasseldg/go-simple/trading/tExchange"

	"github.com/yasseldg/mgm/v4"
)

type Inter interface {
	String() string
	Log()

	Exchange() tExchange.Inter

	Name() string
	Precision() int
	SetPrecision(int)

	IsValid() bool
	Clone() Inter
	Model() InterModel
}
type Inters []Inter

// InterModel

type InterModel interface {
	mgm.ModelDateState

	ExchangeID() string
	Name() string
	Precision() int
}
