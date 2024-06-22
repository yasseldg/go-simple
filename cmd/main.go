package main

import (
	"github.com/yasseldg/go-simple/cmd/components"
	"github.com/yasseldg/go-simple/cmd/indicators"
	"github.com/yasseldg/go-simple/cmd/iters"
	"github.com/yasseldg/go-simple/cmd/repos"
	"github.com/yasseldg/go-simple/trading/tFunc"
	"github.com/yasseldg/go-simple/trading/tSide"

	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/repositorys/rMongo"
	"github.com/yasseldg/go-simple/types/sTime"
)

var (
	_mongo *rMongo.Manager
)

func Init() {
	mongo := rMongo.NewManager()
	_mongo = &mongo
}

func main() {

	clean := sLog.SetByName(sLog.Zap, sLog.LevelInfo, "")
	defer clean()

	Init()

	sLog.Info("Starting...")

	sTime.TimeControl(iters.Tests, "Iters")
}

func testModel() {
	repos.Run(_mongo)
}

func testIndicators() {
	indicators.Run(_mongo)
}

func testRatios() {

	high := 10245.0        //110.0      // 73167.0
	entry_price := 10069.0 // 100.0		// 73061.0
	low := 9952.0          // 95.0      // 72494.0

	println()

	side := tSide.Buy
	tp_price := high
	sl_price := low

	ratio := tFunc.Ratio(entry_price, sl_price, tp_price)

	sLog.Info("Ratio: %f = %f  ..  entry: %f ..  sl: %f  ..  tp: %f  ..  side: %s ", ratio, (tp_price / sl_price), entry_price, sl_price, tp_price, side)

	entry := tFunc.EntryByRatio(ratio, sl_price, tp_price)

	sLog.Info("Ratio: %f  ..  entry: %f ..  sl: %f  ..  tp: %f  ..  side: %s ", ratio, entry, sl_price, tp_price, side)

	take_profit := tFunc.TakeProfitByRatio(entry_price, sl_price, ratio)
	stop_loss := tFunc.StopLossByRatio(entry_price, tp_price, ratio)

	sLog.Info("Ratio: %f  ..  entry: %f ..  sl: %f  ..  tp: %f  ..  side: %s ", ratio, entry, stop_loss, take_profit, side)

	println()

	side = tSide.Sell
	sl_price = high
	tp_price = low

	ratio = tFunc.Ratio(entry_price, sl_price, tp_price)

	sLog.Info("Ratio: %f = %f  ..  entry: %f ..  sl: %f  ..  tp: %f  ..  side: %s ", ratio, (tp_price / sl_price), entry_price, sl_price, tp_price, side)

	entry = tFunc.EntryByRatio(ratio, sl_price, tp_price)

	sLog.Info("Ratio: %f  ..  entry: %f ..  sl: %f  ..  tp: %f  ..  side: %s ", ratio, entry, sl_price, tp_price, side)

	take_profit = tFunc.TakeProfitByRatio(entry_price, sl_price, ratio)
	stop_loss = tFunc.StopLossByRatio(entry_price, tp_price, ratio)

	sLog.Info("Ratio: %f  ..  entry: %f ..  sl: %f  ..  tp: %f  ..  side: %s ", ratio, entry, stop_loss, take_profit, side)

	println()

	risk := 0.784
	reward := 0.149

	sLog.Info("Risk: %f, reward: %f, ratio: %f", risk, reward, tFunc.RatioFromRiskReward(risk, reward))

	sLog.Info("Risk: %f, reward: %f, ratio: %f", reward, risk, tFunc.RatioFromRiskReward(reward, risk))
}

func testComponents() {
	components.Run(_mongo)
}
