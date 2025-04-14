package tCandle

import (
	"fmt"

	"github.com/yasseldg/go-simple/repos/rMongo"
)

func GetCandles(ts_from, ts_to, interval_sec int64, coll rMongo.InterRepo) (Candles, error) {

	var objs Candles

	err := coll.Pipeline(pipCandles(ts_from, ts_to, interval_sec)).Agregates(&objs)
	if err != nil {
		return nil, fmt.Errorf("pipCandles(): %s", err.Error())
	}
	return objs, nil
}

func pipCandles(ts_from, ts_to, interval_sec int64) rMongo.Pipeline {

	m, err := rMongo.FilterFields(rMongo.NewFilter().Ts(ts_from, ts_to))
	if err != nil {
		return nil
	}

	match := rMongo.D{{"$match", m}}
	sort := rMongo.D{{"$sort", rMongo.D{{"ts", 1}}}}
	project := rMongo.D{{"$project", rMongo.D{{"minute", rMongo.D{{"$floor", rMongo.D{{"$divide", rMongo.A{"$ts", interval_sec}}}}}}, {"o", 1}, {"h", 1}, {"l", 1}, {"c", 1}, {"v", 1}}}}
	group := rMongo.D{{"$group", rMongo.D{{"_id", "$minute"}, {"o", rMongo.D{{"$first", "$o"}}}, {"h", rMongo.D{{"$max", "$h"}}}, {"l", rMongo.D{{"$min", "$l"}}}, {"c", rMongo.D{{"$last", "$c"}}}, {"v", rMongo.D{{"$sum", "$v"}}}}}}
	sort_2 := rMongo.D{{"$sort", rMongo.D{{"_id", 1}}}}
	project_2 := rMongo.D{{"$project", rMongo.D{{"ts", rMongo.D{{"$multiply", rMongo.A{"$_id", interval_sec}}}}, {"o", 1}, {"h", 1}, {"l", 1}, {"c", 1}, {"v", 1}, {"_id", 0}}}}

	return rMongo.Pipeline{match, sort, project, group, sort_2, project_2}
}
