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

	if tp_perc == 0 {
		return 0
	}

	adding := true
	if side.IsSell() {
		adding = false
	}

	return sFloats.GetWithPercent(entry_price, tp_perc, adding)
}

func TakeProfitPerc(entry_price, tp_price float64, side tSide.Side) float64 {

	perc := sFloats.GetDiffPercent(entry_price, tp_price)

	if side.IsSell() {
		perc *= -1
	}
	return perc
}

func StopLoss(entry_price, sl_perc float64, side tSide.Side) float64 {

	if sl_perc == 0 {
		return 0
	}

	adding := false
	if side.IsSell() {
		adding = true
	}

	return sFloats.GetWithPercent(entry_price, sl_perc, adding)
}

func StopLossPerc(entry_price, sl_price float64, side tSide.Side) float64 {

	perc := sFloats.GetDiffPercent(entry_price, sl_price)

	if side.IsSell() {
		perc *= -1
	}
	return perc
}

// func (bt *Backtest) Roe(side sTrading.Side, triggers_Coll sMongo.CollManager) float64 {
// 	triggers := bt.Triggers(side, triggers_Coll)

// 	roe := 0.0
// 	for _, trigger := range triggers {
// 		r := trigger.Roe()
// 		sLog.Debug("Backtest.Roe: %s  ..  entry: %.2f  ..  exit: %.2f  ..  roe: %3f", sDate.ForLog(trigger.UnixTs, 0), trigger.Entry.Price, trigger.Close.Price, r)
// 		roe += r
// 	}
// 	return roe * 100
// }

// func (t Trigger) Roe() float64 {
// 	switch t.State {
// 	case sTrading.State_Win:
// 		return (t.TakeProfit.Price - t.Entry.Price) / t.Entry.Price

// 	case sTrading.State_Loss:
// 		return (t.StopLoss.Price - t.Entry.Price) / t.Entry.Price

// 	default:
// 		return 0
// 	}
// }

// func (c *Calcs) update(win, loss int, takeProfit, stopLoss, roe float64) {
// 	total := win + loss
// 	c.WinRate = float64(win) / float64(total)
// 	c.MathHope = (c.WinRate * (takeProfit / stopLoss)) - (float64(loss) / float64(total))
// 	c.Roe = roe
// }
