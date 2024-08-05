package iters

import (
	"github.com/yasseldg/go-simple/trading/tSymbol"
)

func Symbols() {
	iter := tSymbol.NewIterLimited()

	iter.Add(tSymbol.New("Bybit", "BTCUSD"))
	iter.Add(tSymbol.New("Bybit", "ETHUSD"))
	iter.Add(tSymbol.New("Bybit", "XRPUSD"))
	iter.Add(tSymbol.New("Bybit", "EOSUSD"))
	iter.Add(tSymbol.New("Bybit", "LTCUSD"))

	println()
	runSymbols(iter, "Original")

	iter2 := iter.Clone()

	iter2.Reset()
	iter2.Add(tSymbol.New("Bybit", "BNBUSD"))

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
