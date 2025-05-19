package tTrade

import (
	"fmt"

	"github.com/yasseldg/go-simple/trading/tSide"
	"github.com/yasseldg/go-simple/types/sTime"
)

type base struct {
	ts    int64
	oid   string
	price float64
	size  float64
	side  tSide.Side
}

func New(ts int64, oid string, price, size float64, side tSide.Side) *base {
	return &base{
		ts:    ts,
		oid:   oid,
		price: price,
		size:  size,
		side:  side,
	}
}

func (b *base) Ts() int64 {
	return b.ts
}

func (b *base) Oid() string {
	return b.oid
}

func (b *base) Price() float64 {
	return b.price
}

func (b *base) Size() float64 {
	return b.size
}

func (b *base) Side() tSide.Side {
	return b.side
}

func (b *base) String() string {
	return fmt.Sprintf("%s .. oid: %s .. price: %f .. size: %f .. side: %s",
		sTime.ForLog(b.ts, 0), b.oid, b.price, b.size, b.side)
}
