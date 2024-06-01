package indicators

import (
	"fmt"

	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/repositorys/rMongo"
	"github.com/yasseldg/go-simple/trading/tCandle"
	"github.com/yasseldg/go-simple/trading/tIndicator"
	"github.com/yasseldg/go-simple/trading/tInterval"
)

func Run(mongo *rMongo.Manager) {
	indicator := get("SuperTrendIter")

	run(indicator, mongo, "BYBIT_BTCUSDT", tInterval.Interval_D)
}

func get(indicator string) Indicator {
	switch indicator {
	case "RSI":
		return tIndicator.NewRSIcandle(14)

	case "BBands":
		return tIndicator.NewBBcandle(30, 2)

	case "AvgATR":
		return tIndicator.NewAvgATR(10)

	case "SmATR":
		return tIndicator.NewSmATR(14)

	case "ADX":
		return tIndicator.NewADX(14)

	case "SuperTrend":
		return tIndicator.NewSuperTrend(10, 3, false)

	case "SuperTrendIter":
		testBBandSuperTrendIter()
		return nil

	default:
		return nil
	}
}

// private vars

var (
	_coll rMongo.Collection

	_iter tCandle.InterIter
)

type Indicator interface {
	Log()
	Add(candle tCandle.Inter)
}

//  private methods

func run(indicator Indicator, mongo *rMongo.Manager, symbol string, interval tInterval.Interval) {
	if indicator == nil {
		sLog.Info("Indicator is nil")
		return
	}

	err := config(mongo, symbol, interval)
	if err != nil {
		sLog.Info("repos.Run(): config(): %s", err)
		return
	}

	indicator.Log()

	for _iter.Next() {
		indicator.Add(_iter.Item())

		indicator.Log()
	}
}

func config(_mongo *rMongo.Manager, symbol string, interval tInterval.Interval) error {

	coll_name := fmt.Sprintf("%s_%s", "historic_prices", interval.String())

	var err error
	_coll, err = _mongo.GetColl("", "PP_Historic_Trades_R", symbol, coll_name)
	if err != nil {
		return fmt.Errorf("GetColl(): %s", err)
	}
	_coll.Log()

	_iter, err = tCandle.NewIter(rMongo.NewFilter(), _coll)
	if err != nil {
		return fmt.Errorf("tCandle.NewIter(): %s", err)
	}
	_iter.Log("Candles")

	return nil
}
