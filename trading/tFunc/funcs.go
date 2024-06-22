package tFunc

import (
	"github.com/yasseldg/go-simple/trading/tSide"
	"github.com/yasseldg/go-simple/types/sFloats"
)

// TakeProfitStoploss: returns take_profit and stop_loss prices
func TakeProfitStopLoss(entry_price, tp_perc, sl_perc float64, side tSide.Side) (tp_price, sl_price float64) {

	return TakeProfit(entry_price, tp_perc, side), StopLoss(entry_price, sl_perc, side)
}

func TakeProfit(entry_price, tp_perc float64, side tSide.Side) (take_profit float64) {

	adding := true
	if side == tSide.Sell {
		adding = false
	}

	return sFloats.GetWithPercent(entry_price, tp_perc, adding)
}

func TakeProfitPerc(entry_price, tp_price float64, side tSide.Side) float64 {

	perc := sFloats.GetDiffPercent(entry_price, tp_price)

	if side == tSide.Sell {
		perc *= -1
	}
	return perc
}

func StopLoss(entry_price, sl_perc float64, side tSide.Side) float64 {

	adding := false
	if side == tSide.Sell {
		adding = true
	}

	return sFloats.GetWithPercent(entry_price, sl_perc, adding)
}

func StopLossPerc(entry_price, sl_price float64, side tSide.Side) float64 {

	perc := sFloats.GetDiffPercent(entry_price, sl_price)

	if side == tSide.Sell {
		perc *= -1
	}
	return perc
}
