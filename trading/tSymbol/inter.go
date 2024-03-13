package tSymbol

import "github.com/yasseldg/go-simple/trading/tExchange"

type Inter interface {
	Name() string
	IsValid() bool
	String() string
	Exchange() tExchange.Inter
	Precision() int

	SetPrecision(int)
}

type Inters []Inter
