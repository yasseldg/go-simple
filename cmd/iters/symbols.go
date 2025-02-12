package iters

import (
	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/trading/tExchange"
	"github.com/yasseldg/go-simple/trading/tInterval"
	"github.com/yasseldg/go-simple/trading/tSymbol"
)

func Symbols() {

	exchange := tExchange.New("Bybit")

	iter, err := exchange.GetSymbols("BTCUSD", "ETHUSD", "XRPUSD", "EOSUSD", "LTCUSD")
	if err != nil {
		sLog.Error("exchange.GetSymbols(): %s", err)
	}

	println()
	runSymbols(iter, "Original")

	iter2 := iter.Clone()

	iter2.Reset()
	s, err := tSymbol.New("Bybit", "BNBUSD")
	if err != nil {
		sLog.Error("tSymbol.New(): %s", err)
	} else {
		iter2.Add(s)
	}

	for iter2.Next() {
		iter2.Item().SetPrecision(6)
	}

	iter2.Reset()
	runSymbols(iter2, "Clone")

	iter.Reset()
	runSymbols(iter, "Original")

	iter.Reset()
	for iter.Next() {
		iter.Item().SetPrecision(2)
	}

	iter.Reset()
	runSymbols(iter, "Original")

	iter2.Reset()
	runSymbols(iter2, "Clone")
}

func runSymbols(iter tSymbol.InterIterLimited, name string) {
	iter.Log(name)
	for iter.Next() {
		iter.Item().Log()
	}
	println()
}

func runIntervals(iter tInterval.InterIterLimited, name string) {
	iter.Log(name)
	for iter.Next() {
		sLog.Info(iter.Item().String())
	}
	println()
}
