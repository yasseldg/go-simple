package easyway

import (
	"fmt"

	"github.com/yasseldg/go-simple/repos/rFilter"
	"github.com/yasseldg/go-simple/repos/rMongo"
)

type InterRepo interface {
	Log()

	FindTs(ts int64) error
	Item() InterEasyWay
	Ew() InterEwType
}

type Repo struct {
	coll   rMongo.InterColl
	filter rFilter.Filters

	ew_type string

	item InterEasyWay
}

func NewRepo(coll rMongo.InterColl, ew_type string) *Repo {
	filter := rMongo.NewFilter()

	sort := rMongo.NewSort().TsAsc()

	coll.Sorts(sort)
	coll.Limit(500)

	return &Repo{
		coll:    coll,
		filter:  *filter,
		ew_type: ew_type,
	}
}

func (repo *Repo) FindTs(ts int64) error {

	filter := repo.filter.Clone()
	filter.Ts(ts, (ts + 1))

	var item EasyWay
	err := repo.coll.Filters(filter).FindOne(&item)
	if err != nil {
		return fmt.Errorf("coll.FindOne: %s", err)
	}

	repo.item = &item

	return nil
}

func (repo *Repo) Item() InterEasyWay {
	return repo.item
}

func (repo *Repo) Ew() InterEwType {
	return repo.item.Ew(repo.ew_type)
}

func (repo *Repo) Log() {
	repo.item.Log(repo.ew_type)
}
