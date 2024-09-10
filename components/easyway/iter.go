package easyway

import (
	"fmt"

	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/repos/rIter"
	"github.com/yasseldg/go-simple/repos/rMongo"
	"github.com/yasseldg/go-simple/types/sTime"
)

type InterIter interface {
	rIter.Inter

	EwType() string

	Item() InterEasyWay
	Ew() InterEwType

	SetTsFrom(int64)
	SetTsTo(int64)

	Clone() InterIter
}

type Iter struct {
	rIter.Inter

	ew_type string

	ts_from int64
	ts_to   int64

	item  InterEasyWay
	items EasyWays
}

func NewIter(coll rMongo.InterColl, ew_type string) *Iter {

	sort := rMongo.NewSort().TsAsc()

	coll.Sorts(sort)
	coll.Limit(500)

	return &Iter{
		Inter:   rIter.New(coll, nil),
		ew_type: ew_type,
	}
}

func (iter *Iter) String(name string) string {
	return fmt.Sprintf("%s ts_from: %s  ..  ts_to: %s", iter.Inter.String(name), sTime.ForLog(iter.ts_from, 0), sTime.ForLog(iter.ts_to, 0))
}

func (iter *Iter) Log(name string) {
	sLog.Info(iter.String(name))
}

func (iter *Iter) EwType() string {
	return iter.ew_type
}

func (iter *Iter) Item() InterEasyWay {
	return iter.item
}

func (iter *Iter) Ew() InterEwType {
	return iter.item.Ew(iter.ew_type)
}

func (iter *Iter) Next() bool {
	if !iter.Inter.Next() {
		return false
	}

	if len(iter.items) > 0 {
		iter.item = iter.items[0]
		iter.items = iter.items[1:]
		return true
	}

	filter := iter.Filter()
	filter.Ts(iter.ts_from, iter.ts_to)

	// sLog.Warn("next: filter: %v", filter)

	var items EasyWays
	err := iter.Coll().Filters(filter).Find(&items)
	if err != nil {
		iter.SetError(fmt.Errorf("next: coll.Find: %s", err))
		return false
	}

	if len(items) == 0 {
		iter.SetEmpty(true)
		return false
	}

	iter.items = items
	iter.ts_from = items[len(items)-1].Ts() + 1

	return iter.Next()
}

func (iter *Iter) SetTsFrom(ts_from int64) {
	iter.ts_from = ts_from
}

func (iter *Iter) SetTsTo(ts_to int64) {
	iter.ts_to = ts_to
}

func (iter *Iter) Clone() InterIter {
	return NewIter(iter.Coll(), iter.ew_type)
}
