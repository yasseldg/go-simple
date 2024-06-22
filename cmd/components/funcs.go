package components

import (
	"time"

	"github.com/yasseldg/go-simple/components/easyway"
	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/repositorys/rMongo"
	"github.com/yasseldg/go-simple/trading/tInterval"
)

func Run(_mongo *rMongo.Manager) {

	coll, err := _mongo.GetColl("", "PP_Data_Center", "BYBIT_BTCUSDT", "strategies_5") //  PT_Dev_Data_Center	PP_Data_Center
	if err != nil {
		sLog.Error("Failed to get collection: %s", err)
		return
	}
	coll.Log()

	repo(coll)
}

func mix(coll rMongo.Collection) {

	ew_type := "t5000"

	iter := easyway.NewIter(coll, ew_type)
	if iter == nil {
		sLog.Error("Failed to create iter")
		return
	}

	iter.SetTsFrom(1696490400)
	iter.SetTsTo(1696582800)

	repo := easyway.NewRepo(coll, ew_type)
	if repo == nil {
		sLog.Error("Failed to create repo")
		return
	}

	for iter.Next() {
		err := repo.FindTs(iter.Item().Ts())
		if err != nil {
			sLog.Error("Failed to find ts: %s", err)
			return
		}

		sLog.Info("%s  ..  %s", iter.Item().String(ew_type), repo.Ew().String())
	}
}

func iter(coll rMongo.Collection) {

	ew_type := "t5000"

	iter := easyway.NewIter(coll, ew_type)
	if iter == nil {
		sLog.Error("Failed to create iter")
		return
	}

	iter.SetTsFrom(1696490400)
	iter.SetTsTo(1696582800)

	for iter.Next() {
		iter.Item().Log(ew_type)
	}
}

func repo(coll rMongo.Collection) {

	ew_type := "t5000"

	repo := easyway.NewRepo(coll, ew_type)
	if repo == nil {
		sLog.Error("Failed to create repo")
		return
	}

	ts := tInterval.Interval_5m.Prev(time.Now().Unix())
	ts = tInterval.Interval_5m.Prev(ts - 1)

	err := repo.FindTs(ts)
	if err != nil {
		sLog.Error("Failed to find ts( %d ): %s", ts, err)
		return
	}

	repo.Ew().Log()
}
