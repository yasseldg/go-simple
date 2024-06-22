package easyway

import (
	"fmt"

	"github.com/yasseldg/go-simple/logs/sLog"
)

type InterEwType interface {
	String() string
	Log()

	Buy() InterEwValues
	Sell() InterEwValues

	InterValueFuncs
}

type InterValueFuncs interface {
	BuyHead() float64
	BuyTail() float64
	BuyMiddle() float64

	SellHead() float64
	SellTail() float64
	SellMiddle() float64
}

type EwType struct {
	M_buy  EwValues `bson:"b,omitempty" json:"b,omitempty"`
	M_sell EwValues `bson:"s,omitempty" json:"s,omitempty"`
}

func (ew *EwType) String() string {
	return fmt.Sprintf("buy: %s  ..  sell: %s", ew.M_buy.String(), ew.M_sell.String())
}

func (ew *EwType) Log() {
	sLog.Info("Ew: %s", ew.String())
}

func (ew *EwType) Buy() InterEwValues {
	return &ew.M_buy
}

func (ew *EwType) Sell() InterEwValues {
	return &ew.M_sell
}

//  ValueFuncs

func (ew *EwType) BuyHead() float64 {
	return ew.Buy().Head()
}

func (ew *EwType) BuyTail() float64 {
	return ew.Buy().Tail()
}

func (ew *EwType) BuyMiddle() float64 {
	return (ew.BuyHead() + ew.BuyTail()) / 2
}

func (ew *EwType) SellHead() float64 {
	return ew.Sell().Head()
}

func (ew *EwType) SellTail() float64 {
	return ew.Sell().Tail()
}

func (ew *EwType) SellMiddle() float64 {
	return (ew.SellHead() + ew.SellTail()) / 2
}
