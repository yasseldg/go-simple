package tTrade

import "github.com/yasseldg/go-simple/trading/tSide"

type Inter interface {
	Ts() int64
	Oid() string
	Price() float64
	Size() float64
	Side() tSide.Side

	String() string
}
