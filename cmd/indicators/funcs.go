package indicators

import (
	"context"
	"fmt"

	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/repos/rMongo"
	"github.com/yasseldg/go-simple/trading/tCandle"
	"github.com/yasseldg/go-simple/trading/tIndicator"
	"github.com/yasseldg/go-simple/trading/tInterval"
)

func Run(mongo rMongo.Inter) {
	indicator := get("SuperTrend")

	// return

	run(indicator, mongo, "BYBIT_BTCUSDT", tInterval.Interval_4h)
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
	_coll rMongo.InterColl

	_iter tCandle.InterIter
)

type Indicator interface {
	Log()
	Add(candle tCandle.Inter)
}

//  private methods

func run(indicator Indicator, mongo rMongo.Inter, symbol string, interval tInterval.Interval) {
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

func config(_mongo rMongo.Inter, symbol string, interval tInterval.Interval) error {

	coll_name := fmt.Sprintf("%s_%s", "historic_prices", interval.String())

	ctx := context.Background()
	var err error
	_coll, err = _mongo.GetColl(ctx, "", "PP_Historic_Trades_R", symbol, coll_name)
	if err != nil {
		return fmt.Errorf("GetColl(): %s", err)
	}
	_coll.Log()

	_iter = tCandle.NewIter(_coll, nil)
	_iter.Log("Candles")

	return nil
}
