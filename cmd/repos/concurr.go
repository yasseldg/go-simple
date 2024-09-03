package repos

import (
	"fmt"
	"sync"
	"time"

	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/repos/rMongo"
)

func RunConcurr(mongo rMongo.Inter) {

	mongo.SetDebug(true)

	err := config(mongo)
	if err != nil {
		sLog.Error("repos.Run(): config(): %s", err)
		return
	}

	wg := sync.WaitGroup{}

	wg.Add(1)

	go func() {
		defer wg.Done()
		filter := NewFilter().Uuid("uuid_6")

		doc, err := findFilter(filter, 6)
		if err != nil {
			sLog.Error("find(): %s", err)
		}

		if doc != nil {
			doc.Log()
		}
	}()

	time.Sleep(1 * time.Second)

	wg.Add(1)

	go func() {
		defer wg.Done()
		filter := NewFilter().Uuid("uuid_4")

		doc, err := findFilter(filter, 2)
		if err != nil {
			sLog.Error("find(): %s", err)
		}

		if doc != nil {
			doc.Log()
		}
	}()

	wg.Wait()
}

func findFilter(filter *Filters, sleep int64) (*Model, error) {

	coll := _coll.Clone()

	sLog.Warn("find Filter strategie: %s", filter.String())
	coll.Filters(filter).Log()

	time.Sleep(time.Duration(sleep) * time.Second)

	var doc Model
	err := coll.FindOne(&doc)
	if err != nil {
		return nil, fmt.Errorf("coll.FindOne(): %s", err)
	}

	sLog.Warn("found strategie: %s", doc.String())

	return &doc, nil
}
