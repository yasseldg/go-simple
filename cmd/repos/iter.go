package repos

import (
	"context"

	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/repos/rMongo"
	"github.com/yasseldg/go-simple/trading/tCandle"
)

func RunIter(mongo rMongo.Inter) {
	mongo.SetDebug(true)

	ctx := context.Background()

	var err error
	_coll, err = mongo.GetColl(ctx, "", "PP_Bactests_R", "BYBIT_BTCUSDT", "candles_15m")
	if err != nil {
		sLog.Error("GetColl(): %s", err)
		return
	}
	_coll.Log()

	iter := tCandle.NewIter(_coll, nil)

	iter.SetTsFrom(1725475000)
	iter.SetTsTo(1725975000)

	iter.Log("Candle")

	for iter.Next() {
		iter.Item().Log(2)
	}

	if iter.Error() != nil {
		sLog.Error("RunIter(): %s", iter.Error())
	}
}
